package cloudmgmt

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
	cme_client "xi-iot-cli/generated/swagger/client"
	"xi-iot-cli/generated/swagger/models"

	"xi-iot-cli/generated/swagger/client/application"
	"xi-iot-cli/generated/swagger/client/auth"
	"xi-iot-cli/generated/swagger/client/category"
	"xi-iot-cli/generated/swagger/client/certificate"
	"xi-iot-cli/generated/swagger/client/cloud_profile"
	"xi-iot-cli/generated/swagger/client/data_pipeline"
	"xi-iot-cli/generated/swagger/client/data_source"
	"xi-iot-cli/generated/swagger/client/edge"
	"xi-iot-cli/generated/swagger/client/edge_info"
	"xi-iot-cli/generated/swagger/client/function"
	"xi-iot-cli/generated/swagger/client/log"
	"xi-iot-cli/generated/swagger/client/project"
	"xi-iot-cli/generated/swagger/client/runtime_environment"

	"xi-iot-cli/xi-iot/errutils"
	"xi-iot-cli/xi-iot/io_utils"
	xi_models "xi-iot-cli/xi-iot/models"

	httptransport "github.com/go-openapi/runtime/client"
	runtime "github.com/go-openapi/runtime/client"
	"github.com/golang/glog"
)

// Client ...
type Client struct {
	client        *cme_client.Sherlock
	ctx           *xi_models.Context
	mu            *sync.Mutex
	ctxUpdateChan chan<- xi_models.Context
}

// JWTPayload ...
type JWTPayload struct {
	Expiry      int64  `json:"exp"`
	Email       string `json:"email"`
	SpecialRole string `json:"specialRole"`
	TenantID    string `json:"tenantID"`
}

// New returns an instance of cloud mgmt client
func New(cmeURL string, opts ...func(c *Client)) *Client {
	clientConfig := &cme_client.TransportConfig{
		Host:     cmeURL,
		BasePath: "",
		Schemes:  []string{"https"},
	}
	cloudMgmtClient := cme_client.NewHTTPClientWithConfig(nil, clientConfig)
	transport := httptransport.New(clientConfig.Host, clientConfig.BasePath, clientConfig.Schemes)
	transport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cloudMgmtClient.SetTransport(transport)
	c := &Client{client: cloudMgmtClient, mu: &sync.Mutex{}}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithContext(context xi_models.Context) func(c *Client) {
	return func(c *Client) {
		c.ctx = &context
	}
}

func WithCtxUpdateChan(ctxUpdateCh chan<- xi_models.Context) func(c *Client) {
	return func(c *Client) {
		c.ctxUpdateChan = ctxUpdateCh
	}
}

func b64Decode(s string) []byte {
	var err error
	var rtn []byte
	// Decode the middle part(payload/claims)
	if len(s)%4 == 0 {
		rtn, err = base64.URLEncoding.DecodeString(s)
	} else {
		rtn, err = base64.RawURLEncoding.DecodeString(s)
	}
	if err != nil {
		glog.Fatalf("failed to decode: %s. %s", s, err)
	}
	return rtn
}

func ParseJWTPayload(token string) JWTPayload {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		glog.Fatalf("malformed token: %s", token)
	}

	rtn := b64Decode(parts[1])
	payload := JWTPayload{}
	err := json.Unmarshal(rtn, &payload)
	if err != nil {
		glog.Fatalf("failed to unmarshal toke payload while veryfying the expiry. %s", err.Error())
	}
	return payload
}

func IsTokenExpired(token string) bool {
	// Empty token is treated as token expired
	if token == "" {
		return true
	}

	payload := ParseJWTPayload(token)
	// This would not work if the cloud mgmt and client are in different time zones
	// TODO: Cloud mgmt should use utc instead of current local time
	if payload.Expiry < time.Now().Unix() {
		return true
	}
	return false
}

// RefreshToken refreshes token in the context based on expiry
func (c *Client) RefreshToken() {
	// synchronize refresh token call from multiple API's
	c.mu.Lock()
	defer c.mu.Unlock()

	if !IsTokenExpired(c.ctx.Token) {
		return
	}

	token, err := c.Login()
	if err != nil {
		errutils.Exitf("failed to login. %s", *err.Message)
	}
	c.ctx.Token = token
	c.ctx.TenantID = c.ParseTenantID(token)

	// Send the updated context, so that the writing goroutine can update it on disk
	go func(ctx xi_models.Context) {
		c.ctxUpdateChan <- ctx
	}(*c.ctx)
}

