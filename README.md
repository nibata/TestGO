# TestGO
proyecto para experimentar con golang

De momento se usa se está utilizando la librería `ECHO` para crear APIs.

El proyecto cuenta con tres endpoints:

 - `[::]:3000` ccon un mensaje de bienvenida simple y plano.
 - `[::]:3000/usopathparams/{ALGUN_VALOR}` que muestra como se utiliza un endpoint con Path Params.
 - `[::]:3000/usoqueryparams?value={ALGUN_VALOR}` que muestra como se utiliza un endpoint con Query Params.

 Para iniciar el proyecto primero es necesario descargar dependencia con:
`go mod vendor`

Y finalmente se inicia el proyecto con:
`go run sserver.go`

Proximos pasos:
 - estructurar mejor el proyecto de manera que sea editable de mejor manera.
 - dejar funcionando este proyecto en Docker.
