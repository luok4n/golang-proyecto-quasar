# Operacion Fuego de Quasar
## Problema a resolver 

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

* Desafío

Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que retorne la fuente y contenido del mensaje de auxilio. Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave.

Posición de los satélites actualmente en servicio
* Kenobi: [-500, -200]
* Skywalker: [100, -100]
* Sato: [500, 100]

Nivel 1

Crear un programa con las siguientes firmas:

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32)

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string)

Consideraciones:
 * La unidad de distancia en los parámetros de GetLocation es la misma que la que se
 utiliza para indicar la posición de cada satélite.
 * El mensaje recibido en cada satélite se recibe en forma de arreglo de strings.
 * Cuando una palabra del mensaje no pueda ser determinada, se reemplaza por un string
 en blanco en el array.
     - Ejemplo: [“este”, “es”, “”, “mensaje”]
 * Considerar que existe un desfasaje (a determinar) en el mensaje que se recibe en cada satélite.
     - Ejemplo:
         - Kenobi: [“”, “este”, “es”, “un”, “mensaje”] ■ Skywalker: [“este”, “”, “un”, “mensaje”]
         - Sato: [“”, ””, ”es”, ””, ”mensaje”]


Nivel 2
Crear una API REST, hostear esa API en un cloud computing libre (Google App Engine, Amazon AWS, etc), crear el servicio /topsecret/ en donde se pueda obtener la ubicación de la nave y el mensaje que emite.
El servicio recibirá la información de la nave a través de un HTTP POST con un payload con el siguiente formato:

POST  -> /topsecret/

```json
{
    "satellites": [
         {
               "name": "kenobi",
               "distance": 100.0,
               "message": ["este", "", "", "mensaje", ""]
        },
        {
                "name": "skywalker",
                "distance": 115.5
                "message": ["", "es", "", "", "secreto"]
        },
        {
                "name": "sato",
                "distance": 142.7
                "message": ["este", "", "un", "", ""]
        }
    ]
}

```

La respuesta, por otro lado, deberá tener la siguiente forma:

RESPONSE CODE: 200

```json
{
     "position": 
        {
            "x": -100.0,
            "y": 75.5 
        },
     "message": "este es un mensaje secreto"
}
```

En caso que no se pueda determinar la posición o el mensaje, retorna:

RESPONSE CODE: 404

Nivel 3
Considerar que el mensaje ahora debe poder recibirse en diferentes POST al nuevo servicio /topsecret_split/, respetando la misma firma que antes. Por ejemplo:

POST -> /topsecret_split/{satellite_name}

```json
{
     "distance": 100.0,
     "message": ["este", "", "", "mensaje", ""]
}
```

Crear un nuevo servicio /topsecret_split/ que acepte POST y GET. En el GET la respuesta deberá indicar la posición y el mensaje en caso que sea posible determinarlo y tener la misma estructura del ejemplo del Nivel 2. Caso contrario, deberá responder un mensaje de error indicando que no hay suficiente información.


# Solucion Operacion Quasar

Han Solo necesita nuestra ayuda para detectar la posicion y mensaje secreto proveniente de una nave portacarga imperial a la deriva.
Por esta razon hemos cosntruido la siguiente API.

Esta API contiene diferentes metodos para recibir estos llamados tanto en sus tres satelites unidos como cada uno de ellos por separado.

Se creo utilizando el lenguaje Golang y el patron de diseño singleton.




## Requerimientos

* [Golang](https://go.dev/)
* [Visual Studio Code](https://code.visualstudio.com/)
* [GoogleCloud SDK](https://cloud.google.com/appengine?utm_source=google&utm_medium=cpc&utm_campaign=latam-CO-all-es-dr-BKWS-all-all-trial-e-dr-1011454-LUAC0009157&utm_content=text-ad-none-any-DEV_c-CRE_545476760830-ADGP_Hybrid%20%7C%20BKWS%20-%20EXA%20%7C%20Txt%20~%20Compute_App-Engine-KWID_43700042625185793-kwd-372661972204&utm_term=KW_google%20app%20engine-ST_Google%20App%20Engine&gclid=EAIaIQobChMI6MXLpIud-QIVSIBQBh0RsQ49EAAYASAAEgKGX_D_BwE&gclsrc=aw.ds)

## Instalación

Ir a la carpeta main 'MeLi' correr el siguiente comando para su ejecucion en local

```console
C:\Users\..\MeLi> go run main.go
```

Para su despliegue en GoogleCloud es necesario correr el siguiente comando. Es recomendado correrlo en el Google Cloud SDK PowerShell

```console
C:\Users\..\MeLi> gcloud app deploy
```

## Test

Ir a la carpeta main 'unit_tests' correr el siguiente comando para su ejecucion 

```console
C:\Users\..\MeLi\unit_tests> go test
```

Para mas detalles en la ejecucion correr el siguiente comando 

```console
C:\Users\..\MeLi\unit_tests> go test -v
```
    
## API

#### topsecret

```http
  POST /api/topsecret
```

| Body | Type     | Descripcion                |
| :-------- | :------- | :------------------------- |
| `satellites` | `json` | **Requerido**

#### topsecret/{satellite_name}

```http
  GET POST api/topsecret/{satellite_name}
```

| Parametro/Body | Type     | Descripcion                       |
| :-------- | :------- | :-------------------------------- |
| `satellite_name`     | `string` | **Requerido**
| `position` `message` | `json`   | **Requerido**

## DIAGRAMAS

* [Diagramas](https://github.com/luok4n/MeLi/tree/main/Diagrams)
