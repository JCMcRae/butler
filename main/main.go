package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Token string
var TestingChannelID string = "913848153713831964"
var GeneralTalkChannelID = "774000433751392294"
var WelcomeCenterChannelID = "913846519877234724"
var TimeFormat = "RFC822"

func main() {
	//Create a new session using the Bot token.
	//TODO: Move environment variables to a folder outside of any project for universal access.
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	Token = os.Getenv("DISCORD_BUTLER_TOKEN")
	fmt.Println(Token)
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating the Discord session:\n", err)
		return
	}

	//persistence :=

	//Register messageCreate function as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	//Let's get all of our intents clear.
	dg.Identify.Intents = discordgo.IntentsAll

	//Open a websocket connection to Discord and start listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening the connection:\n", err)
		return
	}

	_, err = dg.ChannelMessageSend(TestingChannelID, "Butler is running.")
	//greetUser(dg)

	if err != nil {
		fmt.Println("Error sending the message:\n", err)
	}

	//greetAll(dg)

	//Wait here until CTRL-C or other team signal is received.
	fmt.Println("CYSTech BTLR is running. Press CTRL-C to quit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	//Safely close the Discord connection.
	closeBot(dg)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == session.State.User.ID {
		//We may want to log what the Bot is saying though, just for records.
		return
	}

	if message.Content == "help" {
		userName := message.Author.Username
		response := "Hi " + userName + ". Welcome to The Playground."

		_, err := session.ChannelMessageSend(message.ChannelID, response)
		if err != nil {
			fmt.Println("Error sending the message:\n", err)
		}
	}

	if message.Content == "!quit" {
		//Only I should be able to press shut down the bot.
		if isAdmin(message) {
			closeBot(session)
		} else {

		}
	}
}

func closeBot(session *discordgo.Session) {
	_, err := session.ChannelMessageSend(TestingChannelID, "BTLR is shutting down. Farewell.")
	if err != nil {
		fmt.Println("Error sending the message:\n", err)
	}

	farewell(session)
	session.Close()
}

//func greetUser(session *discordgo.Session) {
//	client := session.Client
//
//}

func farewell(session *discordgo.Session) {

	sneakerCount := "You currently have an unknown amount of sneakers in your collection."
	transactions := "You have no new transactions in RSLL."

	content := "Farewell, Jordan.\n" + sneakerCount + "\n" + transactions

	_, err := session.ChannelMessageSend(TestingChannelID, content)
	if err != nil {
		fmt.Println("Error sending the message:\n", err)
	}
}

func greetAll(session *discordgo.Session) {
	content := "Hi! I'm Butler. I am one of the bots here. Welcome to CYS Playground! This is the exclusive Discord community server for our most engaged and active community members. This purpose of this space is to  "

	channel := discordgo.Channel{ID: WelcomeCenterChannelID}
	channel.Mention()
	_, err := session.ChannelMessageSend(WelcomeCenterChannelID, content)
	if err != nil {
		fmt.Println("Error sending the message:\n", err)
	}
}

func isAdmin(message *discordgo.MessageCreate) (isAdmin bool) {
	adminID := os.Getenv("DISCORD_MASTER_ADMIN_ID")
	isAdmin = message.Author.ID == adminID
	return
}

func greetNewMember() {

}
