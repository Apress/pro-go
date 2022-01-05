package authorization 

import (
    "platform/authorization/identity"
    "platform/services"
    "platform/sessions"
    "context"
)

const USER_SESSION_KEY string = "USER"

func RegisterDefaultSignInService() {
    err := services.AddScoped(func(c context.Context) identity.SignInManager {
        return &SessionSignInMgr{ Context : c}
    })
    if (err != nil) {
        panic(err)
    }
}

type SessionSignInMgr struct {
    context.Context
}

func (mgr *SessionSignInMgr) SignIn(user identity.User) (err error) {
    session, err := mgr.getSession()
    if err == nil {
        session.SetValue(USER_SESSION_KEY, user.GetID())
    }
    return
}

func (mgr *SessionSignInMgr) SignOut(user identity.User) (err error) {
    session, err := mgr.getSession()
    if err == nil {
        session.SetValue(USER_SESSION_KEY, nil)
    }
    return
}

func (mgr *SessionSignInMgr) getSession() (s sessions.Session, err error) {
    err = services.GetServiceForContext(mgr.Context, &s)
    return
}
