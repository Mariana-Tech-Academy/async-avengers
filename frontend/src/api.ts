import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.MODE === 'development' 
    ? 'http://localhost:8080' 
    : 'https://async-avengers.onrender.com', 
});

// Add this "interceptor" to attach the token if it exists
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;