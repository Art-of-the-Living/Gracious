package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/components"
	"testing"
)

func TestPsiBasic(t *testing.T) {
	psi := components.NewEvaluator("test")
	signal := psi.Evoke()
	fmt.Println(signal.Represent())
}
