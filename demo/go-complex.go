package main


import(
	"fmt"
	
	//"sync"
	"time"
	"runtime"
	
	
)

func main(){
		
		
		
		var list = make(chan int, 30)
		var count = make(chan int, 1)
		
	
		fmt.Println("Begin~num:",runtime.NumGoroutine())
		for  idx:= 1; idx<=3;idx++ {
			go func(list chan int, count chan int,  idx int){
			
					defer func(){
							if err := recover();err!=nil {
								fmt.Println("err:", err)
							}
					}()
				
					for {
						it, ok :=<-list
						if ok {
							fmt.Println("idx:list: ", idx, it)
							fmt.Println("num:",runtime.NumGoroutine())
							if it==14 {
								fmt.Println(idx)
								panic("num err:")
							}
							time.Sleep(time.Second)
						}else{
							count<-1
							return
						}
					}
			}(list, count , idx)
			
		}
		//time.Sleep(time.Second*6)
		
			for i:= 1; i<=30; i++ {
		
				list <-i 
		
		}
		
		fmt.Println("End~num:",runtime.NumGoroutine())
		close(list)
		<-count
		
		fmt.Println("--write end--")
		
}