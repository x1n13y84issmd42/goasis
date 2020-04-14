package log

import (
	"fmt"

	"github.com/x1n13y84issmd42/oasis/src/api"
	"github.com/xeipuuv/gojsonschema"
)

// ILogger - interface for execution loggers
type ILogger interface {
	NOMESSAGE(msg string, args ...interface{})

	Usage()
	Error(err error)
	LoadingSpec(path string)

	PrintOperations(ops []*api.Operation)
	TestingProject(p *api.ProjectInfo)
	TestingOperation(res *api.Operation)

	UsingHost(p *api.Host)
	UsingDefaultHost()
	HostNotFound(h string)

	UsingSecurity(sec api.ISecurity)
	SecurityHasNoData(sec api.ISecurity)

	Requesting(method string, url string)

	ParameterHasNoExample(paramName string, in string, container string)
	UsingParameterExample(paramName string, in string, container string)

	HeaderHasNoValue(header *api.Header)
	ResponseHasWrongStatus(schema *api.Response, actualStatus int)
	ResponseHasWrongContentType(schema *api.Response, actualCT string)

	OperationOK(res *api.Operation)
	OperationFail(res *api.Operation)
	OperationNotFound(op string)

	SchemaTesting(schema *api.Schema, data interface{})
	SchemaOK(schema *api.Schema)
	SchemaFail(schema *api.Schema, errors []gojsonschema.ResultError)
}

// Log is a base type for loggers.
type Log struct {
	Level int64
}

// Print prints.
func (log Log) Print(l int64, msg string, args ...interface{}) {
	if l <= log.Level {
		fmt.Printf(msg, args...)
	}
}

// Println prints and adds a newline.
func (log Log) Println(l int64, msg string, args ...interface{}) {
	log.Print(l, msg+"\n", args...)
}

// NOMESSAGE is a default and temporary print function to use when you don't have
// a dedidated message function in the logger. It is meant to be replaced eventually
// with a proper logging method or removed. Henec the indicating name.
func (log Log) NOMESSAGE(msg string, args ...interface{}) {
	log.Println(1, "\t"+msg, args...)
}

// New creates a new logger based on the provided log tyle & leve.
func New(style string, level int64) ILogger {
	switch style {
	case "plain":
		return NewPlain(level)

	case "festive":
		return NewFestive(level)
	}

	fmt.Printf("The \"%s\" log style is unknown.\nAvailable loggers are:\n", style)
	fmt.Println("\tplain - a simple text logger")
	fmt.Println("\tfestive - a nicer colorized logger")

	panic("No way.")
}
