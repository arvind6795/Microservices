package handlers
import(
	"net/http"
	"microservices/product-api/data"
)
// swagger:route GET /products products listProducts
// Return a list of products
// responses:
//	200: productsResponse
func (p *Product) Getproducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET Products")
	lp:=data.GetProducts()
	rw.Header().Set("Content-Type", "application/json")
	err:=lp.ToJSON(rw)
	if err!=nil{
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
}