package model

import "crypto/tls"

const (
	CONN_SECURITY_NONE     = ""
	CONN_SECURITY_PLAIN    = "PLAIN"
	CONN_SECURITY_TLS      = "TLS"
	CONN_SECURITY_STARTTLS = "STARTTLS"

	IMAGE_DRIVER_LOCAL = "local"
	IMAGE_DRIVER_S3    = "amazons3"
)

var ServerTLSSupportedCiphers = map[string]uint16{
	"TLS_RSA_WITH_RC4_128_SHA": tls.TLS_RSA_WITH_RC4_128_SHA,
}

type ServiceSettings struct {
	SiteURL *string `restricted:"true"`
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
	DriverName                  *string  `restricted:"true"`
	DataSource                  *string  `restricted:"true"`
	DataSourceReplicas          []string `restricted:"true"`
	DataSourceSearchReplicas    []string `restricted:"true"`
	MaxIdleConns                *int     `restricted:"true"`
	ConnMaxLifetimeMilliseconds *int     `restricted:"true"`
	MaxOpenConns                *int     `restricted:"true"`
	Trace                       *bool    `restricted:"true"`
	AtRestEncrypt               *string  `restricted:"true"`
	QueryTimeout                *int     `restricted:"true"`
}

type LogSettings struct {
}

type NotificationLogSettings struct {
}

type PasswordSettings struct {
}

type FileSettings struct {
	EnableFileAttachments   *bool
	EnableMobileUpload      *bool
	EnableMobileDownload    *bool
	MaxFileSize             *int64
	DriverName              *string `restricted:"true"`
	Directory               *string `restricted:"true"`
	EnablePublicLink        *bool
	PublicLinkSalt          *string
	InitialFont             *string
	AmazonS3AccessKeyId     *string `restricted:"true"`
	AmazonS3SecretAccessKey *string `restricted:"true"`
	AmazonS3Bucket          *string `restricted:"true"`
	AmazonS3Region          *string `restricted:"true"`
	AmazonS3Endpoint        *string `restricted:"true"`
	AmazonS3SSL             *bool   `restricted:"true"`
	AmazonS3SignV2          *bool   `restricted:"true"`
	AmazonS3SSE             *bool   `restricted:"true"`
	AmazonS3Trace           *bool   `restricted:"true"`
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
	Enable bool
}

type PluginSettings struct {
	Enable             *bool
	EnableUploads      *bool   `restricted:"true"`
	EnableHealthChecks *bool   `restricted:"true"`
	Directory          *string `restricted:"true"`
	ClientDirectory    *string `restricted:"true"`
	Plugins            map[string]map[string]interface{}
	PluginStates       map[string]*PluginState
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
	ServiceSettings         ServiceSettings
	TeamSettings            TeamSettings
	ClientRequirements      ClientRequirements
	SqlSettings             SqlSettings
	LogSettings             LogSettings
	NotificationLogSettings NotificationLogSettings
	PasswordSettings        PasswordSettings
	FileSettings            FileSettings
	EmailSettings           EmailSettings
	RateLimitSettings       RateLimitSettings
	PrivacySettings         PrivacySettings
	SupportSettings         SupportSettings
	AnnouncementSettings    AnnouncementSettings
	ThemeSettings           ThemeSettings
	GitLabSeSettings        SSOSettings
	GoogleSettings          SSOSettings
	Office365Settings       SSOSettings
	LdapSettings            LdapSettings
	ComplianceSettings      ComplianceSettings
	LocalizationSettings    LocalizationSettings
	SamlSettings            SamlSettings
	NativeAppSettings       NativeAppSettings
	ClusterSettings         ClusterSettings
	MetricsSettings         MetricsSettings
	ExperimentalSettings    ExperimentalSettings
	AnalyticsSettings       AnalyticsSettings
	ElasticSearchSettings   ElasticSearchSettings
	DataRetentionSettings   DataRetentionSettings
	MessageExportSettings   MessageExportSettings
	JobSettings             JobSettings
	PluginSettings          PluginSettings
	DisplaySettings         DisplaySettings
	ImageProxySettings      ImageProxySettings
}

func (o *Config) SetDefaults() {

}
