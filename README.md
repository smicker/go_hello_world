# About
Programmet startar en web-server med en sida där den skriver ut "Hello World". Starta programmet och surfa till localhost:8080 för att se resultatet.  

# Build
För att få en executable i Linux så bygger du detta i Linux med  
```go build main.go```  
Du kan också bygga en exe fil till Windows från linux (cross compilation) genom att istället skriva:  
```GOOS=windows GOARCH=amd64 go build main.go```  
Källa: https://stackoverflow.com/questions/41566495/golang-how-to-cross-compile-on-linux-for-windows  

# Test run
Om du bara vill testköra det så skriver du  
```go run main.go```

# Andra källor
https://www.youtube.com/watch?v=1MXIGYrMk80
https://sv.wikipedia.org/wiki/Go_(programspr%C3%A5k)

