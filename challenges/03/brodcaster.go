package main

type Broadcaster struct {
	messagebuss chan int
	listeners   []chan int
}

func (b *Broadcaster) Register() <-chan int {
	var output = make(chan int)
	b.listeners = append(b.listeners, output)
	return output
}

func (b *Broadcaster) Send() chan<- int {
	go func() {
		message := <-b.messagebuss
		for _, listener := range b.listeners {
			listener <- message
		}
	}()
	return b.messagebuss
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{make(chan int), *new([]chan int)}
}
