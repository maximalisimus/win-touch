@echo off
SET PATH=C:\WSL-Win-7\GoLang\go1.16.6.windows-amd64\bin\;%PATH%
SET PATH=C:\WSL-Win-7\GoLang\go1.16.6.windows-amd64\pkg\tool\windows_amd64\;%PATH%
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags=-w -ldflags=-s main.go

