# About
Programmet startar en web-server med en sida där den skriver ut "Hello World". Starta programmet och surfa till localhost:8080 för att se resultatet.  

# Build
För att få en executable i Linux så bygger du detta i Linux med  
```go build main.go```  
Du kan också bygga en exe fil till Windows från linux (cross compilation) genom att istället skriva:  
```GOOS=windows GOARCH=amd64 go build main.go```  

# Test run
Om du bara vill testköra det så skriver du  
```go run main.go```


