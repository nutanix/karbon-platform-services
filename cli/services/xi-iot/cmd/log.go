package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/devtools"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"
	"xi-iot-cli/xi-iot/yaml"

	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log functions",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var (
	logFetchInterval = 500 * time.Millisecond
)

func streamLogFromEndpoint(URL *string) {
	fmt.Println("Start Streaming...")
	urlList := strings.Split(*URL, "/")
	endpoint := urlList[len(urlList)-1]
	host := urlList[2]
	devtoolsClient := devtools.New(host, devtools.WithCMEClient(&CMEClient))
	ticker := time.NewTicker(logFetchInterval)
	var timeStamp = "0"
	for {
		select {
		case <-ticker.C:
			logs, err := devtoolsClient.FetchLogs(endpoint, timeStamp)
			if err != nil {
				errutils.Exitf("Error in fetching logs. %s", err)
			}
			fmt.Print(logs.Logs)
			timeStamp = *logs.LatestTimeStamp
		}
	}
}

func userSelectContainer(length int) int {
	userInput := 0
	for {
		fmt.Printf("Container Selection (1 - %d): ", length)
		line, readErr := io_utils.Readstdin()
		if readErr != nil {
			fmt.Printf("Error reading input %s\n", readErr)
			continue
		}
		userInput, readErr = strconv.Atoi(line)
		if readErr != nil {
			fmt.Printf("Error converting input to number %s\n", readErr)
			continue
		}
		if userInput < 1 || userInput > length {
			fmt.Println("Selection out of range")
			continue
		}
		return userInput
	}
}

func appDeployedOnEdge(edgeID *string, app *models.ApplicationV2) bool {
	for _, e := range app.EdgeIds {
		if e == *edgeID {
			return true
		}
	}
	return false
}

func getAppContainerName(container string) string {
	return strings.Split(container, ":")[0]
}

var logAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Real-time logging for application",
	Example: io_utils.Examples(`
	# Stream real-time logs of an application from an endpoint
	xi-iot log app test-app -e edgeName -c containerFoo`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		edgeName, _ := cmd.Flags().GetString("edge")
		containerName, _ := cmd.Flags().GetString("container")
		appName := args[0]
		edgeID := yaml.ToEdgeID(edgeName, CMEClient)
		apps, err := CMEClient.ListAppsByNames([]string{appName})
		if err != nil {
			errutils.Exitf("Failed to find application %s. %s", appName, *err.Message)
		}
		if len(apps) == 0 {
			errutils.Exitf("Application %s does not exist", appName)
		}
		app := apps[0]
		// check that the app is deployed on this edge on cli (cloud does the check too)
		if !appDeployedOnEdge(edgeID, app) {
			errutils.Exitf("Application %s is not deployed on edge %s", appName, edgeName)
		}
		// get container
		containers, err := CMEClient.ListAppContainers(app.ID, *edgeID)
		if err != nil {
			errutils.Exitf("Failed to get application containers %s", *err.Message)
		}
		if len(containers) == 0 {
			errutils.Exitf("No application container is found")
		}

		containerMap := make(map[string]string)
		for _, c := range containers {
			cName := strings.Split(c, ":")[0]
			containerMap[cName] = c
		}

		if len(containerMap) > 1 && containerName == "" {
			errutils.Exitf("Container name (-c) required if the application has multiple containers.")
		}
		// Set containerName to the first value
		if containerName == "" {
			containerName = strings.Split(containers[0], ":")[0]
		}
		selectedContainerReplica, ok := containerMap[containerName]
		if !ok {
			cNames := make([]string, len(containerMap))
			for key := range containerMap {
				cNames = append(cNames, key)
			}
			errutils.Exitf("Could not find given container %s. Possible options are %v", containerName, cNames)
		}

		// get stream end point
		logStream := models.LogStream{
			ApplicationID:  app.ID,
			DataPipelineID: "",
			EdgeID:         edgeID,
			Container:      &selectedContainerReplica,
		}
		url, err := CMEClient.GetLogEndpoint(&logStream)
		streamLogFromEndpoint(url)
	},
}

