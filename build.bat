go install github.com/tc-hib/go-winres@latest
cd cmd && go-winres simply --icon ../assets/msi.png && go build -ldflags "-H=windowsgui -s -w" -o ../bin/"temp.msi.exe" && del *.syso
del C:\Users\%username%\go\bin\go-winres.exe 
go clean -cache -modcache