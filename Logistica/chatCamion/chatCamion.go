package chatCamion

import (
	"fmt"
	"log"
	"encoding/csv"
	//"strconv"
	"os"
	//"math/rand"
	"golang.org/x/net/context"
	"time"

	//"time"
)

type Server struct {
}

var ColaPrioritaria[]string
var ColaNormal[]string
var ColaRetail[]string

// =========================================== FUNCIONES DE GRPC =======================================================================================

func (s *Server) PedirPaquete(ctx context.Context, mensaje *PeticionPaquete) (*PaqueteRecibido, error) {
	rows := leerFilasRegistro("registroLogistica")
	if len(rows) > 1 {
		for i := 1; i < len(rows); i++ {
			if rows[i][8] == "0" { // intentos == 0
				if rows[i][2] == "normal"{
					if len(ColaNormal) > 0{
						encontrado := 0
						for j := 0; j < len(ColaNormal); j++ {
							if ColaNormal[j] == rows[i][1]{
								encontrado = 1
							}
						}
						if encontrado == 0 {
							ColaNormal = append(ColaNormal, rows[i][1])
						}
					}else{
						ColaNormal = append(ColaNormal, rows[i][1])
					}
				}else if rows[i][2] == "retail" {
					if len(ColaRetail) > 0{
						encontrado := 0
						for j := 0; j < len(ColaRetail); j++ {
							if ColaRetail[j] == rows[i][1]{
								encontrado = 1
							}
						}
						if encontrado == 0 {
							ColaRetail = append(ColaRetail, rows[i][1])
						}
					}else{
						ColaRetail = append(ColaRetail, rows[i][1])
					}
				}else{ // prioritario
					if len(ColaPrioritaria) > 0{
						encontrado := 0
						for j := 0; j < len(ColaPrioritaria); j++ {
							if ColaPrioritaria[j] == rows[i][1]{
								encontrado = 1
							}
						}
						if encontrado == 0 {
							ColaPrioritaria = append(ColaPrioritaria, rows[i][1])
						}
					}else{
						ColaPrioritaria = append(ColaPrioritaria, rows[i][1])
					}
				}
			}
		}
	}else{
		if len(ColaPrioritaria) == 0 || len(ColaNormal) == 0 || len(ColaRetail) == 0 {
			id_r := "0";
			tipo_r := "0";
			valor_r := "0";
			origen_r := "0";
			destino_r := "0";
			intentos_r := "0";
			fechaEntrega_r := "0";
			exito_r := "0";
			fmt.Println("Peticion de paquete, NO hay entregas por repartir")
			
			return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil
		}
	}
	
	// LEE PAQUETES NUEVOS AGREGADOS (CON INTENTOS = 0 QUE NO SE ENCUENTRAN EN LA COLA)
	// PAQUETES NUEVOS SON AGREGADOS A LA COLA

	// PRIMER PAQUETE DE LA COLA ES ENTREGADO AL CAMION Y SE LE AGREGA FECHA ENTREGA
	tipoCamion := mensaje.TipoCamion
	if tipoCamion == "0" { // retail 1 y 2 -----------------------------------------------------------------------------------------
		// entrega paquetes retail primero, luego entrega paquetes prioritarios
		if len(ColaRetail) > 0{ // entrega paquete retail
			// borra paquete de cola retail y envia paquete
			idPaqueteEnviar := ColaRetail[0]
			ColaRetail[0] = "" // Elimina Primer elemento de la cola
			ColaRetail = ColaRetail[1:]
			rows := leerFilasRegistro("registroLogistica")
			for i := 0; i < len(rows); i++ {
				if rows[i][0] == idPaqueteEnviar {
					id_r := rows[i][1]
					tipo_r := rows[i][2]
					valor_r := rows[i][4]
					origen_r := rows[i][5]
					destino_r := rows[i][6]
					intentos_r := rows[i][8]
					rows[i][9] = fechaHoy()
					fechaEntrega_r := rows[i][9]
					exito_r := rows[i][10]
					escribirFilasRegistro("registroLogistica", rows)
					fmt.Println("Camion RETAIL realiza peticion de paquete, paquete RETAIL enviado")
					
					return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil

				}
			}
		} else{ // intenta entregar paquete prioritario
			if len(ColaPrioritaria) > 0{ // entrega paquete prioritario
				// borra paquete de cola prioria y encia paquete
				idPaqueteEnviar := ColaPrioritaria[0]
				ColaPrioritaria[0] = "" // Elimina Primer elemento de la cola
				ColaPrioritaria = ColaPrioritaria[1:]
				rows := leerFilasRegistro("registroLogistica")
				for i := 0; i < len(rows); i++ {
					if rows[i][0] == idPaqueteEnviar {
						id_r := rows[i][1]
						tipo_r := rows[i][2]
						valor_r := rows[i][4]
						origen_r := rows[i][5]
						destino_r := rows[i][6]
						intentos_r := rows[i][8]
						rows[i][9] = fechaHoy()
						fechaEntrega_r := rows[i][9]
						exito_r := rows[i][10]
						escribirFilasRegistro("registroLogistica", rows)
						fmt.Println("Camion RETAIL realiza peticion de paquete, paquete PRIORIORITARIO enviado")
						
						return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil

					}
				}
			} else{ // No hay paquetes por entregar
				id_r := "0";
				tipo_r := "0";
				valor_r := "0";
				origen_r := "0";
				destino_r := "0";
				intentos_r := "0";
				fechaEntrega_r := "0";
				exito_r := "0";
				fmt.Println("Camion RETAIL realiza peticion de paquete, NO hay paquetes retail o prioritarios por entregar")
				
				return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil
	
			}
		}
	} else{ // normal ----------------------------------------------------------------------------------------------
		// entrega prioritarios primero, luego entrega paquetes normales
		if len(ColaPrioritaria) > 0{ // entrega paquete prioritario
			// borra paquete de cola prioritario y envia paquete
			idPaqueteEnviar := ColaPrioritaria[0]
			ColaPrioritaria[0] = "" // Elimina Primer elemento de la cola
			ColaPrioritaria = ColaPrioritaria[1:]
			rows := leerFilasRegistro("registroLogistica")
			for i := 0; i < len(rows); i++ {
				if rows[i][0] == idPaqueteEnviar {
					id_r := rows[i][1]
					tipo_r := rows[i][2]
					valor_r := rows[i][4]
					origen_r := rows[i][5]
					destino_r := rows[i][6]
					intentos_r := rows[i][8]
					rows[i][9] = fechaHoy()
					fechaEntrega_r := rows[i][9]
					exito_r := rows[i][10]
					escribirFilasRegistro("registroLogistica", rows)
					fmt.Println("Camion NORMAL realiza peticion de paquete, paquete PRIORIORITARIO enviado")
					
					return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil

				}
			}
		} else{ // intenta entregar paquete normal
			if len(ColaNormal) > 0{ // entrega paquete normal
					// borra paquete de cola normal y encia paquete
					idPaqueteEnviar := ColaNormal[0]
					ColaNormal[0] = "" // Elimina Primer elemento de la cola
					ColaNormal = ColaNormal[1:]
					rows := leerFilasRegistro("registroLogistica")
					for i := 0; i < len(rows); i++ {
						if rows[i][0] == idPaqueteEnviar {
							id_r := rows[i][1]
							tipo_r := rows[i][2]
							valor_r := rows[i][4]
							origen_r := rows[i][5]
							destino_r := rows[i][6]
							intentos_r := rows[i][8]
							rows[i][9] = fechaHoy()
							fechaEntrega_r := rows[i][9]
							exito_r := rows[i][10]
							escribirFilasRegistro("registroLogistica", rows)
							fmt.Println("Camion NORMAL realiza peticion de paquete, paquete NORMAL enviado")
							
							return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil

						}
					}

				} else{ // No hay paquetes por entregar
					id_r := "0"
					tipo_r := "0"
					valor_r := "0"
					origen_r := "0"
					destino_r := "0"
					intentos_r := "0"
					fechaEntrega_r := "0"
					exito_r := "0"
					fmt.Println("Camion NORMAL realiza peticion de paquete, NO hay paquetes  normales o prioritarios por entregar")
					
					return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil
		
				}
		}
	}
	id_r := "0"
	tipo_r := "0"
	valor_r := "0"
	origen_r := "0"
	destino_r := "0"
	intentos_r := "0"
	fechaEntrega_r := "0"
	exito_r := "0"

	return nil
	//return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil
	
}