func (c *Client) ParseTenantID(token string) string {
	payload := ParseJWTPayload(token)
	return payload.TenantID
}

func errorToAPIError(err error) *models.APIErrorPayload {
	errString := err.Error()
	return &models.APIErrorPayload{Message: &errString}
}

func (c *Client) Login() (string, *models.APIErrorPayload) {
	loginParams := auth.NewLoginCallV2Params()
	loginParams.Request = &models.Credential{Email: &c.ctx.Email, Password: &c.ctx.Password}
	ok, err := c.client.Auth.LoginCallV2(loginParams)
	if err != nil {
		apiErr, castOK := err.(*auth.LoginCallV2Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.Token, nil
}

func getFilterFromNames(names []string) string {
	ids := io_utils.ToSQLList(names)
	filter := fmt.Sprintf("name IN (%s)", ids)
	return filter
}

func (c *Client) ListEdgeInfos(filter string) ([]*models.EdgeUsageInfo, *models.APIErrorPayload) {
	c.RefreshToken()
	edgeParam := edge_info.NewEdgeInfoListV2Params()
	edgeParam.Filter = &filter
	ok, err := c.client.EdgeInfo.EdgeInfoListV2(edgeParam, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*edge_info.EdgeInfoListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.EdgeUsageInfoList, nil
}

// This API is broken, tracker: [ENG-242048]
func (c *Client) ListEdgeInfosProjects(projectID string) ([]*models.EdgeUsageInfo, *models.APIErrorPayload) {
	c.RefreshToken()
	edgeParam := edge_info.NewProjectGetEdgesInfoV2Params()
	edgeParam.ProjectID = projectID
	fmt.Println(projectID)
	ok, err := c.client.EdgeInfo.ProjectGetEdgesInfoV2(edgeParam, runtime.BearerToken(c.ctx.Token))
	fmt.Println(err.Error())
	if err != nil {
		apiErr, castOK := err.(*edge_info.ProjectGetEdgesInfoV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

func (c *Client) ListEdges(filter string) ([]*models.EdgeV2, *models.APIErrorPayload) {
	c.RefreshToken()
	edgeParam := edge.NewEdgeListV2Params()
	edgeParam.Filter = &filter
	ok, err := c.client.Edge.EdgeListV2(edgeParam, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*edge.EdgeListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.EdgeListV2, nil
}

func (c *Client) ListEdgesByNames(names []string) ([]*models.EdgeV2, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.EdgeV2{}, nil
	}
	c.RefreshToken()
	return c.ListEdges(getFilterFromNames(names))
}

func (c *Client) GetDataSourceArtifact(ID string) *models.DataSourceArtifact {
	c.RefreshToken()
	var requestBody bytes.Buffer
	url := fmt.Sprintf("https://%s/v1.0/datasources/%s/artifacts", c.ctx.URL, ID)
	req, err := http.NewRequest("GET", url, &requestBody)
	if err != nil {
		glog.Errorf("failed to create http req for artifacts(%s): %s", url, err.Error())
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ctx.Token))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		glog.Errorf("failed to get resp from artifacts API: %s", err.Error())
	}
	defer response.Body.Close()
	artifact := models.DataSourceArtifact{}
	err = json.NewDecoder(response.Body).Decode(&artifact)
	if err != nil {
		glog.Fatalf("failed to decode artifacts %s: %s", response.Body, err.Error())
	}
	return &artifact
}

func (c *Client) GetDataPipelinesProjects(projectID string) ([]*models.DataPipeline, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewProjectGetDataPipelinesParams()
	params.ProjectID = projectID
	ok, err := c.client.DataPipeline.ProjectGetDataPipelines(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_pipeline.ProjectGetDataPipelinesDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.DataPipelineList, nil
}

func (c *Client) GetDataSource(ID string) (*models.DataSourceV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewDataSourceGetV2Params()
	params.ID = ID
	ok, err := c.client.DataSource.DataSourceGetV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_source.DataSourceGetV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

func (c *Client) ListApps(filter string) ([]*models.ApplicationV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := application.NewApplicationListV2Params()
	params.Filter = &filter
	ok, err := c.client.Application.ApplicationListV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*application.ApplicationListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.ApplicationListV2, nil
}

func (c *Client) ListAppsByNames(names []string) ([]*models.ApplicationV2, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.ApplicationV2{}, nil
	}
	c.RefreshToken()
	return c.ListApps(getFilterFromNames(names))
}

func (c *Client) ListCategories(filter string) ([]*models.Category, *models.APIErrorPayload) {
	c.RefreshToken()
	params := category.NewCategoryListV2Params()
	params.Filter = &filter
	ok, err := c.client.Category.CategoryListV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*category.CategoryListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.CategoryList, nil
}

func (c *Client) ListRuntimeEnvironments(filter string) ([]*models.RuntimeEnvironment, *models.APIErrorPayload) {
	c.RefreshToken()
	params := runtime_environment.NewRuntimeEnvironmentListParams()
	params.Filter = &filter
	ok, err := c.client.RuntimeEnvironment.RuntimeEnvironmentList(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*runtime_environment.RuntimeEnvironmentListDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.RuntimeEnvironmentList, nil
}

func (c *Client) ListFunctionByNames(names []string) ([]*models.Function, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.Function{}, nil
	}
	c.RefreshToken()
	return c.ListFunctions(getFilterFromNames(names))
}

func (c *Client) ListFunctions(filter string) ([]*models.Function, *models.APIErrorPayload) {
	c.RefreshToken()
	params := function.NewFunctionListParams()
	params.Filter = &filter
	ok, err := c.client.Function.FunctionList(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*function.FunctionListDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.FunctionList, nil
}

func (c *Client) ListCategoriesByNames(names []string) ([]*models.Category, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.Category{}, nil
	}
	c.RefreshToken()
	return c.ListCategories(getFilterFromNames(names))
}

func (c *Client) ListProjectByNames(names []string) ([]*models.Project, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.Project{}, nil
	}
	c.RefreshToken()
	return c.ListProjects(getFilterFromNames(names))
}

func (c *Client) ListProjects(filter string) ([]*models.Project, *models.APIErrorPayload) {
	c.RefreshToken()
	params := project.NewProjectListV2Params()
	params.Filter = &filter
	ok, err := c.client.Project.ProjectListV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*project.ProjectListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.ProjectList, nil
}

func (c *Client) ListDataSources(filter string) ([]*models.DataSourceV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewDataSourceListV2Params()
	params.Filter = &filter
	ok, err := c.client.DataSource.DataSourceListV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_source.DataSourceListV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.DataSourceListV2, nil
}

func (c *Client) ListDataSourcesByNames(names []string) ([]*models.DataSourceV2, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.DataSourceV2{}, nil
	}
	c.RefreshToken()
	return c.ListDataSources(getFilterFromNames(names))
}

func (c *Client) ListCloudProfilesByNames(names []string) ([]*models.CloudProfile, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.CloudProfile{}, nil
	}
	c.RefreshToken()
	params := cloud_profile.NewCloudProfileListParams()
	ids := io_utils.ToSQLList(names)
	filter := fmt.Sprintf("name IN (%s)", ids)
	params.Filter = &filter
	ok, err := c.client.CloudProfile.CloudProfileList(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*cloud_profile.CloudProfileListDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.CloudProfileList, nil
}

func (c *Client) ListCloudProfiles(filter string) ([]*models.CloudProfile, *models.APIErrorPayload) {
	c.RefreshToken()
	params := cloud_profile.NewCloudProfileListParams()
	params.Filter = &filter
	ok, err := c.client.CloudProfile.CloudProfileList(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*cloud_profile.CloudProfileListDefault).Payload
	}
	return ok.Payload.CloudProfileList, nil
}

func (c *Client) ListDataPipelines(filter string) ([]*models.DataPipeline, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewDataPipelineListParams()
	params.Filter = &filter
	ok, err := c.client.DataPipeline.DataPipelineList(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_pipeline.DataPipelineListDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.DataPipelineList, nil
}

func (c *Client) ListDataPipelinesByNames(names []string) ([]*models.DataPipeline, *models.APIErrorPayload) {
	if len(names) == 0 {
		return []*models.DataPipeline{}, nil
	}
	c.RefreshToken()
	return c.ListDataPipelines(getFilterFromNames(names))
}

func (c *Client) ListDataSourcesByEdgeID(edgeID string) ([]*models.DataSourceV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewEdgeGetDatasourcesV2Params()
	params.EdgeID = edgeID
	ok, err := c.client.DataSource.EdgeGetDatasourcesV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_source.EdgeGetDatasourcesV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.DataSourceListV2, nil
}

func (c *Client) ListAppContainers(appID string, edgeID string) ([]string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := application.NewGetApplicationContainersParams()
	params.ID = appID
	params.EdgeID = edgeID
	ok, err := c.client.Application.GetApplicationContainers(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*application.GetApplicationContainersDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.ContainerNames, nil
}

func (c *Client) ListPipelineContainers(pipeID string, edgeID string) ([]string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewGetDataPipelineContainersParams()
	params.ID = pipeID
	params.EdgeID = edgeID
	ok, err := c.client.DataPipeline.GetDataPipelineContainers(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_pipeline.GetDataPipelineContainersDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.ContainerNames, nil
}

func (c *Client) GetLogEndpoint(ls *models.LogStream) (*string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := log.NewLogStreamEndpointsParams()
	params.Request = ls
	ok, err := c.client.Log.LogStreamEndpoints(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*log.LogStreamEndpointsDefault)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.URL, nil
}

func (c *Client) CreateCertificates() (*models.Certificates, *models.APIErrorPayload) {
	certParam := certificate.NewCertificatesCreateV2Params()
	ok, err := c.client.Certificate.CertificatesCreateV2(certParam, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*certificate.CertificatesCreateV2Default)
		if castOK {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

func (c *Client) CreateDataSource(ds *models.DataSourceV2) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewDataSourceCreateV2Params()
	params.Body = ds
	ok, err := c.client.DataSource.DataSourceCreateV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_source.DataSourceCreateV2Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) CreateDataPipeline(ds *models.DataPipeline) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewDataPipelineCreateParams()
	params.Body = ds
	ok, err := c.client.DataPipeline.DataPipelineCreate(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_pipeline.DataPipelineCreateDefault)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) CreateCategory(cat *models.Category) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := category.NewCategoryCreateV2Params()
	params.Body = cat
	ok, err := c.client.Category.CategoryCreateV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*category.CategoryCreateV2Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) CreateApp(app *models.ApplicationV2) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := application.NewApplicationCreateV2Params()
	params.Body = app
	ok, err := c.client.Application.ApplicationCreateV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*application.ApplicationCreateV2Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) CreateFunction(f *models.Function) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := function.NewFunctionCreateParams()
	params.Body = f
	ok, err := c.client.Function.FunctionCreate(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*function.FunctionCreateDefault)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) UpdateDataSource(ds *models.DataSourceV2, id string) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewDataSourceUpdateV3Params()
	params.ID = id
	params.Body = ds
	ok, err := c.client.DataSource.DataSourceUpdateV3(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_source.DataSourceUpdateV3Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) UpdateDataPipeline(ds *models.DataPipeline, id string) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewDataPipelineUpdateParams()
	params.ID = id
	params.Body = ds
	ok, err := c.client.DataPipeline.DataPipelineUpdate(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*data_pipeline.DataPipelineUpdateDefault)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) UpdateCategory(cat *models.Category, id string) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := category.NewCategoryUpdateV3Params()
	params.ID = id
	params.Body = cat
	ok, err := c.client.Category.CategoryUpdateV3(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*category.CategoryUpdateV3Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) UpdateApp(app *models.ApplicationV2, id string) (string, *models.APIErrorPayload) {
	c.RefreshToken()
	params := application.NewApplicationUpdateV3Params()
	params.ID = id
	params.Body = app
	ok, err := c.client.Application.ApplicationUpdateV3(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		apiErr, castOK := err.(*application.ApplicationUpdateV3Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.ID, nil
}

func (c *Client) DeleteDataSrc(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := data_source.NewDataSourceDeleteV2Params()
	deleteParams.ID = strings.TrimSpace(ID)
	_, err := c.client.DataSource.DataSourceDeleteV2(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*data_source.DataSourceDeleteV2Default)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) DeleteCategory(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := category.NewCategoryDeleteV2Params()
	deleteParams.ID = strings.TrimSpace(ID)
	_, err := c.client.Category.CategoryDeleteV2(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*category.CategoryDeleteV2Default)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) DeleteApp(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := application.NewApplicationDeleteV2Params()
	deleteParams.ID = strings.TrimSpace(ID)
	_, err := c.client.Application.ApplicationDeleteV2(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*application.ApplicationDeleteV2Default)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) DeleteDataPipeline(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := data_pipeline.NewDataPipelineDeleteParams()
	deleteParams.ID = strings.TrimSpace(ID)
	_, err := c.client.DataPipeline.DataPipelineDelete(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*data_pipeline.DataPipelineDeleteDefault)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) DeleteFunction(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := function.NewFunctionDeleteParams()
	deleteParams.ID = strings.TrimSpace(ID)
	_, err := c.client.Function.FunctionDelete(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*function.FunctionDeleteDefault)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) DeleteEdge(ID string) *models.APIErrorPayload {
	c.RefreshToken()
	deleteParams := edge.NewEdgeDeleteV2Params()
	deleteParams.EdgeID = strings.TrimSpace(ID)
	_, err := c.client.Edge.EdgeDeleteV2(deleteParams, runtime.BearerToken(c.ctx.Token))
	if err == nil {
		return nil
	}
	apiErr, castOK := err.(*edge.EdgeDeleteV2Default)
	if castOK {
		return apiErr.Payload
	}
	return errorToAPIError(err)
}

func (c *Client) GetProject(id string) (*models.Project, *models.APIErrorPayload) {
	c.RefreshToken()
	params := project.NewProjectGetV2Params()
	params.ProjectID = id
	ok, err := c.client.Project.ProjectGetV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*project.ProjectGetV2Default).Payload
	}
	return ok.Payload, nil
}

func (c *Client) GetDataSrc(id string) (*models.DataSourceV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_source.NewDataSourceGetV2Params()
	params.ID = id
	ok, err := c.client.DataSource.DataSourceGetV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*data_source.DataSourceGetV2Default).Payload
	}
	return ok.Payload, nil
}

