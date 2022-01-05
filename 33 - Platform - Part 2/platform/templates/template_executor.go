package templates

import "io"

type TemplateExecutor interface {

    ExecTemplate(writer io.Writer, name string, data interface{}) (err error) 
}
