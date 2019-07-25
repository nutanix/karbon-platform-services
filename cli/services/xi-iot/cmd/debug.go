package cmd

import (
	"context"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"xi-iot-cli/xi-iot/edge"
	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"

	edge_models "xi-iot-cli/generated/edge_swagger/models"
	"xi-iot-cli/generated/swagger/models"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	createNameMaxRetry = 3
	// Check datapipeline every 5 sec
	waitForPipelineTickerInterval = 5
	// Stop checking datapipeline after 50s
	waitForPipelineReadyDeadline = 50
	// Send the debugging information after received 4 times ERROR in pipeState.Status.Status
	maxPipelineErrStateRetry = 4
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug functions",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// generateShortID generates short ID of length `n` from the given letters
// This function is copied from cloudmgmt/services/common/base/utils.go
func generateShortID(n int, letters string) string {
	output := make([]byte, n)

	// Take n bytes, one byte for each character of output.
	randomness := make([]byte, n)
	// read all random
	// Intenationally, not handling errors from rand.Read() as it always returns nil as per documentation.
	crypto_rand.Read(randomness)

	l := len(letters)

	// fill output
	for i := range output {
		// get random item
		random := uint8(randomness[i])
		// Get the position of next char in
		randomPos := random % uint8(l)
		// put into output
		output[i] = letters[randomPos]
	}
	return string(output)
}

func generateRandomName(tenantID string) string {
	isValidPipelineName := regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString
	rand.Seed(time.Now().UnixNano())
	prefix := "short-lived-"
	petName := petname.Generate(2, "-")
	var genName string
	// Sometimes tenantID contains '_' which is not valid for datapipeline name.
	if isValidPipelineName(tenantID) {
		genName = prefix + tenantID + "-" + petName
	} else {
		genName = prefix + petName
	}
	petID := generateShortID(6, strings.Replace(genName, "-", "", -1))
	return genName + "-" + petID
}

var debugPurgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Deletes short lived datapipelines and categories",
	Example: io_utils.Examples(`
	# Deletes all datapipelines and categories whose name start with 'short-lived-'.
	xi-iot debug purge

	# Deletes all datapipelines and categories whose name start with 'short-lived-tenant-id-alex'.
	xi-iot debug purge -p "tenant-id-alex"
	`),
	Run: func(cmd *cobra.Command, args []string) {
		prefix, _ := cmd.Flags().GetString("prefix")
		var regex string
		if prefix == "" {
			regex = "name LIKE 'short-lived-%%'"
		} else {
			regex = fmt.Sprintf("name LIKE 'short-lived-%s%%'", prefix)
		}
		datapipelines, err := CMEClient.ListDataPipelines(regex)
		if err != nil {
			errutils.Exitf("failed to list short-lived pipelines. %s", *err.Message)
		}
		categories, err := CMEClient.ListCategories(regex)
		if err != nil {
			errutils.Exitf("failed to list short-lived categories. %s", *err.Message)
		}
		fmt.Printf("datapipeline count: %d, category count : %d.\n", len(datapipelines), len(categories))
		pipeArg := make([]string, len(datapipelines))
		cateArg := make([]string, len(categories))
		for _, d := range datapipelines {
			pipeArg = append(pipeArg, *d.Name)
		}
		for _, c := range categories {
			cateArg = append(cateArg, *c.Name)
		}
		if len(datapipelines) > 0 {
			deleteDataPipelineCmd.Run(cmd, pipeArg)
		}
		if len(categories) > 0 {
			deleteCategoryCmd.Run(cmd, cateArg)
		}
	},
}

