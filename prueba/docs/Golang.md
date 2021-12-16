# Golang
## Conceptos Basicos
Creacion de API con Golang 
### Protocolo HTTP/HTTP2
Es un estandar para las peticiones y respuestas.

- Request
- Response

### Arquitectura DDD/REST
Todo lo que se mueve en internet es un recurso representado por un formato que debe tener un edintificador unico (URIs). Los cuales son accedidos a traves de los metodos estandar http.

Los recursos pueden tener varias representaciones y pueden realizar comunixaciones sin estado (Staateless). 
### Stateless 
Persistencia de informacion entre peteciones sin respuesta (estado) ex. inicio de secion.
- token
- cookies

### Formatos (XML, JSON, gRPC)
Son las formas de representar los recursos desde el backend
JSON 
gRPC

### Metodos HTTP
- GET
- POST
- PUT
- PATCH
- DELETE
- HEAD
- OPTIONS

### Codigo de Respuesta
 - **1XX** Informativas 
 - **2XX** Exito
 - **3XX** Redireccion
 - **4XX** Error de cliente
 - **5XX** Error de servidor
 - Mayores a **599** respuestas no oficiales en la documentacion.

## Paquete HTTP/HTTP2
### Paquete Net/HTTP
 Sirve para crear:
 - Servidores
 - Clientes

### Servidores y Clientes
 - **ServeMux** Sirve como enrutador  para las peticiones HTTP entrantes con una lista de URIs para ejecutar el correspondiente handler.
 - **Handler** Responsable de escribir la respuesta al cliente. Envia la informacion de los encabezados(headers) y el cuerpo (bodles).
 **Handlers predefinidos**
	 - FileServer
	 - NotFoundHandler
	 - RedirectHandler
	 - StripPrefix
	 - TimeoutHandler

### Servidor web estático
``` go
 package main

import (
  "log"
  "net/http"
)

func main(){
  puerto := ":8080"
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)
  log.Println("Servidor Iniciado", puerto)
  http.ListenAndServe(puerto, nil)
}
```
 
### Primer Handler
 Firma para una funcion Handler:
```go
 func(http.ResponseWriter, *http.Request)
```
funcion `HandleFunc()` de `ServeMux` asi:
```go
package main

import (
  "fmt"
  "net/http"
)

func main(){
	puerto := ":8080"
	http.HandleFunc("/saludar", saludar)
	http.ListenAndServe(puerto, nil)
}
 func saludar(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Hola mundo")
 }
```
### Estructura Request
La estructura `http.Request` contiene informacion de peticion de cliente a un servidor.
```
URL: *url.URL
Method: string
Header: map[string][]string
Body: io.ReadCloser
Form/PostForm: url.Values
MultipartForm: *multipart.Form
```
Metodos 
- NewRequest()
- BasicAuth()
- Context()
- Cookie()
- FormFile()
- FormValue()
- ParseForm()
- ParseMultipartForm()

### Estructura Response y ResponseWriter
 Nos permite dar respuesta a una peticion de un cliente.
 ```go
 Status: string
 StatusCode: string
 Header: map[string][]string
 Body io.ReadCloser
 ```
 Pero no lo usamos directamente, para ello utilizamos ResponseWriter.
 ```go
 Header(): Permite escribir los headers de la respuesta.
 Write(...): Permite escribir el cuerpo de la repuesta.
 WriteHeader(...): Permite escribir el codigo de estado.
 ```
### Estructura Server
 Para personalizar algunas opciones de tu servidor, la estructura `Server` nos ayudara. ***more***
 Campos mas usados:
 ```go
 Addr: string
 Handler: handler
 Readtimeout: time.Duration
 WriteTomiout: time.Duration
 MaxHeaderBytes: int
 TLSConfig: *tls.Config
 Error.log: *log.Logger
 ```
 
## CRUD de una estrucctura
 codigo... 
- Storage create, update, delete, getbyid and getall
- Handler create, get all, update and delete
- Refactorizando las respuestas(codigo que se repite)
- Route and server (main) 
- Postman (erramienta online para realizar peticiones)

Modelos -> casos de usos (metodos)-> handlers(controladores, validadores)->    
Base de datos - 

## Middle wares
**Funciones de primer orden:** Funciones de reciben funciones agregan funcionalidad y devuleven funciones

para registro de logs, autenticacion de tokens...

## JWT autenticacion
un token tiene 3 partes en base 64 y no esta cifrado :
- **Header** Tipo de algoritmo para firmar el token
- **Payload** Informacion de identificacion del token
- **Verify signature** es una firma que verifica si el token no a sido modificado y como a sido contruido.

### Certificados
Permite identificar a una persona de forma unica, generalmente hay 2: 
Generar certificados openssl rsa online generate
- **Publico** Firmar el token
`openssl rsa -in app.rsa -pubout > app.rsa.pub` 
- **Privado** Validar el token
`openssl genrsa -out app.rsa 1024`

### Logica generar y validar token 
todo el codigo en el ejercicio N3
### 
dependencias
db -> modelo  -> casos de uso -> handler -> UI
## Framework GIN/ECHO

## Cliente HTTP


