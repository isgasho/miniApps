FILENAME=excel
all:
	go build -o build/main
	GOOS=windows GOARCH=386 go build -o build/main_win_$(FILENAME).exe
	zip -r build/main_win_$(FILENAME).zip build/main_win_$(FILENAME).exe
