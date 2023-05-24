# About
Programmet startar en web-server med en sida där den skriver ut "Hello World". Starta programmet och surfa till localhost:8080 för att se resultatet.  

# Install go
Download from here: https://go.dev/dl/
Follow instructions here: https://go.dev/doc/install


# Build
För att få en executable i Linux så bygger du detta i Linux med  
```go build main.go```  
Du kan också bygga en exe fil till Windows från linux (cross compilation) genom att istället skriva:  
```GOOS=windows GOARCH=amd64 go build main.go```  
Källa: https://stackoverflow.com/questions/41566495/golang-how-to-cross-compile-on-linux-for-windows  

# Test run
Om du bara vill testköra det så skriver du  
```go run main.go```

# Good GO knowledge
- You can only have one package per directory.
- The name of the package should match the name of the directory, except for main.
- Package names should be all lower case.
- 

# Testing
- Test cases must start with Test*
- Files with tests must be named *_test.go
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing"
- To use helper function in your test code but tell go tester to instead error on the line calling the helper function, add the call ```t.Helper()``` in your helper function and let the helper function take the argument (t testing.TB) instead of (t  *testing.T).

# Benchmarking
För att mäta genomsnittliga tiden det tar att köra en funktion så använder du benchmarking. Benchmarkingfunktioner skrivs i samma filer som testning.
T.ex för att testa din funktion "MyFunctionToTest".
```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunctionToTest("a")
	}
}
```
Starta benchmark med ```go test -bench=.```


# GO Docs
Du kan få dokumentationen för GO localt med hjälp av godoc.
Installera godoc med:
```go install golang.org/x/tools/cmd/godoc@latest```
godoc executable hamnar i $HOME/go/bin så se till att du har följande rad i .zshrc (eller .bashrc) för att kunna köra godoc efter installationen:
```export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH```
Starta godoc med:
```godoc -http :8000```
Se doc genom browsern på url:
http://localhost:8000/pkg

# Andra källor
https://www.youtube.com/watch?v=1MXIGYrMk80  
https://sv.wikipedia.org/wiki/Go_(programspr%C3%A5k)
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

