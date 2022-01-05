package http

import (
    "fmt"
    "sync"
    "net/http"
    "platform/config"
    "platform/logging"
    "platform/pipeline"
)

type pipelineAdaptor struct {
    pipeline.RequestPipeline
}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter, 
    request *http.Request) {
        p.ProcessRequest(request, writer)
}

func Serve(pl pipeline.RequestPipeline, cfg config.Configuration, logger logging.Logger ) *sync.WaitGroup {
    wg := sync.WaitGroup{}

    adaptor := pipelineAdaptor { RequestPipeline: pl }

    enableHttp := cfg.GetBoolDefault("http:enableHttp", true)
    if (enableHttp) {
        httpPort := cfg.GetIntDefault("http:port", 5000)
        logger.Debugf("Starting HTTP server on port %v", httpPort)
        wg.Add(1)
        go func() {
            err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
            if (err != nil) {
                panic(err)
            }
        }()
    }
    enableHttps := cfg.GetBoolDefault("http:enableHttps", false)
    if (enableHttps) {
        httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
        certFile, cfok := cfg.GetString("http:httpsCert")
        keyFile, kfok := cfg.GetString("http:httpsKey")
        if cfok && kfok {
            logger.Debugf("Starting HTTPS server on port %v", httpsPort)
            wg.Add(1)
            go func() {
                err := http.ListenAndServeTLS(fmt.Sprintf(":%v", httpsPort),
                    certFile, keyFile, adaptor)
                if (err != nil) {
                    panic(err)
                }
            }()
        } else {
            panic("HTTPS certificate settings not found")
        }
    }
    return &wg
}
