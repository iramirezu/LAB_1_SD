# LAB_1_SD

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


# Generacion de gRPC:
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative chatCliente/chatCliente.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative chatCamion/chatCamion.proto

# Luego en carpeta que contiene server
go mod init github.com/PrestigioExpress/ServicioDistribuido


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