// waitForPipelineOnEdge will check for datapipeline with id datapipelineID on edge and wait for it.
func waitForPipelineOnEdge(edgeClient *edge.Client, datapipelineID string, interval time.Duration, endTime time.Duration) error {
	var ticker *time.Ticker
	ctxStop, cancelStop := context.WithDeadline(context.Background(), time.Now().Add(endTime))
	defer cancelStop()
	// Capture user cancel signal
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt)
	pipeState, err := edgeClient.GetDataPipeline(datapipelineID)
	if err == nil && pipeState.Status.Status == "RUNNING" {
		return nil
	}
	fmt.Printf("Datapipeline not ready, retry in %s...\n", interval.String())
	ticker = time.NewTicker(interval)
	// When there is syntax error in the function, container will crash, and pipeState.Status.Status will never be running.
	// Thus count for how many times we received ERROR state and stop waiting after received maxPipelineErrorState errors
	containerErrorCount := 0
	// Tick and check
	for {
		select {
		case <-ctrlC:
			return fmt.Errorf("User cancelled")
		case <-ticker.C:
			pipeState, err := edgeClient.GetDataPipeline(datapipelineID)
			if err == nil && pipeState.Status.Status == "RUNNING" {
				return nil
			}
			if err == nil && pipeState.Status.Status == "ERROR" {
				containerErrorCount++
				if containerErrorCount > maxPipelineErrStateRetry {
					return nil
				}
			}
			fmt.Printf("Datapipeline not ready, retry in %s...\n", interval.String())
		case <-ctxStop.Done():
			return fmt.Errorf("Data pipeline %s is not in ready state after %s seconds", datapipelineID, endTime.String())
		}
	}
}

// readFunctionWithArgs read argument input into TransformationArgs
// Input Example: ['func1=arg1:val1', 'func2=arg1:val1,arg2:val2']
func readFunctionWithArgs(functions []string) ([]*models.TransformationArgs, string, error) {
	var functionNameList []string
	for _, t := range functions {
		trans := strings.Split(t, "=")
		functionNameList = append(functionNameList, trans[0])
	}
	functionList, err := CMEClient.ListFunctionByNames(functionNameList)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to get functions %s", *err.Message)
	}
	// function look up map
	functionMap := make(map[string]*models.Function)
	for _, f := range functionList {
		functionMap[*f.Name] = f
	}
	transforms := []*models.TransformationArgs{}
	for _, f := range functions {
		trans := strings.Split(f, "=")
		if function, ok := functionMap[trans[0]]; ok {
			// function exist, fill in the input parameters
			params := []*models.ScriptParamValue{}
			if len(trans) > 1 {
				// arguments look up map
				argsMap := make(map[string]*models.ScriptParam)
				for _, p := range function.Params {
					argsMap[*p.Name] = p
				}
				argVals := strings.Split(trans[1], ",")
				for _, argVal := range argVals {
					argValPair := strings.Split(argVal, ":")
					if len(argValPair) != 2 {
						return nil, "", fmt.Errorf("Function argument list format error. Example: func2=arg1:val1,arg2:val2")
					}
					if arg, ok := argsMap[argValPair[0]]; ok {
						argName := argValPair[0]
						argVal := argValPair[1]
						if *arg.Type == "integer" {
							if _, err := strconv.Atoi(argVal); err != nil {
								return nil, "", fmt.Errorf("Failed to convert integer argument %s in function %s", argName, trans[0])
							}
						}
						params = append(params, &models.ScriptParamValue{Name: &argName, Value: &argVal, Type: arg.Type})
					} else {
						return nil, "", fmt.Errorf("Didn't find argument %s in function %s", argValPair[0], trans[0])
					}
				}
			}
			if len(function.Params) != len(params) {
				return nil, "", fmt.Errorf("Missing or too many arguments for function: %s", trans[0])
			}
			transforms = append(transforms, &models.TransformationArgs{TransformationID: &function.ID, Args: params})
		} else {
			return nil, "", fmt.Errorf("Didn't find function %s", trans[0])
		}
	}
	// read project ID from functions
	projectIDMap := make(map[string]bool)
	for _, f := range functionList {
		projectIDMap[f.ProjectID] = true
	}
	if len(projectIDMap) > 1 {
		return nil, "", fmt.Errorf("Functions are not from the same project")
	}
	if len(projectIDMap) == 0 {
		return nil, "", fmt.Errorf("At least one function should be user created. Debugging with only builtin function is not supported")
	}
	var projectID string
	for p := range projectIDMap {
		projectID = p
	}
	return transforms, projectID, nil
}

