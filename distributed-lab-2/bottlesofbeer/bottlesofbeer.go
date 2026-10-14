package main
import (
	"flag"
	"net/rpc"
	"fmt"
	"time"
	"net"
)

 


type Response struct {
	Message string
}

type Request struct {
	Message string
}

var nextAddr string

func server(thisPort *string){
	rpc.Register(&BottleOperations{})
    listener, _ := net.Listen("tcp", ":"+*thisPort)
    defer listener.Close()
	fmt.Println("Listening on " + *thisPort)
    rpc.Accept(listener)
}

type BottleOperations struct {}

func (s *BottleOperations) BottlesOfBeers(n int, reply *int) error {
       if n <= 0 {
           return nil
       }
	 time.Sleep(1* time.Second)
	 fmt.Println(n, "bottles of beer on the wall, ",n," bottles of beer. Take one down, pass it around...")
	 client, _ := rpc.Dial("tcp", nextAddr)
	 var r int
	 return client.Call("BottleOperations.BottlesOfBeers", (n-1),  &r)
	
}

func main(){
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")

	bottles := flag.Int("n",0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()

	if (*bottles != 0){
		server(thisPort)
	} else {
		go server(thisPort)
		fmt.Println("Starting process")
		ops := BottleOperations{}
		var r int
		ops.BottlesOfBeers(99, &r)
	}
}
