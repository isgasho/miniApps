build:
	rm -rf *.csv
	go build -o main_synergy
	GOOS=windows GOARCH=386 go build -o main_win_synergy.exe
	zip -r main_win_synergy.zip main_win_synergy.exe
