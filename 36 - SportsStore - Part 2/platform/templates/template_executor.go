package templates

import "io"

type TemplateExecutor interface {

    ExecTemplate(writer io.Writer, name string, data interface{}) (err error) 

    ExecTemplateWithFunc(writer io.Writer, name string, 
        data interface{}, handlerFunc InvokeHandlerFunc) (err error) 
}

type InvokeHandlerFunc func(handlerName string, methodName string, 
    args ...interface{}) interface{}
