package cmd

import (
	"time"

	"github.com/itzngga/Lara/src/cmd/constant"
	"github.com/itzngga/Roxy/command"
	"github.com/itzngga/Roxy/embed"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func init() {
	embed.Commands.Add(clear)
}

var clear = &command.Command{
	Name:        "clear",
	Category:    constant.UTILITY_CATEGORY,
	Description: "Clear chat message.",
	RunFunc: func(ctx *command.RunFuncContext) *waProto.Message {
		var textReact string
		if len(ctx.Arguments) >= 1 {
			textReact = ctx.Arguments[0]
		}

		// var patchState = appstate.PatchInfo{
		// 	Timestamp: time.Now(),
		// 	Type: appstate.IndexDeleteChat,
		// }
		// ctx.Client.SendAppState(patchState)

		id := ctx.MessageInfo.ID
		chat := ctx.MessageInfo.Chat
		sender := ctx.MessageInfo.Sender
		key := &waProto.MessageKey{
			FromMe:    proto.Bool(true),
			Id:        proto.String(id),
			RemoteJid: proto.String(chat.String()),
		}

		if !sender.IsEmpty() && sender.User != ctx.Client.Store.ID.String() {
			key.FromMe = proto.Bool(false)
			key.Participant = proto.String(sender.ToNonAD().String())
		}

		return &waProto.Message{
			ReactionMessage: &waProto.ReactionMessage{
				Key:               key,
				Text:              proto.String(textReact),
				SenderTimestampMs: proto.Int64(time.Now().UnixMilli()),
			},
		}
	},
}
