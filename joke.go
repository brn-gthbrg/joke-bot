package main

// import (
// 	"fmt"

// 	"github.com/icelain/jokeapi"
// )

// type JokesResp struct {
// 	Error    bool
// 	Category string
// 	JokeType string
// 	Joke     []string
// 	Flags    map[string]bool
// 	Id       float64
// 	Lang     string
// }

// func printJoke(j JokesResp) string {
// 	fmt.Println(JokesResp)
// }

// func jokeGenerator() {
// 	jt := "single"
// 	blacklist := []string{"nsfw"}
// 	ctgs := []string{"Programming", "Dark"}

// 	api := jokeapi.New()
// 	api.SetParams(&ctgs, &blacklist, &jt)
// 	response := api.Fetch()
// }
