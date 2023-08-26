## Converting Excel Data to SQL Database

### Endpoint
```http
https://localhost:9000/api/customers
https://localhost:9000/api/customer/:id
https://localhost:9000/api/uploads
```

### Upload file using curl
```sh
curl -F document=@yourFileHere.xlsx localhost:9000/api/uploads
```

### Format file
excel data format :
- column name: customer_name & customer_email\
- sheet name: Sheet1
![image](/assets/image.png)