build:
	rm -rf bindata.go main main_windows.exe main_windows.exe.zip
	rm -rf out
	go-bindata templates/COMP/ templates/NONCOMP/
	go build -o main
	GOOS=windows GOARCH=386 go build -o main_windows.exe
