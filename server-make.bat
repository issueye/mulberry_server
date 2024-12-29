go-winres make --in="winres/winres.json" --out="winappres/rsrc"

go build -tags=tempdll -buildmode=exe -ldflags="-s -w -H windowsgui" -o bin/mulberry.exe main.go

upx bin/mulberry.exe

pause