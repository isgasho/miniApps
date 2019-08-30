clean:
	rm -rf unzip wordTemplate out.docx ~$$out.docx
build: clean
	go build
	./wordTemplate
	unzip out.docx -d unzip
all: build
