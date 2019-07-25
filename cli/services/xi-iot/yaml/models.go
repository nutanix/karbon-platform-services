package yaml

// Base ...
type Base struct {
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
}

type Field struct {
	Name  string `yaml:"name"`
	Topic string `yaml:"topic"`
}

type Selector struct {
	CategoryName  string   `yaml:"categoryName"`
	CategoryValue string   `yaml:"categoryValue"`
	Scope         []string `yaml:"scope"`
}

type CatInfo struct {
	Name  string `yaml:"categoryName"`
	Value string `yaml:"categoryValue"`
}

type DataIfcEndpoint struct {
	DataSource string `yaml:"dataSource"`
	TopicName  string `yaml:"topicName"`
	FieldName  string `yaml:"fieldName"`
}

type Port struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type IfcInfo struct {
	Class    string `yaml:"class"`
	Img      string `yaml:"img"`
	Kind     string `yaml:"kind"`
	Protocol string `yaml:"protocol"`
	Ports    []Port `yaml:"ports"`
}

type Category struct {
	Kind        string   `yaml:"kind"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description,omitempty"`
	Values      []string `yaml:"values"`
}

type DataSource struct {
	Kind      string     `yaml:"kind"`
	Name      string     `yaml:"name"`
	Protocol  string     `yaml:"protocol"`
	Edge      string     `yaml:"edge"`
	Fields    []Field    `yaml:"fields"`
	Selectors []Selector `yaml:"selectors"`
	Type      string     `yaml:"type,omitempty"`
	AuthType  string     `yaml:"authType"`
	IfcInfo   *IfcInfo   `yaml:"ifcInfo,omitempty"`
}

// FunctionParam abstracts the params for the function object
type FunctionParam struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

type Function struct {
	Name string            `yaml:"name"`
	Args map[string]string `yaml:"args"`
}

type DataPipeline struct {
	Kind                 string              `yaml:"kind"`
	Name                 string              `yaml:"name"`
	Project              string              `yaml:"project"`
	Description          string              `yaml:"description,omitempty"`
	SamplingIntervalMSec int                 `yaml:"samplingIntervalMsec"`
	Input                *DataPipelineInput  `yaml:"input"`
	Functions            []Function          `yaml:"functions,omitempty"`
	Output               *DataPipelineOutput `yaml:"output"`
}

type Application struct {
	Kind             string            `yaml:"kind"`
	Name             string            `yaml:"name"`
	Description      string            `yaml:"description,omitempty"`
	Project          string            `yaml:"project"`
	DataIfcEndpoints []DataIfcEndpoint `yaml:"dataIfcEndpoints"`
	AppYamlPath      string            `yaml:"appYamlPath,omitempty"`
	AppYaml          string            `yaml:"appYaml"`
	EdgeSelectors    []CatInfo         `yaml:"edgeSelectors,omitempty"`
	Edges            []string          `yaml:"edges,omitempty"`
}

// DataPipelineInput defines the input configuration for a pipeline
type DataPipelineInput struct {
	// TODO: convert map[string][]string to type alias
	CatSels      map[string][]string `yaml:"categorySelectors,omitempty"`
	DataPipeline *string             `yaml:"dataPipeline,omitempty"` // Name of the pipeline to consume from.
}

// RealTimeStream defines the input type real time stream for a pipeline
type RealTimeStream struct {
	DataPipeline string `yaml:"dataPipeline"` // Name of the pipeline to consume from.
}

// DataPipelineOutput defines the output configuration for a pipeline
type DataPipelineOutput struct {
	PublicCloud *PublicCloud `yaml:"publicCloud,omitempty"`
	LocalEdge   *LocalEdge   `yaml:"localEdge,omitempty"`
}

// PublicCloud defines the public cloud output configuration for a pipeline
type PublicCloud struct {
	Type         string `yaml:"type"`         // type of cloud
	Service      string `yaml:"service"`      // service name, example: kafka, kinesis
	Region       string `yaml:"region"`       // example: us-west2, us-east1
	Profile      string `yaml:"profile"`      // cloud profile. For example: may include AWS secret/access keys
	EndpointName string `yaml:"endpointName"` // identifier for the cloud service. example: s3 bucket name, sqs queue name, kinesis stream name
}

// LocalEdge defines the local edge output configuration for a pipeline
type LocalEdge struct {
	Type         string `yaml:"type"`         // type of edge service. Possible options: DataInterface,DataService
	Service      string `yaml:"service"`      // name of the service. example: mqtt, kafka, hls
	EndpointName string `yaml:"endpointName"` // identifier for the edge service. example: kafka topic name, hls stream name, etc
}

// TODO: Change thisb to `Param`
type Params struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// TODO: Change this to Function to avoid confusion
type Transformation struct {
	Name           string   `yaml:"name"`
	Kind           string   `yaml:"kind"`
	Project        string   `yaml:"project"`
	Description    string   `yaml:"description,omitempty"`
	Language       string   `yaml:"language"`
	Environment    string   `yaml:"environment"`
	SourceCodePath string   `yaml:"sourceCodePath,omitempty"`
	Params         []Params `yaml:"params"`
}

// Project object
type Project struct {
	Kind          string   `yaml:"kind"`
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description,omitempty"`
	Edges         []string `yaml:"edges,omitempty"`
	CloudProfiles []string `yaml:"cloudProfiles,omitempty"`
}

// Edge encapsulates edge  metadata
type Edge struct {
	Kind         string              `yaml:"kind"`
	Name         string              `yaml:"name"`
	Description  string              `yaml:"description,omitempty"`
	IPAddress    string              `yaml:"ipAddress"`
	Connected    bool                `yaml:"connected"`
	SerialNumber string              `yaml:"serialNumber"`
	CatSels      map[string][]string `yaml:"categorySelectors,omitempty"`
}
