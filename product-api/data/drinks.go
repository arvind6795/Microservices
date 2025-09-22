package data

import (
	"encoding/json"
	"io"
	"time"
)
type drink struct{
	ID int 				`json:"id"`
	Name string 		`json:"name"`
	Description string  `json:"description"`
	Price float32		`json:"price"`
	SKU string			`json:"sku"`
	CreatedOn string	`json:"-"`
	UpdatedOn string	`json:"-"`
	DeletedOn string	`json:"-"`
}

type drinks []*drink

//direct encode using encode json which is faster than marshal encoding
func (d *drinks) ToJSON(w io.Writer) error{
	e:=json.NewEncoder(w)
	return e.Encode(d)
}

func GetDrinks()drinks{
	return drinkList
}
var drinkList=[]*drink{
	&drink{
		ID: 1,
		Name: "Cola",
		Description: "Carbonated soft drink flavored with vanilla, cinnamon, citrus oils, and other flavorings",
		Price: 2.45,
		SKU: "abc234",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&drink{
		ID: 2,
		Name: "Mountain-Dew",
		Description: "citrus-flavored soft drink",
		Price: 1.99,
		SKU: "fjh234",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}