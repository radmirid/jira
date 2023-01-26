.env file

`
TELEGRAM_TOKEN=telegramToken

TELEGRAM_CHAT_ID=telegramChatId

JIRA_USERNAME=jiraUsername

JIRA_API_TOKEN=jiraApiToken

JIRA_ENDPOINT=https://jira-instance.atlassian.net

JIRA_PROJECT_KEY=jira_project_key

DB_URL=postgres://username:password@host:port/dbname

DB_MIGRATIONS_DIR=migrations
`

### Usage

/create

/list

`
docker-compose --file ./docker-compose.yml up --build
`