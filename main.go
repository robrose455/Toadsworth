package main

import (

  "fmt"
  "github.com/bwmarrin/discordgo"

  )

const token string = "NzY2MjAyMDQ4NTIxNTY4Mjcw.X4f7Qw.MkAkx43xJbgrJBpD1Rq7S9o-rng"
var BotID string
func main() {

  g, err := discordgo.New("Bot " + token)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  u, err := g.User("@me")

  if err != nil {
    fmt.Println(err.Error())
  }

  BotID = u.ID

  err = g.Open()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("Bot is running!")

  <-make(chan struct{})
  return

}
