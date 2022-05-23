package runtime

import (
	"github.com/Art-of-the-Living/gracious"
)

// ReadSignal is a message format for reading a QualitativeSignal from a Service
type ReadSignal struct {
	response chan gracious.QualitativeSignal
}

// WriteSignal is a message format for writing a QualitativeSignal to a Service
type WriteSignal struct {
	value    gracious.QualitativeSignal
	response chan bool
}

// A Provider is a Service component that provides a QualitativeSignal value to requesting basicListener
// instances via a call to "GetFirePattern". GetPattern returns the held pattern.
type Provider interface {
	GetPattern() gracious.QualitativeSignal
}

// A Listener is a Service component that gets a call to be updated during the listen phase.
// A Listener's pattern should be retrieved with GetPattern. This makes a Listener a type
// of Provider. A Listener pulls its value from a specified provider.
type Listener interface {
	Listen()
	GetPattern() gracious.QualitativeSignal
	SetProvider(provider Provider)
}

// A basicListener is a service component that accepts a signal from a Provider instance.
// A basicListener is initialized with a "Start()". A Service may be composed
// of one or more basicListener instances. A basicListener cannot be written too, it only
// reads from a Provider which can be configured via a ConfigureBroadcaster
type basicListener struct {
	heldPattern gracious.QualitativeSignal
	provider    Provider
	configure   chan Provider
}

// newBasicListener creates a new basicListener with the given name
func newBasicListener() *basicListener {
	l := basicListener{
		heldPattern: gracious.NewQualitativeSignal("void"),
		configure:   make(chan Provider),
	}
	return &l
}

func (l *basicListener) GetPattern() gracious.QualitativeSignal {
	return l.heldPattern
}

// Listen the basicListener Service
func (l *basicListener) Listen() {
	// If a request has come to reconfigure the provider then do so;
	// Because this request could come from outside the runtime, it should
	// be implemented as a channel read.
	select {
	case configure := <-l.configure:
		l.provider = configure
	default:
	}
	// Read from the configured Provider when present
	if l.provider != nil {
		l.heldPattern = l.provider.GetPattern()
	}
}

// SetProvider will write the new Provider to the configuration channel. The change will not take effect till the next
// call to Update to take effect.
func (l *basicListener) SetProvider(provider Provider) {
	select {
	case l.configure <- provider:
	default:
	}
}