func readMessageInput(inputFile string) ([]string, []string, error) {
	if inputFile == "-" {
		fmt.Printf("Input Text: ")
		line, err := io_utils.Readstdin()
		if err != nil {
			return nil, nil, err
		}
		return []string{base64.StdEncoding.EncodeToString([]byte(line))}, []string{"stdin"}, nil
	}

	var resp []string
	fileNameList, err := filepath.Glob(inputFile)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Reading input from %d files.\n", len(fileNameList))
	for _, fileName := range fileNameList {
		fileb, err := io_utils.Readfile(fileName)
		if err != nil {
			return nil, nil, err
		}
		resp = append(resp, base64.StdEncoding.EncodeToString(fileb))
	}
	return resp, fileNameList, nil
}

func parseOutputOptions(outputOptions string) (string, string, error) {
	if outputOptions == "-" {
		return "", "", nil
	}
	option := strings.Split(outputOptions, "*")
	if len(option) != 2 {
		return "", "", fmt.Errorf("Output format error. Examples: './output/*.png', '*.jpeg', './*.txt'")
	}
	if option[0] == "" {
		option[0] = "./"
	}
	if _, err := os.Stat(option[0]); os.IsNotExist(err) {
		err := os.MkdirAll(option[0], os.ModePerm)
		if err != nil {
			return "", "", fmt.Errorf("Failed to create output path %s, %s", option[0], err.Error())
		}
		fmt.Printf("Created folder %s\n", option[0])
	}
	return option[0], option[1], nil
}

// cleanDataPipelineAndCategory will be defered executed to clean up temp datapipeline and category
func cleanDataPipelineAndCategory(datapipelineID string, categoryID string) {
	err := CMEClient.DeleteDataPipeline(datapipelineID)
	if err != nil {
		errutils.Exitf("Failed to delete temporary data pipeline %s", *err.Message)
	}
	err = CMEClient.DeleteCategory(categoryID)
	if err != nil {
		errutils.Exitf("Failed to delete temporary category %s", *err.Message)
	}
	fmt.Println("Temporary data pipeline and category deleted...")
}

// printResponses is used to either print responses or store them to files
func printResponses(resp *edge_models.DebugReceiveResponse, startIndex int, outputOptions string) (bool, error) {
	// If there is ALERT, print the alert and return.
	for _, e := range resp.Events {
		if *e.Type == "ALERT" {
			fmt.Printf("Function error:\n\n")
			fmt.Println(*e.Message)
			fmt.Println(e.Context)
			return true, nil
		}
	}
	// Else print message
	outputPath, outputExtension, err := parseOutputOptions(outputOptions)
	if err != nil {
		return false, err
	}
	if len(resp.Messages) == 0 {
		fmt.Println("No message was received...")
	} else {
		if outputOptions == "-" {
			fmt.Printf("Received Messages:\n\n")
			for _, m := range resp.Messages {
				decodedM, err := base64.StdEncoding.DecodeString(m)
				if err != nil {
					return false, fmt.Errorf("Failed to decode response message. %s", err)
				}
				fmt.Println(string(decodedM))
			}
			fmt.Print("\n")
		} else {
			for idx, m := range resp.Messages {
				decodedM, err := base64.StdEncoding.DecodeString(m)
				if err != nil {
					return false, fmt.Errorf("Failed to decode response message. %s", err)
				}
				filePath := outputPath + "resp" + strconv.Itoa(idx+startIndex) + outputExtension
				err = ioutil.WriteFile(filePath, decodedM, 0644)
				if err != nil {
					return false, fmt.Errorf("Failed to write to file %s. %s", filePath, err)
				}
			}
			if len(resp.Messages) == 1 {
				fmt.Printf("Stored Received Messages to: %sresp%d%s\n", outputPath, startIndex, outputExtension)
			} else {
				fmt.Printf("Stored Received Messages to: %sresp[%d-%d]%s\n", outputPath, startIndex, startIndex+len(resp.Messages)-1, outputExtension)
			}
		}
	}
	return false, nil
}

