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
        name: name,
        address: address,
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

function EditBusiness({ business, onUpdated }: { business: any, onUpdated: (b: any) => void }) {
  const [name, setName] = useState(business.name || "");
  const [address, setAddress] = useState(business.address || "");
  const [phone, setPhone] = useState(business.phone || "");
  const [editing, setEditing] = useState(false);

  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.put('/business', {
        name: name,
        address: address,
        phone: phone,
      });
      onUpdated(res.data);
      setEditing(false);
    } catch (err) {
      alert("Failed to update business!");
    }
  };

  if (!editing) {
    return (
      <button onClick={() => setEditing(true)}>Edit Business</button>
    );
  }

  return (
    <form onSubmit={handleUpdate}>
      <input type="text" placeholder="Business Name" value={name} onChange={e => setName(e.target.value)} />
      <input type="text" placeholder="Address" value={address} onChange={e => setAddress(e.target.value)} />
      <input type="text" placeholder="Phone" value={phone} onChange={e => setPhone(e.target.value)} />
      <button type="submit">Save Changes</button>
      <button type="button" onClick={() => setEditing(false)}>Cancel</button>
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
          <h2>Welcome, {business.name}</h2>
          <p>Address: {business.address}</p>
          <p>Phone: {business.phone}</p>
          <EditBusiness business={business} onUpdated={setBusiness} />
        </div>
      )}
    </div>
  );
}

export default App;
