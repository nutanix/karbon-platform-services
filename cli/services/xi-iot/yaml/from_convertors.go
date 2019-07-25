package yaml

import (
	"fmt"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/cloudmgmt"

	"xi-iot-cli/xi-iot/errutils"
)

const (
	ApplicationKind  = "application"
	DataSrcKind      = "dataSource"
	DataPipelineKind = "dataPipeline"
	CategoryKind     = "category"
	ProjectKind      = "project"
	FunctionKind     = "function"
	EdgeKind         = "edge"
)

var (
	projCache         = make(map[string]*models.Project)
	dataSrcCache      = make(map[string]string)
	edgeCache         = make(map[string]string)
	cloudProfileCache = make(map[string]string)
	pipelineCache     = make(map[string]string)
	categoryCache     = make(map[string]string)
	functionCache     = make(map[string]*models.Function)
)

func ToEdgeNames(edgeIDs []string, cache *cloudmgmt.Cache) (edgeNames []string) {
	for _, edge := range cache.FetchEdges(edgeIDs) {
		edgeNames = append(edgeNames, *edge.Name)
	}
	return
}

func fromCategoryInfoToCatInfo(info []*models.CategoryInfo) (catInfo []CatInfo) {
	for _, i := range info {
		catInfo = append(catInfo, CatInfo{Name: *i.ID, Value: *i.Value})
	}
	return
}

func FromApplication(app *models.ApplicationV2, cache *cloudmgmt.Cache) Application {
	rtn := Application{}
	rtn.Name = *app.Name

	if app.ProjectID != "" {
		projects := cache.FetchProjects([]string{app.ProjectID})
		if len(projects) == 0 {
			errutils.Exit(fmt.Sprintf("failed to get projects for application %s", *app.Name))
		}
		rtn.Project = *projects[0].Name
	}
	rtn.Description = app.Description
	rtn.DataIfcEndpoints = make([]DataIfcEndpoint, 0, len(app.DataIfcEndpoints))
	if len(app.EdgeIds) > 0 {
		rtn.Edges = ToEdgeNames(app.EdgeIds, cache)
	} else if len(app.EdgeSelectors) > 0 {
		rtn.EdgeSelectors = fromCategoryInfoToCatInfo(app.EdgeSelectors)
	}
	for _, e := range app.DataIfcEndpoints {
		dataSources := cache.FetchDataSources([]string{*e.ID})
		if len(dataSources) == 0 {
			errutils.Exitf("failed to get data sourdces for application %s", *app.Name)
		}
		rtn.DataIfcEndpoints = append(rtn.DataIfcEndpoints, DataIfcEndpoint{DataSource: *dataSources[0].Name,
			TopicName: *e.Value, FieldName: *e.Name},
		)
	}
	rtn.AppYaml = *app.AppManifest
	rtn.Kind = ApplicationKind
	return rtn
}

// FromCategory converts cloudmgmt category instance into `Category`
func FromCategory(cat *models.Category, cache *cloudmgmt.Cache) Category {
	rtn := Category{}
	rtn.Name = *cat.Name
	rtn.Description = *cat.Purpose
	rtn.Values = cat.Values
	rtn.Kind = CategoryKind
	return rtn
}

func ToCloudProfileNames(ids []string, cache *cloudmgmt.Cache) (cloudProfNames []string) {
	for _, profile := range cache.FetchCloudProfiles(ids) {
		cloudProfNames = append(cloudProfNames, *profile.Name)
	}
	return
}

func FromProject(proj *models.Project, cache *cloudmgmt.Cache) Project {
	rtn := Project{}
	rtn.Name = *proj.Name
	rtn.Kind = ProjectKind
	rtn.Description = *proj.Description

	if len(proj.EdgeIds) > 0 {
		rtn.Edges = ToEdgeNames(proj.EdgeIds, cache)
	}

	if len(proj.CloudCredentialIds) > 0 {
		rtn.CloudProfiles = ToCloudProfileNames(proj.CloudCredentialIds, cache)
	}

	return rtn
}

