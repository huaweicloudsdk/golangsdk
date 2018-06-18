package gophercloud

// AuthOptionsProvider presents the base of an auth options implementation
type AuthOptionsProvider interface {
	GetIdentityEndpoint() string
}
