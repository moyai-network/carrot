package carrot

import (
	"github.com/df-mc/atomic"
	"github.com/rcrowley/go-bson"
	"time"
)

type CoolDownFunc func(cd *CoolDown)

// CoolDown represents a cool-down.
type CoolDown struct {
	expiration atomic.Value[time.Time]

	set   CoolDownFunc
	unSet CoolDownFunc

	c chan struct{}
}

// NewCoolDown returns a new cool-down.
func NewCoolDown(set, unset CoolDownFunc) *CoolDown {
	return &CoolDown{
		set:   set,
		unSet: unset,

		c: make(chan struct{}, 0),
	}
}

// Active returns true if the cool-down is active.
func (c *CoolDown) Active() bool {
	return c.expiration.Load().After(time.Now())
}

// Set sets the cool-down.
func (c *CoolDown) Set(d time.Duration) {
	if c.Active() {
		c.Cancel()
	}
	c.c = make(chan struct{}, 0)

	if c.set != nil {
		c.set(c)
	}

	go func() {
		select {
		case <-time.After(d):
			if c.unSet != nil {
				c.unSet(c)
			}
		case <-c.c:
			return
		}
	}()

	c.expiration.Store(time.Now().Add(d))
}

// Reset resets the cool-down.
func (c *CoolDown) Reset() {
	if c.Active() {
		c.Cancel()
	}

	if c.unSet != nil {
		c.unSet(c)
	}
	c.expiration.Store(time.Time{})
}

// Remaining returns the remaining time of the cool-down.
func (c *CoolDown) Remaining() time.Duration {
	return time.Until(c.expiration.Load())
}

// Cancel cancels the Tag.
func (c *CoolDown) Cancel() {
	c.expiration.Store(time.Time{})
	close(c.c)
}

type coolDownData struct {
	Duration time.Duration
}

// UnmarshalBSON ...
func (c *CoolDown) UnmarshalBSON(b []byte) error {
	d := coolDownData{}
	err := bson.Unmarshal(b, &d)
	c.expiration = *atomic.NewValue(time.Now().Add(d.Duration))
	return err
}

// MarshalBSON ...
func (c *CoolDown) MarshalBSON() ([]byte, error) {
	d := coolDownData{Duration: time.Until(c.expiration.Load())}
	return bson.Marshal(d)
}

// MappedCoolDown represents a cool-down mapped to a key.
type MappedCoolDown[T comparable] map[T]*CoolDown

// NewMappedCoolDown returns a new mapped cool-down.
func NewMappedCoolDown[T comparable]() MappedCoolDown[T] {
	return make(map[T]*CoolDown)
}

// Active returns true if the cool-down is active.
func (m MappedCoolDown[T]) Active(key T) bool {
	coolDown, ok := m[key]
	return ok && coolDown.Active()
}

// Set sets the cool-down.
func (m MappedCoolDown[T]) Set(key T, d time.Duration) {
	coolDown := m.Key(key)
	coolDown.Set(d)
	m[key] = coolDown
}

// Key returns the cool-down for the key.
func (m MappedCoolDown[T]) Key(key T) *CoolDown {
	coolDown, ok := m[key]
	if !ok {
		newCD := &CoolDown{}
		m[key] = newCD
		return newCD
	}
	return coolDown
}

// Reset resets the cool-down.
func (m MappedCoolDown[T]) Reset(key T) {
	delete(m, key)
}

// Remaining returns the remaining time of the cool-down.
func (m MappedCoolDown[T]) Remaining(key T) time.Duration {
	coolDown, ok := m[key]
	if !ok {
		return 0
	}
	return coolDown.Remaining()
}

// All returns all cool-downs.
func (m MappedCoolDown[T]) All() (coolDowns []*CoolDown) {
	for _, coolDown := range m {
		coolDowns = append(coolDowns, coolDown)
	}
	return coolDowns
}

// MarshalBSON ...
func (m MappedCoolDown[T]) MarshalBSON() ([]byte, error) {
	d := map[T]time.Time{}
	for k, cd := range m {
		d[k] = cd.expiration.Load()
	}
	return bson.Marshal(d)
}

// UnmarshalBSON ...
func (m MappedCoolDown[T]) UnmarshalBSON(b []byte) error {
	d := map[T]time.Time{}
	err := bson.Unmarshal(b, &d)
	if err != nil {
		return err
	}

	for k, cd := range d {
		m.Set(k, time.Until(cd))
	}
	return nil
}
