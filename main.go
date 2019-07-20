package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Token string `yaml:"token"`
}

func (c *conf) getConf() (*conf, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall configuration file into yaml: %v", err)
	}

	return c, nil
}

func main() {

	// Token token for the Discord Bot as part of interacting with their API
	var c conf
	
	// read the configuration file, panic if it's not found
	configFile, err := c.getConf()
	if err != nil {
		panic(err)
	}
	
	// this should probably be set to lowercase since we don't usually want
	// to export things in the main package.
	Token := configFile.Token

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that are not prefixed by the bot invocation name
	// e.g. !triton

	if m.Content == "!triton" {
		s.ChannelMessageSendTTS(m.ChannelID, "Please type `!triton help` to see all the available commands")
		return
	}

	strTok := strings.SplitN(m.Content, " ", 2) // Split the string into 2 pieces at the first space
	if strTok[0] != "!triton" {
		s.ChannelMessageSendTTS(m.ChannelID, "Sorry. All interaction with Triton needs to be prefixed with '!triton'")
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "!triton ping" {
		payload := fmt.Sprintf("%s says %s, Triton says pong", m.Author.Username, m.Content)
		s.ChannelMessageSend(m.ChannelID, payload)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "!triton pong" {
		payload := fmt.Sprintf("%s says %s, Triton says ping", m.Author.Username, m.Content)
		s.ChannelMessageSend(m.ChannelID, payload)
	}
}
