package params

import (
	"net/http"
	"reflect"
)

func GetParametersFromRequest(request *http.Request, handlerMethod reflect.Method, 
        urlVals []string) (params []reflect.Value, err error) {
    handlerMethodType := handlerMethod.Type
    params = make([]reflect.Value, handlerMethodType.NumIn() -1)
    if (handlerMethodType.NumIn() == 1) {
        return []reflect.Value {}, nil
    } else if handlerMethodType.NumIn() == 2 && 
            handlerMethodType.In(1).Kind() == reflect.Struct {
        structVal := reflect.New(handlerMethodType.In(1))        
        err = request.ParseForm()
        if err == nil && getContentType(request) == "application/json" {
            err = populateStructFromJSON(structVal, request.Body)
        }        
        if err == nil {
            err = populateStructFromForm(structVal, request.Form)
        }
        return []reflect.Value { structVal.Elem() }, err
    } else {
        return getParametersFromURLValues(handlerMethodType, urlVals)
    } 
}

func getContentType(request *http.Request) (contentType string) {
    headerSlice := request.Header["Content-Type"]
    if headerSlice != nil && len(headerSlice) > 0 {
        contentType = headerSlice[0]
    }
    return
}
