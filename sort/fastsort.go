package fastsort

// func main() {
// 	//测试的随机数
// 	nums := []int{3542, 21, 33, 14, 5, 36, 7, 33, 36, 5, 102, 501, 741, 10, 1561, 151, 123, 212,
// 		1516, 212, 3, 548, 14, 12, 156, 41, 100, 12, 15, 13, 22, 111, 335, 148, 14, 213, 121, 212, 1414, 32, 15,
// 		515, 1515, 213, 21215, 212, 14623, 156, 13515, 135135, 351513, 351533, 15313, 315313, 12121, 211, 515, 1215,
// 		151, 52131, 15312, 17874, 949, 4894, 494984, 89494, 894654, 89454, 8454, 84846, 4548, 48846, 486, 486, 4684,
// 		484, 4646, 51, 321, 3132, 35151, 3554, 6546, 64564, 6546, 45456, 5646, 454, 44545, 2121, 21, 21, 5454, 146845,
// 		1651, 15131, 35153, 12153, 15312, 64846, 546454, 845654, 54847, 4747, 4651, 2151, 5351, 53, 131, 3515, 86, 56,
// 	}

// 	//单核选择单进程(数据量小也应该选择单进程)
// 	FastSort(nums)
// 	fmt.Println(nums)

// 	//多核选择多进程
// 	runtime.GOMAXPROCS(8)
// 	completed := make(chan bool)
// 	go FastSortProcesses(nums, completed)
// 	<-completed
// 	fmt.Println(nums)
// }

//FastSortProcesses 多进程并行快速排序(多核选择)
func FastSortProcesses(slice []int, completed chan bool) {
	allChan := make(chan bool, 2)
	chanNum := 0
	size := len(slice)
	if size == 2 && slice[0] > slice[1] {
		temp := slice[0]
		slice[0] = slice[1]
		slice[1] = temp
	} else if size > 2 {
		var i int
		for i = size - 1; i > 0; i-- {
			if slice[i] < slice[0] {
				var j int
				for j = 1; j < i; j++ {

					if slice[j] > slice[0] {
						temp := slice[i]
						slice[i] = slice[j]
						slice[j] = temp
						break
					}
				}
				if i == j {
					temp := slice[i]
					slice[i] = slice[0]
					slice[0] = temp
					if i > 1 {
						go FastSortProcesses(slice[:i], allChan)
						chanNum++
					}

					if size-i > 2 {
						go FastSortProcesses(slice[i+1:], allChan)
						chanNum++
					}
					break
				}

			}
		}
		if i == 0 {
			go FastSortProcesses(slice[i+1:], allChan)
			chanNum++
		}
	}
	for i := 0; i < chanNum; i++ {
		<-allChan
	}
	completed <- true
}

//FastSort 单线程快速排序(单核选择)
func FastSort(slice []int) {
	size := len(slice)
	if size == 2 && slice[0] > slice[1] {
		temp := slice[0]
		slice[0] = slice[1]
		slice[1] = temp
	} else if size > 2 {
		var i int
		for i = size - 1; i > 0; i-- {
			if slice[i] < slice[0] {
				var j int
				for j = 1; j < i; j++ {

					if slice[j] > slice[0] {
						temp := slice[i]
						slice[i] = slice[j]
						slice[j] = temp
						break
					}
				}
				if i == j {
					temp := slice[i]
					slice[i] = slice[0]
					slice[0] = temp
					if i > 1 {
						FastSort(slice[:i])
					}

					if size-i > 2 {
						FastSort(slice[i+1:])
					}
					break
				}

			}
		}
		if i == 0 {
			FastSort(slice[i+1:])
		}
	}
}
