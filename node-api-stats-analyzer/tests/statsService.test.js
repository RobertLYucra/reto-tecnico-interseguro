const service = require("../src/services/statsService");

describe("StatsService", () => {
  // service is already the instance

  test("calculateStats debería retornar estadísticas correctas para matrices válidas", () => {
    // Arrange
    const qMatrix = [
      [1, 0],
      [0, 1],
    ];
    const rMatrix = [
      [2, 0],
      [0, 3],
    ];

    // Act
    const result = service.calculateStats(qMatrix, rMatrix);

    // Assert
    expect(result.max).toBe(3);
    expect(result.min).toBe(0);
    expect(result.average).toBe(0.875); // (1+0+0+1 + 2+0+0+3) / 8 = 7/8
    expect(result.total_sum).toBe(7);
    expect(result.is_q_diagonal).toBe(true);
    expect(result.is_r_diagonal).toBe(true);
  });

  test("calculateStats debería detectar matriz no diagonal", () => {
    // Arrange
    const qMatrix = [
      [1, 2],
      [3, 4],
    ]; // Non-diagonal
    const rMatrix = [
      [1, 0],
      [0, 1],
    ]; // Diagonal

    // Act
    const result = service.calculateStats(qMatrix, rMatrix);

    // Assert
    expect(result.is_q_diagonal).toBe(false);
    expect(result.is_r_diagonal).toBe(true);
  });

  test("calculateStats debería lanzar error para datos vacíos", () => {
    expect(() => {
      service.calculateStats([], []);
    }).toThrow("No data to process");
  });
});
