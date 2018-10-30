USERNAME ?= $(shell bash -c 'read -s -p "Username: " uname; echo $$uname')
PASSWORD ?= $(shell bash -c 'read -s -p "Password: " pwd; echo $$pwd')
FILE=old
LD_FLAG=-X main.usernameG=$(USERNAME) -X main.passwordG=$(PASSWORD) -X main.version=$$(git rev-parse HEAD)
prime: clean
	go build -ldflags "$(LD_FLAG)" -o build/main	
build: clean prime
	GOOS=windows GOARCH=386 go build -ldflags "$(LD_FLAG)" -o build/main_win_$(FILE).exe
	zip -r build/main_win_$(FILE).zip build/main_win_$(FILE).exe
clean:
	rm -rf out
	rm -rf build
all: build
