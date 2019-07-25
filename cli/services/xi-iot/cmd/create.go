package cmd

import (
	"strings"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"

	iot_yaml "xi-iot-cli/xi-iot/yaml"

	"github.com/go-yaml/yaml"

	"fmt"

	"github.com/spf13/cobra"
)

// DataSources defines the json for data sources input file
type DataSources struct {
	DataSources []models.DataSource `json:"datasources"`
}

// DataPipelines defines the json for data pipeline input file
type DataPipelines struct {
	DataPipelines []models.DataPipeline `json:"datapipelines"`
}

func validateDataSource(ds iot_yaml.DataSource) *errutils.XiErr {
	if ds.Edge == "" {
		return errutils.NewInvalidYamlErr("missing \"edge\"")
	}
	if ds.Name == "" {
		return errutils.NewInvalidYamlErr("missing \"name\"")
	}
	return nil
}

func createDataSource(ds map[interface{}]interface{}, dryRun bool) *errutils.XiErr {
	out, err := yaml.Marshal(ds)
	if err != nil {
		return errutils.NewSerializeErr(dataSourceEntity, "yaml", err.Error())
	}

	dataSrc := iot_yaml.DataSource{}
	err = yaml.Unmarshal(out, &dataSrc)
	if err != nil {
		// Intentionally using Serialize err instead of deserialize because for the end user
		// it is still serialization
		return errutils.NewSerializeErr(dataSourceEntity, "yaml", err.Error())
	}

	errutils.CheckErr(validateDataSource(dataSrc))
	dataSource := iot_yaml.ToDataSrc(dataSrc, Cfg.TenantID, CMEClient)
	if dryRun {
		fmt.Println("\n*******************\ncreating data source\n*******************\n")
		io_utils.PrettyPrintJSON(dataSource)
	} else {
		_, apiErr := CMEClient.CreateDataSource(&dataSource)
		if apiErr != nil {
			return errutils.NewCreateErr(dataSourceEntity, *dataSource.Name, *apiErr.Message)
		}
		fmt.Printf("Successfully created data source: %s\n", *dataSource.Name)
	}
	return nil
}

func validateApplication(app iot_yaml.Application) *errutils.XiErr {
	if app.Name == "" {
		return errutils.NewInvalidYamlErr("missing \"name\"")
	}
	if app.Project == "" {
		return errutils.NewInvalidYamlErr("missing \"project\"")
	}
	if app.Description == "" {
		return errutils.NewInvalidYamlErr("missing \"description\"")
	}
	if app.AppYamlPath == "" {
		return errutils.NewInvalidYamlErr("missing \"appYamlPath\"")
	}
	// validate that only one type of selector is specified and matches project selector
	project := iot_yaml.ToProject(app.Project, CMEClient)
	if *project.EdgeSelectorType == "Explicit" && app.EdgeSelectors != nil {
		errutils.NewInvalidYamlErr("mismatch selection criteria for application and project: project uses explicit edge selection but app specified category based edge selection")
	}
	if *project.EdgeSelectorType == "Category" && app.Edges != nil {
		errutils.NewInvalidYamlErr("mismatch selection criteria for application and project: project uses category based edge selection but app specified explicit edge selection")
	}

	return nil
}

func createApplication(app map[interface{}]interface{}, dryRun bool) *errutils.XiErr {
	out, err := yaml.Marshal(app)
	if err != nil {
		return errutils.NewSerializeErr(applicationEntity, "yaml", err.Error())
	}
	application := iot_yaml.Application{}
	err = yaml.Unmarshal(out, &application)
	if err != nil {
		return errutils.NewSerializeErr(applicationEntity, "yaml", err.Error())
	}
	errutils.CheckErr(validateApplication(application))
	iotApp := iot_yaml.ToApplication(application, CMEClient)
	if dryRun {
		fmt.Println("\n*******************\ncreating application\n*******************\n")
		io_utils.PrettyPrintJSON(iotApp)
	} else {
		_, apiErr := CMEClient.CreateApp(iotApp)
		if apiErr != nil {
			return errutils.NewCreateErr(applicationEntity, *iotApp.Name, *apiErr.Message)
		}
		fmt.Printf("Successfully created application: %s\n", application.Name)
	}
	return nil
}

