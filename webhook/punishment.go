package webhook

import "fmt"

var banHook = NewHook("1144805219159523388", "gSG9jfTV7JKgArWUdF0nkdzQMcSw_dBuAh7q5NcONFjqnW1ejQp5")
var muteHook = NewHook("1144805514425933875", "9Kz6ueRmoOpqQGNyRc6XovycRtX5aiFWY9EMr6tuFIxh9VKX2Cj0yCq3278244sz0gGa")
var oomphHook = NewHook("1144805598161023029", "dzIY1um5AaQnZvrruPMYmf35eHi9KiiEOE5Rq1My")

type Punishment struct {
	p string
	h Hook
}

func BanPunishment() Punishment {
	return Punishment{"Ban", banHook}
}

func UnbanPunishment() Punishment {
	return Punishment{"Unban", banHook}
}

func MutePunishment() Punishment {
	return Punishment{"Mute", muteHook}
}

func UnMutePunishment() Punishment {
	return Punishment{"Unmute", muteHook}
}

func OomphPunishment() Punishment {
	return Punishment{"Oomph", oomphHook}
}

func SendPunishment(staff, victim, reason string, punishment Punishment) {
	var payload Payload
	payload.Embeds = []Embed{
		{
			Title:       punishment.p,
			Description: fmt.Sprintf("**Staff:** %s\n**Victim:** %s\n**Reason:** %s", staff, victim, reason),
		},
	}
	_ = punishment.h.SendMessage(payload)
}
