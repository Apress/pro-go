package placeholder

import (
    "platform/http/actionresults"
    "platform/authorization/identity"
    "fmt"
)

type AuthenticationHandler struct {
    identity.User
    identity.SignInManager
    identity.UserStore
}

func (h AuthenticationHandler) GetSignIn() actionresults.ActionResult {
    return actionresults.NewTemplateAction("signin.html", 
    fmt.Sprintf("Signed in as: %v", h.User.GetDisplayName()))
}

type Credentials struct {
    Username string
    Password string
}

func (h AuthenticationHandler) PostSignIn(creds Credentials) actionresults.ActionResult {
    if creds.Password == "mysecret" {
        user, ok := h.UserStore.GetUserByName(creds.Username)
        if (ok) {
            h.SignInManager.SignIn(user)
            return actionresults.NewTemplateAction("signin.html", 
                fmt.Sprintf("Signed in as: %v", user.GetDisplayName()))
        }
    } 
    return actionresults.NewTemplateAction("signin.html", "Access Denied")
}


func (h AuthenticationHandler) PostSignOut() actionresults.ActionResult {
    h.SignInManager.SignOut(h.User)
    return actionresults.NewTemplateAction("signin.html", "Signed out")
}