// TODO: Check if we can validate the document using cloud schema or something else.
func validateDataPipeline(dp iot_yaml.DataPipeline) *errutils.XiErr {
	if dp.Name == "" {
		return errutils.NewInvalidYamlErr("missing \"name\"")
	}
	if dp.Project == "" {
		return errutils.NewInvalidYamlErr("missing \"project\"")
	}
	if dp.Project == "" {
		return errutils.NewInvalidYamlErr("missing \"project\"")
	}
	if dp.Description == "" {
		return errutils.NewInvalidYamlErr("missing \"description\"")
	}

	if dp.Input == nil {
		return errutils.NewInvalidYamlErr("missing \"input\"")
	}

	if dp.Output == nil {
		return errutils.NewInvalidYamlErr("missing \"output\"")
	}

	if dp.Input.DataPipeline == nil && dp.Input.CatSels == nil || dp.Input.DataPipeline != nil && dp.Input.CatSels != nil {
		return errutils.NewInvalidYamlErr("exactly one out of input.realTimeStream or input.categorySelectors is required")
	}

	if dp.Output.PublicCloud != nil && dp.Output.LocalEdge != nil {
		return errutils.NewInvalidYamlErr("exactly one out of output.externalCloud or output.localEdge is required")
	}

	if dp.Output.PublicCloud != nil && strings.ToLower(dp.Output.PublicCloud.Type) != "aws" && strings.ToLower(dp.Output.PublicCloud.Type) != "gcp" {
		return errutils.NewInvalidYamlErr("supported external clouds: AWS or GCP")
	}

	if dp.Output.LocalEdge != nil && dp.Output.LocalEdge.EndpointName == "" {
		return errutils.NewInvalidYamlErr("missing \"localEdge.endpointName\"")
	}

	if dp.Output.PublicCloud != nil && dp.Output.PublicCloud.EndpointName == "" {
		return errutils.NewInvalidYamlErr("missing \"localEdge.publicCloud\"")
	}
	return nil
	// Validate identifier is specified with localEdge
}

func validateCategory(cat iot_yaml.Category) *errutils.XiErr {
	if cat.Name == "" {
		return errutils.NewInvalidYamlErr("name is a required attribute for a category")
	}
	if len(cat.Values) == 0 {
		return errutils.NewInvalidYamlErr("atleast one value is required for a category")
	}
	return nil
}

func encodeToObj(obj map[interface{}]interface{}, i interface{}) {
	out, err := yaml.Marshal(obj)
	if err != nil {
		errutils.Exitf("failed to marshal yaml: %s", err.Error())
	}

	err = yaml.Unmarshal(out, i)
	if err != nil {
		errutils.Exitf("failed to unmarshal yaml. %s\n", err.Error())
	}
}

func createDataPipeline(dp map[interface{}]interface{}, dryRun bool) *errutils.XiErr {
	dataPipeline := iot_yaml.DataPipeline{}
	encodeToObj(dp, &dataPipeline)
	errutils.CheckErr(validateDataPipeline(dataPipeline))

	iotPipeline := iot_yaml.ToDataPipeline(dataPipeline, CMEClient, Cfg.TenantID)
	if dryRun {
		fmt.Println("\n*******************\ncreating data pipeline\n*******************\n")
		io_utils.PrettyPrintYaml(iotPipeline)
	} else {
		_, apiErr := CMEClient.CreateDataPipeline(iotPipeline)
		if apiErr != nil {
			return errutils.NewCreateErr(dataPipelineEntity, *iotPipeline.Name, *apiErr.Message)
		}
		fmt.Printf("successfully created datapipeline %s\n", *iotPipeline.Name)
	}
	return nil
}

