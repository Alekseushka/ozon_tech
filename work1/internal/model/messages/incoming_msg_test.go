package messages

import (
	"testing"

	"flonsay/workspace_go/tech/ozon_tech/work1/internal/mocks/messages"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OnCommand_Start_Answer_HelloMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mock_messages.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(1))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 1,
	})

	assert.NoError(t, err)
}

func Test_OnCommand_Unknown_Answer_UnknownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mock_messages.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("unknown command", int64(1))
	
	err := model.IncomingMessage(Message{
		Text:	"bubu",
		UserID: 1,
	})
	
	assert.NoError(t, err)
}