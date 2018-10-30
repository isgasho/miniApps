prime: clean
	go build -o build/main	
build: clean prime
	GOOS=windows GOARCH=386 go build  -o build/main_win.exe
	zip -r build/main_win.zip build/main_win.exe
clean:
	rm -rf build
all: build