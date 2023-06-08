package main

// 定义服务

func Sum(args ...float64) float64 {
	sum := 0.0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func Product(args ...float64) float64 {
	product := 1.0
	for i := 0; i < len(args); i++ {
		product *= args[i]
	}
	return product
}

//func Sum(arg1, arg2 float64) float64 {
//	return arg1 + arg2
//}
//
//func Product(arg1, arg2 float64) float64 {
//
//	return arg1 * arg2
//}
