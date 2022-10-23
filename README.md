# go_module_example
This is an example of creating go module and how to use it from another module.
## Create go_module_example
- Create GitHub repo - go_module_example
- Clone the repo locally
- Go to go_module_example folder
- go mod init github.com/rwuniard/go_module_example
- Create a dog folder
- Go to dog folder
- go mod init github.com/rwuniard/go_module_example/dog
- Create a main.go
- Put the code with package dog
- Go up one folder then create hello folder
- go mod init github.com/rwuniard/go_module_example/main
- Create a main.go

## To run the program
- Go to cmd folder
- go run .

## To look at the documentation
- from the go_module_example folder
- godoc -http=:6060
- Point your browser to http://localhost:6060
- search for pkg/dog