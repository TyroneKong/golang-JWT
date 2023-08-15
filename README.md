# golang-JWT

JWT Authentication using golang.


create a .env file in the root with the following 
```bash
DB_HOST= 127.0.0.1                       

DB_DRIVER= mysql                          

DB_USER= root

DB_PASSWORD= root

DB_NAME= golang

DB_PORT= 3306 

API_SECRET *****  replace with your secret

TOKEN_HOUR_LIFESPAN=1
```

## Installation

The following packages will be needed 

```bash
// gin framework
go get -u github.com/gin-gonic/gin
// ORM library
go get -u github.com/jinzhu/gorm
// package that we will be used to authenticate and generate our JWT
go get -u github.com/dgrijalva/jwt-go
// to help manage our environment variables
go get -u github.com/joho/godotenv
// to encrypt our users password
go get -u golang.org/x/crypto
go get -u github.com/go-sql-driver/mysql
```

## Usage

```golang


# to run using the included Makefile
make run



## License

[MIT](https://choosealicense.com/licenses/mit/)
