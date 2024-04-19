build:
	GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bin/hotspotcat main.go

deploy:
	scp -r bin/hotspotcat rpi4:~
