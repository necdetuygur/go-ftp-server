```sh
go get github.com/fclairamb/ftpserverlib
go get github.com/spf13/afero
go mod tidy
go fmt main.go && go build -o ftpserver && ./ftpserver
```

```sh
docker build --output=dist .
git tag v1.0.1
git push origin v1.0.1
```

## v1.0.1 Latest Release
[go-ftp-server-linux-amd64](https://github.com/necdetuygur/go-ftp-server/releases/download/1.0.1/go-ftp-server-linux-amd64)

[go-ftp-server-windows-amd64.exe](https://github.com/necdetuygur/go-ftp-server/releases/download/1.0.1/go-ftp-server-windows-amd64.exe)
