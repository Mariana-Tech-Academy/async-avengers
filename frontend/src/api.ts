import axios from 'axios';

// This is the "bridge" to your Go server
const api = axios.create({
  baseURL: 'http://localhost:8080', // Replace 8080 with the port from your main.go
});

export default api;