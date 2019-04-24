package main

func IntAbs (x int) (y int) { //an integer absolute value function
	if x >0 {
		y = x
	} else {
		y = -1*x
	}
	return
}

func SumSlice (x []float32) (float32) {
    var run_sum float32 = 0
    for i := 0; i < len(x); i++ {
        run_sum = run_sum + x[i]
    }
    return run_sum
}

func Add (x []float32, y float32) ([]float32) {
    ret := x
    for i := 0; i < len(x); i++ {
        ret[i] = x[i] + y
    }
    return ret
}

func Multiply (x []float32, y float32) ([]float32) {
    ret := x
    for i := 0; i < len(x); i++ {
        ret[i] = x[i] * y
    }
    return ret
}
