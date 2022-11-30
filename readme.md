Repo para almacenar el proyecto energy-by-date, que proporciona el consumo de energía por fecha.

# Instrucciones para compilar y ejecutar
Proyecto basado en golang 1.19.3 https://go.dev/doc/install

-> Clonar este repositorio:
```bash
git clone https://github.com/Julian-sUsername/energy-by-date.git
```
-> Abrir el proyecto y copiar en la raíz del proyecto el archivo .env proporcionado vía email.

-> Abrir una terminal posicionada en la raíz del proyecto y correr los siguientes comandos:
```bash
go mod download
```

```bash
go run .
```

-> En otra terminal (bash) ejecutar:
```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period":"daily"}'
```

```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"weekly"}'
```

```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"monthly"}'
```

## Compilación y ejecución local del contenedor con docker

La configuración de Docker permite compilar y ejecutar el proyecto en un contenedor, así como también ser desplegado en la nube. Es necesario instalar y ejecutar Docker. https://docs.docker.com/desktop/

-> abrir una terminal posicionada en la raíz del proyecto y ejecutar los siguientes comandos:
```bash
docker build -t energy-by-date .
```

```bash
docker run -p 8186:8186 energy-by-date
```

-> En otra terminal (bash) ejecutar:
```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "daily"}'
```

```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "weekly"}'
```

```bash
curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "monthly"}'
```

# Despliegues
## Azure
https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/

Puede probarlo abriendo una terminal (bash) y ejecutando cualquiera de los siguientes comandos:
```bash
curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"daily"}'
```
    
```bash
curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"weekly"}'
```

```bash
curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"monthly"}'
```
