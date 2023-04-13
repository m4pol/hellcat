go build -ldflags "-H=windowsgui -s -w" -o "hellcat.exe" hellcat.go

:: If you want to compress the backdoor, don't forget to change the UPX path!
C:\upx\upx.exe -9 "hellcat.exe"