func createCategory(cat map[interface{}]interface{}, dryRun bool) *errutils.XiErr {
	category := iot_yaml.Category{}
	encodeToObj(cat, &category)
	errutils.CheckErr(validateCategory(category))

	iotCat := iot_yaml.ToCategory(category)
	if dryRun {
		fmt.Println("\n*******************\ncreating category\n*******************\n")
		io_utils.PrettyPrintYaml(iotCat)
	} else {
		_, apiErr := CMEClient.CreateCategory(iotCat)
		if apiErr != nil {
			return errutils.NewCreateErr(categoryEntity, *iotCat.Name, *apiErr.Message)
		}
		fmt.Printf("successfully created category %s\n", *iotCat.Name)
	}
	return nil
}

func validateFunction(trans iot_yaml.Transformation) *errutils.XiErr {
	if trans.Name == "" {
		return errutils.NewInvalidYamlErr("missing \"name\"")
	}
	if trans.Project == "" {
		return errutils.NewInvalidYamlErr("missing \"project\"")
	}
	if trans.SourceCodePath == "" {
		return errutils.NewInvalidYamlErr("missing \"sourceCodePath\"")
	}
	if trans.Language == "" {
		return errutils.NewInvalidYamlErr("missing \"language\"")
	}
	if trans.Environment == "" {
		return errutils.NewInvalidYamlErr("missing \"environment\"")
	}
	for _, params := range trans.Params {
		if params.Type != "string" && params.Type != "integer" {
			return errutils.NewInvalidYamlErr("params.type can only be either \"string\" or \"integer\"")
		}
	}
	return nil
}

func createFunction(trans map[interface{}]interface{}, dryRun bool) *errutils.XiErr {
	out, err := yaml.Marshal(trans)
	if err != nil {
		errutils.Exit(fmt.Sprintf("failed to marshal yaml: %s", err.Error()))
	}

	transformation := iot_yaml.Transformation{}
	err = yaml.Unmarshal(out, &transformation)
	if err != nil {
		errutils.Exit(fmt.Sprintf("failed to unmarshal function yaml. %s", err.Error()))
	}

	errutils.CheckErr(validateFunction(transformation))
	function := iot_yaml.ToFunction(transformation, CMEClient)
	if dryRun {
		fmt.Println("\n*******************\ncreating function\n*******************\n")
		io_utils.PrettyPrintJSON(function)
	} else {
		_, apiErr := CMEClient.CreateFunction(function)
		if apiErr != nil {
			return errutils.NewCreateErr(functionEntity, *function.Name, *apiErr.Message)
		}
		fmt.Printf("Successfully created function: %s\n", *function.Name)
	}
	return nil
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates Xi IoT resources such as applications, data sources, categories, data pipelines, functions, and so on.",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input-file")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		objs, err := io_utils.ReadYaml(inputFile)
		errutils.CheckErr(err)
		for _, o := range objs {
			obj := o.(map[interface{}]interface{})
			switch strings.ToLower(obj["kind"].(string)) {
			case strings.ToLower(dataSourceEntity):
				errutils.CheckErr(createDataSource(obj, dryRun))
			case strings.ToLower(applicationEntity):
				errutils.CheckErr(createApplication(obj, dryRun))
			case strings.ToLower(dataPipelineEntity):
				errutils.CheckErr(createDataPipeline(obj, dryRun))
			case strings.ToLower(categoryEntity):
				errutils.CheckErr(createCategory(obj, dryRun))
			case strings.ToLower(functionEntity):
				errutils.CheckErr(createFunction(obj, dryRun))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().Bool("dry-run", false, "whether to perform dry run or not")
	createCmd.Flags().MarkHidden("dry-run")
	createCmd.Flags().StringP("input-file", "f", "", "Input file in YAML format for the resource you want to create")
	createCmd.MarkFlagRequired("input-file")
}
