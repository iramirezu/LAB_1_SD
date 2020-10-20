package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"strconv"
	"encoding/csv"	
	"encoding/json"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func SetupCloseHandler( costos *int, ganancias *int) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		log.Printf("Balance: \nGastos=%d \nIngreso=%d \nTotal=%d \n",*costos,*ganancias,*ganancias-*costos)
		os.Exit(0)
	}()
}

type Json struct {
	idPaquete string `json:"idPaquete"`
	tipo string `json:"tipo"`
	valor string `json:"valor"`
	intentos string `json:"intentos"`
	fechaEntrega string `json:"fechaEntrega"`
	exito string `json:"exito"`
}



func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	var costos int
	var ganancias int	
	SetupCloseHandler(&costos,&ganancias)
	var info map[string]interface{}
	forever := make(chan bool)
	
	go func() {
		for d := range msgs {
						
			err=json.Unmarshal([]byte(d.Body),&info)
			if err!=nil{
				fmt.Println(err)
			}

			log.Printf("Received a message: %s", d.Body)
			f, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			w := csv.NewWriter(f)
			var row []string
			var costo string
			
			id:=fmt.Sprintf("%v", info["idPaquete"])
			row=append(row,id)
			tipo:=fmt.Sprintf("%v", info["tipo"])
			row=append(row,tipo)
			valor:=fmt.Sprintf("%v", info["valor"])
			row=append(row,valor)
			intentos:=fmt.Sprintf("%v", info["intentos"])
			row=append(row,intentos)
			fechaEntrega:=fmt.Sprintf("%v", info["fechaEntrega"])
			row=append(row,fechaEntrega)
			exito:=fmt.Sprintf("%v", info["exito"])
 			row=append(row,exito)
				
			intentos1:=info["intentos"].(float64)
			valor1,_:=strconv.Atoi(valor)
			
			if tipo=="retail"{
  				if info["exito"].(float64)==1 {

          				costos=costos+int(intentos1*10)
					ganancias= ganancias + valor1

			        	costo=strconv.Itoa(-int(intentos1*10)+valor1) //total
					row= append(row,costo)
  				} else{
          				costos=costos+30
          				ganancias= ganancias + valor1
          				costo= strconv.Itoa(valor1-30)
					row= append(row,costo)//total
  				       }
			} else if tipo=="prioritario"{
  				if info["exito"].(float64)==1 {

          				costos=costos+int(intentos1*10)
					valor1= valor1*13/10
					costo=strconv.Itoa(valor1-int(intentos1*10))
          				row= append(row,costo)
          				ganancias= ganancias + valor1
  				} else{
          				costos=costos+int(intentos1*10)
					valor1=valor1*3/10						
					row= append(row,costo)
        				costo=strconv.Itoa(valor1-int(intentos1*10))
          				ganancias= ganancias + valor1
  			              }
			} else if tipo=="normal"{
				if info["exito"].(float64)==1 {
				  	costos=costos+int(intentos1*10)					
                                        costo=strconv.Itoa(valor1-int(intentos1*10))
                                        row= append(row,costo)
					ganancias= ganancias + valor1
				}else{
				   	costo=strconv.Itoa(-int(intentos1*10))
                                        costos=costos+int(intentos1*10)
                                        row= append(row,costo)
                                        ganancias= ganancias + 0 // no se gnaa
				}
			}
						
			w.Write(row)
			w.Flush()			
			   
		}
		
	}()	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	log.Printf("balance=0")
}
