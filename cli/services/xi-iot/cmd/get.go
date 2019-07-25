package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"
	iot_yaml "xi-iot-cli/xi-iot/yaml"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
)

const (
	FormatJSON  = "json"
	FormatYAML  = "yaml"
	FormatTable = "table"
)

var (
	DefaultDataSourceHeader         = []string{"Name", "Edge", "Protocol", "Topics"}
	DefaultEndpointsDataHeader      = []string{"Name", "Endpoints", "client secret"}
	DefaultCategoryHeader           = []string{"Name", "Description", "Values"}
	DefaultEdgeHeader               = []string{"Name", "IP Address", "Connected", "Serial Number"}
	DefaultAppHeader                = []string{"Name", "Project"}
	DefaultProjectHeader            = []string{"Name", "Edges", "CloudProfiles"}
	DefaultDataPipelineHeader       = []string{"Name", "Project", "Input Type", "Output Type", "Functions"}
	DefaultRuntimeEnvironmentHeader = []string{"Name", "ID", "Language", "Created At"}
	DefaultFunctionHeader           = []string{"Name", "Project", "Language", "Environment"}
)

type Endpoint struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// cobra doesn't support nested level of Pre-run methods, hence,
// use putting this one out as a helper.
func getCmdPreRun(cmd *cobra.Command) {
	out, _ := cmd.Flags().GetString("output-format")
	if strings.ToLower(out) != FormatTable && strings.ToLower(out) != FormatYAML {
		errutils.Exitf("allowed output-format(s) are yaml, table(default)")
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:                   "get",
	Short:                 "Get/retrieve Xi IoT resources",
	Run:                   func(cmd *cobra.Command, args []string) {},
	DisableFlagsInUseLine: true,
}

// edgeGetCmd represents the get edge command
var edgeGetCmd = &cobra.Command{
	Use:                   "edge",
	Short:                 "Get Xi IoT edges",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all edges in tabular format
	xi-iot get edge

	# Get all connected edges in tabular format
	xi-iot get edge -c

	# Get all edges in yaml format
	xi-iot get edge -o yaml

	# Get all connected edges in yaml format
	xi-iot get edge -o yaml

	# Get a specific edge in tabular format
	xi-iot get edge edge-sjc

	# Get a specific edge in yaml format
	xi-iot get edge edge-sjc -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		// regex, _ := cmd.Flags().GetString("regex")
		connected, _ := cmd.Flags().GetBool("only-connected")
		out, _ := cmd.Flags().GetString("output-format")
		edges := []*models.EdgeV2{}
		var err *models.APIErrorPayload
		if len(args) == 0 {
			edges, err = CMEClient.ListEdges("")
		} else {
			edges, err = CMEClient.ListEdges(fmt.Sprintf("name LIKE '%s'", args[0]))
		}

		if err != nil {
			errutils.Exitf("failed to fetch edge(s). %s", *err.Message)
		}

		data := [][]string{}
		for _, e := range edges {
			if connected && !e.Connected {
				continue
			}
			iotEdge := iot_yaml.FromEdge(e, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotEdge)
			case FormatTable:
				// []string{"Name", "IP Address", "Connected", "Serial Number"}
				data = append(data, []string{iotEdge.Name, iotEdge.IPAddress, strconv.FormatBool(iotEdge.Connected), iotEdge.SerialNumber})
			}
		}

		if out == FormatTable {
			io_utils.PrintTable(data, DefaultEdgeHeader)
		}
	},
}

