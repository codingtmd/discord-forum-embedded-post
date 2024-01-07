package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gtuk/discordwebhook"
	"github.com/joho/godotenv"
)

func composeMessage(
	creator_userName string,
	game_name string,
	game_play_url string,
	game_intro string,
	game_language string,
	game_category string,
	game_cover_image_url string,
) *discordwebhook.Message {
	// author info
	author := discordwebhook.Author{
		Name: &creator_userName,
	}

	// field info
	var lang = "Lang"
	var inline = true
	field_1 := discordwebhook.Field{
		Name:   &lang,
		Value:  &game_language,
		Inline: &inline,
	}

	var category = "Category"
	field_2 := discordwebhook.Field{
		Name:   &category,
		Value:  &game_category,
		Inline: &inline,
	}

	var free_text = "👇 and 🕹️："
	var not_inline = false
	field_4 := discordwebhook.Field{
		Name:   &free_text,
		Value:  &game_play_url,
		Inline: &not_inline,
	}

	image := discordwebhook.Image{
		Url: &game_cover_image_url,
	}

	var text = "enjoy your daily life :smirk:"
	footer := discordwebhook.Footer{
		Text: &text,
	}

	var color = "16765952"
	embed := discordwebhook.Embed{
		Author:      &author,
		Title:       &game_name,
		Url:         &game_play_url,
		Description: &game_intro,
		Color:       &color,
		Fields:      &[]discordwebhook.Field{field_2, field_1, field_4},
		Image:       &image,
		Footer:      &footer,
	}

	// bot info
	var username = "dalaba"
	var content = "real-time update of creator community"

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
		Embeds:   &[]discordwebhook.Embed{embed},
	}

	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(b))

	return &message
}

func sendMsg(
	url string,
	message *discordwebhook.Message,
) {
	err := discordwebhook.SendMessage(url, *message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	webhook := os.Getenv("webhook_url")

	// showcase use game "Midnight Manor, A Killer Among Us" as example

	message := composeMessage(
		"fisher.lei@gmail.com",
		"Midnight Manor, A Killer Among Us",
		"https://diago-qavdnvfe5a-uc.a.run.app/#/game?gameId=ebc436cd-8a80-4d6f-a69a-2d8cf4e599ad",
		"In the midst of your presence in a castle, a murde",
		"zh",
		"Detective",
		"https://storage.googleapis.com/rpggo-game/rpggo-creator/022a0306-2f75-496f-9959-18d1d7c3b61a/1dcc778b-c156-4e09-ad93-edd61a13f0fb",
	)

	sendMsg(
		webhook,
		message,
	)

}
