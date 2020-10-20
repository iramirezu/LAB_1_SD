# LAB_1_SD

# INTEGRANTES:
# NOMBRES/ROL:
- Marcelo Ramírez U / 201704620-1
- Jorge Sanhueza / 201704575-2


# Maquinas Usadas:
##  Máquina 2 (branch: FINANCIERO) 
ip/hostname: dist127 
contraseña: VeVwrNsz


# Intrucciones Sistema Finanzas: (Dentro de Carpeta "SistemaFinanzas")
- Contiene servidor RabbitMQ que se comunica con sistema de Logistica

- Registro Finanzas: SistemaFinanzas/Registros/data.csv

- Ejecutar: (Dentro de Carpeta "SistemaFinanzas")
    - make run



# Desactivar FireWall
sudo systemctl stop firewalld
sudo systemctl disable firewalld
sudo systemctl mask --now firewalld


# Instalacion de RabbitMQ
go get  github.com/streadway/amqp



