package jira_structs

type Task struct {
	ID       int    `json:"id"`
	Key      string `json:"key"`
	Summary  string `json:"summary"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Assignee string `json:"assignee"`
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Login   string `json:"login"`
	Token   string `json:"token"`
	JiraURL string `json:"jira_url"`
}