func (s *Server) CompletarEntrega(ctx context.Context, mensaje *PaqueteCompletado) (*MensajeReply, error) {
	id_c := mensaje.Id
	tipo_c := mensaje.Tipo
	valor_c := mensaje.Valor
	origen_c := mensaje.Origen
	destino_c := mensaje.Destino
	intentos_c := mensaje.Intentos
	fechaEntrega_c := mensaje.FechaEntrega
	exito_c := mensaje.Exito

	fmt.Println("Entrega Completada recibida desde camion: " + id_c + ", " + tipo_c + ", " + valor_c + ", " + origen_c + ", " + destino_c + ", " + intentos_c + ", " + fechaEntrega_c + ", " + exito_c)
	// actualizar linea a Registro Logistica
	rows := leerFilasRegistro("registroLogistica")
	for i := 0; i < len(rows); i++ {
		if rows[i][0] == id_c {
			rows[i][8] = intentos_c
			rows[i][9] = fechaEntrega_c
			rows[i][10] = exito_c
			escribirFilasRegistro("registroLogistica", rows)
		}
	}
	return &MensajeReply{Respuesta1: "."}, nil
}





// =========================================== FUNCIONES DE AYUDA =======================================================================================

func fechaHoy() string {
	t := time.Now()
	timestamp := fmt.Sprintf("%02d-%02d-%d",
		t.Year(), t.Month(), t.Day())
	return timestamp
}
func cantFilasRegistro(nombreRegistro string) int { // numero de seguimiento sera un autogenerado que se relaciona directamente con la cantidad de registros
    rows := leerFilasRegistro(nombreRegistro)
	cant := (len(rows))	
    return cant
}

func leerFilasRegistro(nombreRegistro string) [][]string {
    f, err := os.Open("Registros/"+nombreRegistro+".csv")
    if err != nil {
        log.Fatal(err)
    }
    rows, err := csv.NewReader(f).ReadAll()
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
    return rows
}

func escribirFilasRegistro(nombreRegistro string, rows [][]string) {
    f, err := os.Create("Registros/"+nombreRegistro+".csv")
    if err != nil {
        log.Fatal(err)
    }
    err = csv.NewWriter(f).WriteAll(rows)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
}

func eliminarFilaRegistro(nombreRegistro string, index int) {
	a := leerFilasRegistro(nombreRegistro)
	copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	escribirFilasRegistro(nombreRegistro, a)
}