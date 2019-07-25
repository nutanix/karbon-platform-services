package cmd

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:                   "delete",
	Short:                 "Delete xi-iot resource",
	DisableFlagsInUseLine: true,
	Run:                   func(cmd *cobra.Command, args []string) {},
}

func SubscribeToErrors(prefixErr, successErr string, errCh <-chan *models.APIErrorPayload, timeout time.Duration) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout*time.Second))
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			glog.Errorf("%s. timed out\n", prefixErr)
			return
		case err := <-errCh:
			if err != nil && !strings.Contains(*err.Message, "[404]") {
				glog.Errorf("%s. %s", prefixErr, *err.Message)
			} else {
				fmt.Println(successErr)
			}
			return
		}
	}
}

// deleteDataSrcCmd deletes the given data source
var deleteDataSourceCmd = &cobra.Command{
	Use:                   "datasource",
	Short:                 "Delete data source(s)",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete data source(s) by names
	xi-iot delete datasrc video-stream youtube-rtmp phone`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		wg := sync.WaitGroup{}
		dataSources, err := CMEClient.ListDataSourcesByNames(names)
		if err != nil {
			errutils.Exitf("failed to list data sources. %s", *err.Message)
		}
		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		for _, ds := range dataSources {
			delete(deleteMap, *ds.Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete datasrc %s", ds.Name),
					fmt.Sprintf("successfully deleted datasrc %s", *ds.Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteDataSrc(ID)
				errCh <- err
			}(ds.ID, errCh)
		}
		wg.Wait()

		for n := range deleteMap {
			fmt.Printf("successfully deleted datasrc %s\n", n)
		}
	},
}

// deleteCategoryCmd deletes the given category
var deleteCategoryCmd = &cobra.Command{
	Use:                   "category",
	Short:                 "Delete categories",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete categories by names
	xi-iot delete category cat1 cat2`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		cats, err := CMEClient.ListCategoriesByNames(names)
		if err != nil {
			errutils.Exitf("failed to list categories %+v. %s", strings.Join(names, ","), *err.Message)
		}
		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		wg := sync.WaitGroup{}
		for _, cat := range cats {
			delete(deleteMap, *cat.Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete category %s", *cat.Name),
					fmt.Sprintf("successfully deleted category %s", *cat.Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteCategory(ID)
				errCh <- err
			}(cat.ID, errCh)
		}

		wg.Wait()
		for n := range deleteMap {
			fmt.Printf("successfully deleted category %s\n", n)
		}

	},
}

// deleteCategory deletes the given category
var deleteAppCmd = &cobra.Command{
	Use:                   "application",
	Short:                 "Delete applications(s)",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete apps by names (names are separated by space)
	xi-iot delete app echoapp demoapp`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		apps, err := CMEClient.ListAppsByNames(names)
		if err != nil {
			errutils.Exitf("failed to list apps: %s", *err.Message)
		}
		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		for _, app := range apps {
			delete(deleteMap, *app.Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete app %s", *app.Name),
					fmt.Sprintf("successfully deleted app %s", *app.Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteApp(ID)
				errCh <- err
			}(app.ID, errCh)
		}
		wg.Wait()
		for n := range deleteMap {
			fmt.Printf("successfully deleted application %s\n", n)
		}
	},
}

// deleteCategory deletes the given category
var deleteEdgeCmd = &cobra.Command{
	Use:                   "edge",
	Short:                 "Delete edge",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete edge by Names
	xi-iot delete edge edge-foo edge-bar`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		wg := sync.WaitGroup{}
		edges, err := CMEClient.ListEdgesByNames(names)
		if err != nil {
			errutils.Exitf("failed to list edges: %s", *err.Message)
		}
		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		// TODO: emit a sucessful message for edges not found
		for _, e := range edges {
			delete(deleteMap, *e.Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete edge %s", *e.Name),
					fmt.Sprintf("successfully deleted edge %s", *e.Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteEdge(ID)
				// TODO: Parse out error object from err and display that
				errCh <- err
			}(e.ID, errCh)
		}

		wg.Wait()
		for n := range deleteMap {
			fmt.Printf("successfully deleted edge %s\n", n)
		}
	},
}

// deleteCategory deletes the given category
var deleteDataPipelineCmd = &cobra.Command{
	Use:                   "datapipeline",
	Short:                 "Delete the given datapipeline(s)",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete edge by Names
	xi-iot delete datapipeline foo bar`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		pipelines, err := CMEClient.ListDataPipelinesByNames(names)
		if err != nil {
			errutils.Exitf("failed to list datapipelines: %s", *err.Message)
		}

		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		// TODO: emit a sucessful message for datapipelines not found
		wg := sync.WaitGroup{}
		for i, pipeline := range pipelines {
			delete(deleteMap, *pipelines[i].Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete datapipeline %s", *pipelines[i].Name),
					fmt.Sprintf("successfully deleted datapipeline %s", *pipelines[i].Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteDataPipeline(ID)
				// TODO: Parse out error object from err and display that
				errCh <- err
			}(pipeline.ID, errCh)
		}
		wg.Wait()

		for n := range deleteMap {
			fmt.Printf("successfully deleted datapipeline %s\n", n)
		}
	},
}

// deleteFunction deletes the given function
var deleteFunctionCmd = &cobra.Command{
	Use:                   "function",
	Short:                 "Delete function(s)",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Delete functions by names (names are separated by space)
	xi-iot delete function echofunction demofunction`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}
		names := make([]string, 0, len(args))
		for _, a := range args {
			names = append(names, strings.TrimSpace(a))
		}
		functions, err := CMEClient.ListFunctionByNames(names)
		if err != nil {
			errutils.Exitf("failed to list functions: %s", *err.Message)
		}
		deleteMap := make(map[string]bool)
		for _, n := range names {
			deleteMap[n] = true
		}

		for _, function := range functions {
			delete(deleteMap, *function.Name)
			errCh := make(chan *models.APIErrorPayload)
			// TODO: Check if 10s is greater than cloud mgmt context or not. this is to make sure
			// that we don't cancel listening to error before API request is cancelled
			go func() {
				defer wg.Done()
				SubscribeToErrors(fmt.Sprintf("failed to delete function %s", *function.Name),
					fmt.Sprintf("successfully deleted function %s", *function.Name), errCh, 10)
			}()
			wg.Add(1)
			go func(ID string, errCh chan<- *models.APIErrorPayload) {
				err := CMEClient.DeleteFunction(ID)
				errCh <- err
			}(function.ID, errCh)
		}
		wg.Wait()
		for n := range deleteMap {
			fmt.Printf("successfully deleted function %s\n", n)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteDataSourceCmd)
	deleteCmd.AddCommand(deleteCategoryCmd)
	deleteCmd.AddCommand(deleteAppCmd)
	deleteCmd.AddCommand(deleteFunctionCmd)
	// Deleting edge is disabled for now.
	// deleteCmd.AddCommand(deleteEdgeCmd)
	deleteCmd.AddCommand(deleteDataPipelineCmd)
}
