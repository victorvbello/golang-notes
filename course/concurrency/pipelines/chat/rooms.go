package chat

import (
	"io"
	"log"
	"sync"
)

type room struct {
	name    string
	msgChan chan string
	clients map[chan<- string]struct{}
	quit    chan struct{}
	*sync.RWMutex
}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	log.Println("[ROOM] Received message", msg)
	for wc := range r.clients {
		go func(lwc chan<- string) {
			log.Printf("[ROOM][broadcastMsg]client[%v] => %s", lwc, msg)
			lwc <- msg
		}(wc)
	}

}

func (r *room) Run() {
	log.Println("[ROOM] Run chant room", r.name)
	go func() {
		for msg := range r.msgChan {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.msgChan, c, r.quit)
	r.clients[wc] = struct{}{}
	log.Printf("[ROOM] Add client to room %s(%d)\n", r.name, r.TotalClients())
	r.Unlock()

	// remove cliente when
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) RemoveClient(wc chan<- string) {
	log.Println("[ROOM] Remove client from room", r.name)
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()

	select {
	case <-r.quit:
		if len(r.clients) == 0 {
			close(r.msgChan)
		}
	default:
	}
}

func (r *room) TotalClients() int {
	return len(r.clients)
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		msgChan: make(chan string),
		clients: make(map[chan<- string]struct{}),
		quit:    make(chan struct{}),
		RWMutex: new(sync.RWMutex),
	}
	r.Run()
	return r
}
