package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"
	xi_models "xi-iot-cli/xi-iot/models"

	"github.com/spf13/cobra"
)

const (
	ContextsFileName = "contexts.json"
)

// Configure ..
type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	URL      string `json:"url"`
	TenantID string `json:"tenantId"`
}

var (
	Cfg                  Config
	DefaultContextHeader = []string{"Name", "URL", "Tenant ID", "Email", "Password"}
)

func SetCfg(curCtx xi_models.Context) {
	Cfg.Email = curCtx.Email
	Cfg.Password = curCtx.Password
	Cfg.Token = curCtx.Token
	Cfg.URL = curCtx.URL
	Cfg.TenantID = curCtx.TenantID
}

func upsertContext(name string, ctx xi_models.Context) error {
	contexts := xi_models.Contexts{Contexts: make(map[string]xi_models.Context)}
	if contextFileExists() {
		err := io_utils.ReadJSON(ContextFilePath(), &contexts)
		if err != nil {
			return err
		}
	}
	contexts.Contexts[name] = ctx
	SetCfg(ctx)
	if contexts.CurrentCtx == "" {
		contexts.CurrentCtx = name
	}

	return io_utils.WriteJSON(ContextFilePath(), contexts)
}

func deleteContext(name string) error {
	contexts := xi_models.Contexts{Contexts: make(map[string]xi_models.Context)}
	if contextFileExists() {
		err := io_utils.ReadJSON(ContextFilePath(), &contexts)
		if err != nil {
			return err
		}
	}
	if _, ok := contexts.Contexts[name]; !ok {
		fmt.Println("Context %s does not exist", name)
		return nil
	}
	delete(contexts.Contexts, name)
	if contexts.CurrentCtx == name {
		return fmt.Errorf("Cannot delete current context: %s", name)
	}

	return io_utils.WriteJSON(ContextFilePath(), contexts)
}

// configureCmd represents the configure command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage xi-iot configuration",
	Long: io_utils.LongDesc(`
	Manage a Xi IoT configuration using xi-iot subcommands

	The xi-iot config subcommand enables you to create and manage named contexts for your Xi IoT resources.
	You can use unique configuration contexts for different tenants, users, and so on.

	For example, you can create a context where the specified user is an infrastructure admin.`),
	Example: io_utils.Examples(`
	# Create or set a new context
	xi-iot config set-context infra-admin -p test -m test@mynutanix.com

	# Get all available contexts
	xi-iot config get-context 

	# Use a context
	xi-iot config use-context infra-admin`),
	Run: func(cmd *cobra.Command, args []string) {},
}

func ContextFilePath() string {
	cur, _ := user.Current()
	return path.Join(cur.HomeDir, ".xi", ContextsFileName)
}

func contextDirExists() bool {
	_, err := os.Stat(filepath.Dir(ContextFilePath()))
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func contextFileExists() bool {
	_, err := os.Stat(ContextFilePath())
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

// configureCmd represents the configure command
var createContextCmd = &cobra.Command{
	Use:                   "create-context [CONTEXT_NAME]",
	Short:                 "Creates a new context for managing your Xi IoT resources",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing positional argument [CONTEXT_NAME]")
		}
		return nil
	},
	Example: io_utils.Examples(`
	# Create a new infra admin context
	xi-iot config set-context infra-admin -p test -m test@mynutanix.com`),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		cmeURL, _ := cmd.Flags().GetString("url")
		if !contextDirExists() {
			err := os.MkdirAll(filepath.Dir(ContextFilePath()), 0770)
			if err != nil {
				panic(err)
			}
		}

		err := upsertContext(name, xi_models.Context{Email: email, Password: password, URL: cmeURL, TenantID: ""})
		if err != nil {
			errutils.Exit(fmt.Sprintf("failed to set context. %s", err.Error()))
		}
	},
}

// configureCmd represents the configure command
var getContextCmd = &cobra.Command{
	Use:                   "get-contexts",
	Short:                 "Lists available contexts",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get available contexts
	xi-iot config get-contexts`),
	Run: func(cmd *cobra.Command, args []string) {
		contexts := xi_models.Contexts{Contexts: make(map[string]xi_models.Context)}
		err := io_utils.ReadJSON(ContextFilePath(), &contexts)
		if os.IsNotExist(err) {
			fmt.Println("No existing context found. Please create a new context by using the set-context subcommand.")
			return
		}
		data := [][]string{}
		for n, c := range contexts.Contexts {
			data = append(data, []string{n, c.URL, c.TenantID, c.Email, c.Password})
		}
		io_utils.PrintTable(data, DefaultContextHeader)
		fmt.Printf("\n\n")
		io_utils.PrintTable([][]string{[]string{contexts.CurrentCtx}}, []string{"Current Context"})
	},
}

// deleteContextCmd represents the deleteContextCmd command
var deleteContextCmd = &cobra.Command{
	Use:   "delete-context [CONTEXT_NAME]",
	Short: "Deletes a given context",
	Example: io_utils.Examples(`
	xi-iot config delete-context infra-admin`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing positional argument [CONTEXT_NAME]")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := deleteContext(name)
		if err != nil {
			errutils.Exit(fmt.Sprintf("Failed to delete. %s", err.Error()))
		}
		fmt.Printf("Context %s deleted", name)
	},
}

var useContextCmd = &cobra.Command{
	Use:   "use-context [CONTEXT_NAME]",
	Short: "Sets the current configuration context to the given context name",
	Example: io_utils.Examples(`
	# Use infra-admin context
	xi-iot config use-context infra-admin`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing positional argument [CONTEXT_NAME]")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		contexts := xi_models.Contexts{Contexts: make(map[string]xi_models.Context)}
		err := io_utils.ReadJSON(ContextFilePath(), &contexts)
		if os.IsNotExist(err) {
			errutils.Exit("No existing context found. Please create a new context by using the create-context subcommand.")
		}

		if _, ok := contexts.Contexts[args[0]]; !ok {
			errutils.Exit(fmt.Sprintf("could not find context %s", args[0]))
		}
		contexts.CurrentCtx = args[0]
		err = io_utils.WriteJSON(ContextFilePath(), contexts)
		if err != nil {
			errutils.Exit(fmt.Sprintf("failed to write contexts json file at %s. %s", ContextFilePath(), err.Error()))
		}
		io_utils.PrintTable([][]string{[]string{contexts.CurrentCtx}}, []string{"Current Context"})
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(getContextCmd)
	configCmd.AddCommand(createContextCmd)
	configCmd.AddCommand(useContextCmd)
	configCmd.AddCommand(deleteContextCmd)

	createContextCmd.Flags().StringP("email", "m", "", "email of the tenant")
	createContextCmd.Flags().StringP("password", "p", "", "password for the tenant")

	createContextCmd.Flags().StringP("url", "u", "iot.nutanix.com", "URL for the Xi IoT cloud management instance")
	createContextCmd.Flags().MarkHidden("url")
	createContextCmd.MarkFlagRequired("email")
	createContextCmd.MarkFlagRequired("password")
}
