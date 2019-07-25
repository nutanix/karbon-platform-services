package yaml

import (
	"io/ioutil"
	"strconv"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/cloudmgmt"
	"xi-iot-cli/xi-iot/errutils"

	"github.com/go-yaml/yaml"

	"github.com/golang/glog"

	"strings"
)

var (
	Unused             = "unused"
	sensor             = "Sensor"
	NoneEdgeStreamType = "None"
	RealTimeStreamType = "RealTimeStreaming"
)

func ToFields(fields []Field) []*models.DataSourceFieldInfoV2 {
	rtn := make([]*models.DataSourceFieldInfoV2, 0, len(fields))
	for _, f := range fields {
		temp := f
		rtn = append(rtn, &models.DataSourceFieldInfoV2{Name: &temp.Name, Topic: &temp.Topic})
	}
	return rtn
}

func ToSelectors(selectors []Selector, client cloudmgmt.Client) []*models.DataSourceFieldSelector {
	rtn := make([]*models.DataSourceFieldSelector, 0, len(selectors))

	selNames := make([]string, 0, len(selectors))
	selByName := make(map[string]int)
	for i, sel := range selectors {
		selNames = append(selNames, sel.CategoryName)
		selByName[sel.CategoryName] = i
	}

	cats, apiErr := client.ListCategoriesByNames(selNames)
	if apiErr != nil {
		errutils.Exitf("failed to find categories for selectors. %s", *apiErr.Message)
	}

	for _, cat := range cats {
		temp := cat
		if sel, ok := selByName[*cat.Name]; ok {
			rtn = append(rtn, &models.DataSourceFieldSelector{ID: &temp.ID, Scope: selectors[sel].Scope, Value: &selectors[sel].CategoryValue})
			delete(selByName, *cat.Name)
		}
	}

	if len(selByName) != 0 {
		for name := range selByName {
			glog.Errorf("Could not find category: %s", name)
		}
		glog.Exit()
	}

	return rtn
}

func ToPorts(ports []Port) []*models.DataSourceIfcPorts {
	rtn := make([]*models.DataSourceIfcPorts, 0, len(ports))
	for _, p := range ports {
		portN := int64(p.Port)
		rtn = append(rtn, &models.DataSourceIfcPorts{Name: &p.Name, Port: &portN})
	}
	return rtn
}

func ToIfcInfo(ifcInfo *IfcInfo) *models.DataSourceIfcInfo {
	if ifcInfo == nil {
		return nil
	}
	rtn := models.DataSourceIfcInfo{}
	rtn.Protocol = &ifcInfo.Protocol
	rtn.ProjectID = Unused
	rtn.Ports = ToPorts(ifcInfo.Ports)
	rtn.Kind = &ifcInfo.Kind
	rtn.DriverID = &Unused
	rtn.Class = &ifcInfo.Class
	rtn.Img = &ifcInfo.Img
	return &rtn
}

func ToEdgeID(edgeName string, client cloudmgmt.Client) *string {
	edges, err := client.ListEdgesByNames([]string{edgeName})
	if err != nil {
		errutils.Exitf("failed to find edge %s. %s", edgeName, *err.Message)
	}
	if len(edges) == 0 {
		errutils.Exitf("edge %q does not exist", edgeName)
	}
	return &edges[0].ID
}

// ToEdgeIDs takes array of edge names as input and returns their corresponding IDs
func ToEdgeIDs(edgeNames []string, client cloudmgmt.Client) (resp []string) {
	edges, err := client.ListEdgesByNames(edgeNames)
	if err != nil {
		errutils.Exitf("failed to find edges %v. %s", edgeNames, *err.Message)
	}
	for _, e := range edges {
		resp = append(resp, e.ID)
	}
	return
}

