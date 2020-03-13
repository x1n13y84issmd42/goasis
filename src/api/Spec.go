package api

// Spec is an interface to access specification data.
type Spec interface {
	GetProjectInfo() *ProjectInfo
	GetOperation(name string) *Operation
	GetOperations() []Operation
	GetHost(name string) *Host
	GetDefaultHost() *Host
	GetSecurity(name string) *Security
	GetSchema(name string) *Schema
}

// Operation is an operation description.
type Operation struct {
	Name      string
	ID        string
	Path      string
	Method    string
	Headers   *HeaderBag
	Security  *Security
	Requests  *[]Request
	Responses *[]Response
}

// ProjectInfo is a generic project information.
type ProjectInfo struct {
	Title       string
	Description string
	Version     string
}

// ParameterLocation describes a parameter location in requests and security settings.
type ParameterLocation = string

// SecurityType is a type of security mechanism used.
type SecurityType = string

// SecurityScheme is specific exclusively to the http security type
// and describes the HTTP authentication mechanism used.
type SecurityScheme = string

// DataType is a schema data type, FFS.
type DataType = string

// DataFormat is an additional rule, specific for each type (string can be emails, as well as dates).
type DataFormat = string

// Types of security mechanisms
const (
	SecurityTypeHTTP   = SecurityType("http")
	SecurityTypeAPIKey = SecurityType("apiKey")
	SecurityTypeOAuth2 = SecurityType("oauth2")
	SecurityTypeOpenID = SecurityType("openIdConnect")
)

// Subtypes of security mechanisms
const (
	SecuritySchemeBasic       = SecurityScheme("basic")
	SecuritySchemeDigest      = SecurityScheme("digest")
	SecuritySchemeBearer      = SecurityScheme("bearer")
	SecuritySchemeHoba        = SecurityScheme("hoba")
	SecuritySchemeMutual      = SecurityScheme("mutual")
	SecuritySchemeNegotiate   = SecurityScheme("negotiate")
	SecuritySchemeOauth       = SecurityScheme("oauth")
	SecuritySchemeScramSHA1   = SecurityScheme("scram-sha-1")
	SecuritySchemeScramSHA256 = SecurityScheme("scram-sha-256")
	SecuritySchemeVapid       = SecurityScheme("vapid")
)

// Parameter locations
const (
	ParameterLocationQuery  = ParameterLocation("query")
	ParameterLocationHeader = ParameterLocation("header")
	ParameterLocationCookie = ParameterLocation("cookie")
)

// Data types
const (
	DataTypeString  = DataType("string")
	DataTypeObject  = DataType("object")
	DataTypeArray   = DataType("array")
	DataTypeBoolean = DataType("boolean")
	DataTypeNumber  = DataType("number")
	DataTypeInteger = DataType("integer")
)

// Data formats
const (
	DataFormatDate     = DataFormat("date")
	DataFormatDateTime = DataFormat("date-time")
)

type AuthFlow struct {
}

// Security is a description of a security mechanism used on some Operation.
type Security struct {
	Name           string
	SecurityType   SecurityType
	SecurityScheme SecurityScheme
	ParamName      string
	In             ParameterLocation
	Example        string
}

// Host is an API host desciption.
type Host struct {
	Name        string
	URL         string
	Description string
}

// Schema is a description of structured data used in requests, responses & security.
// The Properties property must be set if DataType is DataTypeObject.
// The Items property must be set if DataType is DataTypeArray.
type Schema struct {
	Name        string
	Description string
	DataType    DataType
	Properties  *[]Property
	Items       *Schema
	Example     string
}

// Property describes a property of an object being described by a Schema.
type Property struct {
	Name        string
	Description string
	Schema      *Schema
	Required    bool
}

type ParameterStyle string

const (
	ParameterStyleMatrix         = ParameterStyle("matrix")
	ParameterStyleLabel          = ParameterStyle("label")
	ParameterStyleForm           = ParameterStyle("form")
	ParameterStyleSimple         = ParameterStyle("simple")
	ParameterStyleSpaceDelimited = ParameterStyle("spaceDelimited")
	ParameterStylePipeDelimited  = ParameterStyle("pipeDelimited")
	ParameterStyleDeepObject     = ParameterStyle("deepObject")
)

// Header is a description of an HTTP header.
type Header struct {
	Name        string
	Schema      *Schema
	Description string
	Required    bool
	AllowEmpty  bool
	Explode     bool
	Style       ParameterStyle
	Example     string
	Value       string
}

// HeaderBag is a set of headers to used in requests and responses.
type HeaderBag = map[string][]Header

// Request describes a generic HTTP request.
type Request struct {
	ContentType string
	Headers     HeaderBag
	Schema      *Schema
}

// Response describes a generic HTTP response.
type Response struct {
	StatusCode  int
	ContentType string
	Headers     HeaderBag
	Schema      *Schema
	Example     string
}
