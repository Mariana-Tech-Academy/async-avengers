import { useEffect, useState } from 'react';
import api from './api'; // This uses your api.ts file
import './App.css';

function App() {
  const [business, setBusiness] = useState<any>(null);
  const [error, setError] = useState("");

  useEffect(() => {
    // Calling your Go backend
    api.get('/business')
      .then(res => setBusiness(res.data))
      .catch(err => {
        setError("Make sure your Go server is running!");
        console.error(err);
      });
  }, []);

  return (
    <div style={{ padding: '40px' }}>
      <h1>Invoice Dashboard</h1>
      <hr />
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {business ? (
        <div>
          <h3>Connected!</h3>
          <pre>{JSON.stringify(business, null, 2)}</pre>
        </div>
      ) : (
        <p>Loading data from Go...</p>
      )}
    </div>
  );
}

export default App;