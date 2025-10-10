package handlers
import(
	"net/http"
	"microservices/product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)
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