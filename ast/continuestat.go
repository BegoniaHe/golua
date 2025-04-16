package ast

import (
	"github.com/BegoniaHe/golua/token"
)

// ContinueStat is a statement node representing the "continue" statement.
type ContinueStat struct {
	Location
}

var _ Stat = ContinueStat{}

// NewContinueStat returns a ContinueStat instance (the token is needed to record the location of the statement).
func NewContinueStat(tok *token.Token) ContinueStat {
	return ContinueStat{Location: LocFromToken(tok)}
}

// HWrite prints a tree representation of the node.
func (s ContinueStat) HWrite(w HWriter) {
	w.Writef("continue")
}

// ProcessStat uses the given StatProcessor to process the receiver.
func (s ContinueStat) ProcessStat(p StatProcessor) {
	p.ProcessContinueStat(s)
}
