build: clean
	cd ./assets && make
	go build -o build/main
	GOOS=windows GOARCH=386 go build -o build/main_windows.exe
	zip -r build/main_win.zip build/main_windows.exe
clean:
	rm -rf build
all: build
