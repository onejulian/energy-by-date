Repo para almacenar el proyecto energy-by-date, que proporciona el consumo de energía por fecha.

# Instrucciones para compilar y ejecutar
Proyecto basado en golang 1.19.3 https://go.dev/doc/install

-> Clonar este repositorio:

    git clone https://github.com/Julian-sUsername/energy-by-date.git

-> Abrir el proyecto y copiar en la raíz del proyecto el archivo .env proporcionado vía email.

-> Abrir una terminal posicionada en la raíz del proyecto y correr los siguientes comandos:

    go mod download

-----------------------------------------------

    go run .

-> En otra terminal (bash) ejecutar:

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period":"daily"}'

-----------------------------------------------

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"weekly"}'

-----------------------------------------------

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"monthly"}'

# Docker
La configuración de Docker permite compilar y ejecutar el proyecto en un contenedor, así como también ser desplegado en la nube. Es necesario instalar y ejecutar Docker. https://docs.docker.com/desktop/

El contenedor de este proyecto está desplegado en:

https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/

Puede probarlo abriendo una terminal (bash) y ejecutando cualquiera de los siguientes comandos:

    curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"daily"}'
    
-----------------------------------------------

    curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"weekly"}'

-----------------------------------------------

    curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{"date":"2022-10-25","period":"monthly"}'

# Compilación y ejecución local del contenedor

-> En los archivos infraestructure/translateMonth/translateMonth.go y env/env.go es necesario comentar las líneas marcadas al final con '// for local' y descomentar las líneas marcadas al final con '// for docker'; eso con el fin de que la imagen generada pueda leer archivos de configuración.

-> abrir una terminal posicionada en la raíz del proyecto y ejecutar los siguientes comandos:

    docker build -t energy-by-date .

-----------------------------------------------

    docker run -p 8186:8186 energy-by-date

-> En otra terminal (bash) ejecutar:

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "daily"}'

-----------------------------------------------

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "weekly"}'

-----------------------------------------------

    curl -X POST http://localhost:8186/generate-report -H 'Content-Type: application/json' -d '{"date": "2022-10-25","period": "monthly"}'