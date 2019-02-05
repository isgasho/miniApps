FILENAME=excel
clean:
	
prime: clean
	go build -o build/main
all: prime
	GOOS=windows GOARCH=386 go build -o build/main_win_$(FILENAME).exe
	zip -r build/main_win_$(FILENAME).zip build/main_win_$(FILENAME).exe
