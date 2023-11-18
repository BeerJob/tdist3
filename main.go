package main
import (
	"log"
	"fmt"
	"os"
	"bufio"
	"strings"
	//"strconv"
	//"math/rand"
	"time"
	"context"

	"google.golang.org/grpc"
	pb "github.com/BeerJob/tdist3/proto"
)
func main(){
	flag := 1
	for flag==1{
		fmt.Println("Comandos:")
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
		//Conexion con Broker
		conn, err := grpc.Dial("10.6.46.110:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con el Broker Luna!")
		}
		defer conn.Close()
		cliente := pb.NewInformantesClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := cliente.ObtenerServidorConsultarAleatorio(ctx, &pb.ParamamtroVacio{})
		if err != nil{
			log.Print("No hay respuesta del Broker Luna!")
		}else{
			log.Printf("La ip obtenida del Broker Luna es: %s", r.Ip)
		}
		//Conexion con BDs
		switch rawString[0] {
		case "AgregarBase":
			log.Print("se agrego una base")
		case "BorrarBase":
			log.Print("se borro una base")
		}
	}
}