```sh
go get github.com/fclairamb/ftpserverlib
go get github.com/spf13/afero
go mod tidy
go fmt main.go && go build -o ftpserver && ./ftpserver
```

docker build --output=dist .

## v1.0.0 Release
[go-ftp-server-linux-amd64](https://github.com/necdetuygur/go-ftp-server/releases/download/1.0.0/go-ftp-server-linux-amd64)

[go-ftp-server-windows-amd64.exe](https://github.com/necdetuygur/go-ftp-server/releases/download/1.0.0/go-ftp-server-windows-amd64.exe)
