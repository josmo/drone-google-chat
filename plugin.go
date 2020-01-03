package main

import (
	"fmt"
	"github.com/drone/drone-template-lib/template"
	"github.com/josmo/drone-google-chat/google-chat"
	"strings"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Message string
		Status  string
		Link    string
		Started int64
		Created int64
	}

	Config struct {
		Webhook         string
		Template        string
		Token           string
		Key             string
		ConversationKey string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {

	text := message(p.Repo, p.Build)
	if p.Config.Template != "" {
		txt, err := template.RenderTrim(p.Config.Template, p)
		if err != nil {
			return err
		}
		text = txt
	}

	client := google_chat.NewClient(p.Config.Webhook, p.Config.Key, p.Config.Token, p.Config.ConversationKey)

	return client.SendMessage(&google_chat.Message{text})
}

func message(repo Repo, build Build) string {
	return fmt.Sprintf("*%s* <%s|%s/%s#%s> (%s) by %s",
		build.Status,
		build.Link,
		repo.Owner,
		repo.Name,
		build.Commit[:8],
		build.Branch,
		build.Author,
	)
}

func fallback(repo Repo, build Build) string {
	return fmt.Sprintf("%s %s/%s#%s (%s) by %s",
		build.Status,
		repo.Owner,
		repo.Name,
		build.Commit[:8],
		build.Branch,
		build.Author,
	)
}

func color(build Build) string {
	switch build.Status {
	case "success":
		return "good"
	case "failure", "error", "killed":
		return "danger"
	default:
		return "warning"
	}
}

func prepend(prefix, s string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}
