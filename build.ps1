$env:GOOS = "linux"
$env:GOARCH = "arm"
$env:GOARM = "6"

go build -o .\build\screen .\screen
go build -o .\build\ledstrip .\ledstrip
