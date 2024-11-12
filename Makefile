run: build
	./bin/gcreate help 

build: main.go
	go build -o bin/gcreate main.go

install: build
	mkdir -p ~/.config/gcreate/; cp -r templates/ ~/.config/gcreate/; cp ./bin/gcreate ~/.local/bin; echo "DONE";

test: build
	go test ./tests/

clean:
	rm -rfv ~/.config/gcreate/; ~/.local/bin/gcreate; echo "DONE removing from path and config/gcreate";
