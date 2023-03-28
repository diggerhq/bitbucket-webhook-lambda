package lambda

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	eventKeySlice, ok := r.Header["X-Event-Key"]
	if !ok {
		fmt.Fprintf(w, "unknown event")
		return
	}
	if len(eventKeySlice) != 1 {
		fmt.Fprintf(w, "too many values in X-Event-Key")
		return
	}
	eventKey := eventKeySlice[0]

	var pullRequestCommentCreated = BitbucketPullRequestCommentCreated{}
	var pullRequestApproved = BitbucketPullRequestApproved{}
	var repoPush = BitbucketRepoPush{}
	var pullRequestCreated = BitbucketPullRequestCreated{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body %v", err)
		fmt.Fprint(w, "failed to read request body")
		return
	}

	switch eventKey {
	case "pullrequest:comment_created":
		err := json.Unmarshal(bodyBytes, &pullRequestCommentCreated)
		if err != nil {
			fmt.Fprintf(w, "error parsing pullrequest:comment_created event: %v", err)
			log.Printf("error parsing pullrequest:comment_created event: %v", err)
		}
		rawComment := pullRequestCommentCreated.Comment.Content.Raw
		log.Printf("rawComment: %s", rawComment)
		if strings.HasPrefix(strings.TrimSpace(rawComment), "digger plan") {
			log.Print("digger plan")
			err = TriggerPipeline("terraform_plan")
			if err != nil {
				log.Printf("error %s", err.Error())
				fmt.Fprintf(w, "error %s", err.Error())
				return
			}
		}
		if strings.HasPrefix(strings.TrimSpace(rawComment), "digger apply") {
			log.Print("digger plan")
			err = TriggerPipeline("terraform_apply")
			if err != nil {
				log.Printf("error %s", err.Error())
				fmt.Fprintf(w, "error %s", err.Error())
				return
			}
		}
		return
	case "pullrequest:created":
		err := json.Unmarshal(bodyBytes, &pullRequestCreated)
		if err != nil {
			fmt.Fprintf(w, "error parsing pullrequest:created event: %v", err)
		}
		fmt.Fprint(w, "pullrequest:created")
		return
	case "pullrequest:approved":
		err := json.Unmarshal(bodyBytes, &pullRequestApproved)
		if err != nil {
			fmt.Fprintf(w, "error parsing pullrequest:approved event: %v", err)
		}
		fmt.Fprint(w, "pullrequest:approved")
		return
	case "repo:push":
		err := json.Unmarshal(bodyBytes, &repoPush)
		if err != nil {
			fmt.Fprintf(w, "error parsing repo:push event: %v", err)
		}
		fmt.Fprint(w, "repo:push")
		return
	default:
		fmt.Fprintf(w, "unknown event key: %s", eventKey)
		return
	}
}
