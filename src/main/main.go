package main

import(
	"fmt"
	"strconv"
	"net/http"
	"io"
	"os"
)


func GetHttp(url string) (result string, err error){
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	
	defer resp.Body.Close()
	
	buf := make([]byte,4096)
	for{
		n, err2 := resp.Body.Read(buf)
		if n == 0{
			break

		}
		
		if err2 != nil && err2 != io.EOF{
			err = err2
			return
		}
		
		
		result += string(buf[:n])
		
	}
	
	return
	
	
}


func SpiderPage(i int, page chan int){
		url := "https://www.bing.com/search?q=golang&qs=n&sp=-1&pq=golang&sc=8-6&sk=&cvid=28EC02316298494B82D2A9C8C705957D&first="+strconv.Itoa((i-1)*10)+"&FORM=PERE"
		fmt.Println(url)
		result, err:= GetHttp(url)
		if err != nil {
			fmt.Println("Http Get error :", err)
			return
		}
		
		//fmt.Println("result = " + result)
		f, err := os.Create("spider"+ strconv.Itoa(i)+".html")
		if err != nil{
			fmt.Println("Error in create html files", err)
			return
		}
		
		f.WriteString(result)
		f.Close()
		
		page <- i
}

func working(start, end int){
	
	fmt.Printf("Getting from %d to %d...\n", start, end)
	
	page := make(chan int)
	
	for i:=start ;i<=end;i++{
			go SpiderPage(i, page)
	} 
	
	for i:= start; i<= end;i++{
		fmt.Printf("done No. %d\n",<-page)
	}
		
}



func main(){
	var start, end int
	
	fmt.Println("Start at...")
	fmt.Scan(&start)
	fmt.Println("End at...")
	fmt.Scan(&end)
	
	working(start,end)
}