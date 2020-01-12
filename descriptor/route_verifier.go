package descriptor

import "net/url"

type RouteVerifier struct {
	route *Route
	logWriter routeError
}

func newRouteVerifier(r *Route) RouteVerifier {

}

func (r *RouteVerifier) checkUri() {
	_, err := url.ParseRequestURI(r.route.Uri)
	if err != nil {

	}
}


func (r *RouteVerifier) validate() error {

	for _, m := range r.route.Methods {

	}

}
