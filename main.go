package main
import (
	"log"
	"fmt"
	"os"
	"bufio"
	"strings"
	//"strconv"
	//"math/rand"
	//"time"
	//"context"

	//"google.golang.org/grpc"
	//pb "github.com/BeerJob/tdist/proto"
)
func main(){
	flag := 1
	for flag==1{
		fmt.Println("Comandos: ")
		fmt.Println("AgregarBase <nombre_sector> <nombre_base> [valor]")
		fmt.Println("RenombrarBase <nombre_sector> <nombre_base> <nuevo_nombre>")
		fmt.Println("ActualizarValor <nombre_sector> <nombre_base> <nuevo_valor>")
		fmt.Println("BorrarBase <nombre_sector> <nombre_base>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil{
			log.Fatal(err)
		}
		rawString := strings.Fields(scanner.Text())
		switch rawString[0] {
		case "AgregarBase":
			log.Print("se agrego una base")
		case "BorrarBase":
			log.Print("se borro una base")
		}
	}
}