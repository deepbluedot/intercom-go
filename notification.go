package intercom

import (
	"encoding/json"
	"io"
)

// Notification is the object delivered to a webhook.
type Notification struct {
	ID               string        `json:"id,omitempty"`
	AppID            string        `json:"app_id"`
	CreatedAt        int64         `json:"created_at,omitempty"`
	Topic            string        `json:"topic,omitempty"`
	DeliveryAttempts int64         `json:"delivery_attempts,omitempty"`
	FirstSentAt      int64         `json:"first_sent_at,omitempty"`
	RawData          *Data         `json:"data,omitempty"`
	Conversation     *Conversation `json:"-"`
	User             *User         `json:"-"`
	Tag              *Tag          `json:"-"`
	Company          *Company      `json:"-"`
	Event            *Event        `json:"-"`
}

// Data is the data node of the notification.
type Data struct {
	Item json.RawMessage `json:"item,omitempty"`
}

// NewNotification parses a Notification from json read from an io.Reader.
// It may only contain partial objects (such as a single conversation part)
// depending on what is provided by the webhook.
func NewNotification(r io.Reader) (*Notification, error) {
	notification := &Notification{
		RawData: &Data{},
	}
	err := json.NewDecoder(r).Decode(notification)
	if err != nil {
		return nil, err
	}

	switch notification.Topic {
	case "conversation_part.tag.created":

		var binder struct {
			Conversation Conversation `json:"conversation,omitempty"`
			Tag          Tag          `json:"tag,omitempty"`
		}
		json.Unmarshal(notification.RawData.Item, &binder)

		notification.Conversation = &binder.Conversation
		notification.Tag = &binder.Tag

	case "conversation.admin.assigned",
		"conversation.admin.closed",
		"conversation.admin.noted",
		"conversation.admin.open.assigned",
		"conversation.admin.opened",
		"conversation.admin.replied",
		"conversation.admin.single.created",
		"conversation.admin.snoozed",
		"conversation.admin.unsnoozed",
		"conversation.priority.updated",
		"conversation.rating.added",
		"conversation.user.created",
		"conversation.user.replied":
		c := &Conversation{}
		json.Unmarshal(notification.RawData.Item, c)
		notification.Conversation = c

	case "contact.user.updated",
		"contact.lead.updated":
		u := &User{}
		json.Unmarshal(notification.RawData.Item, u)
		notification.User = u

	case "user.created",
		"user.deleted",
		"user.unsubscribed",
		"user.email.updated":
		u := &User{}
		json.Unmarshal(notification.RawData.Item, u)
		notification.User = u

	case "user.tag.created",
		"user.tag.deleted":
		t := &Tag{}
		json.Unmarshal(notification.RawData.Item, t)
		notification.Tag = t

	case "company.created":
		c := &Company{}
		json.Unmarshal(notification.RawData.Item, c)
		notification.Company = c

	case "event.created":
		e := &Event{}
		json.Unmarshal(notification.RawData.Item, e)
		notification.Event = e
	}
	return notification, nil
}
