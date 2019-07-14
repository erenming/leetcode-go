package sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	BubbleSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	InsertSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	SelectSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	MergeSort(arr)
	fmt.Println("排序后：",arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{4,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	QuickSort(arr)
	fmt.Println("排序后：",arr)
}

func TestFindKth(t *testing.T) {
	arr := []int{11, 3, 4, 5, 7}
	k := 5
	fmt.Println("序列：", arr)
	fmt.Printf("第%d大元素为%d\n", k, FindKth(arr, k))
}

func TestBucketSort(t *testing.T) {
	arr := []int{4,1,4,6,3,5,5,1}
	fmt.Println("排序前：",arr)
	BucketSort(arr)
	fmt.Println("排序后：",arr)
}

func TestMerge(t *testing.T)  {
	//arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	arr := [][]int{{1, 4}, {0, 4}}

	fmt.Println("合并前：", arr)
	newarr := Merge(arr)
	fmt.Println("合并后: ", newarr)
}