func checkEdgeReachable(edgeURL string) (*edge.Client, error) {
	edgeClient := edge.New(edgeURL, edge.WithCMEClient(&CMEClient))
	fmt.Printf("Ping %s ... ", edgeURL)
	respCh := make(chan bool, 1)
	errCh := make(chan error, 64)
	ctrlC := make(chan os.Signal, 1)
	go func() {
		_, err := edgeClient.GetInfo()
		if err != nil {
			errCh <- fmt.Errorf("Edge not reachable, %s", err.Error())
		} else {
			respCh <- true
		}
	}()
	select {
	case <-ctrlC:
		return nil, fmt.Errorf("User cancelled")
	case err := <-errCh:
		return nil, err
	case <-respCh:
		fmt.Println("Success")
		return edgeClient, nil
	}
}

// EdgeService stores service name and port number, should be consistent with controller
type EdgeService struct {
	Name string   `json:"Name"`
	Port []string `json:"Port"`
}

// EdgeArtifacts stores edge artifacts, should be consistent with controller
// Controller Edge Artifact Structure:
// {
//    "EdgeExternalAddress": "123.321.123.321"
//    "EdgeServices": [
//       {"Name": "DebugService", "Port": ["30245"]},
//       {"Name": "JupyterService", "Port": ["30203", "30004"]},
//    ]
// }
type EdgeArtifacts struct {
	EdgeExternalAddress string        `json:"EdgeExternalAddress"`
	EdgeServices        []EdgeService `json:"EdgeServices"`
}

// getDebugServiceAddress get edge url from edge artifact and service name
func edgeAritfactsToURL(edgeArtifacts map[string]interface{}, serviceName string) (string, error) {
	artifact := EdgeArtifacts{}
	b, err := json.Marshal(edgeArtifacts)
	if err != nil {
		return "", fmt.Errorf("Error Marshal edgeArtifacts")
	}
	err = json.Unmarshal(b, &artifact)
	if err != nil {
		return "", fmt.Errorf("Error Unmarshal edgeArtifacts: %s", string(b))
	}
	for _, s := range artifact.EdgeServices {
		if s.Name == serviceName {
			if len(s.Port) == 0 {
				return "", fmt.Errorf("No port info for service %s in edge artifacts", serviceName)
			}
			return artifact.EdgeExternalAddress + ":" + s.Port[0], nil
		}
	}
	return "", fmt.Errorf("Didn't find service %s", serviceName)
}

