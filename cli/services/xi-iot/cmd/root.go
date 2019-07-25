// Copyright © 2019 Pankit Thapar <pankit.thapar@nutanix.com>

package cmd

import (
	"os"
	"xi-iot-cli/xi-iot/cloudmgmt"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"
	xi_models "xi-iot-cli/xi-iot/models"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
)

const (
	dataSourceEntity   = "DataSource"
	dataPipelineEntity = "Datapipeline"
	applicationEntity  = "Application"
	categoryEntity     = "Category"
	functionEntity     = "Function"
)

var (
	cfgFile string
	// CMEClient encapsulates the cloud mgmt client
	CMEClient   cloudmgmt.Client
	CMECache    *cloudmgmt.Cache
	nonInitCmds = map[string]bool{
		"create-context [CONTEXT_NAME]": true,
		"get-contexts":                  true,
		"use-context [CONTEXT_NAME]":    true,
		"delete-context [CONTEXT_NAME]": true,
		"help [command]":                true,
		"config":                        true,
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "xi-iot",
	PersistentPreRun: initCMEClient,
	Long: io_utils.LongDesc(`
	A CLI for managing your Xi IoT resources

	Find more information at https://github.com/nutanix/xi-iot/tree/master/cli/overview.md`),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errutils.Exitf(err.Error())
	}
}

func initCMEClient(cmd *cobra.Command, args []string) {
	// check if this command is supposed to be skipped
	if nonInitCmds[cmd.Use] {
		return
	}
	contexts := xi_models.Contexts{}
	err := io_utils.ReadJSON(ContextFilePath(), &contexts)
	if os.IsNotExist(err) {
		errutils.Exitf("No existing context. Use unique configuration contexts for different tenants, users, and so on. Please create a new context using the xi-iot config create-context command.")
	}
	updateChan := make(chan xi_models.Context)

	// TODO: check if current context is set or not
	c := cloudmgmt.New(contexts.Contexts[contexts.CurrentCtx].URL,
		cloudmgmt.WithContext(contexts.Contexts[contexts.CurrentCtx]),
		cloudmgmt.WithCtxUpdateChan(updateChan),
	)

	SetCfg(contexts.Contexts[contexts.CurrentCtx])

	// Start a goroutine to recieve context updates from cloud mgmt
	go func(name string, updateChan <-chan xi_models.Context) {
		for {
			select {
			case ctx := <-updateChan:
				err := upsertContext(name, ctx)
				if err != nil {
					glog.Warningf("Failed to update or set the context: %s. %s", name, err.Error())
				}
				glog.V(5).Infof("Upserted context %s: %+v", name, ctx)
			}
		}
	}(contexts.CurrentCtx, updateChan)

	// Call refresh token once before executing commands
	// This helps in case any of the CME client  API's are not calling refresh token themselves
	c.RefreshToken()
	CMEClient = *c
	CMECache = cloudmgmt.NewCache(c)
}
