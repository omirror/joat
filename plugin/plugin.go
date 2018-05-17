package plugin

type Plugin interface {
	BundleID() string
	Initialize()
}
