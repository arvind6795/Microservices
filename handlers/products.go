package handlers

import (
	"log"
	"microservices/product-api/data"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method==http.MethodPost{
		p.addproduct(rw,r)
		return
	}
	if r.Method==http.MethodPut{
		p.l.Println("PUT")
		// expect id in URI
		reg:=regexp.MustCompile(`/([0-9]+)`)
		g:=reg.FindAllStringSubmatch(r.URL.Path,-1)
		if len(g)!=1{
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}
		if len(g[0])!=2{
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}
		idString:=g[0][1]
		id,err:=strconv.Atoi(idString)
		if err!=nil{
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw,r) 
		return
		// p.l.Println("got id:",id)
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getproducts(rw http.ResponseWriter, r *http.Request){
	lp:=data.GetProducts()
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
}

func (p *Product) addproduct(rw http.ResponseWriter,r *http.Request){
	p.l.Println("Handle POST product")
	prod:=&data.Product{}
	err:=prod.FromJSON(r.Body)
	if err!=nil{
		http.Error(rw,"unable to unmarshal data",http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *Product) updateProducts(id int,rw http.ResponseWriter,r *http.Request){
	p.l.Println("Handle PUT product")
	prod:=&data.Product{}
	err:=prod.FromJSON(r.Body)
	if err!=nil{
		http.Error(rw,"unable to unmarshal data",http.StatusBadRequest)
	}
	err = data.UpdateProduct(id,prod)
	if err==data.ErrProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(rw,"Product not found",http.StatusInternalServerError)
		return
	}
}