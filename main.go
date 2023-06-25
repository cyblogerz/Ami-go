package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sess, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	fmt.Println("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	//Handler is going to handle incoming messages
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//The parameters are passed as pointers , so any change of variables within the function affect the values
		//check if the message posted is not by bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is Online")

	//We nned to make the bot to not exist as sonn as the code is executed
	sc := make(chan os.Signal, 1)
	//This is a channel of typr os.Signal
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//In Go, the <- symbol is used for both sending and receiving values on a channel, depending on its usage.
	//Here,<-sc is used as a receive operation on the sc channel.
	//It blocks the execution of the program until a signal is received on the sc channel.
	//Once a signal is received, the program continues its execution.

}
