package lambda

var bitbucketPullRequestCommentCreatedTest = `{
  "repository": {
    "type": "repository",
    "full_name": "alexey-digger/digger-demo",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo"
      },
      "avatar": {
        "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
      }
    },
    "name": "digger-demo",
    "scm": "git",
    "website": null,
    "owner": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "workspace": {
      "type": "workspace",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "name": "Alexey Digger",
      "slug": "alexey-digger",
      "links": {
        "avatar": {
          "href": "https://bitbucket.org/workspaces/alexey-digger/avatar/?ts=1662992743"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/"
        },
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger"
        }
      }
    },
    "is_private": true,
    "project": {
      "type": "project",
      "key": "DIG",
      "uuid": "{1f09e243-2b09-4c74-8f4a-ae27b5d56076}",
      "name": "digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger/projects/DIG"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/workspace/projects/DIG"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/user/alexey-digger/projects/DIG/avatar/32?ts=1679494939"
        }
      }
    },
    "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
  },
  "actor": {
    "display_name": "Alexey Digger",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
      },
      "avatar": {
        "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
      },
      "html": {
        "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
      }
    },
    "type": "user",
    "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
    "account_id": "631f3963ce3e476e42ac02f8",
    "nickname": "Alexey Digger"
  },
  "pullrequest": {
    "comment_count": 9,
    "task_count": 0,
    "type": "pullrequest",
    "id": 2,
    "title": "readme.txt edited online with Bitbucket",
    "description": "",
    "rendered": {
      "title": {
        "type": "rendered",
        "raw": "readme.txt edited online with Bitbucket",
        "markup": "markdown",
        "html": "<p>readme.txt edited online with Bitbucket</p>"
      },
      "description": {
        "type": "rendered",
        "raw": "",
        "markup": "markdown",
        "html": ""
      }
    },
    "state": "OPEN",
    "merge_commit": null,
    "close_source_branch": false,
    "closed_by": null,
    "author": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "reason": "",
    "created_on": "2023-03-22T16:04:53.186343+00:00",
    "updated_on": "2023-03-23T10:39:59.380253+00:00",
    "destination": {
      "branch": {
        "name": "master"
      },
      "commit": {
        "type": "commit",
        "hash": "f8e2ccf68484",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f8e2ccf68484"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "source": {
      "branch": {
        "name": "test1"
      },
      "commit": {
        "type": "commit",
        "hash": "f55d60f54315",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f55d60f54315"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f55d60f54315"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "reviewers": [],
    "participants": [
      {
        "type": "participant",
        "user": {
          "display_name": "Alexey Digger",
          "links": {
            "self": {
              "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
            },
            "avatar": {
              "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
            },
            "html": {
              "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
            }
          },
          "type": "user",
          "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
          "account_id": "631f3963ce3e476e42ac02f8",
          "nickname": "Alexey Digger"
        },
        "role": "PARTICIPANT",
        "approved": true,
        "state": "approved",
        "participated_on": "2023-03-23T09:42:10.843306+00:00"
      }
    ],
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo/pull-requests/2"
      },
      "commits": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/commits"
      },
      "approve": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/approve"
      },
      "request-changes": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/request-changes"
      },
      "diff": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/alexey-digger/digger-demo:f55d60f54315%0Df8e2ccf68484?from_pullrequest_id=2&topic=true"
      },
      "diffstat": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diffstat/alexey-digger/digger-demo:f55d60f54315%0Df8e2ccf68484?from_pullrequest_id=2&topic=true"
      },
      "comments": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/comments"
      },
      "activity": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/activity"
      },
      "merge": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/merge"
      },
      "decline": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/decline"
      },
      "statuses": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/statuses"
      }
    },
    "summary": {
      "type": "rendered",
      "raw": "",
      "markup": "markdown",
      "html": ""
    }
  },
  "comment": {
    "id": 380817532,
    "created_on": "2023-03-23T10:39:59.380203+00:00",
    "updated_on": "2023-03-23T10:39:59.380253+00:00",
    "content": {
      "type": "rendered",
      "raw": "test comment",
      "markup": "markdown",
      "html": "<p>test comment</p>"
    },
    "user": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "deleted": false,
    "type": "pullrequest_comment",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/comments/380817532"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo/pull-requests/2/_/diff#comment-380817532"
      }
    },
    "pullrequest": {
      "type": "pullrequest",
      "id": 2,
      "title": "readme.txt edited online with Bitbucket",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/digger-demo/pull-requests/2"
        }
      }
    }
  }
}`