func getFunctionMap(pipe *models.DataPipeline) map[string]*models.Function {
	var pipeIDList []string
	for _, f := range pipe.TransformationArgsList {
		pipeIDList = append(pipeIDList, *f.TransformationID)
	}
	sqlList := io_utils.ToSQLList(pipeIDList)
	filter := fmt.Sprintf("id IN (%s)", sqlList)
	functionList, err := CMEClient.ListFunctions(filter)
	if err != nil {
		errutils.Exitf("Failed to get function list %s", *err.Message)
	}
	functionMap := make(map[string]*models.Function)
	for _, f := range functionList {
		functionMap[f.ID] = f
	}
	return functionMap
}

func getTransID(c string, pipe *models.DataPipeline) string {
	transID := strings.Split(c, ":")[0]
	idxStr := strings.Split(transID, "-")[1]
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		errutils.Exitf("Failed to parse transformation index %s", err.Error())
	}
	if idx < 0 || idx >= len(pipe.TransformationArgsList) {
		errutils.Exitf("Parsed transformation index out of range %d", idx)
	}
	return *pipe.TransformationArgsList[idx].TransformationID
}

var logPipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Real-time logging for data pipeline",
	Example: io_utils.Examples(`
	# Stream real-time logs of a datapipeline from an endpoint
	xi-iot log pipeline test-pipe -e edgeName --function func1`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		edgeName, _ := cmd.Flags().GetString("edge")
		functionName, _ := cmd.Flags().GetString("function")
		pipeName := args[0]
		edgeID := yaml.ToEdgeID(edgeName, CMEClient)
		pipes, err := CMEClient.ListDataPipelinesByNames([]string{pipeName})
		if err != nil {
			errutils.Exitf("Failed to find pipeline %s. %s", pipeName, *err.Message)
		}
		if len(pipes) == 0 {
			errutils.Exitf("Pipeline %s does not exist", pipeName)
		}
		pipe := pipes[0]
		// get container
		containers, err := CMEClient.ListPipelineContainers(pipe.ID, *edgeID)
		if err != nil {
			errutils.Exitf("Failed to get application containers %s", *err.Message)
		}
		if len(containers) == 0 {
			errutils.Exitf("No application container is found")
		}
		// get maps from function id to function
		functionMap := getFunctionMap(pipe)

		if len(containers) > 1 && functionName == "" {
			errutils.Exitf("Function name (--function) required if the pipeline has multiple functions.")
		}
		// Set containerName to the first value
		if functionName == "" {
			functionName = strings.Split(containers[0], ":")[0]
		}
		selectedContainer := ""
		functionNames := make([]string, 0, len(containers))
		for _, c := range containers {
			transID := getTransID(c, pipe)
			displayName := functionMap[transID].Name
			functionNames = append(functionNames, *displayName)
			if *displayName == functionName {
				selectedContainer = c
				break
			}
		}
		if selectedContainer == "" {
			errutils.Exitf("Could not find given function %s. Possible options for containers are: %v ", functionName, functionNames)
		}

		// get stream end point
		logStream := models.LogStream{
			ApplicationID:  "",
			DataPipelineID: pipe.ID,
			EdgeID:         edgeID,
			Container:      &selectedContainer,
		}
		url, err := CMEClient.GetLogEndpoint(&logStream)
		streamLogFromEndpoint(url)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.AddCommand(logAppCmd)
	logCmd.AddCommand(logPipelineCmd)

	logAppCmd.Flags().StringP("edge", "e", "", "Name of the edge to stream log from.")
	logAppCmd.Flags().StringP("container", "c", "", "Name of the container to stream log from. Required if the application has multiple containers.")
	logAppCmd.MarkFlagRequired("edge")

	logPipelineCmd.Flags().StringP("edge", "e", "", "Name of the edge to stream log from.")
	logPipelineCmd.Flags().StringP("function", "", "", "Name of the function to stream log from. Required if the pipeline has multiple functions.")
	logPipelineCmd.MarkFlagRequired("edge")

}
