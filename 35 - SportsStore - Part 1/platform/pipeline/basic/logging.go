package basic

import (
    "net/http"
    "platform/logging"
    "platform/pipeline"
    //"platform/services"
)

type LoggingResponseWriter struct {
    statusCode int
    http.ResponseWriter
}

func (w *LoggingResponseWriter) WriteHeader(statusCode int) {
    w.statusCode = statusCode
    w.ResponseWriter.WriteHeader(statusCode)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
    if (w.statusCode == 0) {
        w.statusCode = http.StatusOK
    }
    return w.ResponseWriter.Write(b)
}

type LoggingComponent struct {}

func (lc *LoggingComponent) ImplementsProcessRequestWithServices() {}

func (lc *LoggingComponent) Init() {}

func (lc *LoggingComponent) ProcessRequestWithServices(
    ctx *pipeline.ComponentContext, 
    next func(*pipeline.ComponentContext), 
    logger logging.Logger)  {

    // var logger logging.Logger
    // err := services.GetServiceForContext(ctx.Request.Context(), &logger)
    // if (err != nil) {
    //     ctx.Error(err)
    //     return
    // }

    loggingWriter := LoggingResponseWriter{ 0, ctx.ResponseWriter}
    ctx.ResponseWriter = &loggingWriter

    logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
    next(ctx)
    logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL )
}
