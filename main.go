// Package XuAlgoGo
// Time    : 2021/5/6 2:10 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"fmt"
	"github.com/Achillesxu/XuAlgoGo/models"
	"log"
)

const (
	example1 = "this is a example"
	example2 = "second example"
)

// main project entrance
func main() {
	dbCfg := models.NewDbConfig("mac", "workAholic!4", "tcp", "127.0.0.1:3306", "blog")
	err := models.OpenDb(dbCfg)
	if err != nil {
		panic(err)
	}
	albums, err := models.AlbumsByArtist("John Coltrane")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", albums)

	a, err := models.AlbumByID(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)

	aId, err := models.AddAlbum(models.Album{
		Title:  "My Favorite Things",
		Artist: "John Coltrane",
		Price:  19.88,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", aId)
}
