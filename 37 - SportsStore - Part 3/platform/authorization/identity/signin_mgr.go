package identity

type SignInManager interface {

    SignIn(user User) error
    SignOut(user User) error
}
