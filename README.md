# details
Service to store and modify the product details description of any goods


To create one more "product-details" (client-side validation)
```
curl -X POST -H "Content-Type: application/json" -d '{"group":[["header|main","volume|1l","diameter|10сm","height|20cm"],["header|material","base|ceramic","handle|glass"]],"product-name":"Cup of the lord"}' http://localhost:8080/details/create
```