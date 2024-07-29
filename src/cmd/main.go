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
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
	"github.com/uwine4850/pixarea/src/handlers/hpublication"
	"github.com/uwine4850/pixarea/src/handlers/tmplfilters"
	"github.com/uwine4850/pixarea/src/middlewares/authmddl"
	"github.com/uwine4850/pixarea/src/middlewares/usermddl"
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
	mddl.AsyncHandlerMddl(usermddl.ParseUserCookies)

	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	tmplfilters.RegisterFilters()

	newManager := manager.NewManager(render)
	newManager.Config().DebugConfig().ErrorLoggingFile("src/logs/errors.log")
	newManager.Config().DebugConfig().ErrorLogging(true)
	newManager.Config().DebugConfig().Debug(true)
	newManager.Config().PrintLog(true)
	newManager.Config().Key().Generate32BytesKeys()

	newRouter := router.NewRouter(newManager)
	newRouter.SetMiddleware(mddl)

	newRouter.GetMux().Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))
	newRouter.GetMux().Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("src/media"))))
	newRouter.Get("/explore", handlers.ExploreHNDL)
	newRouter.Get("/profile/<id>", hprofile.ObjectProfileViewHNDL())
	newRouter.Get("/profile-edit/<id>", hprofile.ObjectProfileEditViewHNDL())
	newRouter.Post("/profile-edit-post", hprofile.ProfileEditPostHNDL)
	newRouter.Get("/logout", hauth.LogOutHNDL)
	newRouter.Get("/login", hauth.LoginHNDL)
	newRouter.Post("/login-post", hauth.LoginPostHNDL)
	newRouter.Get("/register", hauth.RegisterHNDL)
	newRouter.Post("/register-post", hauth.RegisterPostHNDL)
	newRouter.Get("/publication/<id>", hpublication.PublicationViewHNDL())
	newRouter.Get("/new-publication", hpublication.NewPublicationPageHNDL)
	newRouter.Post("/new-publication-post", hpublication.NewPublicationHNDL())
	newRouter.Post("/publication-like", hpublication.PublicationLikeHNDL)
	newRouter.Post("/publication-comment", hpublication.PublicationCommentHNDL)

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