// FromFunction convert a cloudmgmt function object to a yaml Function model
func FromFunction(function *models.Function, cache *cloudmgmt.Cache) Transformation {
	rtn := Transformation{}
	rtn.Name = *function.Name
	rtn.Language = *function.Language
	if function.ProjectID != "" {
		projects := cache.FetchProjects([]string{function.ProjectID})
		if len(projects) == 0 {
			errutils.Exit(fmt.Sprintf("failed to get projects for function %s", *function.Name))
		}
		rtn.Project = *projects[0].Name
	}
	rtn.Environment = *function.Environment
	rtn.Params = make([]Params, 0, len(rtn.Params))
	rtn.Description = function.Description
	rtn.Kind = FunctionKind
	for _, p := range function.Params {
		rtn.Params = append(rtn.Params, Params{Name: *p.Name, Type: *p.Type})
		rtn.Name = *function.Name
	}
	return rtn
}

func fromCatInfosToCatSels(catInfos []*models.CategoryInfo, cache *cloudmgmt.Cache) (map[string][]string, error) {
	if len(catInfos) == 0 {
		return nil, nil
	}

	valsByID := make(map[string][]string)

	for _, o := range catInfos {
		valsByID[*o.ID] = append(valsByID[*o.ID], *o.Value)
	}

	rtn := make(map[string][]string)
	for _, cInfo := range catInfos {
		cats := cache.FetchCategories([]string{*cInfo.ID})
		if len(cats) == 0 {
			errutils.Exit(fmt.Sprintf("failed to fetch categories"))
		}
		rtn[*cats[0].Name] = valsByID[cats[0].ID]
	}
	return rtn, nil
}

func ToPipelineNames(ids []string, cache *cloudmgmt.Cache) (pipelineNames []string) {
	for _, pipeline := range cache.FetchDataPipelines(ids) {
		pipelineNames = append(pipelineNames, *pipeline.Name)
	}
	return
}

// FromDatapipeline converts a cloudmgmt pipeline to a yaml model pipeline
func FromDatapipeline(dp *models.DataPipeline, cache *cloudmgmt.Cache) DataPipeline {
	rtn := DataPipeline{}
	rtn.Kind = DataPipelineKind
	rtn.Name = *dp.Name
	rtn.Description = dp.Description
	projects := cache.FetchProjects([]string{dp.ProjectID})
	if len(projects) == 0 {
		errutils.Exit(fmt.Sprintf("failed to get project for data pipeline %s", *dp.Name))
	}
	rtn.Project = *projects[0].Name
	rtn.SamplingIntervalMSec = int(dp.SamplingInterval)
	rtn.Output = &DataPipelineOutput{}
	rtn.Input = &DataPipelineInput{}

	// input
	if *dp.Origin == models.DataPipelineOriginDataStream {
		pipelines := cache.FetchDataPipelines([]string{dp.OriginID})
		if len(pipelines) == 0 {
			errutils.Exit(fmt.Sprintf("failed to find origin pipeline for data pipeline %s", *dp.Name))
		}
		rtn.Input.DataPipeline = pipelines[0].Name
	} else {
		catSels, err := fromCatInfosToCatSels(dp.OriginSelectors, cache)
		if err != nil {
			errutils.Exitf("failed to list categories for pipeline %s. %s", *dp.Name, err.Error())
		}
		rtn.Input.CatSels = catSels
	}

	// output
	switch *dp.Destination {
	case models.DataPipelineDestinationCloud:
		rtn.Output.PublicCloud = &PublicCloud{Type: dp.CloudType}
		switch dp.CloudType {
		case models.DataPipelineCloudTypeAWS:
			rtn.Output.PublicCloud.Region = dp.AWSCloudRegion
			rtn.Output.PublicCloud.Service = dp.AWSStreamType
		case models.DataPipelineCloudTypeGCP:
			rtn.Output.PublicCloud.Region = dp.GCPCloudRegion
			rtn.Output.PublicCloud.Service = dp.GCPStreamType
		default:
			errutils.Exit(fmt.Sprintf("unknown/unsupported cloud type %s", dp.CloudType))
		}
		rtn.Output.PublicCloud.EndpointName = dp.EndPoint
		profiles := cache.FetchCloudProfiles([]string{dp.CloudCredsID})
		if len(profiles) == 0 {
			errutils.Exit(fmt.Sprintf("failed to find cloud profiles for data pipeline %s", *dp.Name))
		}
		rtn.Output.PublicCloud.Profile = *profiles[0].Name

	case models.DataPipelineDestinationDataInterface:
		rtn.Output.LocalEdge = &LocalEdge{Type: models.DataStreamDestinationDataInterface}
		if len(dp.DataIfcEndpoints) == 0 {
			errutils.Exitf("pipeline is connected to local edge service of type DataInterface, but is missing data ifc endpoints")
		}
		dataSources := cache.FetchDataSources([]string{*dp.DataIfcEndpoints[0].ID})
		if len(dataSources) == 0 {
			errutils.Exit(fmt.Sprintf("failed to get LocalEdge.Service service for data pipeline %s", *dp.Name))
		}
		rtn.Output.LocalEdge.Service = *dataSources[0].Name
		rtn.Output.LocalEdge.EndpointName = dp.EndPoint

	case models.DataPipelineDestinationEdge:
		rtn.Output.LocalEdge = &LocalEdge{Type: models.DataPipelineDestinationEdge, EndpointName: dp.EndPoint}
		if dp.EdgeStreamType == "None" {
			rtn.Output.LocalEdge.Service = "RealTimeStreaming"
		} else {
			rtn.Output.LocalEdge.Service = dp.EdgeStreamType
		}
	default:
		errutils.Exitf("unknown destination: %q", *dp.Destination)
	}

	// transformations
	rtn.Functions = make([]Function, 0, len(dp.TransformationArgsList))
	for _, t := range dp.TransformationArgsList {
		funcs := cache.FetchFunctions([]string{*t.TransformationID})
		if len(funcs) == 0 {
			errutils.Exit(fmt.Sprintf("failed to fetch functions for data pipeline %s", *dp.Name))
		}

		args := make(map[string]string)
		for _, arg := range t.Args {
			args[*arg.Name] = *arg.Value
		}
		rtn.Functions = append(rtn.Functions, Function{Name: *funcs[0].Name, Args: args})
	}
	return rtn
}

