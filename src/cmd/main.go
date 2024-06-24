package main

import (
	"errors"
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
)

func main() {
	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	newManager := manager.NewManager(render)
	newRouter := router.NewRouter(newManager)
	newRouter.Get("/home", func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
		w.Write([]byte("ok"))
		return func() {}
	})
	serv := server.NewServer(":8000", newRouter)
	err = serv.Start()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}
