package algorithm

import (
	"fmt"
)

/*
冒泡排序：
平均时间复杂度：O(n^2)
最好情况：O(n)
最坏情况：O(n^2)
空间复杂度：O(1)
稳定性：稳定
 */
func BubbleSort(arr []int, length int) {
	fmt.Println("冒泡排序")
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				temporary := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temporary
			}
		}
	}
}

/*
插入排序：
平均时间复杂度：O(n^2)
最好情况：O(n)
最坏情况：O(n^2)
空间复杂度：O(1)
稳定性：稳定
*/
func InsertSort(arr []int, length int) {
	fmt.Println("插入排序")
	for i := 0; i < length; i++ {
		j := i
		temporary := arr[i]
		for j > 0 && arr[j-1] > temporary {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temporary
	}
}

/*
选择排序：
平均时间复杂度：O(n^2)
最好情况：O(n^2)
最坏情况：O(n^2)
空间复杂度：O(1)
稳定性：不稳定
*/
func SelectSort(arr []int, length int) {
	fmt.Println("选择排序")
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			temporary := arr[min]
			arr[min] = arr[i]
			arr[i] = temporary
		}
	}
}
/*
希尔排序：
平均时间复杂度：O(nlogn)
最好情况：O(n(logn)^2)
最坏情况：O(n(logn)^2)
空间复杂度：O(1)
稳定性：不稳定
*/
func ShellSort(arr []int, length int) {
	fmt.Println("希尔排序")
	for step := length / 2; step > 0; step /= 2 {
		for k := 0; k < step; k++ {
			for i := k + step; i < length; i += step {
				j := i
				temporary := arr[i]
				for j > k && arr[j-step] > temporary {
					arr[j] = arr[j-step]
					j -= step
				}
				arr[j] = temporary
			}
		}
	}
}
/*
归并排序：
平均时间复杂度：O(nlogn)
最好情况：O(nlogn)
最坏情况：O(nlogn)
空间复杂度：O(n)
稳定性：稳定
*/
func MergeSort(arr []int, length int) {
	fmt.Println("归并排序")
	var temp [10]int
	split(arr, 0, length-1, temp[:])
}

func split(arr []int, low int, high int, temp []int) {
	if low < high {
		mid := (high + low) / 2
		split(arr, low, mid, temp)
		split(arr, mid+1, high, temp)
		merge(arr, low, mid, high, temp)

	}
}

func merge(arr []int, low int, mid int, high int, temp []int) {
	lindex := low
	rindex := mid + 1
	tindex := 0
	for lindex <= mid && rindex <= high {
		if arr[lindex] < arr[rindex] {
			temp[tindex] = arr[lindex]
			lindex++
		} else {
			temp[tindex] = arr[rindex]
			rindex++
		}
		tindex++
	}
	for lindex <= mid {
		temp[tindex] = arr[lindex]
		lindex++
		tindex++
	}
	for rindex <= high {
		temp[tindex] = arr[rindex]
		rindex++
		tindex++
	}
	for i := 0; i < tindex; i++ {
		arr[i+low] = temp[i]
	}

}
