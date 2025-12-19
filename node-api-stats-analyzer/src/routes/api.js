const express = require("express");
const router = express.Router();
const statsController = require("../controllers/statsController");

router.post("/stats", (req, res) => statsController.getStats(req, res));

module.exports = router;
