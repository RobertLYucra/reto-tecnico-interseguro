class StatsService {
  calculateStats(qMatrix, rMatrix) {
    if (!this._isValidMatrix(qMatrix) || !this._isValidMatrix(rMatrix)) {
      throw new Error("Invalid or empty matrices");
    }

    // 1. Flatten all values for stats (Declarative & Clean)
    const allValues = [...qMatrix.flat(), ...rMatrix.flat()];

    if (allValues.length === 0) throw new Error("No data to process");

    // 2. Independent Diagonal Checks
    const isQDiagonal = this._isDiagonal(qMatrix);
    const isRDiagonal = this._isDiagonal(rMatrix);

    // 3. Calculate Stats
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
    // "Senior" Approach: Functional .every() instead of nested loops
    // Returns true if ALL non-diagonal elements are effectively 0
    return matrix.every((row, i) =>
      row.every((val, j) => {
        // If it's the diagonal (i===j), any value is fine.
        // If it's NOT the diagonal, value must be 0.
        return i === j || Math.abs(val) <= 1e-9;
      })
    );
  }
}

module.exports = new StatsService();
