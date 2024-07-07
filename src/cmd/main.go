package main

import (
	"errors"
	"net/http"

	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
	"github.com/uwine4850/pixarea/src/handlers"
)

func main() {
	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	newManager := manager.NewManager(render)
	newManager.Config().Debug(true)
	newManager.Config().PrintLog(true)
	newRouter := router.NewRouter(newManager)
	newRouter.GetMux().Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))
	newRouter.Get("/explore", handlers.ExploreHNDL)
	newRouter.Get("/profile", handlers.ProfileHNDL)
	newRouter.Get("/profile/edit", handlers.ProfileEditHNDL)
	newRouter.Get("/login", handlers.LoginHNDL)
	newRouter.Get("/register", handlers.RegisterHNDL)
	newRouter.Get("/publication", handlers.PublicationViewHNDL)
	newRouter.Get("/new-publication", handlers.NewPublicationHNDL)
	serv := server.NewServer(":8000", newRouter)
	err = serv.Start()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}
