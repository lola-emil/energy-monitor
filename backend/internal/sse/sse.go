package sse

type SSEClient chan []byte

type SSEBroker struct {
	clients    map[SSEClient]bool
	register   chan SSEClient
	unregister chan SSEClient
	Broadcast  chan []byte
}

func NewSSEBroker() *SSEBroker {
	return &SSEBroker{
		clients:    make(map[SSEClient]bool),
		register:   make(chan SSEClient),
		unregister: make(chan SSEClient),
		Broadcast:  make(chan []byte),
	}
}

func (b *SSEBroker) Run() {
	for {
		select {
		case client := <-b.register:
			b.clients[client] = true

		case client := <-b.unregister:
			delete(b.clients, client)
			close(client)

		case msg := <-b.Broadcast:
			for client := range b.clients {
				select {
				case client <- msg:
				default:
					// drop slow clients
				}
			}
		}
	}
}
