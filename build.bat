@echo off

set platform=%1

if "%platform%" == "rpi" (
  set GOARM=7
  set GOARCH=arm
  set GOOS=linux
)

if "%platform%" == "rpi" (
  go build -ldflags="-s -w" -o recalbox-manager
) else (
  go build
)
