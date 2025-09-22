package handlers

import (
	"log"
	"microservices/product-api/data"
	"net/http"
)

type Product struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product{
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet{
		p.getproducts(rw,r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getproducts(rw http.ResponseWriter, r *http.Request){
	lp:=data.GetProducts()
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to encode json",http.StatusInternalServerError)
	}
}