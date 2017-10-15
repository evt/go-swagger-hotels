# go-swagger-hotels
Playing with go-swagger
```
iMac-2:go-swagger-hotels jt$ go install ./cmd/hotel-server/
iMac-2:go-swagger-hotels jt$ hotel-server --port=8080
2017/10/16 00:37:52 Serving hotel at http://127.0.0.1:8080
...
```
## GET /hotels/all
```
iMac-2:~ jt$ curl -i http://localhost:8080/hotels/all
HTTP/1.1 200 OK
Content-Type: application/io.goswagger.hotels.v1+json
Date: Sun, 15 Oct 2017 21:40:10 GMT
Transfer-Encoding: chunked

[{"id":"169173","name":"The Crown Borneo"},{"id":"169173","name":"The Crown Borneo"},{"id":"169173","name":"The Crown Borneo"},{"id":"175947","name":"Tune Hotel - 1Borneo Kota Kinabalu"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"151064","name":"Kinabalu Daya Hotel"},{"id":"169174","name":"Megah D'Aru"},{"id":"124720","name":"Cititel Express Kota Kinabalu"},{"id":"124720","name":"Cititel Express Kota Kinabalu"},{"id":"124720","name":"Cititel Express Kota Kinabalu"},{"id":"124720","name":"Cititel Express Kota Kinabalu"},{"id":"74019","name":"The Palace Hotel Kota Kinabalu"},{"id":"74019","name":"The Palace Hotel Kota Kinabalu"},{"id":"187167","name":"Winner Hotel"},{"id":"147333","name":"Celyn Hotel City Mall"},{"id":"147333","name":"Celyn Hotel City Mall"},{"id":"187474","name":"Celyn City Hotel"},{"id":"187474","name":"Celyn City Hotel"},{"id":"195390","name":"Oceania Hotel"},{"id":"195390","name":"Oceania Hotel"},{"id":"195390","name":"Oceania Hotel"},{"id":"195390","name":"Oceania Hotel"},{"id":"74020","name":"Sabah Oriental"},{"id":"169997","name":"Shangri-La"},{"id":"169997","name":"Shangri-La"},{"id":"187230","name":"Zara's Boutique Hotel @ Harbour City"},{"id":"187230","name":"Zara's Boutique Hotel @ Harbour City"},{"id":"71680","name":"Promenade Hotel Sabah"},{"id":"156613","name":"Ming Garden Hotel \u0026 Residences"},{"id":"156613","name":"Ming Garden Hotel \u0026 Residences"},{"id":"156613","name":"Ming Garden Hotel \u0026 Residences"},{"id":"71684","name":"Nexus Resort \u0026 Spa Karambunai"},{"id":"71684","name":"Nexus Resort \u0026 Spa Karambunai"},{"id":"71684","name":"Nexus Resort \u0026 Spa Karambunai"},{"id":"74769","name":"Sutera Harbour - The Pacific Sutera"},{"id":"74769","name":"Sutera Harbour - The Pacific Sutera"},{"id":"71682","name":"Sutera Harbour Resort - Magellan Sutera"},{"id":"71685","name":"Shangri-Laâs Rasa Ria Resort \u0026 Spa"},{"id":"71685","name":"Shangri-Laâs Rasa Ria Resort \u0026 Spa"},{"id":"174326","name":"Gaya Island Resort"},{"id":"74021","name":"Gayana Eco Resort Kota Kinabalu"},{"id":"125457","name":"Bunga Raya Island Resort"},{"id":"125457","name":"Bunga Raya Island Resort"}]
```
