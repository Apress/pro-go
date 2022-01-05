package params

import (
    "reflect"
    "errors"
)

func getParametersFromURLValues(funcType reflect.Type, 
        urlVals []string) (params []reflect.Value, err error) {
    if (len(urlVals) == funcType.NumIn() -1) {
        params = make([]reflect.Value, funcType.NumIn() -1)
        for i := 0; i < len(urlVals); i++ {
            params[i], err = parseValueToType(funcType.In(i + 1), urlVals[i])
            if (err != nil) {
                return
            } 
        }
    } else {
        err = errors.New("Parameter number mismatch")
    }
    return
}
