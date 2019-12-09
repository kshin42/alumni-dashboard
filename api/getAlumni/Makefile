build:
	$(MAKE) clean
	env GOOS=linux go build .
	zip main.zip getAlumni
clean:
	rm -f main.zip getAlumni
push:
	$(MAKE) clean
	$(MAKE) build
	aws lambda update-function-code --function-name sftp-admin-create-user --zip-file fileb://main.zip
