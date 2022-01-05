package handling

import (
    "platform/http/actionresults"
    "platform/services"
    "net/http"
    "reflect"
    "regexp"
    "fmt"
)

func (rc *RouterComponent) AddMethodAlias(srcUrl string, 
        method interface{}, data ...interface{}) *RouterComponent {
    var urlgen URLGenerator
    services.GetService(&urlgen)
    url, err := urlgen.GenerateUrl(method, data...)
    if (err == nil) {
        return rc.AddUrlAlias(srcUrl, url)
    } else {
        panic(err)
    }
}

func (rc *RouterComponent) AddUrlAlias(srcUrl string, 
        targetUrl string) *RouterComponent {
    aliasFunc := func(interface{}) actionresults.ActionResult {
        return actionresults.NewRedirectAction(targetUrl)
    }        
    alias := Route {
        httpMethod: http.MethodGet,
        handlerName: "Alias",
        actionName: "Redirect",
        expression: *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcUrl)),
        handlerMethod: reflect.Method{
            Type: reflect.TypeOf(aliasFunc),
            Func: reflect.ValueOf(aliasFunc),
        },
    }
    rc.routes = append([]Route { alias},  rc.routes... )
    return rc
}
