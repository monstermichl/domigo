package domino

import (
	"github.com/monstermichl/domigo/domino/com"

	"github.com/go-ole/go-ole"
)

type NotesConnector interface {
	Com() com.Com
	Release()
}

/* https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices */
func DispatchSlice[T NotesConnector](conns []T) []*ole.IDispatch {
	s := make([]*ole.IDispatch, len(conns))

	for i, c := range conns {
		s[i] = c.Com().Dispatch()
	}
	return s
}
