# LAB_1_SD

# INTEGRANTES:
# NOMBRES/ROL:
- Marcelo Ramírez U / 201704620-1
- Jorge Sanhueza / 201704575-2


# Maquinas Usadas:
## Máquina 1 (branch: LOGISTICA) 
ip/hostname: dist125 
contraseña: wMGAtb9u

# Intrucciones Sistema Logistica: (Dentro de Carpeta "SistemaLogistica")
- Contiene dos servidores GRPC que se comunican con sistema Clientes y sistema Camiones
- Contiene cliente de RabbitMQ para comunicacion con sistema Finanzas 

- Registro Paquetes: SistemaLogistica/Registros/registroLogistica.csv
- Ejecutar: (Dentro de Carpeta "SistemaLogistica")
    - make run


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



