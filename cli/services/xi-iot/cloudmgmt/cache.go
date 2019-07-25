package cloudmgmt

import (
	"fmt"
	"xi-iot-cli/generated/swagger/models"
	"xi-iot-cli/xi-iot/errutils"
	xi_models "xi-iot-cli/xi-iot/models"
)

// Cache encapsulates cloud mgmt objects for fast retrieval
type Cache struct {
	client            *Client
	projCache         map[string]*models.Project
	dataSrcCache      map[string]*models.DataSourceV2
	dataPipelineCache map[string]*models.DataPipeline
	edgeCache         map[string]*models.EdgeV2
	cloudProfileCache map[string]*models.CloudProfile
	categoryCache     map[string]*models.Category
	functionCache     map[string]*models.Function
}

// NewCache returns a new instance of cache
func NewCache(c *Client) *Cache {
	return &Cache{
		client: c,
	}
}

// Refresh refreshes the cache for the given entity if the cache is empty
// The idea here is to only refresh the cache for the current run of the CLI. That is, ignore
// any changes in cloud mgmt while the command is being executed
func (c *Cache) refresh(entity string) {
	var err *models.APIErrorPayload
	switch entity {
	case xi_models.EntityProject:
		if c.projCache != nil {
			break
		}
		projs, apiErr := c.client.ListProjects("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.projCache = make(map[string]*models.Project)
		for _, p := range projs {
			c.projCache[p.ID] = p
		}
	case xi_models.EntityDataSource:
		if c.dataSrcCache != nil {
			break
		}
		dataSources, apiErr := c.client.ListDataSources("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.dataSrcCache = make(map[string]*models.DataSourceV2)
		for _, d := range dataSources {
			c.dataSrcCache[d.ID] = d
		}
	case xi_models.EntityDataPipeline:
		if c.dataPipelineCache != nil {
			break
		}
		dataPipelines, apiErr := c.client.ListDataPipelines("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.dataPipelineCache = make(map[string]*models.DataPipeline)
		for _, d := range dataPipelines {
			c.dataPipelineCache[d.ID] = d
		}
	case xi_models.EntityFunction:
		if c.functionCache != nil {
			break
		}
		functions, apiErr := c.client.ListFunctions("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.functionCache = make(map[string]*models.Function)
		for _, f := range functions {
			c.functionCache[f.ID] = f
		}
	case xi_models.EntityEdge:
		if c.edgeCache != nil {
			break
		}
		edges, apiErr := c.client.ListEdges("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.edgeCache = make(map[string]*models.EdgeV2)
		for _, e := range edges {
			c.edgeCache[e.ID] = e
		}
	case xi_models.EntityCloudProfile:
		if c.cloudProfileCache != nil {
			break
		}
		profiles, apiErr := c.client.ListCloudProfiles("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.cloudProfileCache = make(map[string]*models.CloudProfile)
		for _, p := range profiles {
			c.cloudProfileCache[p.ID] = p
		}
	case xi_models.EntityCategory:
		if c.categoryCache != nil {
			break
		}
		cats, apiErr := c.client.ListCategories("")
		if apiErr != nil {
			err = apiErr
			break
		}
		c.categoryCache = make(map[string]*models.Category)
		for _, cat := range cats {
			c.categoryCache[cat.ID] = cat
		}
	}
	// TODO: Replace exit with returning error so that caller can throw a useful error
	if err != nil {
		errutils.Exit(fmt.Sprintf("failed to %s. %s", entity, *err.Message))
	}
}

// FetchProjects fetches projects from the cache with the given IDs
func (c *Cache) FetchProjects(ids []string) (rtn []*models.Project) {
	c.refresh(xi_models.EntityProject)
	for _, id := range ids {
		if v, ok := c.projCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchDataSources fetches data sources from the cache with the given IDs
func (c *Cache) FetchDataSources(ids []string) (rtn []*models.DataSourceV2) {
	c.refresh(xi_models.EntityDataSource)
	for _, id := range ids {
		if v, ok := c.dataSrcCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchEdges fetches edges from the cache with the given IDs
func (c *Cache) FetchEdges(ids []string) (rtn []*models.EdgeV2) {
	c.refresh(xi_models.EntityEdge)
	for _, id := range ids {
		if v, ok := c.edgeCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchCategories fetches categories from the cache with the given IDs
func (c *Cache) FetchCategories(ids []string) (rtn []*models.Category) {
	c.refresh(xi_models.EntityCategory)
	for _, id := range ids {
		if v, ok := c.categoryCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchCloudProfiles fetches cloud profiles from the cache with the given IDs
func (c *Cache) FetchCloudProfiles(ids []string) (rtn []*models.CloudProfile) {
	c.refresh(xi_models.EntityCloudProfile)
	for _, id := range ids {
		if v, ok := c.cloudProfileCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchDataPipelines fetches data pipelines from the cache with the given IDs
func (c *Cache) FetchDataPipelines(ids []string) (rtn []*models.DataPipeline) {
	c.refresh(xi_models.EntityDataPipeline)
	for _, id := range ids {
		if v, ok := c.dataPipelineCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}

// FetchFunctions fetches functions from the cache with the given IDs
func (c *Cache) FetchFunctions(ids []string) (rtn []*models.Function) {
	c.refresh(xi_models.EntityFunction)
	for _, id := range ids {
		if v, ok := c.functionCache[id]; ok {
			rtn = append(rtn, v)
		}
	}
	return
}
