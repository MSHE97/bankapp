package deposit
 import (
	"fmt"
	"bank/pkg/bank/types"
 )
func ExampleCalculate () {
	fmt.Println( Calculate(0, types.TJS) )
	fmt.Println( Calculate(0, types.USD) )
	fmt.Println( Calculate(500_000_00, types.TJS) )
	fmt.Println( Calculate(500_000_00, types.USD) )
	fmt.Println( Calculate(1_000_000_00, types.TJS) )
	fmt.Println( Calculate(1_000_000_00, types.USD) )

	// Output:
	// 0 0
	// 0 0
	// 166666 250000
	// 125000 166666
	// 333333 500000
	// 250000 333333
}