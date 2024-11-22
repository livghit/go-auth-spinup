package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateMiddlewareStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := range middlewares {
			x := middlewares[i]
			next = x(next)
		}
		return next
	}
}
