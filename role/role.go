package role

import (
	"github.com/moyai-network/carrot"
	"github.com/rcrowley/go-bson"
	"golang.org/x/exp/slices"
	"sort"
	"sync"
	"time"
)

var (
	// roles contains all registered carrot.Role implementations.
	roles []carrot.Role
	// rolesByName contains all registered carrot.Role implementations indexed by their name.
	rolesByName = map[string]carrot.Role{}
)

// All returns all registered roles.
func All() []carrot.Role {
	return roles
}

// Register registers a role to the roles list. The hierarchy of roles is determined by the order of registration.
func Register(role carrot.Role) {
	roles = append(roles, role)
	rolesByName[role.Name()] = role
}

// ByName returns the role with the given name. If no role with the given name is registered, the second return value
// is false.
func ByName(name string) (carrot.Role, bool) {
	role, ok := rolesByName[name]
	return role, ok
}

// Staff returns true if the role provided is a staff role.
func Staff(role carrot.Role) bool {
	return Tier(role) >= Tier(Mod{}) || role == Operator{}
}

// Tier returns the tier of a role based on its registration hierarchy.
func Tier(role carrot.Role) int {
	return slices.IndexFunc(roles, func(other carrot.Role) bool {
		return role == other
	})
}

type Roles struct {
	roleMu          sync.Mutex
	roles           []carrot.Role
	roleExpirations map[carrot.Role]time.Time
}

// NewRoles creates a new Roles instance.
func NewRoles(roles []carrot.Role, expirations map[carrot.Role]time.Time) *Roles {
	return &Roles{
		roles:           roles,
		roleExpirations: expirations,
	}
}

// Add adds a role to the manager's role list.
func (r *Roles) Add(ro carrot.Role) {
	r.checkExpiry()
	r.roleMu.Lock()
	r.roles = append(r.roles, ro)
	r.roleMu.Unlock()
	r.sortRoles()
}

// Remove removes a role from the manager's role list. Users are responsible for updating the highest role usages if
// changed.
func (r *Roles) Remove(ro carrot.Role) bool {
	r.checkExpiry()
	if _, ok := ro.(Default); ok {
		// You can't remove the default role.
		return false
	}

	r.roleMu.Lock()
	i := slices.IndexFunc(r.roles, func(other carrot.Role) bool {
		return ro == other
	})
	r.roles = slices.Delete(r.roles, i, i+1)
	delete(r.roleExpirations, ro)
	r.roleMu.Unlock()
	r.sortRoles()
	return true
}

// Contains returns true if the manager has any of the given roles. Users are responsible for updating the highest role
// usages if changed.
func (r *Roles) Contains(roles ...carrot.Role) bool {
	r.checkExpiry()
	r.roleMu.Lock()
	defer r.roleMu.Unlock()

	var actualRoles []carrot.Role
	for _, ro := range r.roles {
		r.propagateRoles(&actualRoles, ro)
	}

	for _, r := range roles {
		if i := slices.IndexFunc(actualRoles, func(other carrot.Role) bool {
			return r == other
		}); i >= 0 {
			return true
		}
	}
	return false
}

// Expiration returns the expiration time for a role. If the role does not expire, the second return value will be false.
func (r *Roles) Expiration(ro carrot.Role) (time.Time, bool) {
	r.checkExpiry()
	r.roleMu.Lock()
	defer r.roleMu.Unlock()
	e, ok := r.roleExpirations[ro]
	return e, ok
}

// Expire sets the expiration time for a role. If the role does not expire, the second return value will be false.
func (r *Roles) Expire(ro carrot.Role, t time.Time) {
	r.checkExpiry()
	r.roleMu.Lock()
	defer r.roleMu.Unlock()
	r.roleExpirations[ro] = t
}

// Highest returns the highest role the manager has, in terms of hierarchy.
func (r *Roles) Highest() carrot.Role {
	r.checkExpiry()
	r.roleMu.Lock()
	defer r.roleMu.Unlock()
	return r.roles[len(r.roles)-1]
}

// All returns the user's roles.
func (r *Roles) All() []carrot.Role {
	r.checkExpiry()
	r.roleMu.Lock()
	defer r.roleMu.Unlock()
	return append(make([]carrot.Role, 0, len(r.roles)), r.roles...)
}

type rolesData struct {
	Roles       []string
	Expirations map[string]time.Time
}

// MarshalBSON ...
func (r *Roles) MarshalBSON() ([]byte, error) {
	var d rolesData
	r.roleMu.Lock()
	defer r.roleMu.Unlock()

	for _, rl := range r.roles {
		e, _ := r.roleExpirations[rl]
		if !e.IsZero() {
			d.Expirations[rl.Name()] = e
		}
		d.Roles = append(d.Roles, rl.Name())
	}
	return bson.Marshal(d)
}

// UnmarshalBSON ...
func (r *Roles) UnmarshalBSON(b []byte) error {
	var d rolesData
	err := bson.Unmarshal(b, &d)

	rls := d.Roles
	for _, rl := range rls {
		ro, ok := ByName(rl)
		if ok {
			r.Add(ro)
			e, ok := d.Expirations[rl]
			if ok {
				r.Expire(ro, e)
			}
		}
	}
	return err
}

// propagateRoles propagates roles to the user's role list.
func (r *Roles) propagateRoles(actualRoles *[]carrot.Role, role carrot.Role) {
	*actualRoles = append(*actualRoles, role)
	if h, ok := role.(carrot.HeirRole); ok {
		r.propagateRoles(actualRoles, h.Inherits())
	}
}

// sortRoles sorts the roles in the user's role list.
func (r *Roles) sortRoles() {
	sort.SliceStable(r.roles, func(i, j int) bool {
		return Tier(r.roles[i]) < Tier(r.roles[j])
	})
}

// checkExpirations checks each role the user has and removes the expired ones.
func (r *Roles) checkExpiry() {
	r.roleMu.Lock()
	rl, expirations := r.roles, r.roleExpirations
	r.roleMu.Unlock()

	for _, ro := range rl {
		if t, ok := expirations[ro]; ok && time.Now().After(t) {
			r.Remove(ro)
		}
	}
}

func init() {
	Register(Operator{})
	Register(Default{})

	Register(Plus{})
	Register(Media{})
	Register(Famous{})
	Register(Partner{})

	Register(Mod{})
	Register(Admin{})
	Register(Manager{})
	Register(Owner{})
}
