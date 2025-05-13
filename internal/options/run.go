package options

// RouterOptions contains gin Router settings
type RouterOptions struct {
	Address string
	Port    string
}

// RunOptions
type RunOptions struct {
	RouterOptions RouterOptions

	Parsed bool
}
