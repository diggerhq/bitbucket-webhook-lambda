package lambda

import (
	"encoding/json"
	"fmt"
	"github.com/diggerhq/go-bitbucket"
	"os"
)

import "time"

type Links struct {
	Self struct {
		Href string `json:"href"`
	} `json:"self"`
	Html struct {
		Href string `json:"href"`
	} `json:"html"`
	Avatar struct {
		Href string `json:"href"`
	} `json:"avatar"`
}

type Owner struct {
	DisplayName string `json:"display_name"`
	Links       `json:"links"`
	Type        string `json:"type"`
	Uuid        string `json:"uuid"`
	AccountId   string `json:"account_id"`
	Nickname    string `json:"nickname"`
}

type Workspace struct {
	Type  string `json:"type"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Links `json:"links"`
}

type Project struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Links `json:"links"`
}

type Repository struct {
	Type      string `json:"type"`
	FullName  string `json:"full_name"`
	Links     `json:"links"`
	Name      string      `json:"name"`
	Scm       string      `json:"scm"`
	Website   interface{} `json:"website"`
	Owner     `json:"owner"`
	Workspace `json:"workspace"`
	IsPrivate bool `json:"is_private"`
	Project   `json:"project"`
	Uuid      string `json:"uuid"`
}

type Actor struct {
	DisplayName string `json:"display_name"`
	Links       `json:"links"`
	Type        string `json:"type"`
	Uuid        string `json:"uuid"`
	AccountId   string `json:"account_id"`
	Nickname    string `json:"nickname"`
}

type Author struct {
	DisplayName string `json:"display_name"`
	Links       `json:"links"`
	Type        string `json:"type"`
	Uuid        string `json:"uuid"`
	AccountId   string `json:"account_id"`
	Nickname    string `json:"nickname"`
}

type Rendered struct {
	Title struct {
		Type   string `json:"type"`
		Raw    string `json:"raw"`
		Markup string `json:"markup"`
		Html   string `json:"html"`
	} `json:"title"`
	Description struct {
		Type   string `json:"type"`
		Raw    string `json:"raw"`
		Markup string `json:"markup"`
		Html   string `json:"html"`
	} `json:"description"`
}

type Destination struct {
	Branch struct {
		Name string `json:"name"`
	} `json:"branch"`
	Commit struct {
		Type  string `json:"type"`
		Hash  string `json:"hash"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
	} `json:"commit"`
	Repository struct {
		Type     string `json:"type"`
		FullName string `json:"full_name"`
		Links    struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	} `json:"repository"`
}

type Source struct {
	Branch struct {
		Name string `json:"name"`
	} `json:"branch"`
	Commit struct {
		Type  string `json:"type"`
		Hash  string `json:"hash"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
	} `json:"commit"`
	Repository struct {
		Type     string `json:"type"`
		FullName string `json:"full_name"`
		Links    struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	} `json:"repository"`
}

type Participants []struct {
	Type string `json:"type"`
	User struct {
		DisplayName string `json:"display_name"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Type      string `json:"type"`
		Uuid      string `json:"uuid"`
		AccountId string `json:"account_id"`
		Nickname  string `json:"nickname"`
	} `json:"user"`
	Role           string    `json:"role"`
	Approved       bool      `json:"approved"`
	State          string    `json:"state"`
	ParticipatedOn time.Time `json:"participated_on"`
}

type PullRequest struct {
	CommentCount      int    `json:"comment_count"`
	TaskCount         int    `json:"task_count"`
	Type              string `json:"type"`
	Id                int    `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	Rendered          `json:"rendered"`
	State             string      `json:"state"`
	MergeCommit       interface{} `json:"merge_commit"`
	CloseSourceBranch bool        `json:"close_source_branch"`
	ClosedBy          interface{} `json:"closed_by"`
	Author            `json:"author"`
	Reason            string    `json:"reason"`
	CreatedOn         time.Time `json:"created_on"`
	UpdatedOn         time.Time `json:"updated_on"`
	Destination       `json:"destination"`
	Source            `json:"source"`
	Reviewers         []interface{} `json:"reviewers"`
	Participants      `json:"participants"`
	Links             `json:"links"`
	Summary           struct {
		Type   string `json:"type"`
		Raw    string `json:"raw"`
		Markup string `json:"markup"`
		Html   string `json:"html"`
	} `json:"summary"`
}

type User struct {
	DisplayName string `json:"display_name"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	Type      string `json:"type"`
	Uuid      string `json:"uuid"`
	AccountId string `json:"account_id"`
	Nickname  string `json:"nickname"`
}

type Comment struct {
	Id        int       `json:"id"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Content   struct {
		Type   string `json:"type"`
		Raw    string `json:"raw"`
		Markup string `json:"markup"`
		Html   string `json:"html"`
	} `json:"content"`
	User        `json:"user"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	Links       `json:"links"`
	PullRequest `json:"pullrequest"`
}

type Approval struct {
	Date time.Time `json:"date"`
	User `json:"user"`
}

