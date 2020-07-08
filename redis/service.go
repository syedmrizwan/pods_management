package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// Define a custom struct to hold Album data.
type Album struct {
	Title  string
	Artist string
	Price  float64
	Likes  int
}

func TestRedis() {
	conn := GetPool().Get()
	defer conn.Close()

	album := Album{
		Title:  "Album123",
		Artist: "James11",
		Likes:  22,
		Price:  3333,
	}

	b, err := json.Marshal(&album)
	if err != nil {
		return
	}

	_, err = conn.Do("RPUSH", "mylist 13", string(b))
	if err != nil {
		log.Fatal(err)
	}

	album.Title = "ALBUM2"
	b, err = json.Marshal(&album)
	if err != nil {
		return
	}

	_, err = conn.Do("RPUSH", "mylist 13", string(b))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added")

	reply, _ := redis.ByteSlices(conn.Do("BLPOP", "mylist 13", 5))

	alb := &Album{}
	err = json.Unmarshal(reply[1], alb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(alb)
}
