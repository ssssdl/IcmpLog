package main
import (
	"golang.org/x/net/icmp"
	"log"
	"os"
)
func main() {
	var Interface_IP = os.Args[1]
	//for i,arg := range os.Args{
	//	fmt.Println(arg)
	//	fmt.Println("cmd=",os.Args[i+1])
	//}
	//flag.StringVar(&Interface_IP,"h","0.0.0.0","Attack Server IP")
	if len(Interface_IP) == 0 {
		log.Println(Interface_IP)
		log.Println("icmplog [Attack Server IP]")
		os.Exit(0)
	}
	log.Println(Interface_IP)
	conn, err := icmp.ListenPacket("ip4:icmp", Interface_IP)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var msg []byte
		length, sourceIP, err := conn.ReadFrom(msg)
		if err != nil {
			log.Println(err)
			continue
		}
		if sourceIP !=nil {
			log.Printf("message = '%s', length = %d, source-ip = %s", string(msg), length, sourceIP)
		}
	}
	_ = conn.Close()
}