var dataSourceGetCmd = &cobra.Command{
	Use:                   "datasource",
	Short:                 "Get data sources",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all data sources in tabular format
	xi-iot get datasource

	# Get all datasources in  yaml format
	xi-iot get datasource -o yaml

	# Get all data sources on an edge in tabular format
	xi-iot get datasource -e edge-sjc

	# Get all data sources on an edge in yaml format
	xi-iot get datasource -e edge-sjc -o yaml

	# Get a specific data source in tabular format
	xi-iot get datasource mqtt-datasource

	# Get a specific data source in yaml format
	xi-iot get datasource mqtt-datasource -o yaml

	# Get data sources for an edge by edge name.(in yaml format)
	xi-iot get datasource -e edge-sjc -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		edge, _ := cmd.Flags().GetString("edge")
		out, _ := cmd.Flags().GetString("output-format")
		// hidden option
		showArtifacts, _ := cmd.Flags().GetBool("show-artifacts")
		dataSources := []*models.DataSourceV2{}
		var err *models.APIErrorPayload
		if len(args) != 0 {
			dataSources, err = CMEClient.ListDataSources(fmt.Sprintf("name LIKE '%s'", args[0]))
		} else if edge != "" {
			edges, err := CMEClient.ListEdgesByNames([]string{edge})
			if err != nil {
				errutils.Exitf("failed to find edge %s. %s", edge, *err.Message)
			}
			if len(edges) == 0 {
				errutils.Exitf("edge %s does not exist", edge)
			}
			dataSources, err = CMEClient.ListDataSourcesByEdgeID(edges[0].ID)
		} else {
			// get all data sources
			dataSources, err = CMEClient.ListDataSources(fmt.Sprintf(""))
		}

		if err != nil {
			errutils.Exitf("failed to get data sources. %s", *err.Message)
		}
		data := [][]string{}
		for _, ds := range dataSources {
			iotDataSrc := iot_yaml.FromDataSource(ds, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotDataSrc)
			case FormatTable:
				topics := make([]string, 0, len(iotDataSrc.Fields))
				for _, f := range iotDataSrc.Fields {
					topics = append(topics, f.Topic)
				}
				// []string{"Name", "Edge", Protocol", "Topics"}
				data = append(data, []string{iotDataSrc.Name, iotDataSrc.Edge, iotDataSrc.Protocol, strings.Join(topics, "\n")})
			}
		}
		if out == FormatTable {
			io_utils.PrintTable(data, DefaultDataSourceHeader)
		}
		if !showArtifacts {
			return
		}
		endpointsData := [][]string{}
		data = [][]string{}
		for _, ds := range dataSources {
			artifact := CMEClient.GetDataSourceArtifact(ds.ID)
			topics := []string{}
			for _, t := range ds.FieldsV2 {
				topics = append(topics, *t.Topic)
			}
			b, jsonErr := json.Marshal(artifact.Data["endpoints"])
			if jsonErr != nil {
				errutils.Exitf("failed to marshal endpoints: %s", jsonErr.Error())
			}
			endpoints := []Endpoint{}
			jsonErr = json.Unmarshal(b, &endpoints)
			if jsonErr != nil {
				errutils.Exitf("failed to unmarshal endpoints: %s", jsonErr.Error())
			}
			endpointsStr := []string{}
			for _, e := range endpoints {
				endpointsStr = append(endpointsStr, fmt.Sprintf("%s: %s", e.Name, e.URL))
			}

			clientSecret := ""
			if artifact.Data["clientsecret"] != nil {
				clientSecret = artifact.Data["clientsecret"].(string)
			}
			endpointsData = append(endpointsData, []string{*ds.Name, strings.Join(endpointsStr, "\n"), clientSecret})
			data = append(data, []string{*ds.Name, ds.ID, strings.Join(topics, "\n"), *ds.EdgeID})
		}
		fmt.Printf("\n\n")
		io_utils.PrintTable(endpointsData, DefaultEndpointsDataHeader)
	},
}

