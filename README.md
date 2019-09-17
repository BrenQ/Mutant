## **Challenge Mercadolibre**

 
 
Magneto quiere reclutar la mayor cantidad de mutantes para poder luchar contra los X-Men. 
 
Te ha contratado a ti para que desarrolles un proyecto que detecte si un humano es mutante basándose en su secuencia de ADN. 
 
Para eso te ha pedido crear un programa con un método o función con la siguiente firma (En alguno de los siguiente lenguajes: Java / Golang / C-C++ / Javascript (node) / Python / Ruby): 
 
**boolean isMutant(String[] dna);**   // Ejemplo Java 
 
En donde recibirás como parámetro un array de Strings que representan cada fila de una tabla de (NxN) con la secuencia del ADN. Las letras de los Strings solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN. 

 ![enter image description here](https://lh3.googleusercontent.com/6SjlRO8nnV3ULxnNT05ClzV8mOiq0ZXjEuhuUXu0CWEreIF4wTYycpA08VQGQa1ojzaFx1i1xsgp)

Sabrás si un humano es mutante, si encuentras ​más de una secuencia de cuatro letras iguales​, de forma oblicua, horizontal o vertical. 
 Ejemplo (Caso mutante): 
 
String[] dna = {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"}; 
 
En este caso el llamado a la función isMutant(dna) devuelve “true”. 
 
Desarrolla el algoritmo de la manera más eficiente posible.

 **Desafíos**: 
 
Nivel 1: Programa (en cualquier lenguaje de programación) que cumpla con el método pedido por Magneto. 
 
Nivel 2: Crear una API REST, hostear esa API en un cloud computing libre (Google App Engine, Amazon AWS, etc), crear el servicio “/mutant/” en donde se pueda detectar si un humano es mutante enviando la secuencia de ADN mediante un HTTP POST con un Json el cual tenga el siguiente formato: 
 
POST → /mutant/ { “dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"] } 
 
En caso de verificar un mutante, debería devolver un HTTP 200-OK, en caso contrario un 403-Forbidden 
 
Nivel 3: Anexar una base de datos, la cual guarde los ADN’s verificados con la API.  Solo 1 registro por ADN.  Exponer un servicio extra “/stats” que devuelva un Json con las estadísticas de las verificaciones de ADN: {“count_mutant_dna”:40, “count_human_dna”:100: “ratio”:0.4} 
 

Tener en cuenta que la API puede recibir fluctuaciones agresivas de tráfico (Entre 100 y 1 millón de peticiones por segundo).

## Instalación 

Para la instalación se recomienda tener en cuenta las configuraciones necesarias para correr Go [enter link description here](https://golang.org/doc/install#install)

A continuación , ejecute el siguiente comando :

    go get -u github.com/BrenQ/Mutant

Con este comando podrá obtener el proyecto :) 

## Usabilidad

Para correr el proyecto deberá ejecutar el comando 

    > go run main.go

Una vez ejecutado el siguiente comando se quedará el servidor escuchando en el puerto :5000 

Para verificar que su secuencia de ADN pertenece a un humano o un mutante va a tener que hacer la petición a la siguiente ruta con la estructura que se brinda a continuación:


 - URL : `/mutant`
 - Method: `POST`
 - Auth required : `No`
 - Data required : 
     - `dna`

   Example
	  
>     POST /mutant
>            {
>             		  "dna": [
>             		    "ATGCGA",
>             		    "CAGTGC",
>             		    "TTATGT",
>             		    "AGAAGG",
>             		    "CCCCTA",
>             		    "TCACTG"
>            ]  
>       }
	

## Success response

**Condicion** : Si se verifica que un humano es mutante :
**Code**: `200 - OK`

## Error response

**Condicion** : Si un humano no es mutante:
**Code**: `403-Forbidden `

**Estadisticas**

Tambien la APP permite obtener estadísticas en base a la información almacenada :) 

Puede consultar esta info

- URL : `/stats`
 - Method: `GET`
 - Auth required : `No`
 - Data required : 
     - `No`
     -
    GET /stats
 
	*Response*
	
       {
           “count_mutant_dna”:40, 
           “count_human_dna”:100: 
           “ratio”:0.4
       } 

## Success response

**Condicion** : Devuelve la estadística almacenada :
**Code**: `200 - OK`

