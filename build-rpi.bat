set GOARM=7
set GOARCH=arm
set GOOS=linux

go-bindata templates/...
bindata assets/...

go build -ldflags="-s -w" -o recalbox-manager
