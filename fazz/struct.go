package fazz

import "fmt"

func Matrix() {
	var matrix = [][][]int{
		[][]int{
			[]int{
				1,
			},
			[]int{
				0,
				2,
				3,
			},
		},
		[][]int{
			[]int{
				6,
			},
		},
		[][]int{
			[]int{
				9,
			},
		},
		[][]int{
			[]int{
				100,
				101,
				10,
			},
		},
	}
	fmt.Println(matrix[3][0][2])
	fmt.Println(matrix[0][1][2])
}
