package cmd

import (
	"fmt"

	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply updates to Xi IoT resources such as applications, data sources, categories, data pipelines, and so on. ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
	},
}

var applyDSrcCmd = &cobra.Command{
	Use:   "datasrc",
	Short: "Apply updates to one or more data sources",
	Example: io_utils.Examples(`
	# Apply updates to one data source, as defined in a YAML input file, to add to the specified Xi edge. 
	# Specify the edge by its serial number. To get the serial number, open a web browser and use the edge IP address as part of this URL: http://edge-ip-address:8080/v1/sn
	# Or use <$ xi-iot get edge -p '%'> to get the serial number
	# samples/datasrc-ifc*.yaml contains example YAML input to define a data source.
	xi-iot apply datasrc -u datasrc-id -f datasrc-ifc-out.yaml -e 719f6bf2-8ae0-408f-a6b5-9b1cb27e2ba7`),
	Run: func(cmd *cobra.Command, args []string) {
		// createDSrcCmd.Run(cmd, args)
	},
}

var applyCategory = &cobra.Command{
	Use:   "category",
	Short: "Apply updates to a category with the given values",
	Example: io_utils.Examples(`
	# Apply updates to a category by specifying a name (-n airport) and one or more comma-separated values (--values sjc,sfo,lax).
	xi-iot apply category -u category-id -n airport --values sjc,sfo,lax
	
	# Apply updates to a category by specifying a new name (-n airport).
	xi-iot apply category -u category-id -n city
	
	# Apply updates to a category by specifying one or more comma-separated values (--values sjc,sfo,lax).
	xi-iot apply category -u category-id --values sf,la`),
	Run: func(cmd *cobra.Command, args []string) {
		values, _ := cmd.Flags().GetStringSlice("values")
		name, _ := cmd.Flags().GetString("name")
		applyID, _ := cmd.Flags().GetString("apply-id")
		if applyID == "" {
			glog.Fatal("apply ID must be provided.")
		}
		filter := fmt.Sprintf("id IN ('%s')", applyID)
		categoryList, err := CMEClient.ListCategories(filter)
		if len(categoryList) != 1 || err != nil {
			errutils.Exitf("Error when listing category with provided id: %s", *err.Message)
		}
		cat := models.Category{Name: categoryList[0].Name, Values: categoryList[0].Values, TenantID: &Cfg.TenantID}
		if len(values) != 0 {
			cat.Values = values
		}
		if name != "" {
			cat.Name = &name
		}
		id, err := CMEClient.UpdateCategory(&cat, applyID)
		if err != nil {
			glog.Exit(*err.Message)
		}
		fmt.Printf("successfully apply updates to category %s", id)
	},
}

var applyApplication = &cobra.Command{
	Use:   "app",
	Short: "Apply updates to an application",
	Example: io_utils.Examples(`
	# Apply updates to an application with a given Xi IoT application YAML template and k8s YAML file. 
	xi-iot apply app -u app-id -k <K8s_YAML_FILE> -f <XI_IOT_YAML>

	# Apply updates to an application with a given Xi IoT application YAML template and k8s YAML file.
	# Optionally, you can specify one or more edges by edge name, separated by commas.
	xi-iot apply app -u app-id -k <K8s_YAML_FILE> -f <XI_IOT_YAML> -e <EDGE_NAME1, EDGE_NAME2, ...>`),
	Run: func(cmd *cobra.Command, args []string) {
		// createApplication.Run(cmd, args)
	},
}

var applyDataPipeline = &cobra.Command{
	Use:     "datapipeline",
	Short:   "fix later ",
	Example: io_utils.Examples(`Fix later`),
	Run: func(cmd *cobra.Command, args []string) {
		// createDataPipeline.Run(cmd, args)
	},
}

func init() {
	// Disable apply command for now as it needs to be completely implemented
	// rootCmd.AddCommand(applyCmd)
	applyCmd.AddCommand(applyDSrcCmd)
	applyCmd.AddCommand(applyDataPipeline)
	applyCmd.AddCommand(applyCategory)
	applyCmd.AddCommand(applyApplication)
	applyCmd.PersistentFlags().StringP("input-file", "f", "", "Input file in YAML format for the resource you want to update")
	applyCmd.PersistentFlags().StringP("update-id", "u", "", "ID of the object to be updated")
	applyCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")

	applyCmd.MarkFlagRequired("input-file")
	applyCmd.MarkFlagRequired("update-id")
	applyDSrcCmd.Flags().StringP("edge-id", "e", "", "Edge ID for data source")
	applyDSrcCmd.MarkFlagRequired("edge-id")

	applyCategory.Flags().StringSliceP("values", "", []string{}, "values for the category")
	applyCategory.Flags().StringP("name", "n", "", "name of the category")

	applyApplication.Flags().StringP("k8s-yaml", "k", "", "k8s yaml for the app")
	applyApplication.MarkFlagRequired("k8s-yaml")
	applyApplication.Flags().StringSliceP("edges", "e", nil, "list of edge names where this app will be deployed. This list overrides any edges specified in the input YAML file")
}
