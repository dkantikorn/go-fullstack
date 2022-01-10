## Go 
$ go mod init github.com/dkantikorn/go-fullstack
$ go get github.com/badoux/checkmail
$ go get github.com/jinzhu/gorm
$ go get golang.org/x/crypto/bcrypt
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/gorilla/mux
$ go get github.com/jinzhu/gorm/dialects/mysql //If using mysql 
$ go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
$ go get github.com/joho/godotenv
$ go get gopkg.in/go-playground/assert.v1

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


### DOCKER
$ docker build go-fullstack-app:1.0 .
$ docker-compose up
$ docker-compose up -d

## When you rename or edit docker compose run belowe command instead
$ docker-compose up --remove-orphans

## Delete all container
$ docker rm $(docker ps -a -q)

## Delete all image 
$ docker rmi $(docker images -q)

$ docker-compose down
$ docker-compose down --remove-orphans --volumes
$ docker-compose up -d

$ docker-compose up --build

$ docker inspect <container_id> | grep IPAddress
$ docker inspect 4745278aad30 | grep IPAddress


## Test in docker containner 
$ docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit


## Kubernetes 
$ touch postgres-db-pv.yaml
$ touch postgres-db-pvc.yaml
$ touch postgres-db-deployment.yaml
$ touch postgres-db-service.yaml
$ touch postgres-secret.yaml

## Run minikube
$ minikube start 
$ minikube status 

## Deploy k8s secret 
$ kubectl create -f postgres-secret.yaml [--validate=false] or kubectl apply -f postgres-secret.yaml
$ kubectl get secret 
$ kubectl describe secrets <secret-name>
$ kubectl describe secret/<secret-name>

## Deploy k8s instance
$ kubectl apply -f postgres-db-pv.yaml
$ kubectl apply -f postgres-db-pvc.yaml
$ kubectl apply -f postgres-db-deployment.yaml
$ kubectl apply -f postgres-db-service.yaml

$ kubectl get pods 
$ kubectl describe pod <pod_name>
$ kubectl logs <pod_name>


## MYSQL 
$ kubectl create -f mysql-secret.yaml
$ kubectl apply -f mysql-db-pv.yaml
$ kubectl apply -f mysql-db-pvc.yaml
$ kubectl apply -f mysql-db-deployment.yaml
$ kubectl apply -f mysql-db-service.yaml
$ kubectl get pods
$ kubectl get services

## Docker deploy 
$ docker build -t <app-name> .
$ docker build -t fullstack-kubernetes .

## Docker add tag and publish
$ docker tag <image-name> <dockerhub-username>/<repository-name>:<tag-name>
$ docker tag go-fullstack-kubernetes sarawutt/go-fullstack-kubernetes:1.0

$ docker login 
$ docker push <dockerhub-username>/<repository-name>:<tag_name>
$ docker push sarawutt/go-fullstack-kubernetes:1.0

## Deploying the App to Kubernetes
$ touch app-postgres-deployment.yaml
$ touch app-postgres-service.yaml


$ kubectl apply -f app-postgres-deployment.yaml
$ kubectl apply -f app-postgres-service.yaml

$ kubectl rollout restart deployment go-fullstack-app-postgres


## Get the url exposed using
$ minikube service <service-name> --url
$ minikube service fullstack-app-mysql --url


## Minikube 
$ minikube stop
$ minikube delete