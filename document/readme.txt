go mod init github.com/dkantikorn/go-fullstack
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql //If using mysql 
go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1

## Model test 
$ cd test/modeltests
$ go test -v --run FindAllUsers
$ go test -v --run TestUpdateAPost
$ go test -v #test allcase in current path

## Controller test
$ cd test/controllertests 
$ go test -v --run TestLogin
$ go test -v --run TestCreateUser
$ go test -v --run TestDeletePost

## Test All Controller and model 
$ cd test
$ go test -v ./...