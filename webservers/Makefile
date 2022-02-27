.PHONY: gin chi airborne echo node gorilla fiber fasthttp netcore

start-gin:
	go run ./go/gin/basic.go

start-chi:
	go run ./go/chi/basic.go

start-airborne:
	go run ./go/airborne/basic.go

start-echo:
	go run ./go/echo/basic.go

start-gorilla:
	go run ./go/gorilla/basic.go

start-fiber:
	go run ./go/fiber/basic.go

start-fasthttp:
	go run ./go/fasthttp/basic.go

start-aero:
	go run ./go/aero/basic.go

start-node:
	node ./node_express/index.js

start-netcore:
	dotnet run --project ./netcore/netcore.csproj

#Gin Router
gin-bombardier:
	bombardier -c 1 -d 10s http://localhost:3000
	bombardier -c 64 -d 10s http://localhost:3000

gin-autocannon:
	go run ../autocannon.go -uri=http://localhost:3000 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3000 -duration=10 -connections=64 -pipelining=8

# Chi Router
chi-bombardier:
	bombardier -c 1 -d 10s http://localhost:3001
	bombardier -c 64 -d 10s http://localhost:3001

chi-autocannon:
	go run ../autocannon.go -uri=http://localhost:3001 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3001 -duration=10 -connections=64 -pipelining=8

# Airborne Router
airborne-bombardier:
	bombardier -c 1 -d 10s http://localhost:3002
	bombardier -c 64 -d 10s http://localhost:3002

airborne-autocannon:
	go run ../autocannon.go -uri=http://localhost:3002 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3002 -duration=10 -connections=64 -pipelining=8

# Echo Router
echo-bombardier:
	bombardier -c 1 -d 10s http://localhost:3003
	bombardier -c 64 -d 10s http://localhost:3003

echo-autocannon:
	go run ../autocannon.go -uri=http://localhost:3003 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3003 -duration=10 -connections=64 -pipelining=8

# NodeJS Router
node-bombardier:
	bombardier -c 1 -d 10s http://localhost:3004
	bombardier -c 64 -d 10s http://localhost:3004

node-autocannon:
	go run ../autocannon.go -uri=http://localhost:3004 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3004 -duration=10 -connections=64 -pipelining=8

# Gorilla Router
gorilla-bombardier:
	bombardier -c 1 -d 10s http://localhost:3005
	bombardier -c 64 -d 10s http://localhost:3005

gorilla-autocannon:
	go run ../autocannon.go -uri=http://localhost:3005 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3005 -duration=10 -connections=64 -pipelining=8

# fasthttp Router
fasthttp-bombardier:
	bombardier -c 1 -d 10s http://localhost:3006
	bombardier -c 64 -d 10s http://localhost:3006

fasthttp-autocannon:
	go run ../autocannon.go -uri=http://localhost:3006 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3006 -duration=10 -connections=64 -pipelining=8

# fiber Router
fiber-bombardier:
	bombardier -c 1 -d 10s http://localhost:3007
	bombardier -c 64 -d 10s http://localhost:3007

fiber-autocannon:
	go run ../autocannon.go -uri=http://localhost:3007 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3007 -duration=10 -connections=64 -pipelining=8

# NETCore Router
netcore-bombardier:
	bombardier -c 1 -d 10s http://localhost:3008/WeatherForecast
	bombardier -c 64 -d 10s http://localhost:3008/WeatherForecast

netcore-autocannon:
	go run ../autocannon.go -uri=http://localhost:3008/WeatherForecast -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3008/WeatherForecast -duration=10 -connections=64 -pipelining=8

# aero Router
aero-bombardier:
	bombardier -c 64 -d 10s http://localhost:3009
	bombardier -c 64 -d 10s http://localhost:3009

aero-autocannon:
	go run ../autocannon.go -uri=http://localhost:3009 -duration=10 -connections=1 -pipelining=8
	go run ../autocannon.go -uri=http://localhost:3009 -duration=10 -connections=64 -pipelining=8
