package learning_go

import (
	"testing"
	
	"github.com/stretchr/testify/require"
)

func TestPointerInstanceIsNotMutable(t *testing.T) {
	// You can pass a pointer to a function. The function gets a _copy_ of the pointer. You can mutate
  // the memory it points to but you can't mutate the pointer itself, as it is a copy.

  tryToMutatePointerInstance := func(s *string){
    x := "New String"

    // Try to assign the pointer we passed in to a new pointer instance.
    s  = &x 
  }

  x := "Hello"
  tryToMutatePointerInstance(&x)

  // See that we didn't actually change anything
  require.Equal(t, "Hello", x)

  // But we can mutate the memory it points to
  mutateWhatItPointsTo := func(s *string){
    // Dereferencing. We're putting "New String" in the memory location that the passed in pointer
    // pointed to.
    *s = "New String"
  }

  mutateWhatItPointsTo(&x)

  // Now we've mutated things.
  require.Equal(t, "New String", x)
}

