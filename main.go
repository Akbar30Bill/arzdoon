package main

import (
	"github.com/mamal72/arzdoon/pkg/adapters/bonbast"
	// "github.com/mamal72/arzdoon/pkg/utils"
	"net/http"
	// "fmt"
	"encoding/json"
)

type arzPrice struct {
	value int `json:"value"`
}

type arzData struct {
	postfix string   `json:"postfix"`
	color   string   `json:"color"`
	data    arzPrice `json:"data"`
}

func main() {
	// provider, _ := bonbast.New()
	// prices, _ := provider.GetPriceList()
	// utils.PrintPriceTable(provider.GetAdapterName(), prices)
	http.HandleFunc("/", prices_route)
	http.ListenAndServe(":8080", nil)
}

func prices_route(w http.ResponseWriter, r *http.Request) {
	provider, _ := bonbast.New()
	prices, _ := provider.GetPriceList()
	w.Header().Set("Content-Type", "application/json")
	var price uint64
	for _, priceItem := range *prices {
		// fmt.Printf("%d, %s, %s\n", index, priceItem.Title, priceItem.SellPrice)
		if priceItem.Title == "US Dollar" {
			price = priceItem.SellPrice
		}
	}
	pricess := map[string]interface{}{
		"postfix": "IRR",
		"color":   "green",
		"data": map[string]interface{}{
			"value": price,
		},
	}

	json.NewEncoder(w).Encode(pricess)
	// fmt.Fprintf(w, prices)
}
