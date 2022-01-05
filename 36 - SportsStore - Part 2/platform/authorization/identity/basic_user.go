package identity

import "strings"

var UnauthenticatedUser User = &basicUser{}

func NewBasicUser(id int, name string, roles ...string) User {
    return &basicUser {
        Id: id,
        Name: name,
        Roles: roles,
        Authenticated: true,
    }
}

type basicUser struct {
    Id int
    Name string
    Roles []string
    Authenticated bool
}

func (user *basicUser) GetID() int {
    return user.Id
}

func (user *basicUser) GetDisplayName() string {
    return user.Name
}

func (user *basicUser) InRole(role string) bool {
    for _, r := range user.Roles {
        if strings.EqualFold(r, role) {
            return true
        }
    }
    return false
}

func (user *basicUser) IsAuthenticated() bool {
    return user.Authenticated
}
