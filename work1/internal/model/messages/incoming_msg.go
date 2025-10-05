package messages

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text	string
	UserID	int64
}

func (s *Model) IncomingMessage(msg Message) error {
	var err error
	if (msg.Text == "/start") {
		err = s.tgClient.SendMessage("hello", msg.UserID)
	} else {
		err = s.tgClient.SendMessage("unknown command", msg.UserID)
	}
	return err
}
