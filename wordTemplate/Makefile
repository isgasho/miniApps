clean:
	rm -rf out
build: clean
	mkdir -p out
	cd ./tmpl; go run merge.go
	html-minifier --output=out.min.jsx --output-dir=./out --input-dir=./out --file-ext=jsx --collapse-whitespace --remove-comments --remove-tag-whitespace --keep-closing-slash --remove-comments --case-sensitive
	go build -o ./out/runner 
	cd ./out;./runner;unzip out.docx -d outUnzip
build2: clean 
	go run tableDemo.go
	unzip demo.docx -d demo
all: build
