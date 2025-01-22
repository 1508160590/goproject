package main
import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findval int){
	//判断leftIndex 是否大于 rightIndex
	if leftIndex >rightIndex {
		fmt.Println("我不到")
		return
	}
	
	//先找到 中间的下标
	middle:=(leftIndex+ rightIndex)/2
	
	if(*arr)[middle]>findval{
		//说朋我们要查找的数，应该在1eftIndex---midde1-1
		BinaryFind(arr,leftIndex,middle-1,findval)
	}else if(*arr)[middle]<findval {
		//说明我们要查找的数，应该在midde1+1--->rightIndex
		BinaryFind(arr,middle +1,rightIndex,findval)
	}else {
		//找到了
		fmt.Printf("找到了，下标为%v\n",middle)
	}
}
func main() {
	arr := [6]int{1,8,10,42,51,120}
	BinaryFind(&arr,0,len(arr)-1,510)
}