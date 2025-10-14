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
import(
	"microservices/product-api/data"
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
// swagger:route DELETE /products/{id} products deleteProduct
// Return a list of products
// responses:
//	201: noContent

// DeleteProduct deletes a product from the database
