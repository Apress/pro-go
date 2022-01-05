package services

import (
    "reflect"
    "sync"
)

func AddTransient(factoryFunc interface{}) (err error) {
    return addService(Transient, factoryFunc)
}

func AddScoped(factoryFunc interface{}) (err error) {
    return addService(Scoped, factoryFunc)
}

func AddSingleton(factoryFunc interface{}) (err error) {
    factoryFuncVal := reflect.ValueOf(factoryFunc)
    if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
        var results []reflect.Value
        once := sync.Once{}
        wrapper := reflect.MakeFunc(factoryFuncVal.Type(), 
            func ([]reflect.Value) []reflect.Value {
                once.Do(func() { 
                    results = invokeFunction(nil, factoryFuncVal)
                })
                return results
            })
        err = addService(Singleton, wrapper.Interface())
    }   
    return 
}
