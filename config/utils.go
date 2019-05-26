package config

import (
	"github.com/prince1809/vchat-server/model"
	"strings"
)

// fixConfig patches invalid or missing data in the configuration, returning true if changed.
func fixConfig(cfg *model.Config) bool {
	changed := false

	// Ensure SiteURL has no trailing slash.
	if strings.HasSuffix(*cfg.ServiceSettings.SiteURL, "/") {
		*cfg.ServiceSettings.SiteURL = strings.TrimRight(*cfg.ServiceSettings.SiteURL, "/")
		changed = true
	}

	// Ensure the directory for a local file store has a trailing slash.
	if *cfg.FileSettings.DriverName == model.IMAGE_DRIVER_LOCAL {
		if !strings.HasSuffix(*cfg.FileSettings.Directory, "/") {
			*cfg.FileSettings.Directory += "/"
			changed = true
		}
	}

	if FixInvalidLocales(cfg) {
		changed = true
	}
	return changed
}

// FixInvalidLocales checks and corrects the given config for invalid locale-related settings.
//
// Ideally, this function would be completely internal, but it's currently exposed to allow the cli
// to test the config change before allowing the save.
func FixInvalidLocales(cfg *model.Config) bool {
	return true
}
