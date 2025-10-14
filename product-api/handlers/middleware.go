package handlers
import(
	"microservices/product-api/data"
	"net/http"
	"context"
	"fmt"

)
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