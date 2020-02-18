build:
	cd src && \
	GOOS=linux go build -o ../bin/cloudwatch-sns-to-slack -v
