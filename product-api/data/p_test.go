package data

import "testing"

func TestCheckValidation(t *testing.T){
	p:=&Product{
		Name:"Elliot",
		Price:1.032,
		SKU: "abf-fds-ads",
		}
	err:=p.Validate()
	if err!=nil{
		t.Fatal(err)
	}
}