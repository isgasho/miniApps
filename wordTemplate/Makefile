clean:
	rm -rf src unzip out.docx 
build: clean
	go build
	./src
	unzip out.docx -d unzip
all: build
