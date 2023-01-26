package config

import (
	"log"
	"os"

	"github.com/andygrunwald/go-jira"
)

var jiraClient *jira.Client

// InitJiraClient initializes the Jira client
func InitJiraClient() {
	jiraURL := os.Getenv("JIRA_URL")
	jiraUser := os.Getenv("JIRA_USER")
	jiraToken := os.Getenv("JIRA_TOKEN")

	tp := jira.BasicAuthTransport{
		Username: jiraUser,
		Password: jiraToken,
	}

	client, err := jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		log.Fatal(err)
	}

	jiraClient = client
}

// GetJiraClient returns the Jira client
func GetJiraClient() *jira.Client {
	return jiraClient
}
