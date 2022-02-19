phony: gin chi airborne

start-gin:
	go run ./webservers/gin/basic.go

start-chi:
	go run ./webservers/chi/basic.go

start-airborne:
	go run ./webservers/airborne/basic.go

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