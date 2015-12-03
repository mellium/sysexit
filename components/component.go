package components

// A Component receives stanza's to handle on the given channel, and outputs
// them to the returned channel. This allows it to communicate with a router in
// process.
type Component interface {
	Start(in *chan xmpp.Stanza) (*chan xmpp.Stanza, error) // Dial starts the component and listens for gobs.
	String() string                                        // String returns the component's unique name.
}
