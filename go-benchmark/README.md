# Go Test, Example and Benchmarking
This repo is an example of how to "bet" your code works with benchmarking, 
examples and testing.  

To run the tests, in current directory and all subdirectories
```bash
go test ./...
```
or to test with coverage
```bash
go test -cover ./...
```
finally, to write the coverage to a file and then to view the coverage as HTML
```bash
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out
```

To launch the local "godocs" and to view the examples
```bash
godoc -http:8080
```
Then navigate to the browser `localhost:8080/pkg` then scroll down to 3rd 
Party. Here all the documentation for this package can be found.

Finally, to check how well the code performs you can benchmark variation. 
Change into the directory where the code or functions are then run
```bash
go test -bench .
```
to run all the benchmark tests in that directory.