package router

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

type Route struct {
	path         Path
	methods      map[Method]Handler
	dynamicPaths map[Path]map[Method]Handler
}

func (r *Route) addHandler(path Path, handler Handler, method Method) error {
	if _, hasHandler := r.methods[method]; hasHandler {
		return fmt.Errorf("method already defined for path %s", r.path.String())
	}

	r.methods[method] = handler
	return nil
}

func (r *Route) addDynamicPathHandler(path Path, handler Handler, method Method) error {
	if _, hasDynamicRoute := r.dynamicPaths[path]; !hasDynamicRoute {
		r.dynamicPaths[path] = map[Method]Handler{
			method: handler,
		}
		return nil
	}

	dynamicRoute := r.dynamicPaths[path]
	if _, hasHandler := dynamicRoute[method]; hasHandler {
		return fmt.Errorf("method already defined for path %s", path.String())
	}

	dynamicRoute[method] = handler
	return nil
}

type DynamicRoutePattern struct {
	params     []string
	pathToSort string
	pattern    *regexp.Regexp
	handler    map[Method]Handler
}

func getMethod(stringMethod string) (Method, error) {
	for i, m := range availableMethods {
		if m == stringMethod {
			return Method(i), nil
		}
	}

	return -1, errors.New("no such method")
}

func (r *Route) insertHandlers(serve *http.ServeMux) {
	dynamicPathsSize := len(r.dynamicPaths)

	if dynamicPathsSize > 0 {
		dynamicPathsPatternList := make([]DynamicRoutePattern, len(r.dynamicPaths))

		reg := regexp.MustCompile("(:[^/]+)")
		i := 0
		for path, handler := range r.dynamicPaths {
			matchList := regexp.MustCompile("(:[^/]+)").FindAllString(path.String(), 11)
			matchSize := len(matchList)
			var params = make([]string, matchSize)
			for j := 0; j < matchSize; j++ {
				params[j] = strings.ReplaceAll(matchList[j], ":", "")
			}

			dynamicPathsPatternList[i] = DynamicRoutePattern{
				params:     params,
				pathToSort: reg.ReplaceAllString(path.String(), "|"),
				pattern:    regexp.MustCompile("^" + reg.ReplaceAllString(path.String(), "([^/]+)") + "$"),
				handler:    handler,
			}
			i++
		}
		sort.Slice(dynamicPathsPatternList, func(i, j int) bool {
			return dynamicPathsPatternList[i].pathToSort > dynamicPathsPatternList[j].pathToSort
		})

		serve.HandleFunc(r.path.String(), func(response http.ResponseWriter, request *http.Request) {
			method, err := getMethod(request.Method)

			if err != nil {
				response.WriteHeader(http.StatusMethodNotAllowed)
				response.Write([]byte("Method not allowed"))
				return
			}

			for i = 0; i < len(dynamicPathsPatternList); i++ {
				match := dynamicPathsPatternList[i].pattern.FindStringSubmatch(request.URL.Path)
				matchSize := len(match)

				if matchSize > 1 {
					params := make(map[string]string)
					for j := 1; j < matchSize; j++ {
						params[dynamicPathsPatternList[i].params[j-1]] = match[j]
					}

					if _, hasHandler := dynamicPathsPatternList[i].handler[method]; !hasHandler {
						response.WriteHeader(http.StatusMethodNotAllowed)
						response.Write([]byte("Method not allowed"))
						return
					}

					handler := dynamicPathsPatternList[i].handler[method]
					handler(response, request, &Context{
						Params: params,
					})
					return
				}
			}
			if _, hasHandler := r.methods[method]; !hasHandler {
				response.WriteHeader(http.StatusMethodNotAllowed)
				response.Write([]byte("Method not allowed"))
				return
			}

			handler := r.methods[method]
			handler(response, request, &Context{})
		})

		return
	}

	serve.HandleFunc(r.path.String(), func(response http.ResponseWriter, request *http.Request) {
		method, err := getMethod(request.Method)
		if err != nil {
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("Method not allowed"))
			return
		}
		if _, hasHandler := r.methods[method]; !hasHandler {
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("Method not allowed"))
			return
		}

		handler := r.methods[method]
		handler(response, request, &Context{})
	})
}
