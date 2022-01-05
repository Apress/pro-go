package identity

type AuthorizationCondition interface {

    Validate(user User) bool
}
