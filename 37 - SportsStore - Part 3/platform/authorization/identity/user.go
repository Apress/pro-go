package identity

type User interface {

    GetID() int 

    GetDisplayName() string

    InRole(name string) bool

    IsAuthenticated() bool
}
