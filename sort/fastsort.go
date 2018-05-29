package fastsort

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
