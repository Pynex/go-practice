package main

func main() {
	var a [3]int

	a[0] = 0
	a[1] = 1
	a[2] = 2

	modifyMe(&a)

	for i, el := range a {
		println(i, el)
	}

	q := [...]int{1,3,4}
	println(q[2])

	r := [...]byte{9: 'a'}
	println(r[9], r[8])

	sl := r[3:10]
	println(sl[6], len(sl))

	sl = append(sl, 213)
	println(sl[7], len(sl))
}

func modifyMe(arr *[3]int) {
	arr[1] = 42
}