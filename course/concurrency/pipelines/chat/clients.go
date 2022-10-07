package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type client struct {
	config ChatClientConfig
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

func StartClient(msgChan chan<- string, cn io.ReadWriteCloser, quit chan struct{}) (chan<- string, <-chan struct{}) {
	c := new(client)
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)

	done := make(chan struct{})

	go func(Cc *client) {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			msg := scanner.Text()
			log.Println("[CLIENT] StartClient received", msg)
			if strings.Contains(msg, CHAT_CONFIG_FLAG) {
				err := json.Unmarshal([]byte(strings.Trim(msg, CHAT_CONFIG_FLAG)), &Cc.config)
				if err != nil {
					log.Fatalf("[CLIENT][ERROR][json.Unmarshal] chat config, %v", err.Error())
				}
				continue
			}
			msgChan <- fmt.Sprintf(MSG_LAYOUT, c.config.Name, msg)
		}
		done <- struct{}{}
	}(c)

	c.writeMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()
		case <-done:
		}
	}()

	return c.wc, done
}

func (c *client) writeMonitor() {
	log.Println("[CLIENT][writeMonitor] Init")
	go func() {
		for msg := range c.wc {
			log.Println("[CLIENT][writeMonitor][msg] ===>", msg)
			_, err := c.Writer.WriteString(msg + "\n")
			if err != nil {
				log.Println("[CLIENT][ERROR][c.WriteString] on write msg", msg, err.Error())
			}
			err = c.Writer.Flush()
			if err != nil {
				log.Println("[CLIENT][ERROR][c.Flush] on flush", err.Error())
			}
		}
	}()
}
