build:
	$(MAKE) clean
	env GOOS=linux go build .
	zip main.zip getResume
clean:
	rm -f main.zip getResume
# push:
# 	$(MAKE) clean
# 	$(MAKE) build
# 	aws lambda update-function-code --function-name  --zip-file fileb://main.zip
