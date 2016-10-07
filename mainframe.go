package xsay

type Mainframe struct {
	bubble    *bubble
	collector *textCollector
	replica   *replica
}

func New(asciiArt string, defaultMessage string) *Mainframe {
	return &Mainframe{
		bubble:    newBubble(defaultMessage),
		collector: newTextCollector(),
		replica:   newReplica(asciiArt),
	}
}

func (m *Mainframe) Say() {
	ok, text := m.collector.receivedInput()
	if ok {
		m.bubble.printMessage(text)
	} else {
		m.bubble.printDefaultMessage()
	}
	m.replica.printAscii()
}
