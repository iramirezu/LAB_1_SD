# LAB_1_SD

# INTEGRANTES:
# NOMBRES/ROL:
- Marcelo Ramírez U / 201704620-1
- Jorge Sanhueza / 201704575-2


# Maquinas Usadas:
## Máquina 2 (branch: CAMION) 
ip/hostname: dist126 
contraseña: WEzdgtde


# Intrucciones Sistema Camiones: (Dentro de Carpeta "SistemaCamion")
- Contiene cliente GRPC que se comunica con sistema de Logistica

- Registro Camion Retail 1: SistemaCamion/Registros/registroCamion1.csv
- Registro Camion Retail 2: SistemaCamion/Registros/registroCamion2.csv
- Registro Camion Normal 1: SistemaCamion/Registros/registroCamion3.csv

- Ejecutar: (Dentro de Carpeta "SistemaCamion")
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




