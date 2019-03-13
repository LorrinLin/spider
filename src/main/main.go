package main

import(
	"fmt"
	"strconv"
	"net/http"
	"io"
	"os"
)

// Http Get method using to request the website
func GetHttp(url string) (result string, err error){
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	
	// it will be close after everything done
	// keyword 'defer' is very useful
	defer resp.Body.Close()
	
	buf := make([]byte,4096)
	for{
		
		//read everything from the response body, 
		//and write them into 'result'
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
	//this url is the website I want to spider
		url := "https://www.bing.com/search?q=golang&qs=n&sp=-1&pq=golang&sc=8-6&sk=&cvid=28EC02316298494B82D2A9C8C705957D&first="+strconv.Itoa((i-1)*10)+"&FORM=PERE"
		fmt.Println(url)
		//Useing HTTP Get method to request
		result, err:= GetHttp(url)
		if err != nil {
			fmt.Println("Http Get error :", err)
			return
		}
		
		//create a html file to save the page
		f, err := os.Create("spider"+ strconv.Itoa(i)+".html")
		if err != nil{
			fmt.Println("Error in create html files", err)
			return
		}
		
		//writeing the data to the file and close it
		f.WriteString(result)
		f.Close()
		
		// channel using to make concurrency
		page <- i
}

func working(start, end int){
	fmt.Printf("Getting from %d to %d...\n", start, end)
	
	//make a channel used in concurrency
	page := make(chan int)
	
	for i:=start ;i<=end;i++{
		
		//using go routine to concurrency
			go SpiderPage(i, page)
	} 
	
	//main thread only can be closed until all the go routine done
	for i:= start; i<= end;i++{
		fmt.Printf("done No. %d\n",<-page)
	}
		
}



func main(){
	var start, end int
	//How many pages you want to spider
	//if it succeeds, the files will be save in the current package, together with main.go
	
	fmt.Println("You can choose the pages you want to download from bing with research 'golang'...")
	fmt.Println("Start at...")
	fmt.Scan(&start)
	fmt.Println("End at...")
	fmt.Scan(&end)
	
	working(start,end)
}
