const statsService = require("../services/statsService");
const ResponseUtil = require("../utils/responseUtil");

class StatsController {
  getStats(req, res) {
    try {
      const { q: qMatrix, r: rMatrix } = req.body;

      if (!qMatrix || !rMatrix) {
        return res
          .status(400)
          .json(ResponseUtil.error("Faltan matrices Q o R en la solicitud"));
      }

      const stats = statsService.calculateStats(qMatrix, rMatrix);
      return res.json(
        ResponseUtil.success("Estadísticas calculadas correctamente", stats)
      );
    } catch (error) {
      console.log("Error calculating statistics:", error.message);
      return res
        .status(500)
        .json(
          ResponseUtil.error(error.message || "Error interno del servidor")
        );
    }
  }
}

module.exports = new StatsController();
