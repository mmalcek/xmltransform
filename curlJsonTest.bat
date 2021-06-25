curl.exe -s "https://api.predic8.de/shop/products/" | xmltransform.exe -f json -t "?{{range .products}}\"{{.name}}\",{{.product_url}}{{print \"\n\"}}{{end}}" 


