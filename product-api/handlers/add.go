package handlers
import(
	"net/http"
	"microservices/product-api/data"
)

func (p *Product) Addproduct(rw http.ResponseWriter,r *http.Request){
	p.l.Println("Handle POST product")
	prod:=r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	// p.l.Println(err)
}