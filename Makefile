COMMAND_NAME=imgmd

task:
	go get -u
	go build -o $(COMMAND_NAME) main.go
