
Build Windows executable in Ubuntu, install mingw-w64 first:
```
apt -y update
apt-get install -y mingw-w64
```

Build Windows executable in Ubuntu:
```
go mod init helloworld
go get fyne.io/fyne/v2
go mod tidy
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o helloworld.exe helloworld.go
```
