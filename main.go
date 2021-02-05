package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"gopkg.in/yaml.v2"
)

func doEvery(ctx context.Context, d time.Duration, f func(time.Time)) error {
	ticker := time.Tick(d)
	fmt.Println("Ticker Done")
	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		case x := <- ticker:
			f(x)
		}
	}
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	// defer cancel() // This cancel doesn't really do anything, but you could register a signal channel that could result in canceling the context.
	// doEvery(ctx, 15 * time.Second, helloworld)
	confFile, err := os.Open("config.yml")
	if (err != nil) {
		panic(err.Error())
	}
	defer confFile.Close()

	// Deserialize the config from the file
	var cfg config
	decoder := yaml.NewDecoder(confFile)
	err = decoder.Decode(&cfg)
	if (err != nil) {
		panic(err.Error())
	}

	
}