package options

// RouterOptions contains gin Router settings
type RouterOptions struct {
	Address string
	Port    string
}

// StatusOptions are responsible for containing settings for the status service
type StatusOptions struct {
	ConfigPath string
}

// RunOptions
type RunOptions struct {
	RouterOptions RouterOptions

	StatusOptions StatusOptions

	Parsed bool
}
