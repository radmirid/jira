package main

import (
	"log"

	"github.com/andygrunwald/go-jira"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"github.com/user/project/config"
	"github.com/user/project/db"
	"github.com/user/project/jira_structs"
	"github.com/user/project/migrate"
)

func HandleCreateTask(update tgbotapi.Update, cfg *config.Config) {
	// Parse command arguments
	args := update.Message.CommandArguments()
	if len(args) < 4 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /create_task project summary priority assignee")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}
	project := args[0]
	summary := args[1]
	priority := args[2]
	assignee := args[3]

	// Create task in Jira
	task, err := jira.CreateTask(cfg.Jira, project, summary, priority, assignee)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while creating task in Jira.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Save task in database
	err = db.SaveTask(cfg.DB, task)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while saving task in database.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Task "+task.Key+" has been created.")
	msg.ReplyToMessageID = update.Message.MessageID
	cfg.Telegram.Send(msg)
}

func HandleUpdateTask(update tgbotapi.Update, cfg *config.Config) {
	// Parse command arguments
	args := update.Message.CommandArguments()
	if len(args) < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /update_task task_key status")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}
	key := args[0]
	status := args[1]

	// Retrieve task from database
	task, err := db.GetTask(cfg.DB, key)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while retrieving task from database.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Update task in Jira
	err = jira.UpdateTask(cfg.Jira, task, status)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while updating task in Jira.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Save updated task in database
	err = db.SaveTask(cfg.DB, task)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while saving updated task in database.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Task "+task.Key+" has been updated.")
	msg.ReplyToMessageID = update.Message.MessageID
	cfg.Telegram.Send(msg)
}

func HandleGetTask(update tgbotapi.Update, cfg *config.Config) {
	// Parse command arguments
	key := update.Message.CommandArguments()
	if len(key) < 1 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /get_task task_key")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Retrieve task from database
	task, err := db.GetTask(cfg.DB, key)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while retrieving task from database.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Retrieve task information from Jira
	jiraTask, err := jira.GetTask(cfg.Jira, task.Key)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while retrieving task information from Jira.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Build message with task information
	var message string
	message += "Task Key: " + jiraTask.Key + "\n"
	message += "Summary: " + jiraTask.Fields.Summary + "\n"
	message += "Status: " + jiraTask.Fields.Status.Name + "\n"
	message += "Assignee: " + jiraTask.Fields.Assignee.DisplayName + "\n"
	message += "Priority: " + jiraTask.Fields.Priority.Name + "\n"

	// Send message with task information
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	cfg.Telegram.Send(msg)
}

func HandleGetTaskList(update tgbotapi.Update, cfg *config.Config) {
	// Retrieve tasks from Jira where the user is the author
	jiraTasks, err := jira.GetTasksByAuthor(cfg.Jira, update.Message.From.UserName)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error occurred while retrieving tasks from Jira.")
		msg.ReplyToMessageID = update.Message.MessageID
		cfg.Telegram.Send(msg)
		return
	}

	// Build message with list of tasks
	var message string
	for _, jiraTask := range jiraTasks {
		message += "Task Key: " + jiraTask.Key + "\n"
		message += "Summary: " + jiraTask.Fields.Summary + "\n"
		message += "Status: " + jiraTask.Fields.Status.Name + "\n\n"
	}

	// Send message with list of tasks
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	cfg.Telegram.Send(msg)
}
