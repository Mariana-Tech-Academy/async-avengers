import axios from 'axios';

// This is the "bridge" to your Go server
const api = axios.create({
    // This checks if you're working on your computer or if the app is live
    baseURL: import.meta.env.MODE === 'development'
        ? 'http://localhost:8080'
        : 'https://async-avengers.onrender.com',
});

export default api;