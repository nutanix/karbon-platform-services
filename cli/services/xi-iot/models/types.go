package models

const (
	EntityDataSource   = "DataSource"
	EntityDataPipeline = "DataPipeline"
	EntityApplication  = "Application"
	EntityEdge         = "Edge"
	EntityCloudProfile = "CloudProfile"
	EntityCategory     = "Category"
	EntityProject      = "Project"
	EntityFunction     = "Function"
)

// Context ..
type Context struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	URL      string `json:"url"`
	TenantID string `json:"tenantId"`
}

// Contexts ....
type Contexts struct {
	Contexts   map[string]Context `json:"contexts"`
	CurrentCtx string             `json:"currentCtx"`
}
