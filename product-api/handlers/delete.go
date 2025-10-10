package handlers
import(
	"net/http"
	"microservices/product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)
// swagger:route DELETE /products/{id} products deleteProduct
// Return a list of products
// responses:
//	201: noContent

// DeleteProduct deletes a product from the database
func (p *Product) DeleteProduct(rw http.ResponseWriter,r *http.Request){
	//this will always convert due to router
	vars:=mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])

	p.l.Println("Handle Delete Product",id)

	err:=data.DeleteProduct(id)

	if err==data.ErrProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(rw,"Product not found",http.StatusInternalServerError)
		return
	}
}
