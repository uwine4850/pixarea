package main

import "github.com/uwine4850/foozy/pkg/server/livereload"

func main() {
	reload := livereload.NewReload("src/cmd/main.go", livereload.NewWiretap([]string{"src/cmd", "src/handlers", "src/api", "src/utils", "src/middlewares"},
		[]string{}))
	reload.Start()
}
