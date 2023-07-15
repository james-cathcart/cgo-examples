all: clean build

clean:
	rm -f app

cgotypes: 
	go tool cgo -godefs main.go

build:
	go build -o app main.go