type Changes []struct {
	Old interface{} `json:"old"`
	New struct {
		Name   string `json:"name"`
		Target struct {
			Type   string    `json:"type"`
			Hash   string    `json:"hash"`
			Date   time.Time `json:"date"`
			Author struct {
				Type string `json:"type"`
				Raw  string `json:"raw"`
				User struct {
					DisplayName string `json:"display_name"`
					Links       struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
						Html struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
					Type      string `json:"type"`
					Uuid      string `json:"uuid"`
					AccountId string `json:"account_id"`
					Nickname  string `json:"nickname"`
				} `json:"user"`
			} `json:"author"`
			Message string `json:"message"`
			Summary struct {
				Type   string `json:"type"`
				Raw    string `json:"raw"`
				Markup string `json:"markup"`
				Html   string `json:"html"`
			} `json:"summary"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Parents []struct {
				Type  string `json:"type"`
				Hash  string `json:"hash"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Html struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"parents"`
			Rendered struct {
			} `json:"rendered"`
			Properties struct {
			} `json:"properties"`
		} `json:"target"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Type                 string   `json:"type"`
		MergeStrategies      []string `json:"merge_strategies"`
		DefaultMergeStrategy string   `json:"default_merge_strategy"`
	} `json:"new"`
	Truncated bool `json:"truncated"`
	Created   bool `json:"created"`
	Forced    bool `json:"forced"`
	Closed    bool `json:"closed"`
	Links     struct {
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	Commits []struct {
		Type   string    `json:"type"`
		Hash   string    `json:"hash"`
		Date   time.Time `json:"date"`
		Author struct {
			Type string `json:"type"`
			Raw  string `json:"raw"`
			User struct {
				DisplayName string `json:"display_name"`
				Links       `json:"links"`
				Type        string `json:"type"`
				Uuid        string `json:"uuid"`
				AccountId   string `json:"account_id"`
				Nickname    string `json:"nickname"`
			} `json:"user"`
		} `json:"author"`
		Message string `json:"message"`
		Summary struct {
			Type   string `json:"type"`
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			Html   string `json:"html"`
		} `json:"summary"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Diff struct {
				Href string `json:"href"`
			} `json:"diff"`
			Approve struct {
				Href string `json:"href"`
			} `json:"approve"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
			Patch struct {
				Href string `json:"href"`
			} `json:"patch,omitempty"`
		} `json:"links"`
		Parents []struct {
			Type  string `json:"type"`
			Hash  string `json:"hash"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"parents"`
		Rendered struct {
		} `json:"rendered"`
		Properties struct {
		} `json:"properties"`
	} `json:"commits"`
}

type Push struct {
	Changes `json:"changes"`
}

type BitbucketPullRequestCommentCreated struct {
	Repository  `json:"repository"`
	Actor       `json:"actor"`
	PullRequest `json:"pullrequest"`
	Comment     `json:"comment"`
}

type BitbucketPullRequestApproved struct {
	Repository  `json:"repository"`
	Actor       `json:"actor"`
	PullRequest `json:"pullrequest"`
	Approval    `json:"approval"`
}

type BitbucketRepoPush struct {
	Push       `json:"push"`
	Repository `json:"repository"`
	Actor      `json:"actor"`
}

type BitbucketPullRequestCreated struct {
	Repository  `json:"repository"`
	Actor       `json:"actor"`
	PullRequest `json:"pullrequest"`
}

var requestBody = `{
  "target": {
    "ref_type": "branch",
    "type": "pipeline_ref_target",
    "ref_name": "master",
    "selector": {
      "type": "custom",
      "pattern": "terraform_plan"
    }
  },
  "variables": [
    {
      "key": "DEPLOYMENT_VARIABLE",
      "value": "us-west-1"
    },
    {
      "key": "AWS_ACCESS_KEY_ID",
      "value": "$AWS_ACCESS_KEY_ID",
      "secured": true
    },
    {
      "key": "AWS_SECRET_ACCESS_KEY",
      "value": "$AWS_SECRET_ACCESS_KEY",
      "secured": true
    }
  ]
}`

func TriggerPipeline(pipelineName string) {
	bitbucketUsername := os.Getenv("BITBUCKET_USERNAME")
	bitbucketOwner := os.Getenv("BITBUCKET_REPO_OWNER")
	bitbucketRepoSlug := os.Getenv("BITBUCKET_REPO_SLUG")

	c := bitbucket.NewBasicAuth(bitbucketUsername, os.Getenv("BITBUCKET_PASSWORD"))

	if os.Getenv("BITBUCKET_PASSWORD") != "" {
		fmt.Printf("BITBUCKET_PASSWORD value found\n")
	}

	opt := &bitbucket.PipelinesOptions{
		Owner:    bitbucketOwner,
		RepoSlug: bitbucketRepoSlug,
	}

	var bitbucketTrigerPipelineRequestBody bitbucket.BitbucketTrigerPipelineRequestBody
	bitbucketTrigerPipelineRequestBody.Target.Selector.Pattern = pipelineName

	err := json.Unmarshal([]byte(requestBody), &bitbucketTrigerPipelineRequestBody)
	if err != nil {
		fmt.Printf("error parsing BitbucketTrigerPipelineRequestBody: %v\n", err)
	}

	_, err = c.Repositories.Pipelines.TriggerPipeline(opt, &bitbucketTrigerPipelineRequestBody)
	if err != nil {
		println(err.Error())
	}
}
