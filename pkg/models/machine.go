package models

import (
	"time"

	"github.com/google/uuid"
)

type Machine struct {
	Id           uuid.UUID `json:"id"`
	MachineId    int       `json:"machineId"`
	LastLoggedIn string    `json:"lastloggedIn"`
	SysTime      time.Time `json:"sysTime"`
	Stat         Stat      `json:"stats"`
}
