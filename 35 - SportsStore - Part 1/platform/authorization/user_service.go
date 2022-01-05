package authorization

import (
    "platform/services"
    "platform/sessions"
    "platform/authorization/identity"
)

func RegisterDefaultUserService() {
    err := services.AddScoped(func(session sessions.Session, 
            store identity.UserStore) identity.User {
        userID, found := session.GetValue(USER_SESSION_KEY).(int)
        if found {
            user, userFound := store.GetUserByID(userID)
            if (userFound) {
                return user
            }
        }
        return identity.UnauthenticatedUser
    })
    if (err != nil) {
        panic(err)
    }
}
