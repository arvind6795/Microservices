package handlers

import (
	"log"
	"net/http"
	"microservices/product-api/data"
	"encoding/json"
	
)
type Product struct{
	l *log.Logger
}
func NewProducts(l *log.Logger) *Product {
    return &Product{l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
    p.l.Println("Handle GET Products")

    // get the list from data
    lp := data.GetProducts()  // lp is []*data.Product, not *([]*data.Product)
    // serialize to JSON
    err := json.NewEncoder(rw).Encode(lp)
    if err != nil {
        http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
        return
    }
}

