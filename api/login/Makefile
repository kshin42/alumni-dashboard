build:
	$(MAKE) clean
	env GOOS=linux go build .
	zip main.zip login
clean:
	rm -f main.zip login
# push:
# 	$(MAKE) clean
# 	$(MAKE) build
# 	aws lambda update-function-code --function-name  --zip-file fileb://main.zip
