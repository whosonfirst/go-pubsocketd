package main

import (
       "encoding/json"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"strings"
)

type Message struct {
     Text string `json:text`
}

func main() {

	var url = flag.String("url", "ws://127.0.0.1:8080", "The websocket URL to connect to")
	var origin = flag.String("origin", "", "The origin header to send")

	flag.Parse()

	log.Printf("dialing %s...\n", *url)

	ws, err := websocket.Dial(*url, "", *origin)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to %s and ready to receive new messages\n", *url)

	for {
		var msg = make([]byte, 1024)
		_, err = ws.Read(msg)

		if err != nil {
			log.Fatal(err)
		}

		// https://groups.google.com/forum/#!msg/golang-nuts/77HJlZhWXpk/nyL4XKlnTkUJ
		
		s := string(msg)
		s = strings.Replace(s, "\x00", "", -1)
		b := []byte(s)
		
		var m Message
		err = json.Unmarshal(b, &m)

		if err != nil {
		   log.Println(err)
		   continue
		}
		
		fmt.Println(m.Text)				   
	}
}
