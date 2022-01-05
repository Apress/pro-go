package actionresults

import "net/http"

func NewRedirectAction(url string) ActionResult {
    return &RedirectActionResult{ url: url}
}

type RedirectActionResult struct {
    url string
}

func (action *RedirectActionResult) Execute(ctx *ActionContext) error {
    ctx.ResponseWriter.Header().Set("Location", action.url)
    ctx.ResponseWriter.WriteHeader(http.StatusSeeOther)
    return nil
}
