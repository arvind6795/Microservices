package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)
type Drink struct{
	ID int 				`json:"id"`
	Name string 		`json:"name"`
	Description string  `json:"description"`
	Price float32		`json:"price"`
	SKU string			`json:"sku"`
	CreatedOn string	`json:"-"`
	UpdatedOn string	`json:"-"`
	DeletedOn string	`json:"-"`
}

type drinks []*Drink

// func(d *drink) FromJSON(r io.Reader) error{
// 	e:=json.NewDecoder(r)
// 	return e.Decode(d)
// }
func (d *Drink) FromJSON(r io.Reader) error{
	e:=json.NewDecoder(r)
	return e.Decode(d)
}
//direct encode using encode json which is faster than marshal encoding
func (d *drinks) ToJSON(w io.Writer) error{
	e:=json.NewEncoder(w)
	return e.Encode(d)
}

func GetDrinks()drinks{
	return drinkList
}
func AddDrinks(d *Drink){
	d.ID=getNxtID()
	drinkList=append(drinkList, d)
}
func getNxtID() int{
	lp:=drinkList[len(drinkList)-1]
	return lp.ID + 1
}
func UpdateDrink(id int,d *Drink)error{
	_,pos,err:=findDrink(id)
	if err!=nil{
		return err
	}
	d.ID=id
	drinkList[pos]=d
	return nil
}
var ErrDrinkNotFound=fmt.Errorf("Drink not Found")
func findDrink(id int) (*Drink,int,error){
	for i,d:=range drinkList{
		if d.ID==id{
			return d,i,nil
		}
	}
	return nil,-1,ErrDrinkNotFound
}
var drinkList=[]*Drink{
	&Drink{
		ID: 1,
		Name: "Cola",
		Description: "Carbonated soft drink flavored with vanilla, cinnamon, citrus oils, and other flavorings",
		Price: 2.45,
		SKU: "abc234",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Drink{
		ID: 2,
		Name: "Mountain-Dew",
		Description: "citrus-flavored soft drink",
		Price: 1.99,
		SKU: "fjh234",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}