class StatsService {
  calculateStats(qMatrix, rMatrix) {
    if (!this._isValidMatrix(qMatrix) || !this._isValidMatrix(rMatrix)) {
      throw new Error("Invalid or empty matrices");
    }

    const allValues = [...qMatrix.flat(), ...rMatrix.flat()];

    if (allValues.length === 0) throw new Error("No data to process");

    const isQDiagonal = this._isDiagonal(qMatrix);
    const isRDiagonal = this._isDiagonal(rMatrix);

    const sum = allValues.reduce((a, b) => a + b, 0);

    return {
      max: Math.max(...allValues),
      min: Math.min(...allValues),
      average: sum / allValues.length,
      total_sum: sum,
      is_q_diagonal: isQDiagonal,
      is_r_diagonal: isRDiagonal,
    };
  }

  _isValidMatrix(matrix) {
    return (
      Array.isArray(matrix) && matrix.length > 0 && Array.isArray(matrix[0])
    );
  }

  _isDiagonal(matrix) {
    return matrix.every((row, i) =>
      row.every((val, j) => {
        return i === j || Math.abs(val) <= 1e-9;
      })
    );
  }
}

module.exports = new StatsService();
