clean:
	rm -rf unzip demo wordTemplate out.docx ~$$out.docx
build: clean
	go build
	./wordTemplate
	unzip out.docx -d unzip
build2: clean 
	go run listDemo.go
	unzip demo.docx -d demo
all: build