// FromDataSource converts cloudmgmt DataSource to yaml DataSource model
func FromDataSource(ds *models.DataSourceV2, cache *cloudmgmt.Cache) DataSource {
	rtn := DataSource{Name: *ds.Name, Protocol: *ds.Protocol, AuthType: *ds.AuthType, Type: *ds.Type}
	for _, f := range ds.FieldsV2 {
		rtn.Fields = append(rtn.Fields, Field{Name: *f.Name, Topic: *f.Topic})
	}
	if ds.IfcInfo != nil {
		rtn.IfcInfo = &IfcInfo{Class: *ds.IfcInfo.Class, Img: *ds.IfcInfo.Img,
			Kind: *ds.IfcInfo.Kind, Protocol: *ds.IfcInfo.Protocol}

		for _, p := range ds.IfcInfo.Ports {
			rtn.IfcInfo.Ports = append(rtn.IfcInfo.Ports, Port{Name: *p.Name, Port: int(*p.Port)})
		}
	}
	rtn.Kind = DataSrcKind
	edges := cache.FetchEdges([]string{*ds.EdgeID})
	if len(edges) == 0 {
		errutils.Exitf("failed to find edge for data source %s", *ds.Name)
	}
	rtn.Edge = *edges[0].Name
	for _, s := range ds.Selectors {
		cats := cache.FetchCategories([]string{*s.ID})
		if len(cats) == 0 {
			errutils.Exitf("failed to find categories for data source %s", *ds.Name)
		}
		rtn.Selectors = append(rtn.Selectors, Selector{CategoryName: *cats[0].Name, CategoryValue: *s.Value, Scope: s.Scope})
	}
	return rtn
}

// FromEdge converts cloudmgmt Edge to yaml Edge model
func FromEdge(e *models.EdgeV2, cache *cloudmgmt.Cache) Edge {
	rtn := Edge{Name: *e.Name, Kind: EdgeKind, Description: e.Description,
		Connected: e.Connected, IPAddress: *e.IPAddress, SerialNumber: *e.SerialNumber,
	}
	catSels, err := fromCatInfosToCatSels(e.Labels, cache)
	if err != nil {
		errutils.Exitf("failed to list categories for edge %s. %s", *e.Name, err.Error())
	}
	rtn.CatSels = catSels
	return rtn
}
