package handlers

import (
	"context"
	"fmt"
	"log"
	"microservices/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "regexp"
	// "strconv"
)
type Product struct{
	l *log.Logger
}
func NewProducts(l *log.Logger) *Product{
	return &Product{l}
}
func (p *Product) Getproducts(rw http.ResponseWriter, r *http.Request){
	lp:=data.GetProducts()
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
}

func (p *Product) Addproduct(rw http.ResponseWriter,r *http.Request){
	p.l.Println("Handle POST product")
	prod:=r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	// p.l.Println(err)
}

func (p *Product) UpdateProducts(rw http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id,err:=strconv.Atoi(vars["id"])
	if err!=nil{
		http.Error(rw,"Not convertable id",http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT product",id)
	prod:=r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id,&prod)
	if err==data.ErrProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(rw,"Product not found",http.StatusInternalServerError)
		return
	}
}
type KeyProduct struct{}
func (p *Product) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter,r *http.Request){
		prod:=data.Product{}

		err:=prod.FromJSON(r.Body)
		if err!=nil{
			p.l.Println("[ERROR] deserializing product",err)
			http.Error(rw,"Error Reading Product",http.StatusBadRequest)
			return
		}
		//validate the product
		err=prod.Validate()
		if err!=nil{
			p.l.Println("[ERROR] validating product",err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating Product:%s",err),
				http.StatusBadRequest,
			)
			return
		}
		//add product to the context
		ctx:=context.WithValue(r.Context(),KeyProduct{},prod)
		req:=r.WithContext(ctx)

		//call the nxt handler which can be another middleware in the chain, or final handler
		next.ServeHTTP(rw,req)
	})
}