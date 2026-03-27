import { useState } from 'react';
import api from './api';

function CreateBusiness({ onCreated }: { onCreated: (b: any) => void }) {
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");
  const [phone, setPhone] = useState("");

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/business', {
        buisnessname: name,
        buisnessaddress: address,
        phone: phone,
      });
      onCreated(res.data);
    } catch (err) {
      alert("Failed to create business!");
    }
  };

  return (
    <form onSubmit={handleCreate}>
      <input type="text" placeholder="Business Name" onChange={e => setName(e.target.value)} />
      <input type="text" placeholder="Address" onChange={e => setAddress(e.target.value)} />
      <input type="text" placeholder="Phone" onChange={e => setPhone(e.target.value)} />
      <button type="submit">Create Business</button>
    </form>
  );
}

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [business, setBusiness] = useState<any>(null);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/login', { email, password });
      const token = res.data;
      localStorage.setItem("token", token);
      api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      try {
        const bizRes = await api.get('/business');
        setBusiness(bizRes.data);
      } catch (err: any) {
        if (err.response?.status === 404) {
          setBusiness("none");
        }
      }
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
      ) : business === "none" ? (
        <div>
          <h2>Create Your Business</h2>
          <CreateBusiness onCreated={setBusiness} />
        </div>
      ) : (
        <div>
          <h2>Welcome, {business.buisnessname}</h2>
          <pre>{JSON.stringify(business, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default App;