func ToDataSrc(yamlDs DataSource, tenantID string, client cloudmgmt.Client) models.DataSourceV2 {
	ds := models.DataSourceV2{}
	ds.Name = &yamlDs.Name
	ds.Protocol = &yamlDs.Protocol
	ds.TenantID = &tenantID
	ds.Type = &yamlDs.Type
	ds.AuthType = &yamlDs.AuthType
	ds.EdgeID = ToEdgeID(yamlDs.Edge, client)
	ds.FieldsV2 = ToFields(yamlDs.Fields)
	ds.Selectors = ToSelectors(yamlDs.Selectors, client)
	ds.IfcInfo = ToIfcInfo(yamlDs.IfcInfo)
	ds.Type = &sensor
	return ds
}

func ToDataSourceID(dataSrc string, client cloudmgmt.Client) string {
	dataSrcs, err := client.ListDataSourcesByNames([]string{dataSrc})
	if err != nil {
		errutils.Exitf("failed to find data source %s. %s", dataSrc, *err.Message)
	}
	if len(dataSrcs) == 0 {
		errutils.Exitf("data source %s does not exist", dataSrc)
	}

	return dataSrcs[0].ID
}

func ToDataIfcEndpoints(endpoints []DataIfcEndpoint, client cloudmgmt.Client) []*models.DataIfcEndpoint {
	rtn := make([]*models.DataIfcEndpoint, 0, len(endpoints))
	if len(endpoints) == 0 {
		return nil
	}
	names := make([]string, 0, len(endpoints))
	// Add data source names
	for _, e := range endpoints {
		names = append(names, e.DataSource)
	}

	dataSources, apiErr := client.ListDataSourcesByNames(names)
	if apiErr != nil {
		errutils.Exitf("failed to find data sources: %+v. %s", names, *apiErr.Message)
	}

	dataSrcByName := make(map[string]*models.DataSourceV2)
	for _, ds := range dataSources {
		dataSrcByName[*ds.Name] = ds
	}

	for i, e := range endpoints {
		if _, ok := dataSrcByName[e.DataSource]; !ok {
			errutils.Exitf("failed to find data source %s", e.DataSource)
		}
		rtn = append(rtn, &models.DataIfcEndpoint{Name: &endpoints[i].FieldName,
			ID: &dataSrcByName[e.DataSource].ID, Value: &endpoints[i].TopicName},
		)
	}
	return rtn
}

func ToProjectID(projName string, client cloudmgmt.Client) string {
	projs, apiErr := client.ListProjectByNames([]string{projName})
	if apiErr != nil {
		errutils.Exitf("failed to find project %s. %s", projName, *apiErr.Message)
	}
	if len(projs) == 0 {
		errutils.Exitf("project %s does  not exist", projName)
	}
	return projs[0].ID
}

// ToProject takes project name as input and returns a project object
func ToProject(projName string, client cloudmgmt.Client) *models.Project {
	projs, apiErr := client.ListProjectByNames([]string{projName})
	if apiErr != nil {
		errutils.Exitf("failed to find project %s. %s", projName, *apiErr.Message)
	}
	if len(projs) == 0 {
		errutils.Exitf("project %s does  not exist", projName)
	}
	return projs[0]
}

