**The user interface**, containing the code adapting a delivery mechanism to a use case.

# CMD || Interface de Usuario
Contiene todos los puntos de entrada de nuestra aplicación.
Cada uno de ellos, será un paquete de la carpeta cmd con su respectivo **main.go** además del resto de código necesario como:

- configuraciones
- entornos
- inyección de dependencias
- etc

Es importante recordar que, los paquetes dentro de cmd se corresponderán con el nombre de los puntos de entrada, es decir, con el nombre de los binarios de las aplicaciones correspondientes (por ejemplo,
 /cmd/myapp).