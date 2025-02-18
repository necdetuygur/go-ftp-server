```sh
go get github.com/fclairamb/ftpserverlib
go get github.com/spf13/afero
go mod tidy
go fmt main.go && go build -o ftpserver && ./ftpserver
```
