package cryptographer

import (
	"user-controller/internal/core/interfaces"
)

type cryptographer struct {
	Cost int
}

func New() interfaces.Cryptographer {
	return &cryptographer{
		Cost: 14,
	}
}
