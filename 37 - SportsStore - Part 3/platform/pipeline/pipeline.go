package pipeline

import (
    "net/http"
    "platform/services"
    "reflect"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) { /* do nothing */ }

func CreatePipeline(components ...interface{}) RequestPipeline {
    f := emptyPipeline
    for i := len(components) -1 ; i >= 0; i-- {
        currentComponent := components[i]
        services.Populate(currentComponent)
        nextFunc := f
        if servComp, ok := currentComponent.(ServicesMiddlwareComponent ); ok {
            f = createServiceDependentFunction(currentComponent, nextFunc)
            servComp.Init()
        } else if stdComp, ok := currentComponent.(MiddlewareComponent ); ok {
            f = func(context *ComponentContext) {
                if (context.error == nil) {
                    stdComp.ProcessRequest(context, nextFunc)
                }
            }
            stdComp.Init()
        } else {
            panic("Value is not a middleware component")
        }
    }
    return f
}

func createServiceDependentFunction(component interface{}, 
        nextFunc RequestPipeline) RequestPipeline {
    method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
    if (method.IsValid()) {
        return  func(context *ComponentContext) {
            if (context.error == nil) {
                _, err := services.CallForContext(context.Request.Context(), 
                    method.Interface(), context, nextFunc)
                if (err != nil) {
                    context.Error(err)
                }
            }
        }
    } else {
        panic("No ProcessRequestWithServices method defined")    
    }    
}

func (pl RequestPipeline) ProcessRequest(req *http.Request, 
        resp http.ResponseWriter) error {
    deferredWriter := &DeferredResponseWriter{ ResponseWriter:  resp } 
    ctx := ComponentContext {
        Request: req,
        ResponseWriter: deferredWriter,
    }
    pl(&ctx)
    if (ctx.error == nil) {
        deferredWriter.FlushData()
    }
    return ctx.error    
}
