package store

import "github.com/prince1809/vchat-server/model"

type StoreResult struct {
	Data interface{}
	Err  *model.AppError
}

type StoreChannel chan StoreResult

type Store interface {
	Team() TeamStore
	Channel() ChannelStore
	Post() PostStore
	User() UserStore
	Bot() BotStore
	Audit() AuditStore
	ClusterDiscovery() ClusterDiscoveryStore
	Compliance() ComplianceStore
	Session() SessionStore
	OAuth() OAuthStore
	System() SystemStore
	Webhook() WebhookStore
	Command() CommandStore
	CommandWebhook() CommandWebhookStore
	Preference() PreferenceStore
	License() LicenseStore
	Token() TokenStore
	Emoji() EmojiStore
	Status() StatusStore
	FileInfo() FileInfoStore
	Reaction() ReactionStore
	Role() RoleStore
	Scheme() SchemeStore
	Job() JobStore
	UserAccessToken() UserAccessTokenStore
	ChannelMemberHistory() ChannelMemberHistoryStore
	Plugin() PluginStore
	TermsOfService() TermsOfServiceStore
	Group() GroupStore
	UserTermsOfService() UserAccessTokenStore
	LinkMetadata() LinkMetadataStore
	MarkSystemRanUnitTests()
	Close()
	LockToMaster()
	UnlockFromMaster()
	DropAllTables()
	TotalMasterDbConnections() int
	TotalReadDbConnections() int
	TotalSearchDbConnections() int
}

type TeamStore interface {
}

type ChannelStore interface {
}

type ChannelMemberHistoryStore interface {
}

type PostStore interface {
}

type UserStore interface {
}

type BotStore interface {
}

type SessionStore interface {
}

type AuditStore interface {
}

type ClusterDiscoveryStore interface {
}

type ComplianceStore interface {
}

type OAuthStore interface {
}

type SystemStore interface {
}

type WebhookStore interface {
}

type CommandStore interface {
}

type CommandWebhookStore interface {
}

type PreferenceStore interface {
}

type LicenseStore interface {
}

type TokenStore interface {
}

type EmojiStore interface {
}

type StatusStore interface {
}

type FileInfoStore interface {
}

type ReactionStore interface {
}

type JobStore interface {
}

type UserAccessTokenStore interface {
}

type PluginStore interface {
}

type RoleStore interface {
}

type SchemeStore interface {
}

type TermsOfServiceStore interface {
}

type UserTermsOfServiceStore interface {
}

type GroupStore interface {
}

type LinkMetadataStore interface {
}