// getDebugServiceAddress returns URL for debug service running on edge
func getDebugServiceAddress(projectID string, edgeName string) (*edge.Client, error) {
	project, err := CMEClient.GetProject(projectID)
	if err != nil {
		return nil, fmt.Errorf(*err.Message)
	}
	// Should use ListEdgeInfosProjects blocked by [ENG-242048]
	// edgeInfoList, err := CMEClient.ListEdgeInfosProjects(projectID)
	edgeInfoList, err := CMEClient.ListEdgeInfos("")
	if err != nil {
		return nil, fmt.Errorf(*err.Message)
	}
	// Map of edge ids
	edgeIDMap := make(map[string]bool)
	// No need to check project.EdgeSelectorType, project.EdgeIds always contains all the edges in the project
	for _, e := range project.EdgeIds {
		edgeIDMap[e] = true
	}
	// Map edgeId to edge URL
	edgeURLMap := make(map[string]string)
	var edgeIDList []string
	for _, e := range edgeInfoList {
		// If the edge is in the project and has artifact
		if _, ok := edgeIDMap[*e.EdgeID]; ok && e.EdgeArtifacts != nil {
			edgeURL, err := edgeAritfactsToURL(e.EdgeArtifacts, "DebugService")
			if err != nil {
				return nil, err
			}
			edgeURLMap[*e.EdgeID] = edgeURL
			edgeIDList = append(edgeIDList, *e.EdgeID)
		}
	}
	if len(edgeIDList) == 0 {
		return nil, fmt.Errorf("Project %s has no edge available for debugging", *project.Name)
	}
	// get edge names
	validEdgeIDStr := io_utils.ToSQLList(edgeIDList)
	filter := fmt.Sprintf("id IN (%s)", validEdgeIDStr)
	edgeList, err := CMEClient.ListEdges(filter)
	if err != nil {
		return nil, fmt.Errorf(*err.Message)
	}
	if len(edgeIDList) != len(edgeList) {
		return nil, fmt.Errorf("Looking for %d edges based on edge usage info, but only got %d", len(edgeIDList), len(edgeList))
	}
	// Map edge name to edge id
	edgeNameMap := make(map[string]string)
	var edgeNameList []string
	for _, e := range edgeList {
		edgeNameMap[*e.Name] = e.ID
		edgeNameList = append(edgeNameList, *e.Name)
	}
	// If user specified edge name
	if edgeName != "" {
		selectedEdgeName := edgeName
		selectedEdgeID, ok := edgeNameMap[edgeName]
		if !ok {
			return nil, fmt.Errorf("Project %s does not have %s available for debugging, please specify one edge name from %v through -e", *project.Name, edgeName, edgeNameList)
		}
		edgeClient, err := checkEdgeReachable(edgeURLMap[selectedEdgeID])
		if err != nil {
			return nil, err
		}
		fmt.Printf("Edge %s is selected among %v for debugging...\n", selectedEdgeName, edgeNameList)
		return edgeClient, nil
	}
	// Choose a reachable edge
	for _, selectedEdgeName := range edgeNameList {
		selectedEdgeID := edgeNameMap[selectedEdgeName]
		// check reachable:
		edgeClient, err := checkEdgeReachable(edgeURLMap[selectedEdgeID])
		if err == nil {
			fmt.Printf("Edge %s is selected among %v for debugging...\n", selectedEdgeName, edgeNameList)
			return edgeClient, nil
		}
	}
	return nil, fmt.Errorf("no edge among %v is reachable for debugging", edgeNameList)
}