var dataPipelineGetCmd = &cobra.Command{
	Use:                   "datapipeline",
	Short:                 "Get data pipelines",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all data pipelines in tabular format
	xi-iot get datapipeline

	# Get all data pipelines in yaml format
	xi-iot get datapipeline -o yaml

	# Get all data pipelines for a given project in tabular format
	xi-iot get datapipeline -p project-foo

	# Get all data pipelines for a given project in yaml format
	xi-iot get datapipeline -p project-foo -o yaml

	# Get a specific data  pipeline in tabular format
	xi-iot get datapipeline detect-pipeline

	# Get a specific data  pipeline in yaml format
	xi-iot get datapipeline detect-pipeline`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		project, _ := cmd.Flags().GetString("project")
		out, _ := cmd.Flags().GetString("output-format")

		var err *models.APIErrorPayload

		var datapipelines []*models.DataPipeline
		if len(args) != 0 {
			datapipelines, err = CMEClient.ListDataPipelines(fmt.Sprintf("name = '%s'", args[0]))
		} else if project != "" {
			projects, err := CMEClient.ListProjectByNames([]string{project})
			if err != nil {
				errutils.Exitf("failed to find project %s. %s", project, *err.Message)
			}
			if len(projects) == 0 {
				errutils.Exitf("project %s does not exist", project)
			}
			datapipelines, err = CMEClient.GetDataPipelinesProjects(projects[0].ID)
		} else {
			datapipelines, err = CMEClient.ListDataPipelines("")
		}

		if err != nil {
			errutils.Exitf("failed to get data pipelines. %s", *err.Message)
		}
		data := [][]string{}
		for _, dp := range datapipelines {
			iotDataPipeline := iot_yaml.FromDatapipeline(dp, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotDataPipeline)
			case FormatTable:
				// TODO: Convert these to consts
				// []string{"Name", "Project", "InputType", "OutputType", "Functions"}
				inputType := "dataPipeline"
				if iotDataPipeline.Input.CatSels != nil {
					inputType = "categorySelectors"
				}
				outputType := "publicCloud"
				if iotDataPipeline.Output.LocalEdge != nil {
					outputType = "localEdge"
				}
				functions := make([]string, 0, len(iotDataPipeline.Functions))
				for _, f := range iotDataPipeline.Functions {
					functions = append(functions, f.Name)
				}
				data = append(data, []string{iotDataPipeline.Name, iotDataPipeline.Project, inputType, outputType, strings.Join(functions, "\n")})
			}
		}

		if out == FormatTable {
			io_utils.PrintTable(data, DefaultDataPipelineHeader)
		}
	},
}

var runtimeEnvironmentGetCmd = &cobra.Command{
	Use:                   "runtime",
	Short:                 "Get runtimes",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all runtimes where the runtime name starts with the given prefix python-trans1.
	xi-iot get runtime -r "python-trans1%"
	
	# Get all runtimes where the runtime name is python-trans1.
	xi-iot get runtime -r "python-trans1"`),
	Run: func(cmd *cobra.Command, args []string) {
		regex, _ := cmd.Flags().GetString("regex")
		if regex == "" {
			glog.Exit("missing prefix option")
		}
		var err *models.APIErrorPayload
		runtimes, err := CMEClient.ListRuntimeEnvironments(fmt.Sprintf("name LIKE '%s%s'", regex, "%"))
		if err != nil {
			errutils.Exitf("failed to  get categories. %s", *err.Message)
		}
		out, _ := cmd.Flags().GetString("output-format")
		if strings.ToLower(out) == FormatJSON {
			for _, r := range runtimes {
				io_utils.PrettyPrintJSON(*r)
			}
		} else {
			data := [][]string{}
			for _, r := range runtimes {
				data = append(data, []string{*r.Name, r.ID, *r.Language, r.CreatedAt.String()})
			}
			io_utils.PrintTable(data, DefaultRuntimeEnvironmentHeader)
		}
	},
}

var functionGetCmd = &cobra.Command{
	Use:                   "function",
	Short:                 "Get functions/script",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all functions in tabular format
	xi-iot get function

	# Get all functions in yaml format
	xi-iot get function -o yaml

	# Get a specific function in tabular format
	xi-iot get function object-detect-function

	# Get a specific function in yaml format
	xi-iot get function object-detect-function -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		out, _ := cmd.Flags().GetString("output-format")
		var err *models.APIErrorPayload
		var functions []*models.Function
		if len(args) != 0 {
			functions, err = CMEClient.ListFunctions(fmt.Sprintf("name = '%s'", args[0]))
		} else {
			functions, err = CMEClient.ListFunctions("")
		}

		if err != nil {
			errutils.Exitf("failed to get functions. %s", *err.Message)
		}

		data := [][]string{}
		for _, f := range functions {
			iotFunc := iot_yaml.FromFunction(f, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotFunc)
			case FormatTable:
				// 	DefaultFunctionHeader = []string{"Name", "Project", "Language", "Environment"}
				data = append(data, []string{iotFunc.Name, iotFunc.Project, iotFunc.Language, iotFunc.Environment})
			}
		}

		if out == FormatTable {
			io_utils.PrintTable(data, DefaultFunctionHeader)
		}
	},
}

