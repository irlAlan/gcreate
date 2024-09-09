run: build
	./bin/cppcreate help 

build: main.go
	go build -o bin/cppcreate main.go

install: build
	mkdir -p ~/.config/cppcreate/; cp -r templates/ ~/.config/cppcreate/; cp ./bin/cppcreate ~/.local/bin; echo "DONE";

test: all_test.go
	go test
