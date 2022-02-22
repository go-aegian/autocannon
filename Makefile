phony: gin chi airborne echo node gorilla fiber fasthttp netcore

start-gin:
	go run ./webservers/gin/basic.go

start-chi:
	go run ./webservers/chi/basic.go

start-airborne:
	go run ./webservers/airborne/basic.go

start-echo:
	go run ./webservers/echo/basic.go

start-gorilla:
	go run ./webservers/gorilla/basic.go

start-fiber:
	go run ./webservers/fiber/basic.go

start-fasthttp:
	go run ./webservers/fasthttp/basic.go

start-node:
	node ./webservers/node_express/index.js

start-netcore:
	dotnet run --project ./webservers/netcore/netcore.csproj

#Gin Router
gin:
	go run autocannon.go -uri=http://localhost:3000 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3000 -connections=64 -pipelining=8

# Chi Router
chi:
	go run autocannon.go -uri=http://localhost:3001 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3001 -connections=64 -pipelining=8

# Airborne Router
airborne:
	go run autocannon.go -uri=http://localhost:3002 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3002 -connections=64 -pipelining=8

# Echo Router
echo:
	go run autocannon.go -uri=http://localhost:3003 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3003 -connections=64 -pipelining=8

# NodeJS Router
node:
	go run autocannon.go -uri=http://localhost:3004 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3004 -connections=64 -pipelining=8

# Gorilla Router
gorilla:
	go run autocannon.go -uri=http://localhost:3005 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3005 -connections=64 -pipelining=8

# fasthttp Router
fasthttp:
	go run autocannon.go -uri=http://localhost:3006 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3006 -connections=64 -pipelining=8

# fiber Router
fiber:
	go run autocannon.go -uri=http://localhost:3007 -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3007 -connections=64 -pipelining=8

# NETCore Router
netcore:
	go run autocannon.go -uri=http://localhost:3008/WeatherForecast -connections=1 -pipelining=8
	go run autocannon.go -uri=http://localhost:3008/WeatherForecast -connections=64 -pipelining=8