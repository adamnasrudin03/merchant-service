# merchant-service
# How to run this project

1. Create an .env file whose contents are based on the .env.example file
2. Create a database named db based on DB_NAME in the .env file
3. run the project

```
go run main.go
```

# Testing api with postman

```
import the file in Postman located at ./merchant service.postman_collection.json
```


## Tools used this project

1. Gin-gonic
2. Gorm
3. MySQL
4. godotenv
5. jwt



# REST API Specification

### GROUP: Auth

- [1] - Login
- [POST] : {root.api}/api/v1/auth/login

```json

Request Body:
{
    "username": "admin1",
    "password": "admin1"
}

Response:
{
    "meta": {
        "message": "login Success",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "name": "Admin 1",
        "user_name": "admin1",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY5NDQyMDkyMCwiaWF0IjoxNjYyODg0OTIwLCJpc3MiOiJhZGFtbmFzcnVkaW4ifQ.-mcosNqdVT_I1-FLMnAGqO5gOgrcAHaZRNnBeH4pGfw",
        "created_at": "2022-09-10T15:49:25+07:00",
        "created_by": 1,
        "updated_at": "2022-09-10T15:49:25+07:00",
        "updated_by": 1
    }
}

```

### GROUP: Transactions
- [2] - List transaction by merchant id
- [GET] : {root.api}/api/v1/transaction/merchant/:id_merchant

```json

"header": [
  {
    "key": "Authorization",
    "value": "Bearer token",
    "type": "text"
  }
]

Request Param:
  -page(Required)
  -limit(Required)
  -start_at(Required)
  -end_at(Required)


Response:
{
    "meta": {
        "message": "List of transaction",
        "code": 200,
        "status": "success"
    },
    "data": {
        "total_data": 5,
        "limit": 10,
        "page": 1,
        "last_page": 1,
        "data": [
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 1,
                "outlet_name": "Outlet 1",
                "omset_total": 4500,
                "transaction_date": "2021-11-01T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 1,
                "outlet_name": "Outlet 1",
                "omset_total": 6000,
                "transaction_date": "2021-11-02T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 3,
                "outlet_name": "Outlet 3",
                "omset_total": 2500,
                "transaction_date": "2021-11-03T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 1,
                "outlet_name": "Outlet 1",
                "omset_total": 6000,
                "transaction_date": "2021-11-04T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 1,
                "outlet_name": "Outlet 1",
                "omset_total": 14000,
                "transaction_date": "2021-11-05T00:00:00+07:00"
            }
        ]
    }
}
```


- [3] - List transaction by outlet id
- [GET] : {root.api}/api/v1/transaction/outlet/:id_outlet

```json

"header": [
  {
    "key": "Authorization",
    "value": "Bearer token",
    "type": "text"
  }
]

Request Param:
  -page(Required)
  -limit(Required)
  -start_at(Required)
  -end_at(Required)

Response:
{
    "meta": {
        "message": "List of transaction",
        "code": 200,
        "status": "success"
    },
    "data": {
        "total_data": 4,
        "limit": 10,
        "page": 1,
        "last_page": 1,
        "data": [
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 3,
                "outlet_name": "Outlet 3",
                "omset_total": 2000,
                "transaction_date": "2021-11-02T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 3,
                "outlet_name": "Outlet 3",
                "omset_total": 2500,
                "transaction_date": "2021-11-03T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 3,
                "outlet_name": "Outlet 3",
                "omset_total": 5000,
                "transaction_date": "2021-11-04T00:00:00+07:00"
            },
            {
                "merchant_id": 1,
                "merchant_name": "Merchant 1",
                "outlet_id": 3,
                "outlet_name": "Outlet 3",
                "omset_total": 7000,
                "transaction_date": "2021-11-05T00:00:00+07:00"
            }
        ]
    }
}
```