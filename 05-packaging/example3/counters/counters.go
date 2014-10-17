// All material is licensed under the GNU Free Documentation License

// Package counters provides alert counter support.
package counters

// alertCounter is an unexported named type that
// contains an integer counter for alerts.
type alertCounter int

// New creates and returns objects of
// the unexported type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
