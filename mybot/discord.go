package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token = "TOKEN"
)

func ConnectToDiscord() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("MADE BY GH")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateGameStatus(0, "-help")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from bot himself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "-status" {
		handleStatus(s, m)
		return
	}
	if m.Content == "-help" {
		handleHelp(s, m)
		return
	}
	if m.Content == "-methods" {
		handleMethods(s, m)
		return
	}

	if m.Content == "-ovh" {
		handleOvh(s, m)
		return
	}

	if m.Content == "-hetzner" {
		handleHetzner(s, m)
		return
	}
	if m.Content == "-homekill" {
		handleHomekill(s, m)
		return
	}
	if m.Content == "-fivem" {
		handleFivem(s, m)
		return
	}
	if m.Content == "-cf-bypass" {
		handleCf(s, m)
		return
	}
	if m.Content == "-https" {
		handleHttps(s, m)
		return
	}
	if m.Content == "-100up" {
		handleUp(s, m)
		return
	}
	if m.Content == "-ark" {
		handleArk(s, m)
		return
	}
	if m.Content == "-minecraft" {
		handleMinecraft(s, m)
		return
	}
	if m.Content == "-tcp" {
		handleTcp(s, m)
		return
	}
	if m.Content == "-udp" {
		handleTcp(s, m)
		return
	}
	if m.Content == "-ssh" {
		handleTcp(s, m)
		return
	}
	if m.Content == "-nfo" {
		handleNfo(s, m)
		return
	}
	if m.Content == "-game" {
		handleGame(s, m)
		return
	}
	if m.Content == "-gamevip" {
		handleGame(s, m)
		return
	}
	if m.Content == "-csgovip" {
		handleCsgo(s, m)
		return
	}
	if m.Content == "-codrape" {
		handleCod(s, m)
		return
	}
	if m.Content == "-mv-slap" {
		handleMv(s, m)
		return
	}

}

func handleStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "SERVER STATUS :\n "
	msg = "\n"
	msg = "\n"
	msg = "\n"

	msg = fmt.Sprintf("%s\t%s\n", msg, "server is alive ✅")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "commands:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "> -help")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> -status")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> -methods")

	s.ChannelMessageSend(m.ChannelID, msg)
}
func handleMethods(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "methods:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-ovh**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-nfo**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-https**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-cf-bypass**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-homekill**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-ssh**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-100up**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-hetzner**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-ark**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-fivem**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-minecraft**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-tcp**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-udp**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-https**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-game**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-gamevip**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-csgovip**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-codrape**")
	msg = fmt.Sprintf("%s\t%s\n", msg, "> **-mv-slap**")
	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleHomekill(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "homekill method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "for homes ip")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : 80**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleMinecraft(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "Minecraft method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "java bypass | for minecraft servers")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleOvh(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "ovh method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "ovh bypass")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}
func handleFivem(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "fivem / fivemvip method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "for fivem servers")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}
func handleHttps(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "https method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, " standard website flood | urls only")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : 443 , 80**")

	s.ChannelMessageSend(m.ChannelID, msg)
}
func handleUp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "100up method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "for servers that host by 100up")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}
func handleArk(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "ark method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "for ark servers")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleCf(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "cf-bypass method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "for websites protected by cloudflare")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port: 443**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleTcp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "tcp/tcpkill/tcpbypass method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "tcp base flood / high pps flood")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port: depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleUdp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "udp/udpkill/udpbypass method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "udp base flood / high pps flood")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleSsh(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "ssh method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "high connection flood")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port usually 22**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleNfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "nfo method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "nfo bypass      ")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "game method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "generic game attack")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleCsgo(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "game method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "csgo community server bypass")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : usually 27015**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleCod(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "game method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "call of duty freeze / crash")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleMv(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "game method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "call of duty modern warfare freeze / crash")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleHetzner(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "hetzner method:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "layer3 protocol | for servers that hosted on hetzne")
	msg = fmt.Sprintf("%s\t%s\n", msg, "**port : depend from the server**")

	s.ChannelMessageSend(m.ChannelID, msg)
}
