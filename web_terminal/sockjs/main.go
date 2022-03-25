package main

import (
	"flag"
	"github.com/igm/pubsub"
	"github.com/igm/sockjs-go/v3/sockjs"
	"log"
	"net/http"
	"time"
)

var (
	chat          pubsub.Publisher
	serverKeyPath string
	serverCrtPath string
)

func main() {
	flag.StringVar(&serverKeyPath, "key", "./cert/server.key", "https server key")
	flag.StringVar(&serverCrtPath, "cert", "./cert/server.crt", "https server crt")
	go startHttp()
	startHttps()
}

func startHttp() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Println("Server started on port: 8081")
	s := &http.Server{
		Addr:           ":8081",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func startHttps() {
	s := &http.Server{
		Addr: ":443",
	}

	log.Println("Server started on port: 443")
	err := s.ListenAndServeTLS(serverCrtPath, serverKeyPath)
	if err != nil {
		log.Fatal("ListenAndServeTLS err:", err)
	}
}

func echoHandler(session sockjs.Session) {
	log.Println("new sockjs session established")
	var closedSession = make(chan struct{})
	chat.Publish("[info] new participant joined chat, " + session.ID())
	defer chat.Publish("[info] participant left chat, " + session.ID())
	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(session.ID() + ": " + msg.(string)); err != nil {
					return
				}
			}

		}
	}()

	for {
		if msg, err := session.Recv(); err == nil {
			chat.Publish(msg)
			continue
		}
		break
	}
	close(closedSession)
	log.Println("sockjs session closed")
}
