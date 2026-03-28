import { useState, useEffect } from 'react';
import api from './api';

// ─── Types ────────────────────────────────────────────────────────────────────
type View = 'login' | 'register' | 'business-setup' | 'app';
type Section = 'dashboard' | 'business' | 'clients' | 'products' | 'invoices';

interface Business {
  ID: number; name: string; address: string; phone: string;
  email: string; Logo: string; vat_number: string; tax_rate: number;
}
interface Client {
  ID: number; name: string; email: string; phone: string; address: string;
}
interface Product {
  ID: number; name: string; description: string; price: number; unit: string;
}
interface Invoice {
  ID: number; client_id: number; status: string; total: number;
  CreatedAt: string; due_date: string; notes: string;
  items?: InvoiceItem[];
}
interface InvoiceItem {
  product_id: number; quantity: number; price: number;
}

// ─── Global Styles ────────────────────────────────────────────────────────────
const STYLES = `
  @import url('https://fonts.googleapis.com/css2?family=Quicksand:wght@400;500;600;700&family=Exo+2:wght@300;400;600;700;800&display=swap');

  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  :root {
    --deep: #04041a;
    --navy: #080824;
    --card: #0e0e35;
    --card2: #12123d;
    --sidebar: #090928;
    --purple: #7c3aed;
    --purple2: #9f67ff;
    --purple-soft: #c4b5fd;
    --purple-glow: rgba(124, 58, 237, 0.35);
    --purple-glow2: rgba(159, 103, 255, 0.15);
    --gold: #fbbf24;
    --pink: #f472b6;
    --cyan: #67e8f9;
    --text: #ede9fe;
    --text2: #a89fd4;
    --text3: #6b63a0;
    --border: rgba(124, 58, 237, 0.25);
    --border2: rgba(124, 58, 237, 0.5);
    --success: #34d399;
    --danger: #f87171;
    --warning: #fbbf24;
  }

  body {
    font-family: 'Quicksand', sans-serif;
    background: var(--deep);
    color: var(--text);
    min-height: 100vh;
    overflow-x: hidden;
  }

  /* Stars background */
  body::before {
    content: '';
    position: fixed; inset: 0;
    background-image:
      radial-gradient(1px 1px at 10% 20%, rgba(255,255,255,0.8) 0%, transparent 100%),
      radial-gradient(1px 1px at 30% 60%, rgba(255,255,255,0.6) 0%, transparent 100%),
      radial-gradient(1.5px 1.5px at 50% 10%, rgba(255,255,255,0.9) 0%, transparent 100%),
      radial-gradient(1px 1px at 70% 40%, rgba(255,255,255,0.7) 0%, transparent 100%),
      radial-gradient(1px 1px at 85% 80%, rgba(255,255,255,0.5) 0%, transparent 100%),
      radial-gradient(1.5px 1.5px at 15% 75%, rgba(255,255,255,0.8) 0%, transparent 100%),
      radial-gradient(1px 1px at 55% 55%, rgba(255,255,255,0.6) 0%, transparent 100%),
      radial-gradient(1px 1px at 90% 15%, rgba(255,255,255,0.9) 0%, transparent 100%),
      radial-gradient(1px 1px at 40% 90%, rgba(255,255,255,0.7) 0%, transparent 100%),
      radial-gradient(2px 2px at 25% 35%, rgba(196,181,253,0.6) 0%, transparent 100%),
      radial-gradient(1.5px 1.5px at 75% 65%, rgba(196,181,253,0.5) 0%, transparent 100%),
      radial-gradient(1px 1px at 60% 25%, rgba(255,255,255,0.8) 0%, transparent 100%),
      radial-gradient(1px 1px at 5% 50%, rgba(255,255,255,0.6) 0%, transparent 100%),
      radial-gradient(1px 1px at 95% 45%, rgba(255,255,255,0.7) 0%, transparent 100%),
      radial-gradient(1.5px 1.5px at 45% 70%, rgba(255,255,255,0.5) 0%, transparent 100%);
    pointer-events: none;
    z-index: 0;
  }

  h1, h2, h3, h4 { font-family: 'Exo 2', sans-serif; }

  /* Scrollbar */
  ::-webkit-scrollbar { width: 6px; }
  ::-webkit-scrollbar-track { background: var(--navy); }
  ::-webkit-scrollbar-thumb { background: var(--purple); border-radius: 3px; }

  /* Auth Screen */
  .auth-wrap {
    min-height: 100vh; display: flex; align-items: center; justify-content: center;
    position: relative; z-index: 1; padding: 24px;
  }
  .auth-card {
    background: var(--card);
    border: 1px solid var(--border2);
    border-radius: 24px;
    padding: 48px 40px;
    width: 100%; max-width: 420px;
    box-shadow: 0 0 60px var(--purple-glow), 0 0 120px rgba(124,58,237,0.1);
    animation: fadeUp 0.6s ease;
  }
  .auth-logo {
    text-align: center; margin-bottom: 32px;
  }
  .auth-logo h1 {
    font-size: 28px; font-weight: 800;
    background: linear-gradient(135deg, var(--purple-soft), var(--gold));
    -webkit-background-clip: text; -webkit-text-fill-color: transparent;
    background-clip: text;
  }
  .auth-logo p { color: var(--text2); font-size: 14px; margin-top: 6px; }
  .auth-logo .star { font-size: 40px; display: block; margin-bottom: 12px; animation: spin 8s linear infinite; }

  /* Form Elements */
  .form-group { margin-bottom: 16px; }
  .form-group label { display: block; font-size: 13px; font-weight: 600; color: var(--purple-soft); margin-bottom: 6px; }
  .form-input {
    width: 100%; padding: 12px 16px;
    background: var(--card2); border: 1px solid var(--border);
    border-radius: 12px; color: var(--text);
    font-family: 'Quicksand', sans-serif; font-size: 14px;
    transition: all 0.2s;
    outline: none;
  }
  .form-input:focus { border-color: var(--purple2); box-shadow: 0 0 0 3px var(--purple-glow2); }
  .form-input::placeholder { color: var(--text3); }
  select.form-input option { background: var(--card2); }

  /* Buttons */
  .btn {
    display: inline-flex; align-items: center; gap: 8px;
    padding: 11px 22px; border-radius: 12px; font-size: 14px;
    font-weight: 700; font-family: 'Quicksand', sans-serif;
    cursor: pointer; border: none; transition: all 0.2s; text-decoration: none;
  }
  .btn-primary {
    background: linear-gradient(135deg, var(--purple), var(--purple2));
    color: white; box-shadow: 0 4px 20px var(--purple-glow);
  }
  .btn-primary:hover { transform: translateY(-2px); box-shadow: 0 6px 28px var(--purple-glow); }
  .btn-ghost {
    background: transparent; color: var(--purple-soft);
    border: 1px solid var(--border2);
  }
  .btn-ghost:hover { background: var(--purple-glow2); }
  .btn-danger { background: rgba(248,113,113,0.15); color: var(--danger); border: 1px solid rgba(248,113,113,0.3); }
  .btn-danger:hover { background: rgba(248,113,113,0.25); }
  .btn-success { background: rgba(52,211,153,0.15); color: var(--success); border: 1px solid rgba(52,211,153,0.3); }
  .btn-success:hover { background: rgba(52,211,153,0.25); }
  .btn-full { width: 100%; justify-content: center; padding: 14px; font-size: 15px; }
  .btn-sm { padding: 7px 14px; font-size: 12px; border-radius: 8px; }
  .btn:disabled { opacity: 0.5; cursor: not-allowed; transform: none !important; }

  /* App Layout */
  .app-layout { display: flex; height: 100vh; position: relative; z-index: 1; }

  /* Sidebar */
  .sidebar {
    width: 240px; min-width: 240px;
    background: var(--sidebar);
    border-right: 1px solid var(--border);
    display: flex; flex-direction: column;
    padding: 24px 0;
    box-shadow: 4px 0 30px rgba(0,0,0,0.4);
  }
  .sidebar-logo {
    padding: 0 20px 24px;
    border-bottom: 1px solid var(--border);
    margin-bottom: 16px;
  }
  .sidebar-logo h2 {
    font-size: 16px; font-weight: 800;
    background: linear-gradient(135deg, var(--purple-soft), var(--gold));
    -webkit-background-clip: text; -webkit-text-fill-color: transparent;
    background-clip: text;
  }
  .sidebar-logo p { font-size: 11px; color: var(--text3); margin-top: 2px; }
  .sidebar-logo .logo-icon { font-size: 28px; display: block; margin-bottom: 8px; }
  .nav-item {
    display: flex; align-items: center; gap: 12px;
    padding: 12px 20px; cursor: pointer;
    color: var(--text2); font-size: 14px; font-weight: 600;
    transition: all 0.2s; border-left: 3px solid transparent;
    margin: 2px 0;
  }
  .nav-item:hover { color: var(--text); background: var(--purple-glow2); }
  .nav-item.active {
    color: var(--purple-soft); background: var(--purple-glow2);
    border-left-color: var(--purple2);
  }
  .nav-item .icon { font-size: 18px; width: 24px; text-align: center; }
  .sidebar-bottom { margin-top: auto; padding: 16px 20px; border-top: 1px solid var(--border); }
  .user-badge {
    display: flex; align-items: center; gap: 10px;
    padding: 10px 12px; background: var(--card);
    border-radius: 12px; border: 1px solid var(--border);
    margin-bottom: 12px;
  }
  .user-avatar {
    width: 32px; height: 32px; border-radius: 50%;
    background: linear-gradient(135deg, var(--purple), var(--purple2));
    display: flex; align-items: center; justify-content: center;
    font-size: 14px; font-weight: 700; flex-shrink: 0;
  }
  .user-email { font-size: 11px; color: var(--text2); white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

  /* Main content */
  .main-content {
    flex: 1; overflow-y: auto;
    background: linear-gradient(180deg, var(--deep) 0%, var(--navy) 100%);
    padding: 32px;
  }
  .page-header { margin-bottom: 28px; }
  .page-header h2 {
    font-size: 26px; font-weight: 800;
    background: linear-gradient(135deg, var(--text), var(--purple-soft));
    -webkit-background-clip: text; -webkit-text-fill-color: transparent;
    background-clip: text;
  }
  .page-header p { color: var(--text2); font-size: 14px; margin-top: 4px; }
  .page-header-row { display: flex; align-items: center; justify-content: space-between; }

  /* Cards */
  .card {
    background: var(--card); border: 1px solid var(--border);
    border-radius: 16px; padding: 24px;
    transition: border-color 0.2s;
  }
  .card:hover { border-color: var(--border2); }
  .card-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(280px, 1fr)); gap: 16px; }
  .stats-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(180px, 1fr)); gap: 16px; margin-bottom: 28px; }
  .stat-card {
    background: var(--card); border: 1px solid var(--border);
    border-radius: 16px; padding: 20px;
  }
  .stat-card .stat-icon { font-size: 28px; margin-bottom: 8px; }
  .stat-card .stat-num { font-size: 28px; font-weight: 800; font-family: 'Exo 2', sans-serif; color: var(--purple-soft); }
  .stat-card .stat-label { font-size: 13px; color: var(--text2); margin-top: 2px; }

  /* Item Cards */
  .item-card {
    background: var(--card); border: 1px solid var(--border);
    border-radius: 14px; padding: 18px 20px;
    display: flex; align-items: center; justify-content: space-between;
    transition: all 0.2s;
  }
  .item-card:hover { border-color: var(--border2); transform: translateY(-1px); box-shadow: 0 4px 20px var(--purple-glow2); }
  .item-card-info h4 { font-size: 15px; font-weight: 700; margin-bottom: 4px; }
  .item-card-info p { font-size: 13px; color: var(--text2); }
  .item-card-actions { display: flex; gap: 8px; }

  /* Modal */
  .modal-overlay {
    position: fixed; inset: 0; z-index: 100;
    background: rgba(4,4,26,0.8); backdrop-filter: blur(8px);
    display: flex; align-items: center; justify-content: center; padding: 24px;
    animation: fadeIn 0.2s ease;
  }
  .modal {
    background: var(--card); border: 1px solid var(--border2);
    border-radius: 20px; padding: 32px;
    width: 100%; max-width: 520px;
    box-shadow: 0 0 60px var(--purple-glow);
    animation: fadeUp 0.3s ease;
    max-height: 90vh; overflow-y: auto;
  }
  .modal h3 { font-size: 20px; font-weight: 800; margin-bottom: 24px; }
  .modal-actions { display: flex; gap: 12px; justify-content: flex-end; margin-top: 24px; }

  /* Status Badges */
  .badge {
    display: inline-flex; align-items: center; gap: 4px;
    padding: 4px 10px; border-radius: 20px; font-size: 12px; font-weight: 700;
  }
  .badge-draft { background: rgba(107,99,160,0.2); color: var(--text2); border: 1px solid rgba(107,99,160,0.3); }
  .badge-sent { background: rgba(103,232,249,0.15); color: var(--cyan); border: 1px solid rgba(103,232,249,0.3); }
  .badge-paid { background: rgba(52,211,153,0.15); color: var(--success); border: 1px solid rgba(52,211,153,0.3); }
  .badge-overdue { background: rgba(248,113,113,0.15); color: var(--danger); border: 1px solid rgba(248,113,113,0.3); }

  /* Table */
  .table-wrap { overflow-x: auto; border-radius: 14px; border: 1px solid var(--border); }
  table { width: 100%; border-collapse: collapse; }
  th {
    padding: 14px 18px; text-align: left; font-size: 12px;
    font-weight: 700; text-transform: uppercase; letter-spacing: 0.05em;
    color: var(--text2); background: var(--card2);
    border-bottom: 1px solid var(--border);
  }
  td { padding: 14px 18px; font-size: 14px; border-bottom: 1px solid var(--border); }
  tr:last-child td { border-bottom: none; }
  tr:hover td { background: var(--purple-glow2); }

  /* Empty state */
  .empty-state {
    text-align: center; padding: 60px 24px;
    color: var(--text2);
  }
  .empty-state .empty-icon { font-size: 48px; margin-bottom: 16px; }
  .empty-state h3 { font-size: 18px; font-weight: 700; margin-bottom: 8px; color: var(--text); }
  .empty-state p { font-size: 14px; margin-bottom: 24px; }

  /* Error / Alert */
  .alert { padding: 12px 16px; border-radius: 10px; font-size: 14px; margin-bottom: 16px; }
  .alert-error { background: rgba(248,113,113,0.1); border: 1px solid rgba(248,113,113,0.3); color: var(--danger); }
  .alert-success { background: rgba(52,211,153,0.1); border: 1px solid rgba(52,211,153,0.3); color: var(--success); }

  /* Divider */
  .divider { border: none; border-top: 1px solid var(--border); margin: 20px 0; }

  /* Invoice detail */
  .invoice-detail-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 16px; margin-bottom: 20px; }
  .detail-field label { font-size: 12px; font-weight: 700; color: var(--text2); text-transform: uppercase; letter-spacing: 0.05em; margin-bottom: 4px; display: block; }
  .detail-field p { font-size: 14px; color: var(--text); }

  /* Auth toggle */
  .auth-toggle { text-align: center; margin-top: 20px; font-size: 14px; color: var(--text2); }
  .auth-toggle button { background: none; border: none; color: var(--purple-soft); font-weight: 700; cursor: pointer; font-family: 'Quicksand', sans-serif; font-size: 14px; }
  .auth-toggle button:hover { color: var(--text); }

  /* Loading spinner */
  .spinner {
    width: 20px; height: 20px; border: 2px solid var(--border2);
    border-top-color: var(--purple2); border-radius: 50%;
    animation: spin 0.8s linear infinite; display: inline-block;
  }

  /* Animations */
  @keyframes fadeUp { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: translateY(0); } }
  @keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
  @keyframes spin { to { transform: rotate(360deg); } }

  .form-row { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }
  @media (max-width: 600px) { .form-row { grid-template-columns: 1fr; } .invoice-detail-grid { grid-template-columns: 1fr; } }
`;

