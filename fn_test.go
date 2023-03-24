package lambda

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBitbucketPullRequestCommentCreated(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(bitbucketPullRequestCommentCreatedTest))
	r.Header.Set("X-Event-Key", "pullrequest:comment_created")
	WebhookHandler(w, r)

	rw := w.Result()
	body, err := io.ReadAll(rw.Body)
	assert.NoError(t, err)
	assert.Equal(t, "pullrequest:comment_created", string(body))
}

func TestBitbucketPullRequestApproved(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(bitbucketPullRequestApprovedTest))
	r.Header.Set("X-Event-Key", "pullrequest:approved")
	WebhookHandler(w, r)

	rw := w.Result()
	body, err := io.ReadAll(rw.Body)
	assert.NoError(t, err)
	assert.Equal(t, "pullrequest:approved", string(body))
}

func TestBitbucketRepoPush(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(bitbucketRepoPushTest))
	r.Header.Set("X-Event-Key", "repo:push")
	WebhookHandler(w, r)

	rw := w.Result()
	body, err := io.ReadAll(rw.Body)
	assert.NoError(t, err)
	assert.Equal(t, "repo:push", string(body))
}

func TestBitbucketPullRequestCreated(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(bitbucketPullRequestCreatedTest))
	r.Header.Set("X-Event-Key", "pullrequest:created")
	WebhookHandler(w, r)

	rw := w.Result()
	body, err := io.ReadAll(rw.Body)
	assert.NoError(t, err)
	assert.Equal(t, "pullrequest:created", string(body))
}

func TestBitbucketUnknownEvent(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(bitbucketUnknownEventTest))
	r.Header.Set("X-Event-Key", "test:comment_created")
	WebhookHandler(w, r)

	rw := w.Result()
	body, err := io.ReadAll(rw.Body)
	assert.NoError(t, err)
	assert.Equal(t, "unknown event key: test:comment_created", string(body))
}

func TestBitbucketTrigger(t *testing.T) {
	TriggerPipeline("terraform_plan")
}