func debugFunc(cmd *cobra.Command, args []string) error {
	datapipelineName, _ := cmd.Flags().GetString("datapipeline")
	edgeName, _ := cmd.Flags().GetString("edge")
	edgeURL, _ := cmd.Flags().GetString("edge-url")
	inputFile, _ := cmd.Flags().GetString("input")
	repeatInput, _ := cmd.Flags().GetInt("repeat")
	repeatInterval, _ := cmd.Flags().GetString("repeat-interval")
	outputOptions, _ := cmd.Flags().GetString("output")
	duration, _ := cmd.Flags().GetString("time")
	functions, _ := cmd.Flags().GetStringSlice("functions")
	keep, _ := cmd.Flags().GetBool("keep")
	var edgeClient *edge.Client

	durationDecode, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("Duration format error. Example: 2h3m1s")
	}
	repeatIntervalDecode, err := time.ParseDuration(repeatInterval)
	if err != nil {
		return fmt.Errorf("Repeat interval format error. Example: 3m1s")
	}
	// Read input into base64 encoded string
	messageList, fileNameList, err := readMessageInput(inputFile)
	if err != nil {
		return err
	}
	// Validate output option format
	_, _, err = parseOutputOptions(outputOptions)
	if err != nil {
		return err
	}
	if repeatInput < 1 {
		return fmt.Errorf("Repeat input times cannot be smaller than 1")
	}

	var datapipelineID string
	// TODO: use function yaml when we implement creation of function.
	if len(functions) != 0 {
		// Create transfomations object and infer project id from function names
		transforms, projectID, err := readFunctionWithArgs(functions)
		if err != nil {
			return err
		}
		// Get edge client from project id
		if edgeURL == "" {
			edgeClient, err = getDebugServiceAddress(projectID, edgeName)
		} else {
			edgeClient, err = checkEdgeReachable(edgeURL)
		}
		if err != nil {
			return err
		}

		retry := 0
		for ; retry < createNameMaxRetry; retry++ {
			// Generate random name for the temp pipeline
			name := generateRandomName(Cfg.TenantID)
			// Create temporary input selector in the same name
			category := models.Category{Name: &name, Values: []string{name}}
			categoryID, err := CMEClient.CreateCategory(&category)
			if err != nil {
				fmt.Printf("Failed to create category %s, with name %s. Retry with a different name", *err.Message, name)
				continue
			}
			categoryInfo := models.CategoryInfo{ID: &categoryID, Value: &name}
			categoryInfoList := []*models.CategoryInfo{&categoryInfo}
			// Create temporary data pipeline object.
			destination := models.DataPipelineDestinationEdge
			origin := models.DataPipelineOriginDataSource
			tempPipeline := models.DataPipeline{
				Destination:            &destination,
				EdgeStreamType:         "None",
				EndPoint:               name,
				Name:                   &name,
				Origin:                 &origin,
				ProjectID:              projectID,
				TenantID:               &Cfg.TenantID,
				TransformationArgsList: transforms,
				OriginSelectors:        categoryInfoList,
			}

			datapipelineID, err = CMEClient.CreateDataPipeline(&tempPipeline)
			if err != nil {
				fmt.Printf("Failed to create data pipeline %s, with name %s. Retry with a different name\n", *err.Message, name)
				err = CMEClient.DeleteCategory(categoryID)
				if err != nil {
					return fmt.Errorf("Failed to delete temporary category %s", *err.Message)
				}
				continue
			} else {
				fmt.Printf("Created pipeline with name: %s\n", name)
				if !keep {
					defer cleanDataPipelineAndCategory(datapipelineID, categoryID)
				}
				break
			}
		}
		if retry == createNameMaxRetry {
			return fmt.Errorf("Failed to create random name for too many times")
		}
	} else if datapipelineName != "" {
		dataPipelineList, apiErr := CMEClient.ListDataPipelinesByNames([]string{datapipelineName})
		if apiErr != nil {
			return fmt.Errorf("Failed to get data pipeline. %s", *apiErr.Message)
		}
		if len(dataPipelineList) != 1 {
			return fmt.Errorf("Data pipeline not found or multiple data pipeline has the same name")
		}
		// Get edge client from project id
		if edgeURL == "" {
			edgeClient, err = getDebugServiceAddress(dataPipelineList[0].ProjectID, edgeName)
		} else {
			edgeClient, err = checkEdgeReachable(edgeURL)
		}
		if err != nil {
			return err
		}

		datapipelineID = dataPipelineList[0].ID
	} else {
		return fmt.Errorf("Either --functions or --datapipeline-id is required")
	}
	// Wait for data pipeline being created
	fmt.Println("Waiting for datapipeline to be created on edge...")
	waitInterval := time.Duration(waitForPipelineTickerInterval * time.Second)
	waitEnd := time.Duration(waitForPipelineReadyDeadline * time.Second)
	waitErr := waitForPipelineOnEdge(edgeClient, datapipelineID, waitInterval, waitEnd)
	if waitErr != nil {
		return waitErr
	}
	// Receive response
	fmt.Printf("Subscribe to response with no output timeout of %s...\n", durationDecode.String())
	respCh := make(chan *edge_models.DebugReceiveResponse, 64)
	errCh := make(chan error, 64)

	messageIndex := 0
	for i := 0; i < repeatInput; i++ {
		for idx, message := range messageList {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				receive, edgeErr := edgeClient.DebugReceive(datapipelineID, duration)
				if edgeErr != nil {
					errCh <- edgeErr
				} else {
					respCh <- receive
				}
			}()
			go func() {
				defer wg.Done()
				// Delay sending to make sure subscription to NATs is success.
				time.Sleep(1 * time.Second)
				edgeErr := edgeClient.DebugSend(datapipelineID, message)
				if edgeErr != nil {
					fmt.Printf("Error: Failed to send message.%s\n", edgeErr.Error())
				} else {
					fmt.Printf("Debug data from %s sent...\n", fileNameList[idx])
				}
			}()
			ctrlC := make(chan os.Signal, 1)
			signal.Notify(ctrlC, os.Interrupt)
			select {
			case <-ctrlC:
				return fmt.Errorf("User cancelled")
			case err := <-errCh:
				return err
			case resp := <-respCh:
				alertEvent, err := printResponses(resp, messageIndex, outputOptions)
				if err != nil {
					return err
				}
				if alertEvent {
					return nil
				}
				messageIndex += len(resp.Messages)
			}
			wg.Wait()
			time.Sleep(repeatIntervalDecode)
		}
	}

	return nil
}

