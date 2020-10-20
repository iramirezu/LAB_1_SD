# LAB_1_SD

# Maquinas Usadas:
## Máquina 1 (branch: LOGISTICA) 
ip/hostname: dist125 
contraseña: wMGAtb9u

## Máquina 2 (branch: CAMION) 
ip/hostname: dist126 
contraseña: WEzdgtde

##  Máquina 2 (branch: FINANCIERO) 
ip/hostname: dist127 
contraseña: VeVwrNsz

## Máquina 4 (branch: CLIENTE) 
ip/hostname: dist128 
contraseña: ztwNgh9j

# Desactivar FireWall
sudo systemctl stop firewalld
sudo systemctl disable firewalld
sudo systemctl mask --now firewalld

# Instalacion de GRPC
export GO111MODULE=on  
go get github.com/golang/protobuf/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin

# Instalacion de RabbitMQ
go get  github.com/streadway/amqp

# Generacion de gRPC:
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative chatCliente/chatCliente.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative chatCamion/chatCamion.proto

# Luego en carpeta que contiene server
go mod init github.com/PrestigioExpress/ServicioDistribuido


# Intrucciones Sistema Logistica:
- Contiene dos servidores GRPC que se comunican con sistema Clientes y sistema Camiones
- Contiene cliente de RabbitMQ para comunicacion con sistema Finanzas 
Registro Paquetes: Registros/registroLogistica.csv
Ejecutar:
    go run server.go

# Intrucciones Sistema Cliente:
- Contiene cliente GRPC que se comunica con sistema de Logistica
Registro intrucciones Retail: Registros/retail.csv
Registro intrucciones Pymes: Registros/pymes.csv

Clientes: 
    - Existen Clientes tipo Pymes y Retail
    - Se pueden agregar clientes como "go routines" dentro del main
    - Estos usuarios se reparten los registros de intrucicones dependiendo del tipo de usuario
Ejecutar:
    go run clienteCliente.go

# Intrucciones Sistema Camiones:
- Contiene cliente GRPC que se comunica con sistema de Logistica
Registro Camion Retail 1: Registros/registroCamion1.csv
Registro Camion Retail 2: Registros/registroCamion2.csv
Registro Camion Normal 1: Registros/registroCamion3.csv
Ejecutar:
    go run clienteCamion.go

# Intrucciones Sistema Finanzas:
