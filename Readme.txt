> Proyecto Magneto Reclutador:

El siguiente proyecto está  basado en una prueba para demostrar habilidades técnicas a la hora de desarrollar.
El proyecto consiste en diseñar e implementar una app web donde se recibe un patrón de ADN y por medio de este
debemos verificar si cumple con las condiciones requeridas para ser aceptado como mutante y posteriormente ser
reclutado por Magneto.

 


## Tabla de contenido
* [Información general]
* [Tecnologías utilizadas]
* [Características]
* [Uso]
* [Estado del proyecto]
* [Contacto]



## Información general

Este proyecto se lleva a cabo con la finalidad de presentar mis habilidades en desarrollo a los evaluadores y 
reclutadores de la empresa Pulzo.
 



## Tecnologías utilizadas
- Lenguaje base - Golang - versión go1.18.3 for Windows
- Editor de Codigo - Visual Studio Code - versión 1.68.0 for Windws
- Libreria -github.com/gorilla/mux - versión 1.8.0


## Características

El proyecto esta implementado por 4 funciones base. Se  inicia haciendo una petición por medio de un api rest
para traer los datos en formato json los cuales serán decodificados y pasados a una segunda función que la vez
haciendo uso de otras 2 se encargará de determinar si el patrón recibido cumple con las especificaciones.





## Uso
El proyecto está integrado de tal manera que al momento de descargar el código de github está listo para usarse.
-El código podría ser abierto en un editor de código como lo es visual studio code.
-Una vez esté abierto el código pasaremos a abrir la terminal de dicho editor.
-La terminal debe hacer referencia a que debe estar parada sobre la carpeta src del proyecto
-Una vez seguros de estar en esa carpeta vamos a ejecutar le comando go run main.go
-Esto iniciará un servidor donde se correrá la app
-Por último, para probar el método post vamos a utilizar software llamado postman, 
el cual se encargará de hacer el envió de datos a nuestra app




## Estado del proyecto
El proyecto está: _en progreso_ . Solo hace falta poder pasar el array de bytes que genera el archivo json al la función isMutant.



## Contacto
Creado por [https://github.com/Rafa-Garcia-Dev] rafaelgarciaocampo@gmail.com
