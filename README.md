# TestGO
proyecto para experimentar con golang

Para este proyecto me guío por las guías que están en (en especial para la instalación de Swagger):
 
 - https://github.com/swaggo/echo-swagger
 - https://medium.com/geekculture/tutorial-generate-swagger-specification-and-swaggerui-for-echo-go-web-framework-3ac33afc77e2

De momento se usa se está utilizando la librería `ECHO` para crear APIs.

El proyecto cuenta con tres endpoints:

 - `[::]:3000` con un mensaje de bienvenida simple y plano.
 - `[::]:3000/docs/index.html` contiene la documentación SWAGGER.
 - `[::]:3000/usopathparams/{ALGUN_VALOR}` que muestra como se utiliza un endpoint con Path Params.
 - `[::]:3000/usoqueryparams?value={ALGUN_VALOR}` que muestra como se utiliza un endpoint con Query Params.

 Para iniciar el proyecto primero es necesario descargar dependencia con:
`go mod vendor`

Y finalmente se inicia el proyecto con:
`go run server.go`

Para actualizar la documentación:
`swag init -g EchoApp/server.go --output docs/EchoApp`

Proximos pasos:
 - estructurar mejor el proyecto de manera que sea editable de mejor manera.
 - dejar funcionando este proyecto en Docker.
