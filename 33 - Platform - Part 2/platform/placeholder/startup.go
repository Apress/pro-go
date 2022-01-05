package placeholder

import (
    "platform/http"
    "platform/pipeline"
    "platform/pipeline/basic"
    "platform/services"
    "sync"
    "platform/http/handling"
)

func createPipeline() pipeline.RequestPipeline {
    return pipeline.CreatePipeline(
        &basic.ServicesComponent{},
        &basic.LoggingComponent{},
        &basic.ErrorComponent{},
        &basic.StaticFileComponent{},
        //&SimpleMessageComponent{},
        handling.NewRouter(
            handling.HandlerEntry{ "",  NameHandler{}},
        ),        
    )
}

func Start() {
    results, err := services.Call(http.Serve, createPipeline())
    if (err == nil) {
        (results[0].(*sync.WaitGroup)).Wait()
    } else {
        panic(err)
    }
}
