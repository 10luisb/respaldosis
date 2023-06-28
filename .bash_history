sudo apt-get update
protoc --version
ls
cd ejemplo_despliegue/
ls
cd redisServer/
ls
go mod init redisMain
go mod tidy
go get github.com/redis/go-redis/v9
go get -u github.com/redis/go-redis/v9
ls
cd ejemplo_despliegue/
clear
ls
cd ..
ls
mkdir redisSever
mkdir ejemplo_despliegue
go mod init redisMain
go mod tidy
go get github.com/redis/go-redis/v9
mkdir mysqlServer
ls
go mod init mySqlMai
go mod tidy
ls
go get gorm.io/gorm
go get gorm.io/driver/mysql
ls
go mod init Main
go mod tidy
ls
go mod tidy
go get google.golang.org/grpc
protoc --version
go get github.com/gofiber/fiber/v2
go mod tidy
go mod download
ls
cd grpcClient/
ls
protoc --go_out=./ --go-grpc_out=./ client.proto
go run Main.go
cd ..
clear
ls
go run Main.go
go get -u github.com/gofiber/fiber/v2
go run Main.go
go mod tidy
go mod download
go run Main.go
go get -u google.golang.org/grpc
go get github.com/gofiber/fiber/v2
go run Main.go
ls
export PATH=$PATH:$(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH)/bin
go run Main.go
kubectl create namespace ejemplo
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project alpine-alpha-391204
kubectl create namespace ejemplo
go mod tidy
go mod download
ls
cd prueba
ls
go mod init hola
go mod tidy
ls
cd ejemplo_despliegue/
ls
go mod init Main
go mod tidy
ls
cd ..
ls
go work init ./ejemplo_despliegue
ls
cd ejemplo_despliegue/
ls
go run main.go
go run Main.go
go mod tidy
go run Main.go
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project alpine-alpha-391204
ls
mkdir redisServer
ls
go mod init redisMain
go mod tidy
ls
go run redisMain.go 
go get github.com/redis/go-redis/v9
ls
cd ejemplo_despliegue/
ls
go get google.golang.org/grpc
go get github.com/gofiber/fiber/v2
ls
cd grpcClient/
ls
clear
ls
protoc --go_out=./ --go-grpc_out=./ client.proto
ls
cd ..
ls
clear
ls
cd ..
ls
cd redisServer/
ls
clear
ls
mkdir grpcServerRedis
ls
cd grpcServerRedis/
ls
clear
protoc --go_out=./ --go-grpc_out=./ serverRedis.proto
cd ..
clear
ls
go get google.golang.org/grpc
docker built -t gcr.io/alpine-alpha-391204/grpcredis .
ls
docker built -t gcr.io/alpine-alpha-391204/grpcredis .
gomod dowload
go mod tidy
docker built -t gcr.io/alpine-alpha-391204/grpcredis .
docker build -t gcr.io/alpine-alpha-391204/grpcredis .
docker push gcr.io/alpine-alpha-391204/grpcredis 
docker build -t gcr.io/alpine-alpha-391204/grpcredis .
docker push gcr.io/alpine-alpha-391204/grpcredis 
kubectl attach redisbox -c redisbox -i -t
kubectl run -i -tty --image=gcr.io/google_container/redis:v1 -- sh
kubectl run -i --tty --image=gcr.io/google_container/redis:v1 -- sh
kubectl attach redisbox -c redisbox -i -t
gcloud container clusters get-credentials proyecto2 --zone=us-central1-a --project=alpine-alpha-391204
kubctl attach redisbox -c redisbox -i -t
kubectl attach redisbox -c redisbox -i -t
kubectl run -i --tty --image=gcr.io/google_container/redis:v1 -- sh
kubectl run -i --tty redisbox  --image=gcr.io/google_container/redis:v1 -- sh
kubctl attach redisbox -c redisbox -i -t
kubectl attach redisbox -c redisbox -i -t
kubectl run -i --tty redisbox  --image=gcr.io/google_container/redis:v1 -- sh
kubectl attach redisbox -c redisbox -i -t
kubectl run -i --tty busybox --image=busybox -- sh
redis-cli -h 10.153.60.51 info
kubectl run -i --tty busybox --image=busybox -- sh
kubectl run -i --ty redisbox --image=gcr.io/google_containers/redis:v1 -- sh
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project alpine-alpha-391204
kubectl run -i --ty redisbox --image=gcr.io/google_containers/redis:v1 -- sh
kubectl run -i --tty redisbox --image=gcr.io/google_containers/redis:v1 -- sh
kubectl attach redisbox -c redisbox -1 -t
kubectl attach redisbox -c redisbox -i -t
kubectl run -i --tty redisbox --image=gcr.io/google_containers/redis:v1 -- sh
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project alpine-alpha-391204
kubectl attach redisbox -c redisbox -i -t
kubectl get pods
kubectl namespace
kubectl namespaces
kubectl get namespaces
kubectl run -i --tty busybox --image=gcr.io/google_containers/redis:v1 -- sh
kubectl attach busybox -c redisbox -i -t
kubectl attach redisbox -c redisbox -i -t
kubectl run -i --tty redisbox --image=gcr.io/google_containers/redis:v1 -- sh
kubectl attach redisbox -c redisbox -i -t
clear
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project alpine-alpha-391204
kubectl create deploymento grpcredis --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
kubectl create deployment grpcredis --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
docker build -t gcr.io/alpine-alpha-391204/grpcredis .
docker push gcr.io/alpine-alpha-391204/grpcredis 
kubectl expose deployment grcpredis --type=LoadBalancer --port 3001 -n=ejemplo
kubectl expose deployment grpcredis --type=LoadBalancer --port 3001 -n=ejemplo
ls
docker build -t gcr.io/alpine-alpha-391204/grpclient .
docker push gcr.io/alpine-alpha-391204/grpclient 
kubectl create deployment grpclient --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
kubectl expose deployment grpclient --type=LoadBalancer --port 3000 -n=ejemplo
LS
ls
go run Main.go
ls
go run Main.go
ls
docker build -t gcr.io/alpine-alpha-391204/grpclient .
docker push gcr.io/alpine-alpha-391204/grpclient 
kubectl attach redisbox -c redisbox -i -t
kubectl create deployment grpclient --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
kubectl expose deployment grpclient --type=LoadBalancer --port 3000 -n=ejemplo
kubectl logs
kubectl get pods
kubectl logs -l app=grpclient
kubectl get namespace
ls
go run Main.go
docker build -t gcr.io/alpine-alpha-391204/grpclient .
docker push gcr.io/alpine-alpha-391204/grpclient 
kubectl create deployment grpclient --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
kubectl expose deployment grpclient --type=LoadBalancer --port 3000 -n=ejemplo
go run Main.go
docker build -t gcr.io/alpine-alpha-391204/grpclient .
docker push gcr.io/alpine-alpha-391204/grpclient 
kubectl create deployment grpclient --image=gcr.io/alpine-alpha-391204/grpcredis -n=ejemplo
kubectl expose deployment grpclient --type=LoadBalancer --port 3000 -n=ejemplo
ls
cd ..
ls
git init
sudo git init
