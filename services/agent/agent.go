// Package agent holds on services
// that build all together the business flow,
// agent only depends on package core.
package agent

import (
	"github.com/go-devs-ua/octagon/core"
)

type Agent struct {
	Repo core.Repository
	// Log
}

func NewUserAgent(r core.Repository) *Agent {
	return &Agent{Repo: r}
}
