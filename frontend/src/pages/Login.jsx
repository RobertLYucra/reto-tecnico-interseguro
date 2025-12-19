import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await api.post('/login', { username, password });
      if (response.data.success) {
        localStorage.setItem('token', response.data.data.token);
        navigate('/dashboard');
      } else {
        setError(response.data.message);
      }
    } catch (err) {
      if (err.response && err.response.data) {
        const errorMsg = err.response.data.error 
            ? (Array.isArray(err.response.data.error) ? err.response.data.error[0] : err.response.data.error) 
            : err.response.data.message;
        setError(errorMsg || 'Login fallido');
      } else {
        setError('Error de conexión');
      }
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900 text-white overflow-hidden relative">
      {/* Background Decor */}
      <div className="absolute top-0 left-0 w-full h-full overflow-hidden z-0">
          <div className="absolute w-96 h-96 bg-blue-500/20 rounded-full blur-3xl -top-20 -left-20 animate-pulse"></div>
          <div className="absolute w-96 h-96 bg-purple-500/20 rounded-full blur-3xl bottom-0 right-0 animate-pulse delay-1000"></div>
      </div>

      <div className="bg-white/10 backdrop-blur-lg border border-white/20 p-8 rounded-2xl shadow-2xl w-full max-w-md z-10 relative">
        <h2 className="text-3xl font-bold mb-2 text-center text-white tracking-tight">Interseguro Matrix</h2>
        <p className="text-center text-blue-200 text-sm mb-8 font-light">Acceso Seguro al Sistema</p>
        
        {error && (
          <div className="bg-red-500/80 backdrop-blur-sm text-white p-3 rounded-lg mb-6 text-sm text-center shadow-lg border border-red-400/50 animate-bounce-short">
            {error}
          </div>
        )}
        
        <form onSubmit={handleLogin} className="space-y-6">
          <div className="space-y-1">
            <label className="block text-xs font-semibold uppercase tracking-wider text-blue-100 ml-1">Usuario</label>
            <input
              type="text"
              className="w-full bg-slate-900/50 border border-slate-600/50 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:border-blue-400 focus:ring-1 focus:ring-blue-400 transition-all duration-300 hover:bg-slate-900/70"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Ej: admin"
            />
          </div>
          <div className="space-y-1">
            <label className="block text-xs font-semibold uppercase tracking-wider text-blue-100 ml-1">Contraseña</label>
            <input
              type="password"
              className="w-full bg-slate-900/50 border border-slate-600/50 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:border-blue-400 focus:ring-1 focus:ring-blue-400 transition-all duration-300 hover:bg-slate-900/70"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
            />
          </div>
          <button
            type="submit"
            className="w-full bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 text-white font-bold py-3 px-4 rounded-xl shadow-lg transform transition duration-200 hover:scale-[1.02] active:scale-[0.98] border border-white/10 mt-2"
          >
            Iniciar Sesión
          </button>
        </form>
      </div>
    </div>
  );
};

export default Login;