// ─── Helpers ──────────────────────────────────────────────────────────────────
const statusBadge = (status: string) => {
  const map: Record<string, string> = { draft: 'badge-draft', sent: 'badge-sent', paid: 'badge-paid', overdue: 'badge-overdue' };
  const icons: Record<string, string> = { draft: '✏️', sent: '📤', paid: '✅', overdue: '⚠️' };
  return <span className={`badge ${map[status] || 'badge-draft'}`}>{icons[status] || '📄'} {status}</span>;
};

const fmt = (n: number) => `£${(n || 0).toFixed(2)}`;

// ─── Modals ───────────────────────────────────────────────────────────────────
function Modal({ title, onClose, children }: { title: string; onClose: () => void; children: React.ReactNode }) {
  return (
    <div className="modal-overlay" onClick={e => { if (e.target === e.currentTarget) onClose(); }}>
      <div className="modal">
        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 24 }}>
          <h3 style={{ margin: 0 }}>{title}</h3>
          <button onClick={onClose} className="btn btn-ghost btn-sm">✕</button>
        </div>
        {children}
      </div>
    </div>
  );
}

// ─── Business Section ─────────────────────────────────────────────────────────
function BusinessSection({ business, onUpdate }: { business: Business | null; onUpdate: (b: Business) => void }) {
  const [editing, setEditing] = useState(!business);
  const [form, setForm] = useState({
    name: business?.name || '', address: business?.address || '',
    phone: business?.phone || '', email: business?.email || '',
    vat_number: business?.vat_number || '', tax_rate: business?.tax_rate?.toString() || '0',
  });
  const [loading, setLoading] = useState(false);
  const [msg, setMsg] = useState('');

  const set = (k: string) => (e: React.ChangeEvent<HTMLInputElement>) => setForm(f => ({ ...f, [k]: e.target.value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); setLoading(true); setMsg('');
    try {
      const payload = { ...form, tax_rate: parseFloat(form.tax_rate) || 0 };
      const res = business ? await api.put('/business', payload) : await api.post('/business', payload);
      onUpdate(res.data);
      setMsg('✨ Business saved successfully!');
      setEditing(false);
    } catch { setMsg('❌ Failed to save. Please try again.'); }
    setLoading(false);
  };

  return (
    <div>
      <div className="page-header">
        <div className="page-header-row">
          <div>
            <h2>🏢 My Business</h2>
            <p>Your business profile appears on all invoices</p>
          </div>
          {!editing && <button className="btn btn-primary" onClick={() => setEditing(true)}>✏️ Edit Business</button>}
        </div>
      </div>
      {msg && <div className={`alert ${msg.startsWith('✨') ? 'alert-success' : 'alert-error'}`}>{msg}</div>}
      {!editing && business ? (
        <div className="card" style={{ maxWidth: 600 }}>
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 20 }}>
            {[
              { label: 'Business Name', value: business.name },
              { label: 'Email', value: business.email },
              { label: 'Phone', value: business.phone },
              { label: 'Address', value: business.address },
              { label: 'VAT Number', value: business.vat_number },
              { label: 'Tax Rate', value: `${business.tax_rate}%` },
            ].map(({ label, value }) => (
              <div className="detail-field" key={label}>
                <label>{label}</label>
                <p>{value || <span style={{ color: 'var(--text3)' }}>Not set</span>}</p>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <div className="card" style={{ maxWidth: 600 }}>
          <form onSubmit={handleSubmit}>
            <div className="form-row">
              <div className="form-group"><label>Business Name ✨</label><input className="form-input" value={form.name} onChange={set('name')} placeholder="Async Cakes" required /></div>
              <div className="form-group"><label>Email</label><input className="form-input" type="email" value={form.email} onChange={set('email')} placeholder="hello@asynccakes.com" /></div>
            </div>
            <div className="form-group"><label>Address</label><input className="form-input" value={form.address} onChange={set('address')} placeholder="123 Space Lane, London" /></div>
            <div className="form-row">
              <div className="form-group"><label>Phone</label><input className="form-input" value={form.phone} onChange={set('phone')} placeholder="+44 123 456 789" /></div>
              <div className="form-group"><label>VAT Number</label><input className="form-input" value={form.vat_number} onChange={set('vat_number')} placeholder="GB123456789" /></div>
            </div>
            <div className="form-group"><label>Tax Rate (%)</label><input className="form-input" type="number" value={form.tax_rate} onChange={set('tax_rate')} placeholder="20" /></div>
            <div className="modal-actions">
              {business && <button type="button" className="btn btn-ghost" onClick={() => setEditing(false)}>Cancel</button>}
              <button type="submit" className="btn btn-primary" disabled={loading}>{loading ? <span className="spinner" /> : business ? '💾 Save Changes' : '🚀 Create Business'}</button>
            </div>
          </form>
        </div>
      )}
    </div>
  );
}

// ─── Clients Section ──────────────────────────────────────────────────────────
function ClientsSection() {
  const [clients, setClients] = useState<Client[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [editing, setEditing] = useState<Client | null>(null);
  const [form, setForm] = useState({ name: '', email: '', phone: '', address: '' });
  const [saving, setSaving] = useState(false);
  const [msg, setMsg] = useState('');

  useEffect(() => { loadClients(); }, []);

  const loadClients = async () => {
    setLoading(true);
    try { const res = await api.get('/clients/user'); setClients(res.data || []); }
    catch { setClients([]); }
    setLoading(false);
  };

  const openAdd = () => { setEditing(null); setForm({ name: '', email: '', phone: '', address: '' }); setShowModal(true); };
  const openEdit = (c: Client) => { setEditing(c); setForm({ name: c.name, email: c.email, phone: c.phone, address: c.address }); setShowModal(true); };
  const set = (k: string) => (e: React.ChangeEvent<HTMLInputElement>) => setForm(f => ({ ...f, [k]: e.target.value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); setSaving(true);
    try {
      if (editing) { await api.put(`/clients/${editing.ID}`, form); }
      else { await api.post('/clients', form); }
      await loadClients();
      setShowModal(false);
    } catch { setMsg('Failed to save client.'); }
    setSaving(false);
  };

  return (
    <div>
      <div className="page-header">
        <div className="page-header-row">
          <div><h2>👥 Clients</h2><p>{clients.length} client{clients.length !== 1 ? 's' : ''} in your space</p></div>
          <button className="btn btn-primary" onClick={openAdd}>✨ Add Client</button>
        </div>
      </div>
      {loading ? <div style={{ textAlign: 'center', padding: 40 }}><span className="spinner" /></div> : clients.length === 0 ? (
        <div className="empty-state"><div className="empty-icon">🌌</div><h3>No clients yet</h3><p>Add your first client to get started</p><button className="btn btn-primary" onClick={openAdd}>✨ Add Client</button></div>
      ) : (
        <div className="card-grid">
          {clients.map(c => (
            <div className="item-card" key={c.ID}>
              <div className="item-card-info">
                <h4>👤 {c.name}</h4>
                <p>📧 {c.email}</p>
                {c.phone && <p>📞 {c.phone}</p>}
              </div>
              <div className="item-card-actions">
                <button className="btn btn-ghost btn-sm" onClick={() => openEdit(c)}>✏️</button>
              </div>
            </div>
          ))}
        </div>
      )}
      {showModal && (
        <Modal title={editing ? '✏️ Edit Client' : '✨ Add Client'} onClose={() => setShowModal(false)}>
          {msg && <div className="alert alert-error">{msg}</div>}
          <form onSubmit={handleSubmit}>
            <div className="form-group"><label>Name</label><input className="form-input" value={form.name} onChange={set('name')} placeholder="Jane Smith" required /></div>
            <div className="form-group"><label>Email</label><input className="form-input" type="email" value={form.email} onChange={set('email')} placeholder="jane@example.com" /></div>
            <div className="form-row">
              <div className="form-group"><label>Phone</label><input className="form-input" value={form.phone} onChange={set('phone')} placeholder="+44 123 456 789" /></div>
              <div className="form-group"><label>Address</label><input className="form-input" value={form.address} onChange={set('address')} placeholder="London, UK" /></div>
            </div>
            <div className="modal-actions">
              <button type="button" className="btn btn-ghost" onClick={() => setShowModal(false)}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? <span className="spinner" /> : editing ? '💾 Save' : '✨ Add Client'}</button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  );
}

// ─── Products Section ─────────────────────────────────────────────────────────
function ProductsSection() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [editing, setEditing] = useState<Product | null>(null);
  const [form, setForm] = useState({ name: '', description: '', price: '', unit: '' });
  const [saving, setSaving] = useState(false);

  useEffect(() => { loadProducts(); }, []);

  const loadProducts = async () => {
    setLoading(true);
    try { const res = await api.get('/products/user'); setProducts(res.data || []); }
    catch { setProducts([]); }
    setLoading(false);
  };

  const openAdd = () => { setEditing(null); setForm({ name: '', description: '', price: '', unit: '' }); setShowModal(true); };
  const openEdit = (p: Product) => { setEditing(p); setForm({ name: p.name, description: p.description, price: p.price?.toString(), unit: p.unit }); setShowModal(true); };
  const set = (k: string) => (e: React.ChangeEvent<HTMLInputElement>) => setForm(f => ({ ...f, [k]: e.target.value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); setSaving(true);
    try {
      const payload = { ...form, price: parseFloat(form.price) || 0 };
      if (editing) { await api.put(`/products/${editing.ID}`, payload); }
      else { await api.post('/products', payload); }
      await loadProducts(); setShowModal(false);
    } catch { alert('Failed to save product.'); }
    setSaving(false);
  };

  return (
    <div>
      <div className="page-header">
        <div className="page-header-row">
          <div><h2>🎂 Products & Services</h2><p>{products.length} item{products.length !== 1 ? 's' : ''} in your catalogue</p></div>
          <button className="btn btn-primary" onClick={openAdd}>✨ Add Product</button>
        </div>
      </div>
      {loading ? <div style={{ textAlign: 'center', padding: 40 }}><span className="spinner" /></div> : products.length === 0 ? (
        <div className="empty-state"><div className="empty-icon">🛸</div><h3>No products yet</h3><p>Add your cakes, services and products here</p><button className="btn btn-primary" onClick={openAdd}>✨ Add Product</button></div>
      ) : (
        <div className="card-grid">
          {products.map(p => (
            <div className="item-card" key={p.ID}>
              <div className="item-card-info">
                <h4>🎂 {p.name}</h4>
                <p>{p.description}</p>
                <p style={{ color: 'var(--purple-soft)', fontWeight: 700, marginTop: 4 }}>{fmt(p.price)}{p.unit ? ` / ${p.unit}` : ''}</p>
              </div>
              <div className="item-card-actions">
                <button className="btn btn-ghost btn-sm" onClick={() => openEdit(p)}>✏️</button>
              </div>
            </div>
          ))}
        </div>
      )}
      {showModal && (
        <Modal title={editing ? '✏️ Edit Product' : '✨ Add Product'} onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="form-group"><label>Name</label><input className="form-input" value={form.name} onChange={set('name')} placeholder="Galaxy Chocolate Cake" required /></div>
            <div className="form-group"><label>Description</label><input className="form-input" value={form.description} onChange={set('description')} placeholder="Rich dark chocolate sponge..." /></div>
            <div className="form-row">
              <div className="form-group"><label>Price (£)</label><input className="form-input" type="number" step="0.01" value={form.price} onChange={set('price')} placeholder="45.00" required /></div>
              <div className="form-group"><label>Unit</label><input className="form-input" value={form.unit} onChange={set('unit')} placeholder="slice / cake / hour" /></div>
            </div>
            <div className="modal-actions">
              <button type="button" className="btn btn-ghost" onClick={() => setShowModal(false)}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? <span className="spinner" /> : editing ? '💾 Save' : '✨ Add'}</button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  );
}

// ─── Invoices Section ─────────────────────────────────────────────────────────
function InvoicesSection() {
  const [invoices, setInvoices] = useState<Invoice[]>([]);
  const [clients, setClients] = useState<Client[]>([]);
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [showCreate, setShowCreate] = useState(false);
  const [selected, setSelected] = useState<Invoice | null>(null);
  const [items, setItems] = useState<{ product_id: number; quantity: number }[]>([{ product_id: 0, quantity: 1 }]);
  const [form, setForm] = useState({ client_id: '', due_date: '', notes: '', status: 'draft' });
  const [saving, setSaving] = useState(false);
  const [statusMsg, setStatusMsg] = useState('');

  useEffect(() => { loadAll(); }, []);

  const loadAll = async () => {
    setLoading(true);
    try {
      const [invRes, cliRes, proRes] = await Promise.all([
        api.get('/invoices/user').catch(() => ({ data: [] })),
        api.get('/clients/user').catch(() => ({ data: [] })),
        api.get('/products/user').catch(() => ({ data: [] })),
      ]);
      setInvoices(invRes.data || []);
      setClients(cliRes.data || []);
      setProducts(proRes.data || []);
    } catch { }
    setLoading(false);
  };

  const set = (k: string) => (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => setForm(f => ({ ...f, [k]: e.target.value }));

  const addItem = () => setItems(i => [...i, { product_id: 0, quantity: 1 }]);
  const removeItem = (idx: number) => setItems(i => i.filter((_, j) => j !== idx));
  const updateItem = (idx: number, k: string, v: string) => setItems(i => i.map((it, j) => j === idx ? { ...it, [k]: k === 'quantity' ? parseInt(v) || 1 : parseInt(v) || 0 } : it));

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault(); setSaving(true);
    try {
      await api.post('/invoices', {
        client_id: parseInt(form.client_id),
        due_date: form.due_date,
        notes: form.notes,
        status: form.status,
        items: items.filter(i => i.product_id > 0),
      });
      await loadAll(); setShowCreate(false);
      setItems([{ product_id: 0, quantity: 1 }]);
      setForm({ client_id: '', due_date: '', notes: '', status: 'draft' });
    } catch { alert('Failed to create invoice.'); }
    setSaving(false);
  };

  const updateStatus = async (id: number, status: string) => {
    try {
      await api.put(`/invoices/${id}/status`, { status });
      setStatusMsg(`✅ Status updated to ${status}!`);
      await loadAll();
      if (selected?.ID === id) setSelected(s => s ? { ...s, status } : null);
      setTimeout(() => setStatusMsg(''), 3000);
    } catch { alert('Failed to update status.'); }
  };

  const downloadPDF = async (id: number) => {
    try {
      const res = await api.get(`/invoices/${id}/pdf`, { responseType: 'blob' });
      const url = window.URL.createObjectURL(new Blob([res.data]));
      const a = document.createElement('a'); a.href = url;
      a.download = `invoice-${id}.pdf`; a.click();
      window.URL.revokeObjectURL(url);
    } catch { alert('Failed to download PDF.'); }
  };

  const getClientName = (id: number) => clients.find(c => c.ID === id)?.name || `Client #${id}`;

  return (
    <div>
      <div className="page-header">
        <div className="page-header-row">
          <div><h2>🧾 Invoices</h2><p>{invoices.length} invoice{invoices.length !== 1 ? 's' : ''} total</p></div>
          <button className="btn btn-primary" onClick={() => setShowCreate(true)}>✨ New Invoice</button>
        </div>
      </div>
      {statusMsg && <div className="alert alert-success">{statusMsg}</div>}
      {loading ? <div style={{ textAlign: 'center', padding: 40 }}><span className="spinner" /></div> : invoices.length === 0 ? (
        <div className="empty-state"><div className="empty-icon">🪐</div><h3>No invoices yet</h3><p>Create your first invoice to start billing clients</p><button className="btn btn-primary" onClick={() => setShowCreate(true)}>✨ New Invoice</button></div>
      ) : (
        <div className="table-wrap">
          <table>
            <thead><tr><th>#</th><th>Client</th><th>Status</th><th>Date</th><th>Total</th><th>Actions</th></tr></thead>
            <tbody>
              {invoices.map(inv => (
                <tr key={inv.ID}>
                  <td style={{ fontWeight: 700, color: 'var(--purple-soft)' }}>#{inv.ID}</td>
                  <td>{getClientName(inv.client_id)}</td>
                  <td>{statusBadge(inv.status)}</td>
                  <td style={{ color: 'var(--text2)' }}>{inv.CreatedAt ? new Date(inv.CreatedAt).toLocaleDateString() : '—'}</td>
                  <td style={{ fontWeight: 700 }}>{fmt(inv.total)}</td>
                  <td>
                    <div style={{ display: 'flex', gap: 8 }}>
                      <button className="btn btn-ghost btn-sm" onClick={() => setSelected(inv)}>👁️ View</button>
                      <button className="btn btn-ghost btn-sm" onClick={() => downloadPDF(inv.ID)}>📄 PDF</button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {/* Create Invoice Modal */}
      {showCreate && (
        <Modal title="✨ Create Invoice" onClose={() => setShowCreate(false)}>
          <form onSubmit={handleCreate}>
            <div className="form-row">
              <div className="form-group">
                <label>Client</label>
                <select className="form-input" value={form.client_id} onChange={set('client_id')} required>
                  <option value="">Select client...</option>
                  {clients.map(c => <option key={c.ID} value={c.ID}>{c.name}</option>)}
                </select>
              </div>
              <div className="form-group">
                <label>Status</label>
                <select className="form-input" value={form.status} onChange={set('status')}>
                  <option value="draft">Draft</option>
                  <option value="sent">Sent</option>
                  <option value="paid">Paid</option>
                </select>
              </div>
            </div>
            <div className="form-group">
              <label>Due Date</label>
              <input className="form-input" type="date" value={form.due_date} onChange={set('due_date')} />
            </div>
            <hr className="divider" />
            <div style={{ marginBottom: 12 }}>
              <label style={{ fontSize: 13, fontWeight: 700, color: 'var(--purple-soft)' }}>🛒 Line Items</label>
            </div>
            {items.map((item, idx) => (
              <div key={idx} style={{ display: 'flex', gap: 8, marginBottom: 10, alignItems: 'center' }}>
                <select className="form-input" style={{ flex: 2 }} value={item.product_id} onChange={e => updateItem(idx, 'product_id', e.target.value)}>
                  <option value="0">Select product...</option>
                  {products.map(p => <option key={p.ID} value={p.ID}>{p.name} — {fmt(p.price)}</option>)}
                </select>
                <input className="form-input" style={{ flex: 1, width: 70 }} type="number" min="1" value={item.quantity} onChange={e => updateItem(idx, 'quantity', e.target.value)} placeholder="Qty" />
                {items.length > 1 && <button type="button" className="btn btn-danger btn-sm" onClick={() => removeItem(idx)}>✕</button>}
              </div>
            ))}
            <button type="button" className="btn btn-ghost btn-sm" onClick={addItem} style={{ marginBottom: 16 }}>+ Add Item</button>
            <div className="form-group">
              <label>Notes</label>
              <input className="form-input" value={form.notes} onChange={set('notes')} placeholder="Thank you for your order! 🌟" />
            </div>
            <div className="modal-actions">
              <button type="button" className="btn btn-ghost" onClick={() => setShowCreate(false)}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? <span className="spinner" /> : '🚀 Create Invoice'}</button>
            </div>
          </form>
        </Modal>
      )}

      {/* View Invoice Modal */}
      {selected && (
        <Modal title={`🧾 Invoice #${selected.ID}`} onClose={() => setSelected(null)}>
          <div className="invoice-detail-grid">
            <div className="detail-field"><label>Client</label><p>{getClientName(selected.client_id)}</p></div>
            <div className="detail-field"><label>Status</label><p>{statusBadge(selected.status)}</p></div>
            <div className="detail-field"><label>Date</label><p>{selected.CreatedAt ? new Date(selected.CreatedAt).toLocaleDateString() : '—'}</p></div>
            <div className="detail-field"><label>Total</label><p style={{ fontWeight: 700, color: 'var(--purple-soft)' }}>{fmt(selected.total)}</p></div>
          </div>
          {selected.notes && <div className="detail-field" style={{ marginBottom: 16 }}><label>Notes</label><p>{selected.notes}</p></div>}
          <hr className="divider" />
          <div style={{ marginBottom: 16 }}>
            <label style={{ fontSize: 12, fontWeight: 700, color: 'var(--text2)', textTransform: 'uppercase', letterSpacing: '0.05em', display: 'block', marginBottom: 8 }}>Update Status</label>
            <div style={{ display: 'flex', gap: 8, flexWrap: 'wrap' }}>
              {['draft', 'sent', 'paid', 'overdue'].map(s => (
                <button key={s} className={`btn btn-sm ${selected.status === s ? 'btn-primary' : 'btn-ghost'}`} onClick={() => updateStatus(selected.ID, s)}>{s}</button>
              ))}
            </div>
          </div>
          <div style={{ display: 'flex', gap: 8, justifyContent: 'flex-end' }}>
            <button className="btn btn-success" onClick={() => downloadPDF(selected.ID)}>📄 Download PDF</button>
            <button className="btn btn-ghost" onClick={() => setSelected(null)}>Close</button>
          </div>
        </Modal>
      )}
    </div>
  );
}

// ─── Dashboard Overview ───────────────────────────────────────────────────────
function Dashboard({ business, onNavigate }: { business: Business | null; onNavigate: (s: Section) => void }) {
  const [stats, setStats] = useState({ clients: 0, products: 0, invoices: 0, revenue: 0 });

  useEffect(() => {
    Promise.all([
      api.get('/clients/user').catch(() => ({ data: [] })),
      api.get('/products/user').catch(() => ({ data: [] })),
      api.get('/invoices/user').catch(() => ({ data: [] })),
    ]).then(([c, p, i]) => {
      const invs = i.data || [];
      const revenue = invs.filter((inv: Invoice) => inv.status === 'paid').reduce((sum: number, inv: Invoice) => sum + (inv.total || 0), 0);
      setStats({ clients: (c.data || []).length, products: (p.data || []).length, invoices: invs.length, revenue });
    });
  }, []);

  return (
    <div>
      <div className="page-header">
        <h2>🌌 Dashboard</h2>
        <p>Welcome back to {business?.name || 'your Invoice System'} ✨</p>
      </div>
      <div className="stats-grid">
        <div className="stat-card" onClick={() => onNavigate('clients')} style={{ cursor: 'pointer' }}>
          <div className="stat-icon">👥</div>
          <div className="stat-num">{stats.clients}</div>
          <div className="stat-label">Clients</div>
        </div>
        <div className="stat-card" onClick={() => onNavigate('products')} style={{ cursor: 'pointer' }}>
          <div className="stat-icon">🎂</div>
          <div className="stat-num">{stats.products}</div>
          <div className="stat-label">Products</div>
        </div>
        <div className="stat-card" onClick={() => onNavigate('invoices')} style={{ cursor: 'pointer' }}>
          <div className="stat-icon">🧾</div>
          <div className="stat-num">{stats.invoices}</div>
          <div className="stat-label">Invoices</div>
        </div>
        <div className="stat-card">
          <div className="stat-icon">💰</div>
          <div className="stat-num">{fmt(stats.revenue)}</div>
          <div className="stat-label">Revenue (Paid)</div>
        </div>
      </div>
      <div className="card" style={{ maxWidth: 500 }}>
        <h3 style={{ marginBottom: 16, fontSize: 16 }}>🚀 Quick Actions</h3>
        <div style={{ display: 'flex', gap: 12, flexWrap: 'wrap' }}>
          <button className="btn btn-primary" onClick={() => onNavigate('invoices')}>✨ New Invoice</button>
          <button className="btn btn-ghost" onClick={() => onNavigate('clients')}>👤 Add Client</button>
          <button className="btn btn-ghost" onClick={() => onNavigate('products')}>🎂 Add Product</button>
        </div>
      </div>
    </div>
  );
}

// ─── Main App ─────────────────────────────────────────────────────────────────
export default function App() {
  const [view, setView] = useState<View>('login');
  const [section, setSection] = useState<Section>('dashboard');
  const [business, setBusiness] = useState<Business | null>(null);
  const [userEmail, setUserEmail] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  // Inject styles
  useEffect(() => {
    const style = document.createElement('style');
    style.textContent = STYLES;
    document.head.appendChild(style);
    return () => { document.head.removeChild(style); };
  }, []);

  // Check existing token
  useEffect(() => {
    const token = localStorage.getItem('token');
    const saved = localStorage.getItem('userEmail');
    if (token) {
      api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      if (saved) setUserEmail(saved);
      loadBusiness();
    }
  }, []);

  const loadBusiness = async () => {
    try {
      const res = await api.get('/business');
      setBusiness(res.data);
      setView('app');
    } catch (err: any) {
      if (err.response?.status === 404) { setBusiness(null); setView('app'); }
    }
  };

  const handleAuth = async (e: React.FormEvent, mode: 'login' | 'register') => {
    e.preventDefault(); setLoading(true); setError('');
    try {
      if (mode === 'register') await api.post('/register', { email, password });
      const res = await api.post('/login', { email, password });
      const token = res.data;
      localStorage.setItem('token', token);
      localStorage.setItem('userEmail', email);
      api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      setUserEmail(email);
      await loadBusiness();
    } catch { setError(mode === 'login' ? 'Login failed! Check your credentials.' : 'Registration failed! Email may already exist.'); }
    setLoading(false);
  };

  const handleLogout = () => {
    localStorage.removeItem('token'); localStorage.removeItem('userEmail');
    delete api.defaults.headers.common['Authorization'];
    setBusiness(null); setUserEmail(''); setView('login');
  };

  // ── Auth screens ──────────────────────────────────────────────────────────
  if (view === 'login' || view === 'register') {
    const isLogin = view === 'login';
    return (
      <div className="auth-wrap">
        <div className="auth-card">
          <div className="auth-logo">
            <span className="star">🚀</span>
            <h1>Async Cakes</h1>
            <p>Invoice System ✨</p>
          </div>
          {error && <div className="alert alert-error">{error}</div>}
          <form onSubmit={e => handleAuth(e, isLogin ? 'login' : 'register')}>
            <div className="form-group">
              <label>Email</label>
              <input className="form-input" type="email" placeholder="hello@asynccakes.com" value={email} onChange={e => setEmail(e.target.value)} required autoFocus />
            </div>
            <div className="form-group">
              <label>Password</label>
              <input className="form-input" type="password" placeholder="••••••••••" value={password} onChange={e => setPassword(e.target.value)} required />
            </div>
            <button type="submit" className="btn btn-primary btn-full" style={{ marginTop: 8 }} disabled={loading}>
              {loading ? <span className="spinner" /> : isLogin ? '🚀 Launch In' : '✨ Create Account'}
            </button>
          </form>
          <div className="auth-toggle">
            {isLogin ? <>No account? <button onClick={() => { setView('register'); setError(''); }}>Create one ✨</button></> : <>Already have one? <button onClick={() => { setView('login'); setError(''); }}>Log in 🚀</button></>}
          </div>
        </div>
      </div>
    );
  }

  // ── Main App layout ───────────────────────────────────────────────────────
  const navItems: { id: Section; icon: string; label: string }[] = [
    { id: 'dashboard', icon: '🌌', label: 'Dashboard' },
    { id: 'business', icon: '🏢', label: 'My Business' },
    { id: 'clients', icon: '👥', label: 'Clients' },
    { id: 'products', icon: '🎂', label: 'Products' },
    { id: 'invoices', icon: '🧾', label: 'Invoices' },
  ];

  return (
    <div className="app-layout">
      <aside className="sidebar">
        <div className="sidebar-logo">
          <span className="logo-icon">🚀</span>
          <h2>Async Cakes</h2>
          <p>Invoice System</p>
        </div>
        {navItems.map(n => (
          <div key={n.id} className={`nav-item ${section === n.id ? 'active' : ''}`} onClick={() => setSection(n.id)}>
            <span className="icon">{n.icon}</span> {n.label}
          </div>
        ))}
        <div className="sidebar-bottom">
          <div className="user-badge">
            <div className="user-avatar">{userEmail?.[0]?.toUpperCase() || '🌟'}</div>
            <div className="user-email">{userEmail}</div>
          </div>
          <button className="btn btn-ghost" style={{ width: '100%', justifyContent: 'center' }} onClick={handleLogout}>👋 Log Out</button>
        </div>
      </aside>
      <main className="main-content">
        {section === 'dashboard' && <Dashboard business={business} onNavigate={setSection} />}
        {section === 'business' && <BusinessSection business={business} onUpdate={(b) => { setBusiness(b); setSection('dashboard'); }} />}
        {section === 'clients' && <ClientsSection />}
        {section === 'products' && <ProductsSection />}
        {section === 'invoices' && <InvoicesSection />}
      </main>
    </div>
  );
}
