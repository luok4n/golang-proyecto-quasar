
# MeLI: Operacion Fuego de Quasar

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

