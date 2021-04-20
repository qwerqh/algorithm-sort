package algorithm

import (
	"fmt"
)

func BubbleSort(arr []int,length int)  {
	fmt.Println("冒泡排序")
	for i:=0;i<length-1;i++{
		for j:=0;j<length-1-i;j++ {
			if arr[j]>arr[j+1] {
				temporary:=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temporary
			}
		}
	}
}

func InsertSort(arr []int,length int)  {
	fmt.Println("插入排序")
	for i := 0; i < length; i++ {
		j:=i
		temporary:=arr[i]
		for j>0&&arr[j-1]>temporary {
			arr[j]=arr[j-1]
			j--
		}
		arr[j]=temporary
	}
}

func SelectSort(arr []int,length int)  {
	fmt.Println("选择排序")
	for i:=0;i<length;i++ {
		 min:=i
		for j:=i+1;j<length;j++ {
			if arr[min]>arr[j] {
				min=j
			}
		}
		if i!=min {
			temporary :=arr[min]
			arr[min] =arr[i]
			arr[i]=temporary
		}
	}
}