var bitbucketPullRequestApprovedTest = `{
  "repository": {
    "type": "repository",
    "full_name": "alexey-digger/digger-demo",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo"
      },
      "avatar": {
        "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
      }
    },
    "name": "digger-demo",
    "scm": "git",
    "website": null,
    "owner": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "workspace": {
      "type": "workspace",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "name": "Alexey Digger",
      "slug": "alexey-digger",
      "links": {
        "avatar": {
          "href": "https://bitbucket.org/workspaces/alexey-digger/avatar/?ts=1662992743"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/"
        },
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger"
        }
      }
    },
    "is_private": true,
    "project": {
      "type": "project",
      "key": "DIG",
      "uuid": "{1f09e243-2b09-4c74-8f4a-ae27b5d56076}",
      "name": "digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger/projects/DIG"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/workspace/projects/DIG"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/user/alexey-digger/projects/DIG/avatar/32?ts=1679494939"
        }
      }
    },
    "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
  },
  "actor": {
    "display_name": "Alexey Digger",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
      },
      "avatar": {
        "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
      },
      "html": {
        "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
      }
    },
    "type": "user",
    "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
    "account_id": "631f3963ce3e476e42ac02f8",
    "nickname": "Alexey Digger"
  },
  "pullrequest": {
    "comment_count": 8,
    "task_count": 0,
    "type": "pullrequest",
    "id": 2,
    "title": "readme.txt edited online with Bitbucket",
    "description": "",
    "rendered": {
      "title": {
        "type": "rendered",
        "raw": "readme.txt edited online with Bitbucket",
        "markup": "markdown",
        "html": "<p>readme.txt edited online with Bitbucket</p>"
      },
      "description": {
        "type": "rendered",
        "raw": "",
        "markup": "markdown",
        "html": ""
      }
    },
    "state": "OPEN",
    "merge_commit": null,
    "close_source_branch": false,
    "closed_by": null,
    "author": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "reason": "",
    "created_on": "2023-03-22T16:04:53.186343+00:00",
    "updated_on": "2023-03-23T09:42:05.854559+00:00",
    "destination": {
      "branch": {
        "name": "master"
      },
      "commit": {
        "type": "commit",
        "hash": "f8e2ccf68484",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f8e2ccf68484"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "source": {
      "branch": {
        "name": "test1"
      },
      "commit": {
        "type": "commit",
        "hash": "f55d60f54315",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f55d60f54315"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f55d60f54315"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "reviewers": [],
    "participants": [
      {
        "type": "participant",
        "user": {
          "display_name": "Alexey Digger",
          "links": {
            "self": {
              "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
            },
            "avatar": {
              "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
            },
            "html": {
              "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
            }
          },
          "type": "user",
          "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
          "account_id": "631f3963ce3e476e42ac02f8",
          "nickname": "Alexey Digger"
        },
        "role": "PARTICIPANT",
        "approved": true,
        "state": "approved",
        "participated_on": "2023-03-23T09:42:10.843306+00:00"
      }
    ],
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo/pull-requests/2"
      },
      "commits": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/commits"
      },
      "approve": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/approve"
      },
      "request-changes": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/request-changes"
      },
      "diff": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/alexey-digger/digger-demo:f55d60f54315%0Df8e2ccf68484?from_pullrequest_id=2&topic=true"
      },
      "diffstat": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diffstat/alexey-digger/digger-demo:f55d60f54315%0Df8e2ccf68484?from_pullrequest_id=2&topic=true"
      },
      "comments": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/comments"
      },
      "activity": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/activity"
      },
      "merge": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/merge"
      },
      "decline": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/decline"
      },
      "statuses": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/2/statuses"
      }
    },
    "summary": {
      "type": "rendered",
      "raw": "",
      "markup": "markdown",
      "html": ""
    }
  },
  "approval": {
    "date": "2023-03-23T09:42:10.843306+00:00",
    "user": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    }
  }
}`

