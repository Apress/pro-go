package services

import (
    "reflect"
    "errors"
    "context"
)

func Populate(target interface{}) error {
    return PopulateForContext(context.Background(), target)
}

func PopulateForContext(c context.Context, target interface{}) (err error) {
    return PopulateForContextWithExtras(c, target, 
        make(map[reflect.Type]reflect.Value))
}

func PopulateForContextWithExtras(c context.Context, target interface{}, 
        extras map[reflect.Type]reflect.Value) (err error) {
    targetValue := reflect.ValueOf(target)
    if targetValue.Kind() == reflect.Ptr && 
            targetValue.Elem().Kind() == reflect.Struct {
        targetValue = targetValue.Elem()
        for i := 0; i < targetValue.Type().NumField(); i++ {
            fieldVal := targetValue.Field(i)
            if fieldVal.CanSet() {
                if extra, ok := extras[fieldVal.Type()]; ok {
                    fieldVal.Set(extra)
                } else {
                    resolveServiceFromValue(c,  fieldVal.Addr() )
                }
            }
            
        }
    } else {
        err = errors.New("Type cannot be used as target")
    }
    return
}
