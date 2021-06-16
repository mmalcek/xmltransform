@echo OFF
go build -trimpath -ldflags "-s -w" -o ./dist/xmltransform.exe