// convertToYamlBytes converts an input of bytes to a prettified yaml if
// the given input can be converted to yaml. Example: a yaml file
func convertToYamlBytes(src []byte) ([]byte, error) {
	var m interface{}
	err := yaml.Unmarshal(src, &m)
	if err != nil {
		return nil, err

	}
	o, err := yaml.Marshal(m)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// toCatInfoMap takes an array of CatInfo as input and returns a map from category name to value
func toCatInfoMap(cats []CatInfo) (resp map[string][]string) {
	for _, c := range cats {
		resp[c.Name] = []string{c.Value}
	}
	return
}

func ToApplication(app Application, client cloudmgmt.Client) *models.ApplicationV2 {
	rtn := models.ApplicationV2{}
	rtn.Name = &app.Name
	project := ToProject(app.Project, client)
	rtn.ProjectID = project.ID
	rtn.DataIfcEndpoints = ToDataIfcEndpoints(app.DataIfcEndpoints, client)
	rtn.Description = app.Description
	if *project.EdgeSelectorType == "Explicit" && app.Edges != nil {
		rtn.EdgeIds = ToEdgeIDs(app.Edges, client)
	} else if *project.EdgeSelectorType == "Category" && app.EdgeSelectors != nil {
		catInfoMap := toCatInfoMap(app.EdgeSelectors)
		rtn.EdgeSelectors = ToCatInfo(catInfoMap, client)
	}
	b, err := ioutil.ReadFile(app.AppYamlPath)
	if err != nil {
		errutils.Exitf("failed to read app yaml. %s", err.Error())
	}
	o, err := convertToYamlBytes(b)
	if err != nil {
		errutils.Exitf("failed to parse application yaml from %s. %s", app.AppYamlPath, err.Error())
	}
	appManiFest := string(o)
	rtn.AppManifest = &appManiFest
	return &rtn
}

func ToDataPipelineID(name string, client cloudmgmt.Client) string {
	dps, err := client.ListDataPipelinesByNames([]string{name})
	if err != nil {
		errutils.Exitf("failed to find data pipeline %s. %s", name, *err.Message)
	}
	if len(dps) == 0 {
		errutils.Exitf("datapipeline %s does not exist", name)
	}
	return dps[0].ID
}

func ToCatInfo(catVals map[string][]string, client cloudmgmt.Client) []*models.CategoryInfo {
	rtn := make([]*models.CategoryInfo, 0, len(catVals))
	valsByNames := make(map[string][]string)
	selNames := make([]string, 0, len(catVals))
	for cat, vals := range catVals {
		selNames = append(selNames, cat)
		valsByNames[cat] = append(valsByNames[cat], vals...)
	}

	cats, apiErr := client.ListCategoriesByNames(selNames)
	if apiErr != nil {
		errutils.Exitf("failed to find categories. %s", *apiErr.Message)
	}

	catsByName := make(map[string]*models.Category)
	for _, c := range cats {
		catsByName[*c.Name] = c

		// Assert that category has all the values required
		availableVals := make(map[string]bool)
		for _, v := range catsByName[*c.Name].Values {
			availableVals[v] = true
		}
		for _, v := range valsByNames[*c.Name] {
			if _, ok := availableVals[v]; !ok {
				errutils.Exitf("Cateogory %s does not have value %s. Possible values: %s", *c.Name, v, strings.Join(c.Values, ","))
			}
		}
	}

	for name, vals := range catVals {
		if _, ok := catsByName[name]; !ok {
			errutils.Exitf("category %s does not exist", name)
		}
		for i := range vals {
			rtn = append(rtn, &models.CategoryInfo{ID: &catsByName[name].ID, Value: &vals[i]})
		}
	}

	return rtn
}

func ToCloudProfileID(cloudProfile string, client cloudmgmt.Client) string {
	profs, err := client.ListCloudProfilesByNames([]string{cloudProfile})
	if err != nil {
		errutils.Exitf("failed to find cloud profile %s. %s", cloudProfile, *err.Message)
	}
	if len(profs) == 0 {
		errutils.Exitf("cloud profile %s does not exist", cloudProfile)
	}
	return profs[0].ID
}

// TODO: run this in a goroutine
func ToTransform(function Function, client cloudmgmt.Client) *models.TransformationArgs {
	runtimes, err := client.ListFunctionByNames([]string{function.Name})
	if err != nil {
		errutils.Exitf("failed to find runtime %s. %s", function.Name, *err.Message)
	}
	if len(runtimes) == 0 {
		errutils.Exitf("function %s does not exist", function.Name)
	}
	argTypes := make(map[string]string)

	for _, a := range runtimes[0].Params {
		argTypes[*a.Name] = *a.Type
	}

	params := []*models.ScriptParamValue{}
	for k, v := range function.Args {
		if _, ok := argTypes[k]; !ok {
			errutils.Exitf("arg %s does not exist for function %s", k, function.Name)
		}
		valType := "string"
		if argTypes[k] == "integer" {
			valType = "integer"
			_, err := strconv.Atoi(v)
			if err != nil {
				errutils.Exitf("expected type integer, for arg %s", k)
			}
		}
		params = append(params, &models.ScriptParamValue{Name: &k, Value: &v, Type: &valType})
	}

	return &models.TransformationArgs{TransformationID: &runtimes[0].ID, Args: params}
}

func ToCategory(cat Category) *models.Category {
	return &models.Category{Name: &cat.Name, Purpose: &cat.Description, Values: cat.Values}
}

func ToDataPipeline(dp DataPipeline, client cloudmgmt.Client, tenantID string) *models.DataPipeline {
	rtn := models.DataPipeline{}
	rtn.Name = &dp.Name
	rtn.TenantID = &tenantID
	rtn.ProjectID = ToProjectID(dp.Project, client)
	if dp.SamplingIntervalMSec > 0 {
		rtn.SamplingInterval = float64(dp.SamplingIntervalMSec)
		samplingEnabled := true
		rtn.EnableSampling = &samplingEnabled
	}

	rtn.Description = dp.Description
	origin := models.DataPipelineOriginDataSource

	if dp.Input.DataPipeline == nil {
		rtn.OriginSelectors = ToCatInfo(dp.Input.CatSels, client)
	} else {
		origin = models.DataPipelineOriginDataStream
		rtn.OriginID = ToDataPipelineID(*dp.Input.DataPipeline, client)
	}
	rtn.Origin = &origin

	dest := models.DataPipelineDestinationEdge
	if dp.Output.PublicCloud != nil {
		dest = models.DataPipelineDestinationCloud
		rtn.CloudType = dp.Output.PublicCloud.Type
		rtn.EndPoint = dp.Output.PublicCloud.EndpointName
		switch rtn.CloudType {
		case "AWS":
			rtn.AWSCloudRegion = dp.Output.PublicCloud.Region
			rtn.AWSStreamType = dp.Output.PublicCloud.Service
		case "GCP":
			rtn.GCPCloudRegion = dp.Output.PublicCloud.Region
			rtn.GCPStreamType = dp.Output.PublicCloud.Service
		}
		rtn.CloudCredsID = ToCloudProfileID(dp.Output.PublicCloud.Profile, client)
	} else {
		rtn.EndPoint = dp.Output.LocalEdge.EndpointName
		if dp.Output.LocalEdge.Type == "DataInterface" {
			dest = models.DataStreamDestinationDataInterface
			dataSrcID := ToDataSourceID(dp.Output.LocalEdge.Service, client)
			ifcEndpoint := models.DataIfcEndpoint{ID: &dataSrcID, Name: &dp.Output.LocalEdge.EndpointName, Value: &dp.Output.LocalEdge.EndpointName}
			rtn.DataIfcEndpoints = []*models.DataIfcEndpoint{&ifcEndpoint}
		} else {
			// Real time streams do not need anything special. Edge always outputs it
			// to real tim stream with name of the stream being identifier.
			if dp.Output.LocalEdge.Service != "RealTimeStreaming" {
				rtn.EdgeStreamType = dp.Output.LocalEdge.Service
			} else {
				rtn.EdgeStreamType = "None"
			}
		}
	}
	rtn.Destination = &dest
	rtn.TransformationArgsList = []*models.TransformationArgs{}
	for _, f := range dp.Functions {
		rtn.TransformationArgsList = append(rtn.TransformationArgsList, ToTransform(f, client))
	}
	return &rtn
}

func ToScriptParams(params []Params) []*models.ScriptParam {
	var scriptParam []*models.ScriptParam
	for _, p := range params {
		sp := models.ScriptParam{}
		sp.Name = &p.Name
		sp.Type = &p.Type
		scriptParam = append(scriptParam, &sp)
	}
	return scriptParam
}

func ToFunction(trans Transformation, client cloudmgmt.Client) *models.Function {
	function := models.Function{}
	b, err := ioutil.ReadFile(trans.SourceCodePath)
	if err != nil {
		errutils.Exitf("Failed to read function source code file. %s", err.Error())
	}
	code := string(b)
	functionType := "Transformation"
	function.Name = &trans.Name
	function.ProjectID = ToProjectID(trans.Project, client)
	function.Language = &trans.Language
	function.Environment = &trans.Environment
	function.Code = &code
	function.Type = &functionType
	function.Params = ToScriptParams(trans.Params)
	function.Description = trans.Description
	return &function
}
