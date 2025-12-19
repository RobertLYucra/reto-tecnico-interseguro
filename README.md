# Matrix QR Processor & Statistics Analyzer

Este proyecto implementa una solución completa de Microservicios para el procesamiento de matrices y cálculo de estadísticas, cumpliendo con los requisitos del desafío técnico.

## 🚀 Arquitectura

El sistema consta de 3 servicios contenerizados:

1.  **Backend Go (`go-matrix-processor`)**:

    - **Puerto**: 8080
    - **Responsabilidad**: Recibe la matriz original, realiza una **Rotación de 90° Anti-Horaria**, y calcula la Factorización QR (Matrices Q y R).
    - **Seguridad**: Protegido con JWT.

2.  **Backend Node.js (`node-stats-analyzer`)**:

    - **Puerto**: 3000
    - **Responsabilidad**: Recibe las matrices Q y R (ya procesadas) y calcula estadísticas (Máximo, Mínimo, Promedio, Suma, Diagonalidad).
    - **Nota**: Las estadísticas se calculan sobre el conjunto total de valores de Q y R.

3.  **Frontend (`frontend`)**:
    - **Puerto**: 5173
    - **Tecnología**: React + Vite + TailwindCSS.
    - **Funcionalidad**: Dashboard interactivo para ingresar matrices y visualizar resultados (Tablas y Stats).

## 🛠️ Cómo Ejecutar

Requisitos: Docker y Docker Compose.

1.  Clonar el repositorio (o ubicar la carpeta raíz).
2.  Ejecutar el siguiente comando para construir y levantar los servicios:

```bash
docker compose up --build
```

3.  Acceder al Frontend en: **[http://localhost:5173](http://localhost:5173)**

## 🔑 Credenciales

Para iniciar sesión en la aplicación:

- **Usuario**: `admin`
- **Contraseña**: `admin`

## ✅ Endpoints Principales

- `POST /api/v1/login`: Generación de Token JWT.
- `POST /api/v1/process`: Procesa la matriz (Requiere Header `Authorization: Bearer <token>`).

## 🧪 Tests

Para ejecutar las pruebas unitarias y de integración:

**Go:**

```bash
cd go-matrix-processor
go test ./...
```

**Node.js:**

```bash
cd node-stats-analyzer
npm test
```
