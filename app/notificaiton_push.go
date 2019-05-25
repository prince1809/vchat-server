package app

type NotificationType string

type PushNotificationHub struct {
	Channels []chan PushNotification
}

type PushNotification struct {
	id               string
	notificationType NotificationType
	CurrentSessionId string
	userId           string
	channelId        string
}
