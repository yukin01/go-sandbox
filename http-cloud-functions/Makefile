NAME=HelloWorld

deploy:
	gcloud functions deploy $(NAME) \
	--runtime go111 \
	--trigger-http

delete:
	gcloud functions delete $(NAME)
