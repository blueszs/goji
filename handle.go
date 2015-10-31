package goji

import "net/http"

/*
Handle adds a new route to the Mux. Requests that match the given Pattern will
be dispatched to the given http.Handler. If the http.Handler also supports
Handler, that interface will be used instead.

Routing is performed in the order in which routes are added: the first route
with a matching Pattern will be used. In particular, Goji guarantees that
routing is performed in a manner that is indistinguishable from the following
algorithm:

	// Assume routes is a slice that every call to Handle appends to
	for route := range routes {
		// For performance, Patterns can opt out of this call to Match.
		// See the documentation for Pattern for more.
		if ctx2 := route.pattern.Match(ctx, r); ctx2 != nil {
			route.handler.ServeHTTPC(ctx2, w, r)
			break
		}
	}

It is not safe to concurrently register routes from multiple goroutines.
*/
func (m *Mux) Handle(p Pattern, h http.Handler) {
}

/*
HandleC adds a context-aware handler to the Mux. See the documentation for
Handle for more information about the semantics of routing.

It is not safe to concurrently register routes from multiple goroutines.
*/
func (m *Mux) HandleC(p Pattern, h Handler) {
}
