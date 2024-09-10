run: build
	./bin/cppcreate help 

build: main.go
	go build -o bin/gcreate main.go

install: build
	mkdir -p ~/.config/gcreate/; cp -r templates/ ~/.config/gcreate/; cp ./bin/gcreate ~/.local/bin; echo "DONE";

test: build
	go test ./tests/
