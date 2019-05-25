package model

const (
	POST_EMBED_IMAGE              PostEmbedType = "image"
	POST_EMBED_MESSAGE_ATTACHMENT PostEmbedType = "message_attachment"
	POST_EMBED_OPENGRAPH          PostEmbedType = "opengraph"
)

type PostEmbedType string

type PostEmbed struct {
	Type PostEmbedType `json:"type"`

	// The URL of the embedded content. Used for image and OpenGraph embeds.
	URL string `json:"url"`

	Data interface{} `json:"data,omitempty"`
}
