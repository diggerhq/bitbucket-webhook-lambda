# bitbucket-webhook-lambda

Unlike Gihub actions, Bitbucket doesn't support triggering pipelines from various events such as [PR approval](https://jira.atlassian.com/browse/BCLOUD-21695), [PR comments](https://jira.atlassian.com/browse/BCLOUD-14026) [etc](https://jira.atlassian.com/browse/BCLOUD-14026). 
This repo contains a GCP function which can listen to those events via webhooks and then trigger pipeline via Bitbucket API. 
We will port it to AWS and Azure functions in the future. 
Feel free to use it for your usecases and if you find it useful consider contributing!

## How to deploy
```
GOOGLE_PROJECT_NAME=<gcp project name> GOOGLE_CREDENTIALS_JSON_FILE=<path to application_default_credentials.json> serverless deploy --stage prod

```


allow unauthenticated requests
``` 
gcloud functions add-iam-policy-binding <functio name> --member="allUsers" --role="roles/cloudfunctions.invoker" --project=<project name> --region=us-central1
```