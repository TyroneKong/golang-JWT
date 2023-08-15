# golang-JWT


**SETUP**


Prerequesites: MYSQL

 - Configure .env 

DB_HOST= 127.0.0.1                       
DB_DRIVER= mysql                          
DB_USER= root
DB_PASSWORD= root
DB_NAME=golang
DB_PORT=3306 
API_SECRET= **** replace with your secret
TOKEN_HOUR_LIFESPAN=1

The following packages are needed

// gin framework
go get -u github.com/gin-gonic/gin
// ORM library
go get -u github.com/jinzhu/gorm
// package that we will be used to authenticate and generate our JWT
go get -u github.com/dgrijalva/jwt-go
// to help manage our environment variables
go get -u github.com/joho/godotenv
// to encrypt our user's password
go get -u golang.org/x/crypto


**Running **
**make run
