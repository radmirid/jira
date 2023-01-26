package jira

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/radmirid/jira/jira_structs"
)

// Client is a struct that contains a Jira client
type Client struct {
	Client *jira.Client
}

// NewClient returns a new Jira client
func NewClient(url string, username string, password string) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client, err := jira.NewClient(tp.Client(), url)
	if err != nil {
		return nil, err
	}

	return &Client{Client: client}, nil
}

// CreateTask creates a new task in Jira
func (c *Client) CreateTask(projectKey string, summary string, priority string, assignee string) (*jira.Issue, error) {
	issue := &jira.Issue{
		Fields: &jira.IssueFields{
			Project: jira.Project{
				Key: projectKey,
			},
			Summary:     summary,
			Description: "",
			Assignee: &jira.User{
				Name: assignee,
			},
			Priority: &jira.Priority{
				Name: priority,
			},
		},
	}

	newIssue, _, err := c.Client.Issue.Create(issue)
	if err != nil {
		return nil, err
	}

	return newIssue, nil
}

// GetTaskByKey retrieves a task by its key in Jira
func (c *Client) GetTaskByKey(key string) (*jira.Issue, error) {
	issue, _, err := c.Client.Issue.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// GetTaskList returns a list of tasks where the user is the author
func GetTaskList(client *jira.Client, username string) ([]jira_structs.Task, error) {
	jql := "author = " + username

	issues, _, err := client.Issue.Search(context.Background(), jql, nil)
	if err != nil {
		return nil, err
	}

	var taskList []jira_structs.Task
	for _, issue := range issues {
		task := jira_structs.Task{
			Key:      issue.Key,
			Summary:  issue.Fields.Summary,
			Assignee: issue.Fields.Assignee.Name,
			Priority: issue.Fields.Priority.Name,
			Status:   issue.Fields.Status.Name,
		}
		taskList = append(taskList, task)
	}

	return taskList, nil
}
