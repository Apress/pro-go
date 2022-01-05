package authorization

import (
    "net/http"
    "platform/authorization/identity"
    "platform/config"
    "platform/http/handling"
    "platform/pipeline"
    "strings"
    "regexp"
)

func NewAuthComponent(prefix string, condition identity.AuthorizationCondition,
        requestHandlers ...interface{}) *AuthMiddlewareComponent {    
    
    entries := []handling.HandlerEntry {}
    for _, handler := range requestHandlers {
        entries = append(entries, handling.HandlerEntry{prefix, handler})
    }
    
    router := handling.NewRouter(entries...)

    return &AuthMiddlewareComponent{
        prefix: "/" + prefix ,
        condition:  condition, 
        RequestPipeline: pipeline.CreatePipeline(router),
        fallbacks: map[*regexp.Regexp]string {},
    }
}

type AuthMiddlewareComponent struct {
    prefix string
    condition identity.AuthorizationCondition
    pipeline.RequestPipeline
    config.Configuration
    authFailURL string
    fallbacks map[*regexp.Regexp]string
}

func (c *AuthMiddlewareComponent) Init() {
        c.authFailURL, _ = c.Configuration.GetString("authorization:failUrl")
}

func (*AuthMiddlewareComponent) ImplementsProcessRequestWithServices() {}

func (c *AuthMiddlewareComponent) ProcessRequestWithServices(
        context *pipeline.ComponentContext, 
        next func(*pipeline.ComponentContext),
        user identity.User) {

    if strings.HasPrefix(context.Request.URL.Path, c.prefix) {
        for expr, target := range c.fallbacks {
            if expr.MatchString(context.Request.URL.Path) {
                http.Redirect(context.ResponseWriter, context.Request, 
                    target, http.StatusSeeOther)
                    return
            }
        }
        if c.condition.Validate(user) {
            c.RequestPipeline.ProcessRequest(context.Request, context.ResponseWriter)
        } else {
            if c.authFailURL != "" {
                http.Redirect(context.ResponseWriter, context.Request, 
                    c.authFailURL, http.StatusSeeOther)
            } else if user.IsAuthenticated() {
                context.ResponseWriter.WriteHeader(http.StatusForbidden)
            } else {
                context.ResponseWriter.WriteHeader(http.StatusUnauthorized)
            }
        }
    } else {
        next(context)
    }
}

func (c *AuthMiddlewareComponent) AddFallback(target string, 
        patterns ...string) *AuthMiddlewareComponent {
    for _, p := range patterns {
        c.fallbacks[regexp.MustCompile(p)] = target
    }
    return c
}
