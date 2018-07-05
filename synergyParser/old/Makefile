prime: clean
	go build -o build/main	
build: clean prime
	GOOS=windows GOARCH=386 go build -o build/main_windows.exe
	zip -r build/main_win_old.zip build/main_windows.exe
clean:
	rm -rf out
	rm -rf build
all: build
