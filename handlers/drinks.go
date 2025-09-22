package handlers

import (
	"log"
	"microservices/product-api/data"
	"net/http"
)

type drink struct{
	l *log.Logger
}

func NewDrinks(l *log.Logger) *drink{
	return &drink{l}
}


func (d *drink)ServeHTTP(rw http.ResponseWriter,r *http.Request){
	if r.Method==http.MethodGet{
		d.getdrinks(rw,r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (d *drink) getdrinks(rw http.ResponseWriter,r *http.Request){
	lp:=data.GetDrinks()
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to encode drinks data",http.StatusInternalServerError)
	}
}