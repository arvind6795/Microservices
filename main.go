package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)
func main(){
	l:=log.New(os.Stdout,"/",log.LstdFlags)
	ph:=handlers.NewProducts(l)
	// dh:=handlers.NewDrinks(l)
	sm:=http.NewServeMux()
	sm.Handle("/",ph)
	// sm.Handle("/drinks/",dh)
	s:=&http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}//manually creating server in go
	go func(){
		err:=s.ListenAndServe()
		if err!=nil{
			l.Fatal(err)
		}
	}()
	sigchan:=make(chan os.Signal)
	signal.Notify(sigchan,os.Interrupt)
	signal.Notify(sigchan,os.Kill)
	sig:=<-sigchan
	l.Println("Recieved terminate, graceful shutdown",sig)
	tc,_:=context.WithTimeout(context.Background(),30*time.Second)
	s.Shutdown(tc)//shutdown server after certain task is being completed
	// http.ListenAndServe(":9090",sm)//Ip-addresss string and handler(2 param in this func)
}//This 26 lines of code is for simple Web Server in Go