package main

import (
	"context"
	"log"
	"microservices/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "/", log.LstdFlags)
	// create the handlers
	ph := handlers.NewProducts(l)
	// dh:=handlers.NewDrinks(l)
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.Getproducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.Addproduct)
	postRouter.Use(ph.MiddlewareProductValidation)
	// sm.Handle("/drinks/",dh)
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	ops:=middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh:=middleware.Redoc(ops,nil)
	getRouter.Handle("/docs",sh)
	getRouter.Handle("/swagger.yaml",http.FileServer(http.Dir("./")))
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	} //manually creating server in go
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc) //shutdown server after certain task is being completed
	// http.ListenAndServe(":9090",sm)//Ip-addresss string and handler(2 param in this func)
} //This 26 lines of code is for simple Web Server in Go
