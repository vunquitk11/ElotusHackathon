# ElotusHackathon VuNguyen

##Build App
```
docker compose up -d
make migrateup
make gen-orm-models
```

##Run App
```
go run cmd/*.go
```

##APIs

###REST API server

###Authenticated: http://localhost:3000/authenticated

###Public: http://localhost:3000/public

### Register
```http request method
POST http://localhost:3000/public/register
```

- Request body
```json
{
  "username": "username1",
  "password": "password1"
}
```

- Response body:
```json
{
  "message": "success"
}
```

### Login
```http request method
POST http://localhost:3000/public/login
```

- Request body
```json
{
  "username": "username1",
  "password": "password1"
}
```

- Response body:
```json
{
  "message": "success"
}
```

### Logout
```http request method
POST http://localhost:3000/authenticated/logout
```

- Empty Request body

- Response body:
```json
{
  "message": "success"
}
```

### Upload File
```http request method
POST http://localhost:3000/authenticated/upload
```

- Request body
```form-data
file
```

- Response body:
```json
{
  "message": "success"
}
```

### Get files of login user
```http request method
GET http://localhost:3000/authenticated/files
```

- Empty Request body

- Response body:
```json
[
  {
    "id": 10,
    "user_id": 1,
    "name": "viet_travel.png",
    "type": "image/png",
    "size": 301517,
    "data": "abcxyz"
  }
]
```

## Project architecture
- Workflow: Request => Handler => Controller => Repository => Database

- Three layers model:
    + Handler: Get request from httpRequest, decode, validate, call controllers, write httpResponse
    + Controller: Handle business logic, call repositories
    + Repository: Data access layer 
    
## Testing
- Get data from file record
- Use this link for checking if generate correct image from base64 data: https://codebeautify.org/base64-to-image-converter
- Postman collection in /postman folder