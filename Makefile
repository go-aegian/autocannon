phony: gin chi airborne

start-gin:
	go run ./webservers/gin/basic.go

start-chi:
	go run ./webservers/chi/basic.go

start-airborne:
	go run ./webservers/airborne/basic.go

start-echo:
	go run ./webservers/echo/basic.go

start-node:
	node ./webservers/node_express/index.js

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