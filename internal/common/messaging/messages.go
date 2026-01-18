package messaging

import (
	"time"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/common/mimetypes"
)

type Type string

type Source string

type Message struct {
	Specversion     string             `json:"specversion"`
	Type            Type               `json:"type"`
	Source          Source             `json:"source"`
	Id              uuid.UUID          `json:"id"`
	Time            time.Time          `json:"timestamp"`
	Subject         string             `json:"subject"`
	Datacontenttype mimetypes.MimeType `json:"datacontenttype"`
	Data            any                `json:"data"`
}
