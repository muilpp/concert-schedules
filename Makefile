test:
	go test -v ./...

compile:
	echo "Compiling for Linux"
	GOOS=linux go build -o bin/concerts

deploy:
	scp /home/marc/workspace/marc/concert-schedules/bin/concerts marc-digital-ocean:/home/marc/concerts/

run:
	test compile
	pwd
	cd .. && ./bin/concerts

all: test compile