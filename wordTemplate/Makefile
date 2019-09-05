clean:
	rm -rf unzip demo demo.docx wordTemplate out.docx ~$$out.docx 
build: clean
	go build
	./wordTemplate
	unzip out.docx -d unzip
build2: clean 
	go run tableDemo.go
	unzip demo.docx -d demo
all: build