// GetDataPipeline fetches data pipeline for the given ID
func (c *Client) GetDataPipeline(id string) (*models.DataPipeline, *models.APIErrorPayload) {
	c.RefreshToken()
	params := data_pipeline.NewDataPipelineGetParams()
	params.ID = id
	ok, err := c.client.DataPipeline.DataPipelineGet(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*data_pipeline.DataPipelineGetDefault).Payload
	}
	return ok.Payload, nil
}

// GetFunction fetches function for the given ID
func (c *Client) GetFunction(id string) (*models.Function, *models.APIErrorPayload) {
	c.RefreshToken()
	params := function.NewFunctionGetParams()
	params.ID = id
	ok, err := c.client.Function.FunctionGet(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*function.FunctionGetDefault).Payload
	}
	return ok.Payload, nil
}

// GetCloudProfile fetches cloud profile for the given ID
func (c *Client) GetCloudProfile(id string) (*models.CloudProfile, *models.APIErrorPayload) {
	c.RefreshToken()
	params := cloud_profile.NewCloudProfileGetParams()
	params.ID = id
	ok, err := c.client.CloudProfile.CloudProfileGet(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*cloud_profile.CloudProfileGetDefault).Payload
	}
	return ok.Payload, nil
}

// GetEdge fetches edge for the given ID
func (c *Client) GetEdge(id string) (*models.EdgeV2, *models.APIErrorPayload) {
	c.RefreshToken()
	params := edge.NewEdgeGetV2Params()
	params.EdgeID = id
	ok, err := c.client.Edge.EdgeGetV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*edge.EdgeGetV2Default).Payload
	}
	return ok.Payload, nil
}

// GetCategory fetches category for the given ID
func (c *Client) GetCategory(id string) (*models.Category, *models.APIErrorPayload) {
	c.RefreshToken()
	params := category.NewCategoryGetV2Params()
	params.ID = id
	ok, err := c.client.Category.CategoryGetV2(params, runtime.BearerToken(c.ctx.Token))
	if err != nil {
		return nil, err.(*category.CategoryGetV2Default).Payload
	}
	return ok.Payload, nil
}
