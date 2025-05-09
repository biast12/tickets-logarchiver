package v1

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/channel/embed"
)

type Message struct {
	Author      User                 `json:"author"`
	Content     string               `json:"content"`
	Timestamp   time.Time            `json:"timestamp"`
	Embeds      []embed.Embed        `json:"embeds,omitempty"`
	Attachments []channel.Attachment `json:"attachments,omitempty"`
}
