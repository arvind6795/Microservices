package handlers

import (
	"log"
	"net/http"
)

type First struct{
	l *log.Logger
}

func Newfirst(l *log.Logger) *First{
	return &First{l}
}

func (f *First)ServeHTTP(rw http.ResponseWriter,r *http.Request){
	rw.Write([]byte("Hi I'm FirstHandler ;)"))
}