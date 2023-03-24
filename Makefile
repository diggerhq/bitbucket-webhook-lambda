.PHONY: gomodgen deploy delete

gomodgen:
	GO111MODULE=on go mod init

deploy:
	gcloud functions deploy lambda --entry-point WebhookHandler --runtime go111 --trigger-http

delete:
	gcloud functions delete lambda --entry-point WebhookHandler --runtime go111 --trigger-http
