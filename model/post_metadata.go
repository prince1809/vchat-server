package model

type PostMetadata struct {
	// Embeds holds information requirement to render content embedded in the post. The includes the OpenGraph metadata
	// for links in the post.
	Embeds []*PostEmbed `json:"embeds,omitempty"`

	// Emojis holds all custom emojis, used in the post or used in reaction to the post
	Emojis []*Emoji `json:"emojis,omitempty"`
}