var categoryGetCmd = &cobra.Command{
	Use:   "category",
	Short: "Get categories",
	Example: io_utils.Examples(`
	# Get all categories in tabular format
	xi-iot get category

	# Get all categories in yaml format
	xi-iot get category -o yaml

	# Get a specific category in tabular format
	xi-iot get category sjc-cat

	# Get a specific category in yaml format
	xi-iot get category sjc-cat -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		out, _ := cmd.Flags().GetString("output-format")
		var categories []*models.Category
		var err *models.APIErrorPayload

		if len(args) != 0 {
			categories, err = CMEClient.ListCategories(fmt.Sprintf("name = '%s'", args[0]))
		} else {
			categories, err = CMEClient.ListCategories("")
		}
		if err != nil {
			errutils.Exitf("failed to  get categories. %s", *err.Message)
		}
		data := [][]string{}
		for _, cat := range categories {
			iotCat := iot_yaml.FromCategory(cat, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotCat)
			case FormatTable:
				data = append(data, []string{iotCat.Name, iotCat.Description, strings.Join(iotCat.Values, "\n")})
			}
		}
		if out == FormatTable {
			io_utils.PrintTable(data, DefaultCategoryHeader)
		}
	},
}

var projectGetCmd = &cobra.Command{
	Use:                   "project",
	Short:                 "Get projects",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all projects in tabular format
	xi-iot get project

	# Get all projects in yaml format
	xi-iot get project -o yaml

	# Get a specific project in tabular format
	xi-iot get project project-inference

	# Get a specific project in yaml format
	xi-iot get project project-inference -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		out, _ := cmd.Flags().GetString("output-format")

		var err *models.APIErrorPayload
		var projects []*models.Project
		if len(args) != 0 {
			projects, err = CMEClient.ListProjects(fmt.Sprintf("name = '%s'", args[0]))
		} else {
			projects, err = CMEClient.ListProjects("")
		}

		if err != nil {
			errutils.Exitf("failed to get projects. %s", *err.Message)
		}
		data := [][]string{}
		for _, p := range projects {
			iotProject := iot_yaml.FromProject(p, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotProject)
			case FormatTable:
				data = append(data, []string{iotProject.Name, strings.Join(iotProject.Edges, "\n"), strings.Join(iotProject.CloudProfiles, "\n")})
			}
		}

		if strings.ToLower(out) == FormatTable {
			io_utils.PrintTable(data, DefaultProjectHeader)
		}
	},
}

var appGetCmd = &cobra.Command{
	Use:                   "application",
	Short:                 "Get applications",
	DisableFlagsInUseLine: true,
	Example: io_utils.Examples(`
	# Get all applications in tabular format
	xi-iot get application

	# Get all applications in yaml format
	xi-iot get application -o yaml

	# Get a specific application in tabular format
	xi-iot get application nginx-webserver

	# Get a specific application in yaml format
	xi-iot get application nginx-webserver -o yaml`),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdPreRun(cmd)
		out, _ := cmd.Flags().GetString("output-format")
		var apps []*models.ApplicationV2
		var err *models.APIErrorPayload

		if len(args) != 0 {
			apps, err = CMEClient.ListApps(fmt.Sprintf("name = '%s'", args[0]))
		} else {
			apps, err = CMEClient.ListApps("")
		}

		if err != nil {
			errutils.Exitf("failed to get apps. %s", *err.Message)
		}

		data := [][]string{}
		for _, app := range apps {
			iotApp := iot_yaml.FromApplication(app, CMECache)
			switch strings.ToLower(out) {
			case FormatYAML:
				io_utils.PrettyPrintYaml(iotApp)
			case FormatTable:
				data = append(data, []string{iotApp.Name, iotApp.Project})
			}
		}

		if strings.ToLower(out) == FormatTable {
			io_utils.PrintTable(data, DefaultAppHeader)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(edgeGetCmd)
	getCmd.AddCommand(dataSourceGetCmd)
	getCmd.AddCommand(categoryGetCmd)
	getCmd.AddCommand(dataPipelineGetCmd)
	getCmd.AddCommand(projectGetCmd)
	getCmd.AddCommand(appGetCmd)
	// Disabled for now
	// getCmd.AddCommand(runtimeEnvironmentGetCmd)
	getCmd.AddCommand(functionGetCmd)

	getCmd.PersistentFlags().StringP("output-format", "o", FormatTable, "output format. Example: yaml, table")

	edgeGetCmd.Flags().BoolP("only-connected", "c", false, "Show only connected edges")

	dataSourceGetCmd.Flags().StringP("edge", "e", "", "edge name")
	dataSourceGetCmd.Flags().Bool("show-artifacts", false, "whether to show data source artifacts or not")
	dataSourceGetCmd.Flags().MarkHidden("show-artifacts")
	dataPipelineGetCmd.Flags().StringP("project", "p", "", "project name of the datapipeline")
}