var bitbucketRepoPushTest = `{
  "push": {
    "changes": [
      {
        "old": null,
        "new": {
          "name": "test_branch22",
          "target": {
            "type": "commit",
            "hash": "f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138",
            "date": "2023-03-22T16:18:08+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "bitbucket-pipelines.yml edited online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "bitbucket-pipelines.yml edited online with Bitbucket",
              "markup": "markdown",
              "html": "<p>bitbucket-pipelines.yml edited online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "cd45307435d357b4caf2740e2964ad277f94d869",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/cd45307435d357b4caf2740e2964ad277f94d869"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          },
          "links": {
            "self": {
              "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/refs/branches/test_branch22"
            },
            "commits": {
              "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commits/test_branch22"
            },
            "html": {
              "href": "https://bitbucket.org/alexey-digger/digger-demo/branch/test_branch22"
            }
          },
          "type": "branch",
          "merge_strategies": [
            "merge_commit",
            "squash",
            "fast_forward"
          ],
          "default_merge_strategy": "merge_commit"
        },
        "truncated": true,
        "created": true,
        "forced": false,
        "closed": false,
        "links": {
          "commits": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commits?include=f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/branch/test_branch22"
          }
        },
        "commits": [
          {
            "type": "commit",
            "hash": "f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138",
            "date": "2023-03-22T16:18:08+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "bitbucket-pipelines.yml edited online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "bitbucket-pipelines.yml edited online with Bitbucket",
              "markup": "markdown",
              "html": "<p>bitbucket-pipelines.yml edited online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              },
              "diff": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              },
              "approve": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138/approve"
              },
              "comments": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138/comments"
              },
              "statuses": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138/statuses"
              },
              "patch": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/patch/f8e2ccf68484c9428aa60b8f2c6ebb3a4acea138"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "cd45307435d357b4caf2740e2964ad277f94d869",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/cd45307435d357b4caf2740e2964ad277f94d869"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          },
          {
            "type": "commit",
            "hash": "cd45307435d357b4caf2740e2964ad277f94d869",
            "date": "2023-03-22T16:10:23+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "bitbucket-pipelines.yml edited online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "bitbucket-pipelines.yml edited online with Bitbucket",
              "markup": "markdown",
              "html": "<p>bitbucket-pipelines.yml edited online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/cd45307435d357b4caf2740e2964ad277f94d869"
              },
              "diff": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/cd45307435d357b4caf2740e2964ad277f94d869"
              },
              "approve": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869/approve"
              },
              "comments": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869/comments"
              },
              "statuses": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/cd45307435d357b4caf2740e2964ad277f94d869/statuses"
              },
              "patch": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/patch/cd45307435d357b4caf2740e2964ad277f94d869"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "5a3a6361cb30f96b22a8ed1c71393db080660cb2",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/5a3a6361cb30f96b22a8ed1c71393db080660cb2"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/5a3a6361cb30f96b22a8ed1c71393db080660cb2"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          },
          {
            "type": "commit",
            "hash": "5a3a6361cb30f96b22a8ed1c71393db080660cb2",
            "date": "2023-03-22T16:02:42+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "Merged in test1 (pull request #1)\n\nreadme.txt created online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "Merged in test1 (pull request #1)\n\nreadme.txt created online with Bitbucket",
              "markup": "markdown",
              "html": "<p>Merged in test1 (<a href=\"https://bitbucket.org/alexey-digger/digger-demo/pull-requests/1/readmetxt-created-online-with-bitbucket\" rel=\"nofollow\" class=\"ap-connect-link\">pull request #1</a>)</p>\n<p>readme.txt created online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/5a3a6361cb30f96b22a8ed1c71393db080660cb2"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/5a3a6361cb30f96b22a8ed1c71393db080660cb2"
              },
              "diff": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/5a3a6361cb30f96b22a8ed1c71393db080660cb2"
              },
              "approve": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/5a3a6361cb30f96b22a8ed1c71393db080660cb2/approve"
              },
              "comments": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/5a3a6361cb30f96b22a8ed1c71393db080660cb2/comments"
              },
              "statuses": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/5a3a6361cb30f96b22a8ed1c71393db080660cb2/statuses"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "ec4d61035e55f0a2405c5c761a538b1f0347cc9e",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
                  }
                }
              },
              {
                "type": "commit",
                "hash": "431d5860c6ca1546b9ec593e67d818ac80d5a99a",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          },
          {
            "type": "commit",
            "hash": "ec4d61035e55f0a2405c5c761a538b1f0347cc9e",
            "date": "2023-03-22T15:59:12+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "bitbucket-pipelines.yml edited online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "bitbucket-pipelines.yml edited online with Bitbucket",
              "markup": "markdown",
              "html": "<p>bitbucket-pipelines.yml edited online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
              },
              "diff": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
              },
              "approve": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/ec4d61035e55f0a2405c5c761a538b1f0347cc9e/approve"
              },
              "comments": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/ec4d61035e55f0a2405c5c761a538b1f0347cc9e/comments"
              },
              "statuses": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/ec4d61035e55f0a2405c5c761a538b1f0347cc9e/statuses"
              },
              "patch": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/patch/ec4d61035e55f0a2405c5c761a538b1f0347cc9e"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "f103ac8763b9dc3d327c7aced69f839346e0e4dd",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f103ac8763b9dc3d327c7aced69f839346e0e4dd"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f103ac8763b9dc3d327c7aced69f839346e0e4dd"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          },
          {
            "type": "commit",
            "hash": "431d5860c6ca1546b9ec593e67d818ac80d5a99a",
            "date": "2023-03-22T15:37:32+00:00",
            "author": {
              "type": "author",
              "raw": "Alexey Digger <alexey@digger.dev>",
              "user": {
                "display_name": "Alexey Digger",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
                  },
                  "avatar": {
                    "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
                  },
                  "html": {
                    "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
                  }
                },
                "type": "user",
                "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
                "account_id": "631f3963ce3e476e42ac02f8",
                "nickname": "Alexey Digger"
              }
            },
            "message": "readme.txt created online with Bitbucket",
            "summary": {
              "type": "rendered",
              "raw": "readme.txt created online with Bitbucket",
              "markup": "markdown",
              "html": "<p>readme.txt created online with Bitbucket</p>"
            },
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
              },
              "html": {
                "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
              },
              "diff": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
              },
              "approve": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/431d5860c6ca1546b9ec593e67d818ac80d5a99a/approve"
              },
              "comments": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/431d5860c6ca1546b9ec593e67d818ac80d5a99a/comments"
              },
              "statuses": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/431d5860c6ca1546b9ec593e67d818ac80d5a99a/statuses"
              },
              "patch": {
                "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/patch/431d5860c6ca1546b9ec593e67d818ac80d5a99a"
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "f103ac8763b9dc3d327c7aced69f839346e0e4dd",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f103ac8763b9dc3d327c7aced69f839346e0e4dd"
                  },
                  "html": {
                    "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f103ac8763b9dc3d327c7aced69f839346e0e4dd"
                  }
                }
              }
            ],
            "rendered": {},
            "properties": {}
          }
        ]
      }
    ]
  },
  "repository": {
    "type": "repository",
    "full_name": "alexey-digger/digger-demo",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo"
      },
      "avatar": {
        "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
      }
    },
    "name": "digger-demo",
    "scm": "git",
    "website": null,
    "owner": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "workspace": {
      "type": "workspace",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "name": "Alexey Digger",
      "slug": "alexey-digger",
      "links": {
        "avatar": {
          "href": "https://bitbucket.org/workspaces/alexey-digger/avatar/?ts=1662992743"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/"
        },
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger"
        }
      }
    },
    "is_private": true,
    "project": {
      "type": "project",
      "key": "DIG",
      "uuid": "{1f09e243-2b09-4c74-8f4a-ae27b5d56076}",
      "name": "digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger/projects/DIG"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/workspace/projects/DIG"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/user/alexey-digger/projects/DIG/avatar/32?ts=1679494939"
        }
      }
    },
    "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
  },
  "actor": {
    "display_name": "Alexey Digger",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
      },
      "avatar": {
        "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
      },
      "html": {
        "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
      }
    },
    "type": "user",
    "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
    "account_id": "631f3963ce3e476e42ac02f8",
    "nickname": "Alexey Digger"
  }
}`

