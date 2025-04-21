run: build
	./bin/cppcreate help 

build: main.go
	go build -o bin/gcreate main.go

install: build
	mkdir -p ~/.config/gcreate/; cp -r templates/ ~/.config/gcreate/; cp ./bin/gcreate ~/.local/bin; echo "DONE";

test: build
	go test ./tests/

# replace with a script call that safely cleans up
# clean:
	# rm -rfv ~/.config/gcreate/; rm -rfv ~/.local/bin/gcreate
