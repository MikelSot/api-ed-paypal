# PAGOS CON PAYPAL

### ARQUITECTURA HEXAGONAL

Es una arquitectura de software en la que se busca es separar el core 
logico de la aplicacion, dejarlo en el centro totalmente aislado del exterior,
del cliente y de otras interacciones.

![](.img/hexagonal.png)

- **Infraestructura**: son los elementos externos con los que se comunica la aplicación, tanto de entrada como de salida.
    Puntos de entrada son una API con REST o GraphQL, mensajería con RabbitMQ o mediante línea de comandos.
    Puntos de salida son una base de datos relacional con PostgreSQL, no relacional con MongoDB, o también envío de mensajes con RabbitMQ. A los puntos de entrada se les denomina puertos y a los puntos de salida adaptadores.
  
  - **Puertos**: una aplicación puede ofrecer diferentes formas de comunicación al mismo tiempo, ya sea con una API REST o GraphQL, mensajes o mediante línea de comandos.
  - **Adaptadores**: igualmente una aplicación puede utilizar diferentes bases de datos o sistemas de comunicación con el exterior.
- **Aplicación**: son los servicios que definen la API pública del dominio e independiza al dominio de cualesquiera 
              elementos de infraestructura actuales o en el futuro.
- **Dominio**: contiene la lógica de negocio de la aplicación. Esta puede ser implementada usando los principios de DDD.


`Extructura del proyecto`

```bash
├── cmd
├── domain
│   └── invoice
│       └── invoice.go
│       └── usecase.go
│
├── infrastructure
│   ├── handler
│   │   ├── request
│   │   ├── response
│   │   ├── invoice
│   │       ├── handler.go
│   │       └── route.go
│   └── postgres
│       └── invoice
│           └── invoice.go
├── model
    └── invoice.go
```
