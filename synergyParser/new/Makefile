build:
	rm -rf *.csv
	go build -o main_recruiterx
	GOOS=windows GOARCH=386 go build -o main_win_recruiterx.exe
	zip -r main_win_recruiterx.zip main_win_recruiterx.exe
