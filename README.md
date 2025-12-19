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

## 📡 Documentación de la API

### 1. Autenticación (Login)

Genera un Token JWT para acceder a los endpoints protegidos.

- **URL**: `/api/v1/login`
- **Método**: `POST`

**Body (JSON):**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Respuesta Exitosa (200 OK):**

```json
{
  "success": true,
  "message": "Login exitoso",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

### 2. Procesar Matriz

Rota la matriz 90° anti-horaria, realiza la factorización QR y calcula estadísticas.

- **URL**: `/api/v1/process`
- **Método**: `POST`
- **Headers**:
  - `Authorization`: `Bearer <TU_TOKEN_JWT>`

**Body (JSON):**

```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

**Respuesta Exitosa (200 OK):**

```json
{
  "success": true,
  "message": "Matriz procesada correctamente",
  "data": {
    "q": [
      [ -0.333, 0.666, -0.666 ],
      ...
    ],
    "r": [
      [ -9.0, 0.0, 0.0 ],
      ...
    ],
    "max": 9,
    "min": 1,
    "average": 5,
    "total_sum": 45,
    "is_q_diagonal": false,
    "is_r_diagonal": false
  }
}
```

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

## 🔧 Instalación y Ejecución Manual (Sin Docker)

Si prefieres ejecutar cada servicio individualmente en tu máquina local:

### 1. Go Backend (Puerto 8080)

```bash
cd go-matrix-processor
go mod tidy       # Instalar dependencias
go run cmd/api/main.go
```

### 2. Node.js Backend (Puerto 3000)

```bash
cd node-stats-analyzer
npm install       # Instalar dependencias
npm start
```

### 3. Frontend (Puerto 5173)

```bash
cd frontend
npm install       # Instalar dependencias
npm run dev
```