var bitbucketPullRequestCreatedTest = `{
  "repository": {
    "type": "repository",
    "full_name": "alexey-digger/digger-demo",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo"
      },
      "avatar": {
        "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
      }
    },
    "name": "digger-demo",
    "scm": "git",
    "website": null,
    "owner": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "workspace": {
      "type": "workspace",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "name": "Alexey Digger",
      "slug": "alexey-digger",
      "links": {
        "avatar": {
          "href": "https://bitbucket.org/workspaces/alexey-digger/avatar/?ts=1662992743"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/"
        },
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger"
        }
      }
    },
    "is_private": true,
    "project": {
      "type": "project",
      "key": "DIG",
      "uuid": "{1f09e243-2b09-4c74-8f4a-ae27b5d56076}",
      "name": "digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/workspaces/alexey-digger/projects/DIG"
        },
        "html": {
          "href": "https://bitbucket.org/alexey-digger/workspace/projects/DIG"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/user/alexey-digger/projects/DIG/avatar/32?ts=1679494939"
        }
      }
    },
    "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
  },
  "actor": {
    "display_name": "Alexey Digger",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
      },
      "avatar": {
        "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
      },
      "html": {
        "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
      }
    },
    "type": "user",
    "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
    "account_id": "631f3963ce3e476e42ac02f8",
    "nickname": "Alexey Digger"
  },
  "pullrequest": {
    "comment_count": 0,
    "task_count": 0,
    "type": "pullrequest",
    "id": 3,
    "title": "readme.txt edited online with Bitbucket",
    "description": "",
    "rendered": {
      "title": {
        "type": "rendered",
        "raw": "readme.txt edited online with Bitbucket",
        "markup": "markdown",
        "html": "<p>readme.txt edited online with Bitbucket</p>"
      },
      "description": {
        "type": "rendered",
        "raw": "",
        "markup": "markdown",
        "html": ""
      }
    },
    "state": "OPEN",
    "merge_commit": null,
    "close_source_branch": false,
    "closed_by": null,
    "author": {
      "display_name": "Alexey Digger",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/%7B94bac741-62e2-4223-a45f-91773254971b%7D"
        },
        "avatar": {
          "href": "https://secure.gravatar.com/avatar/2fbee1042b15f82c532c9f52002cb2d1?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FAS-2.png"
        },
        "html": {
          "href": "https://bitbucket.org/%7B94bac741-62e2-4223-a45f-91773254971b%7D/"
        }
      },
      "type": "user",
      "uuid": "{94bac741-62e2-4223-a45f-91773254971b}",
      "account_id": "631f3963ce3e476e42ac02f8",
      "nickname": "Alexey Digger"
    },
    "reason": "",
    "created_on": "2023-03-23T11:05:28.605356+00:00",
    "updated_on": "2023-03-23T11:05:29.400610+00:00",
    "destination": {
      "branch": {
        "name": "master"
      },
      "commit": {
        "type": "commit",
        "hash": "f8e2ccf68484",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/f8e2ccf68484"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/f8e2ccf68484"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "source": {
      "branch": {
        "name": "test_branch22"
      },
      "commit": {
        "type": "commit",
        "hash": "3e9e3ca745ea",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/commit/3e9e3ca745ea"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo/commits/3e9e3ca745ea"
          }
        }
      },
      "repository": {
        "type": "repository",
        "full_name": "alexey-digger/digger-demo",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
          },
          "html": {
            "href": "https://bitbucket.org/alexey-digger/digger-demo"
          },
          "avatar": {
            "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
          }
        },
        "name": "digger-demo",
        "uuid": "{96fc099c-d073-424e-9109-1aae34b1dd08}"
      }
    },
    "reviewers": [],
    "participants": [],
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo/pull-requests/3"
      },
      "commits": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/commits"
      },
      "approve": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/approve"
      },
      "request-changes": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/request-changes"
      },
      "diff": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diff/alexey-digger/digger-demo:3e9e3ca745ea%0Df8e2ccf68484?from_pullrequest_id=3&topic=true"
      },
      "diffstat": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/diffstat/alexey-digger/digger-demo:3e9e3ca745ea%0Df8e2ccf68484?from_pullrequest_id=3&topic=true"
      },
      "comments": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/comments"
      },
      "activity": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/activity"
      },
      "merge": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/merge"
      },
      "decline": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/decline"
      },
      "statuses": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo/pullrequests/3/statuses"
      }
    },
    "summary": {
      "type": "rendered",
      "raw": "",
      "markup": "markdown",
      "html": ""
    }
  }
}`

var bitbucketUnknownEventTest = `{
  "aaaaaaa": {
    "type": "repository",
    "full_name": "alexey-digger/digger-demo",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/alexey-digger/digger-demo"
      },
      "html": {
        "href": "https://bitbucket.org/alexey-digger/digger-demo"
      },
      "avatar": {
        "href": "https://bytebucket.org/ravatar/%7B96fc099c-d073-424e-9109-1aae34b1dd08%7D?ts=default"
      }
    }
  }`
