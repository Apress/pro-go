package handling

import (
    "reflect"
    "regexp"
    "strings"
    "net/http"
)

type HandlerEntry struct {
    Prefix string
    Handler interface{}
}

type Route struct {
    httpMethod string
    prefix string
    handlerName string
    actionName string
    expression regexp.Regexp
    handlerMethod reflect.Method
}

var httpMethods = []string { http.MethodGet, http.MethodPost, 
    http.MethodDelete, http.MethodPut }

func generateRoutes(entries ...HandlerEntry) []Route {
    routes := make([]Route, 0, 10)
    for _, entry := range entries {
        handlerType := reflect.TypeOf(entry.Handler)
        promotedMethods := getAnonymousFieldMethods(handlerType)

        for i := 0; i < handlerType.NumMethod(); i++ {
            method := handlerType.Method(i)
            methodName := strings.ToUpper(method.Name)
            for _, httpMethod := range httpMethods {
                if strings.Index(methodName, httpMethod) == 0 {
                    if (matchesPromotedMethodName(method, promotedMethods)) {
                        continue
                    }
                    route := Route{
                        httpMethod: httpMethod,
                        prefix: entry.Prefix,
                        handlerName: strings.Split(handlerType.Name(), "Handler")[0],
                        actionName: strings.Split(methodName, httpMethod)[1],
                        handlerMethod: method,
                    }
                    generateRegularExpression(entry.Prefix, &route)                    
                    routes = append(routes, route)
                }
            }        
        }
    }
    return routes
}

func matchesPromotedMethodName(method reflect.Method, 
       methods []reflect.Method) bool {
    for _, m := range methods {
        if m.Name == method.Name {
            return true
        } 
    }
    return false
}

func getAnonymousFieldMethods(target reflect.Type) []reflect.Method {
    methods := []reflect.Method {}
    for i := 0; i < target.NumField(); i++ {
        field := target.Field(i)
        if (field.Anonymous && field.IsExported()) {
            for j := 0; j < field.Type.NumMethod(); j++ {
                method := field.Type.Method(j)
                if (method.IsExported()) {
                    methods = append(methods, method)
                }
            }
        }
    }
    return methods
}

func generateRegularExpression(prefix string, route *Route) {
    if (prefix != "" && !strings.HasSuffix(prefix, "/")) {
        prefix += "/"
    }
    pattern := "(?i)" + "/" + prefix + route.actionName
    if (route.httpMethod == http.MethodGet) {
        for i := 1; i < route.handlerMethod.Type.NumIn(); i++ {
            if route.handlerMethod.Type.In(i).Kind() == reflect.Int {
                pattern += "/([0-9]*)"
            } else {
                pattern += "/([A-z0-9]*)"
            }
        }
    }
    pattern = "^" + pattern + "[/]?$"
    route.expression = *regexp.MustCompile(pattern)        
}
