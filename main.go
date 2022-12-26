package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func haltOn(err error) {
	if err != nil {
		log.Fatal("Error here :>> ", err)
	}
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	haltOn(err)
	defer conn.Close()

	inter, err := conn.Do(
		"HMSET",
		"podcast:1",
		"title", "Tech Over Tea",
		"creater", "Brodie Reberstone",
		"category", "techology",
		"membership_fee", 9.99,
	)
	haltOn(err)
	fmt.Println("inter :>> ", inter)

	title, err := redis.String(conn.Do("HGET", "podcast:1", "title"))
	haltOn(err)
	fmt.Println("title :>> ", title)

	membershipFee, err := redis.Float64(conn.Do("HGET", "podcast:1", "membership_fee"))
	fmt.Println("membershipFee :>> ", membershipFee)
}
