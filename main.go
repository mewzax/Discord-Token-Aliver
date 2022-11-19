package main

import (
	"fmt"
	"bufio"
	"strings"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zenthangplus/goccm"
	"github.com/mewzax/discordgo"
)

func getTokens(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tokens []string
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tokens
}

func run(token string) {
  var Token string
  if strings.Contains(token, ":") {
	  Token = strings.Split(token, ":")[2]
  } else {
    Token = token
  }
	dg, err := discordgo.New(Token)
	err = dg.Open()
	if err != nil {
		log.Println(fmt.Sprintf("Error creation session for %s", Token))
		return
	}
	log.Println(fmt.Sprintf("Connected to %s#%s | %s", dg.State.User.Username, dg.State.User.Discriminator, Token))
}

func main() {
  // input tokens file and number of threads here
	tokens := getTokens("tokens.txt")
	threads := goccm.New(50)

    for _, token := range tokens {
        threads.Wait()
        go func(token string) {
			run(token)
			threads.Done()
		}(token)
    }
    threads.WaitAllDone()

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})
	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
