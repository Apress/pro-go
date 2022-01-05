package sessions

import (
    "context"
    "platform/services"
    gorilla "github.com/gorilla/sessions"
)

const SESSION__CONTEXT_KEY string = "pro_go_session"

func RegisterSessionService() {
    err := services.AddScoped(func(c context.Context) Session {
        val := c.Value(SESSION__CONTEXT_KEY)
        if s, ok := val.(*gorilla.Session); ok {
            return &SessionAdaptor{ gSession: s}
        } else {
            panic("Cannot get session from context ")
        }
    })
    if (err != nil) {
        panic(err)
    }
}

type Session interface {
    GetValue(key string) interface{}
    GetValueDefault(key string, defVal interface{}) interface{}
    SetValue(key string, val interface{})
}

type SessionAdaptor struct {
    gSession *gorilla.Session
}

func (adaptor *SessionAdaptor) GetValue(key string) interface{} {
    return adaptor.gSession.Values[key]
}

func (adaptor *SessionAdaptor) GetValueDefault(key string, 
        defVal interface{}) interface{} {
    if val, ok := adaptor.gSession.Values[key]; ok {
        return val
    }
    return defVal
}

func (adaptor *SessionAdaptor) SetValue(key string, val interface{}) {
    if val == nil {
        adaptor.gSession.Values[key] = nil
    } else {
        switch typedVal := val.(type) {
            case int, float64, bool, string:
                adaptor.gSession.Values[key] = typedVal
            default:
                panic("Sessions only support int, float64, bool, and string values")
        }
    }
}
