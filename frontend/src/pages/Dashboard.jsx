import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api';

const Dashboard = () => {
  const navigate = useNavigate();
  const [rows, setRows] = useState(3);
  const [cols, setCols] = useState(3);
  const [matrixData, setMatrixData] = useState([]); // 2D array
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  // Protect route
  useEffect(() => {
    if (!localStorage.getItem('token')) {
      navigate('/');
    }
  }, [navigate]);

  // Initialize/Resize Matrix inputs
  useEffect(() => {
    const newMatrix = Array(rows).fill(0).map(() => Array(cols).fill(0));
    setMatrixData(newMatrix);
  }, [rows, cols]);

  const handleInputChange = (r, c, val) => {
    const newData = [...matrixData];
    newData[r] = [...newData[r]];
    newData[r][c] = parseFloat(val) || 0;
    setMatrixData(newData);
  };

  const handleProcess = async () => {
    setLoading(true);
    setError('');
    setResult(null);
    try {
      const response = await api.post('/process', { data: matrixData });
      if (response.data.success) {
        setResult(response.data.data);
      } else {
        setError(response.data.message);
      }
    } catch (err) {
      if (err.response && err.response.status === 401) {
        localStorage.removeItem('token');
        navigate('/');
      } else {
        const errorMsg = err.response && err.response.data && err.response.data.error
            ? (Array.isArray(err.response.data.error) ? err.response.data.error[0] : err.response.data.message)
            : 'Error al procesar';
        setError(errorMsg);
      }
    } finally {
      setLoading(false);
    }
  };

  const MatrixView = ({ data, title }) => (
    <div className="bg-white/10 backdrop-blur-md border border-white/10 p-4 rounded-xl shadow-lg">
      <h3 className="text-xl font-bold mb-4 text-blue-200 text-center tracking-wide">{title}</h3>
      <div className="overflow-x-auto">
        <table className="min-w-full text-center border-collapse">
          <tbody>
            {data.map((row, i) => (
              <tr key={i}>
                {row.map((val, j) => (
                  <td key={j} className="border border-white/10 p-2 text-sm text-gray-200 font-mono">
                    {val.toFixed(4)}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900 text-white p-6 relative overflow-hidden">
      {/* Background Decor */}
      <div className="absolute top-0 left-0 w-full h-full overflow-hidden z-0 pointer-events-none">
          <div className="absolute w-[800px] h-[800px] bg-blue-500/10 rounded-full blur-3xl -top-40 -left-40 animate-pulse"></div>
          <div className="absolute w-[600px] h-[600px] bg-purple-500/10 rounded-full blur-3xl bottom-0 right-0 animate-pulse delay-700"></div>
      </div>

      <div className="max-w-6xl mx-auto relative z-10">
        {/* Header */}
        <div className="flex justify-between items-center mb-8 bg-white/5 backdrop-blur-md p-4 rounded-xl border border-white/10">
          <h1 className="text-2xl md:text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-300 to-purple-300 tracking-tight">
            Matrix Processor QR
          </h1>
          <button
            onClick={() => { localStorage.removeItem('token'); navigate('/'); }}
            className="bg-red-500/20 hover:bg-red-500/40 text-red-200 border border-red-500/30 px-4 py-2 rounded-lg text-sm transition-all duration-300"
          >
            Cerrar Sesión
          </button>
        </div>

        {/* Configuration & Input */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
          {/* Config */}
          <div className="lg:col-span-1 bg-white/10 backdrop-blur-lg border border-white/10 p-6 rounded-2xl shadow-xl h-fit">
            <h2 className="text-lg font-semibold mb-4 text-blue-100 border-b border-white/10 pb-2">Configuración</h2>
            <div className="flex gap-4 mb-6">
              <div className="flex-1">
                <label className="block text-xs mb-2 text-gray-400 uppercase tracking-wider">Filas</label>
                <input
                  type="number"
                  min="2"
                  className="w-full bg-slate-900/50 border border-slate-600/50 rounded-lg px-3 py-2 text-white focus:border-blue-400 focus:outline-none focus:ring-1 focus:ring-blue-400 transition-colors"
                  value={rows}
                  onChange={(e) => setRows(parseInt(e.target.value) || 2)}
                />
              </div>
              <div className="flex-1">
                <label className="block text-xs mb-2 text-gray-400 uppercase tracking-wider">Columnas</label>
                <input
                  type="number"
                  min="2"
                  className="w-full bg-slate-900/50 border border-slate-600/50 rounded-lg px-3 py-2 text-white focus:border-blue-400 focus:outline-none focus:ring-1 focus:ring-blue-400 transition-colors"
                  value={cols}
                  onChange={(e) => setCols(parseInt(e.target.value) || 2)}
                />
              </div>
            </div>
            
            <div className="flex gap-3">
              <button
                onClick={handleProcess}
                disabled={loading}
                className={`flex-1 py-3 rounded-xl font-bold shadow-lg transform transition-all duration-200 hover:scale-[1.02] active:scale-[0.98] ${
                  loading 
                    ? 'bg-gray-600/50 cursor-not-allowed text-gray-400' 
                    : 'bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 text-white border border-white/10'
                }`}
              >
                {loading ? 'Procesando...' : 'Calcular'}
              </button>
              
              <button
                onClick={() => {
                   const newMatrix = Array(rows).fill(0).map(() => Array(cols).fill(0));
                   setMatrixData(newMatrix);
                   setResult(null);
                   setError('');
                }}
                disabled={loading}
                className="px-4 py-3 rounded-xl font-bold bg-white/5 hover:bg-white/10 text-white border border-white/10 transition-colors"
                title="Reiniciar Matriz"
              >
                ↻
              </button>
            </div>
            
            {error && (
              <div className="mt-6 bg-red-500/20 text-red-100 p-3 rounded-lg text-xs border border-red-500/30 text-center animate-bounce-short">
                {error}
              </div>
            )}
          </div>

          {/* Matrix Input Grid */}
          <div className="lg:col-span-2 bg-white/10 backdrop-blur-lg border border-white/10 p-6 rounded-2xl shadow-xl">
            <h2 className="text-lg font-semibold mb-4 text-blue-100 border-b border-white/10 pb-2">Matriz de Entrada</h2>
            <div className="overflow-auto max-h-[500px] scrollbar-thin scrollbar-thumb-blue-500/30 scrollbar-track-transparent pr-2">
                <div 
                  className="grid gap-3"
                  style={{ gridTemplateColumns: `repeat(${cols}, minmax(70px, 1fr))` }}
                >
                  {matrixData.map((row, i) => (
                    row.map((val, j) => (
                      <input
                        key={`${i}-${j}`}
                        type="number"
                        step="any"
                        className="bg-slate-900/60 border border-slate-700/50 rounded-lg px-2 py-3 text-center text-sm focus:border-blue-400 focus:bg-slate-800 focus:outline-none transition-all duration-200 font-mono hover:bg-slate-800/50"
                        value={matrixData[i][j]} 
                        onFocus={(e) => e.target.select()}
                        onChange={(e) => handleInputChange(i, j, e.target.value)}
                      />
                    ))
                  ))}
                </div>
            </div>
          </div>
        </div>

        {/* Results */}
        {result && (
          <div className="space-y-8 animate-fade-in-up">
             <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                <div className="bg-white/5 backdrop-blur-md p-5 rounded-xl border border-white/10 hover:bg-white/10 transition-colors">
                    <p className="text-blue-300 text-xs uppercase tracking-wider mb-1">Máximo</p>
                    <p className="text-3xl font-bold text-white">{result.max.toFixed(4)}</p>
                </div>
                <div className="bg-white/5 backdrop-blur-md p-5 rounded-xl border border-white/10 hover:bg-white/10 transition-colors">
                    <p className="text-blue-300 text-xs uppercase tracking-wider mb-1">Mínimo</p>
                    <p className="text-3xl font-bold text-white">{result.min.toFixed(4)}</p>
                </div>
                <div className="bg-white/5 backdrop-blur-md p-5 rounded-xl border border-white/10 hover:bg-white/10 transition-colors">
                    <p className="text-blue-300 text-xs uppercase tracking-wider mb-1">Promedio</p>
                    <p className="text-3xl font-bold text-white">{result.average.toFixed(4)}</p>
                </div>
                <div className="bg-gradient-to-br from-green-500/20 to-emerald-500/20 backdrop-blur-md p-5 rounded-xl border border-green-500/30">
                    <p className="text-green-300 text-xs uppercase tracking-wider mb-1">Suma Total</p>
                    <p className="text-3xl font-bold text-white">{result.total_sum.toFixed(4)}</p>
                </div>
             </div>

             <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                {result.q && <MatrixView data={result.q} title="Matriz Q (Ortogonal)" />}
                {result.r && <MatrixView data={result.r} title="Matriz R (Triangular Superior)" />}
             </div>

             <div className="flex justify-center gap-8 text-sm bg-black/20 p-4 rounded-full w-fit mx-auto backdrop-blur-sm border border-white/5">
                 <span className={`flex items-center gap-2 ${result.is_q_diagonal ? "text-green-400" : "text-gray-400"}`}>
                    <span className={`w-2 h-2 rounded-full ${result.is_q_diagonal ? "bg-green-400" : "bg-gray-400"}`}></span>
                    Q Diagonal: {result.is_q_diagonal ? 'SÍ' : 'NO'}
                 </span>
                 <div className="w-px bg-white/10"></div>
                 <span className={`flex items-center gap-2 ${result.is_r_diagonal ? "text-green-400" : "text-gray-400"}`}>
                    <span className={`w-2 h-2 rounded-full ${result.is_r_diagonal ? "bg-green-400" : "bg-gray-400"}`}></span>
                    R Diagonal: {result.is_r_diagonal ? 'SÍ' : 'NO'}
                 </span>
             </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Dashboard;
