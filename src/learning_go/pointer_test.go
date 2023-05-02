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
    s  = &x // Try to assign the pointer we passed in to a new pointer instance
  }

  x := "Hello"
  tryToMutatePointerInstance(&x)

  require.Equal(t, "Hello", x)
}

