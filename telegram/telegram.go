package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/radmirid/jira/jira_structs"
)

// Client is a struct that contains a Telegram client
type Client struct {
	Client *tgbotapi.BotAPI
}

// NewClient returns a new Telegram client
func NewClient(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Client{Client: bot}, nil
}

// SendTaskList sends a list of tasks to the user through Telegram
func (c *Client) SendTaskList(chatID int64, tasks []jira_structs.Task) {
	for _, task := range tasks {
		msg := tgbotapi.NewMessage(chatID, task.Key+" - "+task.Summary+" - "+task.Status)
		_, err := c.Client.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

// SendMessage sends a message to the user through Telegram
func (c *Client) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := c.Client.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
