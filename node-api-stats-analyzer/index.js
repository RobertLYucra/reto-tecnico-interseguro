const express = require("express");
const bodyParser = require("body-parser");
const apiRoutes = require("./src/routes/api");

const app = express();

// Middleware
app.use(bodyParser.json({ limit: "10mb" }));

// Rutas
app.use("/", apiRoutes);

// Manejo de errores 404
app.use((req, res) => {
  res.status(404).json({ error: "Ruta no encontrada" });
});

const PORT = process.env.PORT || 3000;

if (require.main === module) {
  app.listen(PORT, () => {
    console.log(`Servidor Node.js corriendo en el puerto ${PORT}`);
  });
}

module.exports = app;
