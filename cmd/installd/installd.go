package main

import (
	"context"
	"flag"
	"github.com/4thel00z/installd/pkg/v1/libinstall"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/4thel00z/libhttp"
)

var (
	templatesPath = flag.String("templates", "templates", "Specify the dir that contains fails that map to OS names.")
	addr          = flag.String("addr", ":1337", "Specify what addr to bind to")
)

func main() {
	flag.Parse()
	templates, err := libinstall.TemplatesFromPath(*templatesPath)
	if err != nil {
		flag.Usage()
		os.Exit(64) // EX USAGE
	}
	router := libhttp.Router{}

	router.GET("/", libinstall.OSHandler(templates))

	svc := router.Serve().
		Filter(libhttp.ErrorFilter).
		Filter(libhttp.H2cFilter)
	srv, err := libhttp.Listen(svc, *addr)
	if err != nil {
		panic(err)
	}
	log.Printf("👋  Listening on %v", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
