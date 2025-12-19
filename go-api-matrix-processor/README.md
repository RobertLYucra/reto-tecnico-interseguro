# Go Matrix Processor API

Servicio backend encargado de recibir la matriz, rotarla y realizar la factorización QR.

## 🛠️ Requisitos

- Go 1.21+

## 📥 Instalación

1.  Ubicarse en la carpeta del proyecto:

    ```bash
    cd go-matrix-processor
    ```

2.  Descargar dependencias:
    ```bash
    go mod tidy
    ```

## 🚀 Ejecución (Local)

Para levantar el servidor en el puerto 8080:

```bash
go run cmd/api/main.go
```

## 🧪 Testing

Ejecutar todas las pruebas unitarias:

```bash
go test ./...
```
