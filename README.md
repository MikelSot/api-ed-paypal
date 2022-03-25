# PAGOS CON PAYPAL

### ARQUITECTURA HEXAGONAL

Es una arquitectura de software en la que se busca es separar el core 
logico de la aplicacion, dejarlo en el centro totalmente aislado del exterior,
del cliente y de otras interacciones.

![](.img/hexagonal.png)

Extructura del proyecto

```bash
├── cmd
├── domain
│   └── interface_storage.go
│   └── interface_usecase.go
│   └── usecase.go
│
├── infrastructure
│   ├── handler
│   │   ├── request
│   │   ├── response
│   │   ├── paypal
│   │       ├── handler.go
│   │       └── route.go
│   └── postgres
│       └── userrole
│           └── userrole.go
├── model
    └── product.go
```
