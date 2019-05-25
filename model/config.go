package model

import "crypto/tls"

const (
	CONN_SECURITY_NONE     = ""
	CONN_SECURITY_PLAIN    = "PLAIN"
	CONN_SECURITY_TLS      = "TLS"
	CONN_SECURITY_STARTTLS = "STARTTLS"
)

var ServerTLSSupportedCiphers = map[string]uint16{
	"TLS_RSA_WITH_RC4_128_SHA": tls.TLS_RSA_WITH_RC4_128_SHA,
}

type ServiceSettings struct {
}

type ClusterSettings struct {
}

type MetricsSettings struct {
}

type ExperimentalSettings struct {
}

type AnalyticsSettings struct {
}

type SSOSettings struct {
}

type SqlSettings struct {
}

type LogSettings struct {
}

type NotificationLogSettings struct {
}

type PasswordSettings struct {
}

type FileSettings struct {
}

type EmailSettings struct {
}

type RateLimitSettings struct {
}

type PrivacySettings struct {
}

type SupportSettings struct {
}

type AnnouncementSettings struct {
}

type ThemeSettings struct {
}

type TeamSettings struct {
}

type ClientRequirements struct {
}

type LdapSettings struct {
}

type ComplianceSettings struct {
}

type LocalizationSettings struct {
	DefaultServerLocale *string
	DefaultClientLocale *string
	AvailableLocales    *string
}

func (s *LocalizationSettings) SetDefaults() {
	if s.DefaultServerLocale == nil {
		s.DefaultServerLocale = NewString(DEFAULT_LOCALE)
	}

	if s.DefaultClientLocale == nil {
		s.DefaultClientLocale = NewString(DEFAULT_LOCALE)
	}

	if s.AvailableLocales == nil {
		s.AvailableLocales = NewString("")
	}
}

type SamlSettings struct {
	// Basic
	Enable                        *bool
	EnableSyncWithLdap            *bool
	EnableSyncWithLdapIncludeAuth *bool

	Verify  *bool
	Encrypt *bool
}
type NativeAppSettings struct {
}

type ElasticSearchSettings struct {
}

type DataRetentionSettings struct {
}

type JobSettings struct {

}

type PluginState struct {

}

type PluginSettings struct {

}

type GlobalRelayMessageExportSettings struct {

}

type MessageExportSettings struct {

}

type DisplaySettings struct {

}

type ImageProxySettings struct {

}

type ConfigFunc func() *Config

type Config struct {

}
