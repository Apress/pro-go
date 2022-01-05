package handling

import (
    "fmt"
    "net/http"
    "strings"
    "errors"
    "reflect"
)

type URLGenerator interface {

    GenerateUrl(method interface{}, data ...interface{}) (string, error)

    GenerateURLByName(handlerName, methodName string, 
        data ...interface{}) (string, error)

    AddRoutes(routes []Route)
}

type routeUrlGenerator struct {
    routes []Route
}

func (gen *routeUrlGenerator) AddRoutes(routes []Route) {
    if gen.routes == nil {
        gen.routes = routes
    } else {
        gen.routes = append(gen.routes, routes...)
    }
}

func (gen *routeUrlGenerator) GenerateUrl(method interface{}, 
        data ...interface{}) (string, error) {
    methodVal := reflect.ValueOf(method)
    if methodVal.Kind() == reflect.Func && 
            methodVal.Type().In(0).Kind() == reflect.Struct {
        for _, route := range gen.routes {
            if route.handlerMethod.Func.Pointer() == methodVal.Pointer() {
                return generateUrl(route, data...)
            } 
        }
    } 
    return "", errors.New("No matching route")
}

func (gen *routeUrlGenerator) GenerateURLByName(handlerName, methodName string, 
        data ...interface{}) (string, error) {
    for _, route := range gen.routes {
        if strings.EqualFold(route.handlerName, handlerName) &&
                strings.EqualFold(route.httpMethod + route.actionName, methodName) {
            return generateUrl(route, data...)
        }
    }
    return "", errors.New("No matching route")
}

func generateUrl(route Route, data ...interface{}) (url string, err error) {
    url = "/" + route.prefix
    if (!strings.HasPrefix(url, "/")) {
        url = "/" + url
    }
    if (!strings.HasSuffix(url, "/")) {
        url += "/"
    }
    url+= strings.ToLower(route.actionName)
    if len(data) > 0 && !strings.EqualFold(route.httpMethod, http.MethodGet) {
        err = errors.New("Only GET handler can have data values")                
    } else if strings.EqualFold(route.httpMethod, http.MethodGet) &&
            len(data) != route.handlerMethod.Type.NumIn() -1 {
        err = errors.New("Number of data values doesn't match method params")
    } else {
        for _, val := range data {
            url = fmt.Sprintf("%v/%v", url, val)
        }
    }
    return
}
