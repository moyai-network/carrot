package carrot

import (
	"github.com/df-mc/atomic"
	"time"
)

// TagFunc is a function that is called when a Tag is set.
type TagFunc func(t *Tag)

// Tag is a Tag that can be used to cancel a task.
type Tag struct {
	expiration atomic.Value[time.Time]

	Tag   TagFunc
	unTag TagFunc

	c chan struct{}
}

// NewTag returns a new Tag.
func NewTag(t TagFunc, f TagFunc) *Tag {
	return &Tag{
		Tag:   t,
		unTag: f,

		c: make(chan struct{}),
	}
}

// Active returns true if the Tag is active.
func (t *Tag) Active() bool {
	return t.expiration.Load().After(time.Now())
}

// Remaining returns the remaining time of the Tag.
func (t *Tag) Remaining() time.Duration {
	return time.Until(t.expiration.Load())
}

// Set adds a duration to the Tag.
func (t *Tag) Set(d time.Duration) {
	if t.Tag != nil {
		t.Tag(t)
	}

	if t.Active() {
		t.Cancel()
	}
	t.c = make(chan struct{})

	go func() {
		select {
		case <-time.After(d):
			if t.unTag != nil {
				t.unTag(t)
			}
		case <-t.c:
			return
		}
	}()
	t.expiration.Store(time.Now().Add(d))
}

// Reset resets the Tag.
func (t *Tag) Reset() {
	if t.unTag != nil {
		t.unTag(t)
	}

	if t.Active() {
		t.Cancel()
	}
	t.expiration.Store(time.Time{})
}

// C returns the channel of the Tag.
func (t *Tag) C() <-chan struct{} {
	return t.c
}

// Cancel cancels the Tag.
func (t *Tag) Cancel() {
	t.expiration.Store(time.Time{})
	close(t.c)
}
