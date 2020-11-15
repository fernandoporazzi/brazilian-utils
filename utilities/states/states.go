package states

import (
	"sort"

	"github.com/fernandoporazzi/brazilian-utils/data"
)

// State holds the Code and the Name of a state
type State struct {
	Code string
	Name string
}

// NameSorter sorts states by name
type NameSorter []State

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// GetStates returns all Brazilian states
func GetStates() []State {
	var output []State

	for key, state := range data.States {
		s := State{key, state.Name}
		output = append(output, s)
	}

	sort.Sort(NameSorter(output))

	return output
}
