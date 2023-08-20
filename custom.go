package carrot

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	_ "unsafe"
)

// SendCustomSound sends a custom sound to the user, or it's viewers as well.
func SendCustomSound(p *player.Player, sound string, volume, pitch float64, public bool) {
	pos := p.Position()
	pk := &packet.PlaySound{
		SoundName: sound,
		Position:  vec64To32(pos),
		Volume:    float32(volume),
		Pitch:     float32(pitch),
	}

	v := []world.Viewer{player_session(p)}
	if public {
		v = viewers(p)
	}

	for _, v := range v {
		if s, ok := v.(*session.Session); ok {
			session_writePacket(s, pk)
		}
	}
}

// viewers returns a list of all viewers of the Player.
func viewers(p *player.Player) []world.Viewer {
	s := player_session(p)
	viewers := p.World().Viewers(p.Position())
	for _, v := range viewers {
		if v == s {
			return viewers
		}
	}
	return append(viewers, s)
}

// vec64To32 converts a mgl64.Vec3 to a mgl32.Vec3.
func vec64To32(vec3 mgl64.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{float32(vec3[0]), float32(vec3[1]), float32(vec3[2])}
}

// noinspection ALL
//
//go:linkname player_session github.com/df-mc/dragonfly/server/player.(*Player).session
func player_session(*player.Player) *session.Session

// noinspection ALL
//
//go:linkname session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
func session_writePacket(*session.Session, packet.Packet)
