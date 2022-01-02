package router

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	path   Path
	routes map[Path]*Route
}

func New(stringPath string) (*Router, error) {
	path := Path(stringPath)
	if isValid, err := path.IsRouteValid(); !isValid {
		return nil, err
	}

	return &Router{
		path:   path,
		routes: map[Path]*Route{},
	}, nil
}

func (r *Router) AddHandlersIntoServeMux(serve *http.ServeMux) {
	for _, route := range r.routes {
		route.insertHandlers(serve)
	}
}

func (r *Router) All(stringPath string, handler Handler) error {
	return r.addRouteHandler(stringPath, handler, Method(ALL))
}

func (r *Router) Get(stringPath string, handler Handler) error {
	return r.addRouteHandler(stringPath, handler, Method(GET))
}

func (r *Router) Post(stringPath string, handler Handler) error {
	return r.addRouteHandler(stringPath, handler, Method(POST))
}

func (r *Router) Put(stringPath string, handler Handler) error {
	return r.addRouteHandler(stringPath, handler, Method(PUT))
}

func (r *Router) Delete(stringPath string, handler Handler) error {
	return r.addRouteHandler(stringPath, handler, Method(DELETE))
}

func (r *Router) Use(stringPath string, router *Router) error {
	path := Path(stringPath)
	if isValid, err := path.IsRouteValid(); !isValid {
		return err
	}
	rootPath := r.path.Merge(path.Merge(router.path))

	for _, route := range router.routes {
		mergedPath := rootPath.Merge(route.path)
		newRoute := r.GetRoute(mergedPath)
		for method, handler := range route.methods {
			newRoute.methods[method] = handler
		}
		for dynamicPath, methodHandler := range route.dynamicPaths {
			dynamicMergedPath := rootPath.Merge(dynamicPath)
			newRoute.dynamicPaths[dynamicMergedPath] = methodHandler
		}
	}

	return nil
}

func (r *Router) GetRoute(path Path) *Route {
	if route, hasRoute := r.routes[path]; hasRoute {
		return route
	}

	r.routes[path] = &Route{
		path:         path,
		methods:      map[Method]Handler{},
		dynamicPaths: map[Path]map[Method]Handler{},
	}
	return r.routes[path]
}

func (r *Router) addRouteHandler(stringPath string, handler Handler, method Method) error {
	path := Path(stringPath)
	if valid, err := path.IsValid(); err != nil || !valid {
		return err
	}

	if path.IsRoot() {
		route := r.GetRoute(path)
		route.addHandler(path, handler, method)
	}

	if path.IsParam() {
		route := r.GetRoute(Path("/"))
		fullPath := r.path.Merge(path)
		err := route.addDynamicPathHandler(fullPath, handler, method)
		if err != nil {
			return err
		}
	}

	if isValid, _ := path.IsRouteValid(); isValid {
		fullPath := r.path.Merge(path)
		route := r.GetRoute(fullPath)
		route.addHandler(Path("/"), handler, method)
	} else {
		rootStringPath := strings.Split(stringPath, "/:")
		rootPath := Path(rootStringPath[0] + "/")
		if valid, _ := rootPath.IsRouteValid(); !valid {
			log.Fatal("something went wrong with the dynamic path rework")
		}

		fullPath := r.path.Merge(rootPath)
		route := r.GetRoute(fullPath)
		err := route.addDynamicPathHandler(r.path.Merge(path), handler, method)
		if err != nil {
			return err
		}
	}

	return nil
}
