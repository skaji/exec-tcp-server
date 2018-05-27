run: example/example
	sudo ./exec-tcp-server --addr 0.0.0.0:80 --user nobody example/example

example/example: example/example.go
	cd example && go build -o example
