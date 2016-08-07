package models

import "errors"

type Datastore interface {
	GetApp(appName string) (*App, error)
	GetApps(*AppFilter) ([]*App, error)
	StoreApp(*App) (*App, error)
	RemoveApp(appName string) error

	GetRoute(appName, routeName string) (*Route, error)
	GetRoutes(*RouteFilter) (routes []*Route, err error)
	StoreRoute(*Route) (*Route, error)
	RemoveRoute(appName, routeName string) error
}

var (
	ErrDatastoreEmptyAppName   = errors.New("Missing app name")
	ErrDatastoreEmptyRouteName = errors.New("Missing route name")
	ErrDatastoreEmptyApp       = errors.New("Missing app")
	ErrDatastoreEmptyRoute     = errors.New("Missing route")
)

func ApplyAppFilter(app *App, filter *AppFilter) bool {
	return true
}

func ApplyRouteFilter(route *Route, filter *RouteFilter) bool {
	return (filter.Path == "" || route.Path == filter.Path) &&
		(filter.AppName == "" || route.AppName == filter.AppName)
}