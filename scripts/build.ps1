# $env:GOOS = "windows"
# go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main ../cmd/main.go
~\Go\Bin\build-lambda-zip.exe -output ../dist/main.zip main