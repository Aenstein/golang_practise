package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"time"

	"phonebook/logger"
	"phonebook/book"
)

func handleCommand(cmd func([]string, book.PhoneBook) error, args []string, phoneBook book.PhoneBook) {
	if err := cmd(args, phoneBook); err != nil {
		logger.Warn(err, "command failed")
	}
}

func doAdd(args []string, phonebook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("mising parameters for 'add' command. Use add name=number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalide format. Use: add name=number")
	}

	name, number := kv[0], kv[1]
	err := phonebook.Add(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Added an entry: %s -> %s", name, number))

	return nil
}

func doGet(args []string, phonebook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("mising parameters for 'get' command. Use get name")
	}

	name := args[0]

	numberData, err := phonebook.Get(name)
	if err != nil {
		return err
	}

	unixUpdateAt := time.Unix(numberData.LastUpdatAt, 0)

	logger.Info(
		fmt.Sprintf("Number for %s is %s (last updated at %s)\n",
		name,
		numberData.Number,
		unixUpdateAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doList(_ []string, phonebook book.PhoneBook) error {
	if len(phonebook) == 0 {
		logger.Info("phonebook is empty")
	} else {
		results := ""
		for name, number := range phonebook {
			results += fmt.Sprintf("%s -> %s\n", name, number.Number)
		}

		logger.Info(results)
	}

	return nil
}

func doUpdate(args []string, phonebook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("mising parameters for 'update' command. Use update name=number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalide format. Use: update name=number")
	}

	name, number := kv[0], kv[1]

	err := phonebook.Update(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Update an entry: %s -> %s", name, number))

	return nil
}

func doDelete(args []string, phonebook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("mising parameters for 'delete' command. Use delete name")
	}

	name := args[0]

	err := phonebook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Deleted entry for %s\n", name))

	return nil
}

func main() {
	phonebook := make(book.PhoneBook)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Phonebook")
	fmt.Println("Available commands: add, get, delete, update, list, exit")

	for {
		fmt.Println("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "add":
			handleCommand(doAdd, args, phonebook)
		case "get":
			handleCommand(doGet, args, phonebook)
		case "delete":
			handleCommand(doDelete, args, phonebook)
		case "update":
			handleCommand(doUpdate, args, phonebook)
		case "list":
			handleCommand(doList, args, phonebook)
		case "exit":
			logger.Info("Exiting phonebook...")
			return
		default:
			logger.Warn(errors.New("unsupported command. Try 'add', 'get', 'delete', 'update', 'list', or 'exit'"))
		}
	}
}
