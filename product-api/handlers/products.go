// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"microservices/product-api/data"
	"net/http"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct{
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct{
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}
// swagger:response noContent
type productsNoContent struct{

}
type Product struct{
	l *log.Logger
}
func NewProducts(l *log.Logger) *Product{
	return &Product{l}
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