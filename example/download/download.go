package main

/*
Convert below curl call to go code
curl 'https://www.axis.com/api/pia/v2/items?categories=audio,cameras,encdec,modular,networkswitches,pac,recorders1&fields=none,category,name&orderBy=name&state=40&type=ProductVariant' \
 -H 'Referer: https://www.axis.com/en-us/products/product-selector' \
 -H 'Authorization: apikey da6cac02-e554-44c5-8125-1281982c3cdb'
*/

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// Prepare API request
	v := url.Values{}
	v.Set("categories", "audio,cameras,encdec,modular,networkswitches,pac,recorders1")
	v.Set("fields", "none,category,name")
	v.Set("orderBy", "name")
	v.Set("state", "40")
	v.Set("type", "ProductVariant")
	host := "https://www.axis.com/api/pia/v2/items"
	r, _ := http.NewRequest("GET", host+"?"+v.Encode(), nil)
	r.Header.Set("Authorization", "apikey da6cac02-e554-44c5-8125-1281982c3cdb")
	r.Header.Set("Referer", "https://www.axis.com/en-us/products/product-selector")

	// Call api
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Tidy response
	var raw, nice bytes.Buffer
	io.Copy(&raw, resp.Body)
	json.Indent(&nice, raw.Bytes(), "", "  ")

	// Write out
	io.Copy(os.Stdout, &nice)
}
