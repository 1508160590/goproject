package main

func main(){
	var i = 3
	switch{
		case true:
			println("111")
		if(i == 2){
			break
		}

		fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1

		case false:
			println("222")
		case false:
			println("333")
	}
}