var debugFunctionCmd = &cobra.Command{
	Use:   "function",
	Short: "Triggers data pipelines by fake inputs and collects results",
	Example: io_utils.Examples(`
	# 6-10 cloud API calls + 3-11 edge API calls

	# Creates a short-lived datapipeline with transforms test and testcopy on edge edge-name-6fo4n.
	# Sends contents within the input file './test.txt'.
	# Sets no output timeout of receiving pipeline output to 5 min 30 sec. Prints output to stdout.
	xi-iot debug function -e edge-name-6fo4n -f test,testcopy -i './test.txt' -o '-' -t 5m30s

	# Creates a short-lived datapipeline with transforms test and testcopy.
	# Reads input text from stdin and repeats sending the same text every 3s for 5 times.
	# Sets no output timeout of receiving pipeline output to 1 min 30 sec. Prints output to stdout.
	xi-iot debug function -f test,testcopy -t 1m30s -r 5 --repeat-interval 3s
	
	# Creates a short-lived datapipeline with transforms test and testcopy on edge edge-name-6fo4n.
	# Reads contents from files whose name matches './input/test*.png'. 
	# Contents are send to short-lived datapipeline in alphabetic order of their file names.
	# Responses will be stored in './output/resp0.png', './output/resp1.png', ...
	xi-iot debug function -e edge-name-6fo4n -f test,testcopy -i './input/test*.png' -o './output/*.png'

	# Creates a short-lived datapipeline with transforms func1 with arg1:val1 and func2 with arg1:val1,arg2:val2 
	#		on edge edge-name-6fo4n.
	# Reads input text from stdin. Sends the text to the short-lived data pipeline. Prints output to stdout.
	xi-iot debug function -e edge-name-6fo4n -f func1=arg1:val1 -f func2=arg1:val1,arg2:val2 -i '-' -o '-'
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// This function is added for error handling because cobra does not allow error returning
		// If glog.Exit is called, all defered functions will not be execute.
		err := debugFunc(cmd, args)
		if err != nil {
			glog.Exit(err.Error())
		}
	},
}

var debugPipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Triggers data pipelines by fake inputs and collects results",
	Example: io_utils.Examples(`
	# 1 cloud API call + 2-10 edge API calls

	# Triggers data pipeline with name testPipe on edge at edge-name-6fo4n with contents within the input file './test.txt'.
	# Sets no output timeout of receiving pipeline output to 5 min 30 sec. Prints output to stdout.
	xi-iot debug pipeline -e edge-name-6fo4n -d testPipe -i './test.txt' -o '-' -t 5m30s

	# Triggers data pipeline with name testPipe with text read from stdin. 
	# Repeats sending the same text every 3s for 5 times.
	# Sets no output timeout of receiving pipeline output to 1 min 30 sec. Prints output to stdout.
	xi-iot debug pipeline -d testPipe -t 1m30s -r 5 --repeat-interval 3s
	
	# Triggers data pipeline with name testPipe on edge at edge-name-6fo4n with contents read from files 
	#		whose name matches './input/test*.png'. Contents are send in alphabetic order of their file names.
	# Responses will be stored in './output/resp0.png', './output/resp1.png', ...
	xi-iot debug pipeline -e edge-name-6fo4n -d testPipe -i './input/test*.png' -o './output/*.png'

	# Triggers data pipeline with name testPipe on edge at edge-name-6fo4n with text read from stdin.
	# Receives text from stdin, then sends the text to the tmp data pipeline. Then print output to stdout.
	xi-iot debug pipeline -e edge-name-6fo4n -d testPipe -i '-' -o '-'
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// This function is added for error handling because cobra does not allow error returning
		// If glog.Exit is called, all defered functions will not be execute.
		err := debugFunc(cmd, args)
		if err != nil {
			glog.Exit(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
	debugCmd.AddCommand(debugPurgeCmd)
	debugCmd.AddCommand(debugFunctionCmd)
	debugCmd.AddCommand(debugPipelineCmd)

	debugPurgeCmd.Flags().StringP("prefix", "p", "", "Prefix of short lived datapipelines and categories to delete. No need to include 'short-lived-' in prefix.")

	debugFunctionCmd.Flags().StringP("edge", "e", "", "Debug edge name. Randomly pick an edge for debugging if not specified.")
	debugFunctionCmd.Flags().StringP("edge-url", "", "", "Debug edge url. Overwrite the edge url from cloud management. For advanced users.")
	debugFunctionCmd.Flags().StringP("input", "i", "-", "Input file to be send to datapipeline, default '-' means read from stdin.")
	debugFunctionCmd.Flags().StringP("output", "o", "-", "Output folder and file extension to receive output from datapipeline, default '-' means output to stdout. Example: ./output/*.png")
	debugFunctionCmd.Flags().IntP("repeat", "r", 1, "Repeat sending messages for 'repeat' times.")
	debugFunctionCmd.Flags().StringP("repeat-interval", "", "0s", "Repeat sending messages every 'repeat-interval' time")
	debugFunctionCmd.Flags().StringP("time", "t", "30s", "Timeout of waiting for output, in case there is no output for a certain input. Example: 2h4m23s means 2 hours 4 minutes 23 seconds. Valid range [5s, 1h].")
	debugFunctionCmd.Flags().StringSliceP("functions", "f", []string{}, "Comma separated list of function names. Example: -f func1=arg1:val1,arg2:val2 -f func2=arg3,arg4 .")
	debugFunctionCmd.Flags().BoolP("keep", "k", false, "Keep the generated short-lived datapipeline instead of deleting it.")

	debugFunctionCmd.Flags().MarkHidden("edge-url")

	debugPipelineCmd.Flags().StringP("edge", "e", "", "Debug edge name. Randomly pick an edge for debugging if not specified.")
	debugPipelineCmd.Flags().StringP("edge-url", "", "", "Debug edge url. Overwrite the edge url from cloud management. For advanced users.")
	debugPipelineCmd.Flags().StringP("datapipeline", "d", "", "Datapipeline name.")
	debugPipelineCmd.Flags().StringP("input", "i", "-", "Input file to be send to datapipeline, default '-' means read from stdin.")
	debugPipelineCmd.Flags().StringP("output", "o", "-", "Output folder and file extension to receive output from datapipeline, default '-' means output to stdout. Example: ./output/*.png")
	debugPipelineCmd.Flags().IntP("repeat", "r", 1, "Repeat sending messages for 'repeat' times.")
	debugPipelineCmd.Flags().StringP("repeat-interval", "", "0s", "Repeat sending messages every 'repeat-interval' time")
	debugPipelineCmd.Flags().StringP("time", "t", "30s", "Timeout of waiting for output, in case there is no output for a certain input. Example: 2h4m23s means 2 hours 4 minutes 23 seconds. Valid range [5s, 1h].")

	debugPipelineCmd.Flags().MarkHidden("edge-url")
}
