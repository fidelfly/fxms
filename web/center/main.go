package main

import (
        "github.com/micro/go-micro/util/log"
	"net/http"

        "github.com/micro/go-micro/web"
        "github.com/fidelfly/fxms/center/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("com.fxms.web.center"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/center/call", handler.CenterCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
