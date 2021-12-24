## Channels

This code show how to use channel control flow

```go
package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	fmt.Println("send 1")
	time.Sleep(1 * time.Second)

	c <- 42
	fmt.Println("send 2")
}

func receive(c <-chan int) {
	fmt.Println("receive 1")
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	fmt.Println("receive 2")
}

func main() {
	c := make(chan int)

	go send(c) // func send change the channel type to (chan<- int), this only permit push value to the channel

	receive(c) // func receive change the channel type to (<-chan int), this only permit pull value from the channel

	fmt.Println("end")
}

```
**Code:** https://go.dev/play/p/jXM_D4IpCam

---

This code show how to implement fan in pattern using channels

```go
package main

import (
	"fmt"
	"time"
)

type Notify interface {
	Send()
}

type MailNotify struct {
	Email string
	Body  string
}

func (mN *MailNotify) Send() {
	fmt.Printf("Sending mail to: <%s>, %s\n", mN.Email, mN.Body)
}

type SmsNotify struct {
	Tfl string
	Msg string
}

func (mN *SmsNotify) Send() {
	fmt.Printf("Sending sms to: <%s>, %s\n", mN.Tfl, mN.Msg)
}

// fan-in function
func sendNotification(cMail <-chan MailNotify, cSms <-chan SmsNotify) <-chan Notify {
	nc := make(chan Notify)
	go func() {
		for {
			var n MailNotify
			n = <-cMail
			nc <- &n
		}
	}()

	go func() {
		for {
			var n SmsNotify
			n = <-cSms
			nc <- &n
		}
	}()
	return nc
}

func sendRandomEmail() <-chan MailNotify {
	nMc := make(chan MailNotify)
	go func() {
		for i := 1; ; i++ {
			nMc <- MailNotify{
				Email: fmt.Sprintf("test-email-%d@test.com", i),
				Body:  fmt.Sprintf("Hi this is a test body for email %d", i),
			}
			time.Sleep(time.Duration(800 * time.Millisecond))
		}
		close(nMc)
	}()
	return nMc
}

func sendRandomSms() <-chan SmsNotify {
	nSc := make(chan SmsNotify)
	go func() {
		for i := 1; ; i++ {
			nSc <- SmsNotify{
				Tfl: fmt.Sprintf("111-11-111-%d", i),
				Msg: fmt.Sprintf("Hi this is a test msg for sms %d", i),
			}
			time.Sleep(time.Duration(500 * time.Millisecond))
		}
		close(nSc)
	}()
	return nSc
}

func main() {
	notifyC := sendNotification(sendRandomEmail(), sendRandomSms())
	for i := 0; i < 10; i++ {
		n := <-notifyC
		n.Send()
	}
}

```

**Code:** https://go.dev/play/p/TNZNawCu5N3