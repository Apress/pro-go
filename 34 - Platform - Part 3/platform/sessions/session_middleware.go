package sessions

import (
	"context"
    "time"
	"platform/config"
	"platform/pipeline"
	gorilla "github.com/gorilla/sessions"
)

type SessionComponent struct {
    store *gorilla.CookieStore
    config.Configuration
}

func (sc *SessionComponent) Init() {
    cookiekey, found := sc.Configuration.GetString("sessions:key")
    if !found {
        panic("Session key not found in configuration")
    }    
    if sc.GetBoolDefault("sessions:cyclekey", true) {
        cookiekey += time.Now().String()
    }
    sc.store = gorilla.NewCookieStore([]byte(cookiekey))    
}

func (sc *SessionComponent) ProcessRequest(ctx *pipeline.ComponentContext, 
        next func(*pipeline.ComponentContext)) {
    session, _ := sc.store.Get(ctx.Request, SESSION__CONTEXT_KEY)
    c := context.WithValue(ctx.Request.Context(), SESSION__CONTEXT_KEY, session)
    ctx.Request = ctx.Request.WithContext(c)    
    next(ctx)
    session.Save(ctx.Request, ctx.ResponseWriter)
}
