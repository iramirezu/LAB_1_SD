# LAB_1_SD

# INTEGRANTES:
# NOMBRES/ROL:
- Marcelo Ramírez U / 201704620-1
- Jorge Sanhueza / 201704575-2


# Maquinas Usadas:
## Máquina 4 (branch: CLIENTE) 
ip/hostname: dist128 
contraseña: ztwNgh9j


# Intrucciones Sistema Cliente: (Dentro de Carpeta "SistemaCliente")
- Contiene cliente GRPC que se comunica con sistema de Logistica

- Registro intrucciones Retail: SistemaCliente/Registros/retail.csv
- Registro intrucciones Pymes: SistemaCliente/Registros/pymes.csv

- Clientes: 
    - Existen Clientes tipo Pymes y Retail
    - Se pueden agregar clientes como "go routines" dentro del main
    - Estos usuarios se reparten los registros de intrucicones dependiendo del tipo de usuario
- Ejecutar: (Dentro de Carpeta "SistemaCliente")
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


