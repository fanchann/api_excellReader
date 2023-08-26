## Converting Excel Data to SQL Database

### Endpoint
GET all data
```sh
curl -X GET https://localhost:9000/api/customers
```

GET data by id
```sh
curl -X GET https://localhost:9000/api/customer/:id
```

Upload data
```sh
curl -X POST -F document=@yourFileHere.xlsx https://localhost:9000/api/uploads
```

### Upload file using curl
```sh
curl -X POST -F document=@yourFileHere.xlsx localhost:9000/api/uploads
```

### Format file
excel data format :
- column name: customer_name & customer_email\
- sheet name: Sheet1
![image](/assets/image.png)