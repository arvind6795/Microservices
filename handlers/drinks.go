package handlers

import (
	"log"
	"microservices/product-api/data"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method==http.MethodPost{
		d.adddrinks(rw,r)
		return
	}
	if r.Method==http.MethodPut{
		d.l.Println("Put")
		reg:=regexp.MustCompile(`/([0-9]+)`)
		g:=reg.FindAllStringSubmatch(r.URL.Path,-1)
		if len(g)!=1{
			http.Error(rw,"Invalid URI for len 1",http.StatusBadRequest)
			return
		}
		if len(g[0])!=2{
			http.Error(rw,"Invalid URI for len 2",http.StatusBadRequest)
			return
		}
		idString:=g[0][1]
		id,err:=strconv.Atoi(idString)
		if err!=nil{
			http.Error(rw,"Invalid URI for err",http.StatusBadRequest)
			return
		}
		d.updateDrinks(id,rw,r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (d *drink) getdrinks(rw http.ResponseWriter,r *http.Request){
	lp:=data.GetDrinks()
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to encode data",http.StatusInternalServerError)
	}
}
func (d *drink) adddrinks(rw http.ResponseWriter,r *http.Request){
	d.l.Println("Handle Post Drinks")
	drin:=&data.Drink{}
	err:=drin.FromJSON(r.Body)
	if err!=nil{
		http.Error(rw,"unable to unmarshal data",http.StatusBadRequest)
	}
	data.AddDrinks(drin)
	d.l.Println("drink:%#v",drin)
}
func (d *drink) updateDrinks(id int,rw http.ResponseWriter,r *http.Request){
	d.l.Println("Handle PUT Drinks")
	drin:=&data.Drink{}
	err:=drin.FromJSON(r.Body)
	if err!=nil{
		http.Error(rw,"unable to unmarshal data",http.StatusBadRequest)
	}
	err = data.UpdateDrink(id,drin)
	if err==data.ErrDrinkNotFound{
		http.Error(rw,"Drink not found",http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(rw,"Drink not found",http.StatusInternalServerError)
		return
	}
}