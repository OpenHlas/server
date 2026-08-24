package app

import (
	router "github.com/gouef/router"
	"github.com/gouef/web-project/controllers"
)

func RouterFactory() *router.RouteList {
	app := controllers.DefaultController{}
	node := controllers.NodeController{}
	l := router.NewRouteList()
	v1 := router.NewRouteList()
	l.Add("home", "/", app.Index, router.Get)
	l.Add("ping", "/ping", app.Ping, router.Get)
	l.Add("users:detail", "/users/:id", app.Index, router.Get)
	v1.Add("api:health", "/api/v1/health", node.Health, router.Get)
	v1.Add("api:info", "/api/v1/info", node.Info, router.Get)
	l.AddChild(v1)
	l.Add("websocket", "/ws", node.WebSocket, router.Get)

	return l
}
