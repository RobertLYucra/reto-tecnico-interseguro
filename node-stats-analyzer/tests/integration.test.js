const request = require("supertest");
const app = require("../index");

describe("POST /stats (Integración)", () => {
  test("Debería calcular estadísticas correctamente para entrada válida", async () => {
    const payload = {
      q: [
        [1, 0],
        [0, 1],
      ],
      r: [
        [2, 2],
        [0, 2],
      ],
    };

    const response = await request(app)
      .post("/stats")
      .send(payload)
      .expect("Content-Type", /json/)
      .expect(200);

    expect(response.body.success).toBe(true);
    expect(response.body.data.max).toBe(2);
    expect(response.body.data.total_sum).toBe(8);
  });

  test("Debería retornar 400 si falta q o r", async () => {
    const response = await request(app)
      .post("/stats")
      .send({ q: [[1]] }) // Missing r
      .expect(400);

    expect(response.body.success).toBe(false);
    expect(response.body.message).toContain("Faltan matrices");
  });

  test("Debería retornar 500 para datos de matriz inválidos", async () => {
    const response = await request(app)
      .post("/stats")
      .send({
        q: [[1, 2], [3]], // Invalid shape
        r: [
          [1, 0],
          [0, 1],
        ],
      })
      .expect(500);

    expect(response.body.success).toBe(false);
  });
});
