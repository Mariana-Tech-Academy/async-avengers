import { useState } from 'react';
import api from './api';

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [business, setBusiness] = useState<any>(null);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      // 1. Call the login route
      const res = await api.post('/login', { email, password });
      console.log("Login response:", res.data);
      const token = res.data;
      console.log("Token:", token);

      // 2. Save token to LocalStorage
      localStorage.setItem("token", token);

      // 3. Tell axios to use this token for all future calls
      api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      // 4. Now try to get the protected data
      const bizRes = await api.get('/business');
      setBusiness(bizRes.data);
    } catch (err) {
      alert("Login failed!");
    }
  };

  return (
    <div style={{ padding: '40px' }}>
      <h1>Invoice System</h1>
      {!business ? (
        <form onSubmit={handleLogin}>
          <input type="email" placeholder="Email" onChange={e => setEmail(e.target.value)} />
          <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
          <button type="submit">Login</button>
        </form>
      ) : (
        <div>
          <h2>Welcome, {business.name}</h2>
          <pre>{JSON.stringify(business, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default App;