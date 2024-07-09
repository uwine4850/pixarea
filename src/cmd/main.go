package main

import (
	"errors"
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/bglobalflow"
	"github.com/uwine4850/foozy/pkg/builtin/builtin_mddl"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
	"github.com/uwine4850/foozy/pkg/server/globalflow"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/handlers"
	"github.com/uwine4850/pixarea/src/handlers/hauth"
	"github.com/uwine4850/pixarea/src/middlewares/authmddl"
)

func main() {
	db := connectToDb()
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	mddl := middlewares.NewMiddleware()
	mddl.HandlerMddl(0, authmddl.UpdKeys(db))
	mddl.HandlerMddl(1, builtin_mddl.GenerateAndSetCsrf)
	mddl.HandlerMddl(2, authmddl.AuthPermissions)

	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}

	newManager := manager.NewManager(render)
	newManager.Config().Debug(true)
	newManager.Config().PrintLog(true)
	newManager.Config().Generate32BytesKeys()

	newRouter := router.NewRouter(newManager)
	newRouter.SetMiddleware(mddl)

	newRouter.GetMux().Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))
	newRouter.Get("/explore", handlers.ExploreHNDL)
	newRouter.Get("/profile", handlers.ProfileHNDL)
	newRouter.Get("/profile/edit", handlers.ProfileEditHNDL)
	newRouter.Get("/logout", hauth.LogOutHNDL)
	newRouter.Get("/login", hauth.LoginHNDL)
	newRouter.Post("/login-post", hauth.LoginPostHNDL)
	newRouter.Get("/register", hauth.RegisterHNDL)
	newRouter.Post("/register-post", hauth.RegisterPostHNDL)
	newRouter.Get("/publication", handlers.PublicationViewHNDL)
	newRouter.Get("/new-publication", handlers.NewPublicationHNDL)

	gf := globalflow.NewGlobalFlow(10)
	gf.AddNotWaitTask(bglobalflow.KeyUpdater(3600))
	gf.Run(newManager)

	serv := server.NewServer(":8000", newRouter)
	err = serv.Start()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}

func connectToDb() *database.Database {
	db := database.NewDatabase(cnf.DB_ARGS)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	return db
}
