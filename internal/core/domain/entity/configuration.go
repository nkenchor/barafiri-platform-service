package entity

type Configuration struct {
	PlatformCountries         string `redis:"platform-countries" json:"platform-countries,omitempty"`
	PlatformCurrencies        string `redis:"platform-currencies" json:"platform-currencies,omitempty"`
	PlatformNotificationTypes string `redis:"platform-notification-types" json:"platform-notification-types,omitempty"`
	PlatformOtpTypes          string `redis:"platform-otp-types" json:"platform-otp-types,omitempty"`
	PlatformCategories        string `redis:"platform-categories" json:"platform-categories,omitempty"`
	PlatformIndustries        string `redis:"platform-industries" json:"platform-industries,omitempty"`
	PlatformTtlConfiguration  string `redis:"platform-ttl-configurations" json:"platform-ttl-configuration"`
}
