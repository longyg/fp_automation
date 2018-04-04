package nidd

import (
	"net/url"
	"../httputil"
)

func AddProduct(postUrl string, domain string, productAbbr string, productName string,
	counterAbbrUni string, counterPiidUni string,
	productDesc string) error  {

	values := make(url.Values)
	values.Set("domain", domain)
	values.Set("productAbbr", productAbbr)
	values.Set("productName", productName)
	values.Set("counterAbbrUni", counterAbbrUni)
	values.Set("counterPiidUni", counterPiidUni)
	values.Set("productDesc", productDesc)

	return httputil.PostRequest(postUrl, values)
}
