import axios from 'axios';

// Если приложение запускается внутри Capacitor (WebView), берем URL из window
// Если React в браузере (npm start), можно использовать fallback
const API_BASE_URL = 'http://localhost:8000/api';

console.log('Using API base URL:', API_BASE_URL);

export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: false,
});

// 🔹 Сервисы для работы с API
export const authAPI = {
  register: (userData) => api.post('/register', userData),
  login: (credentials) => api.post('/login', credentials),
  forgotPassword: (email) => api.post('/forgot-password', { email }),
  getProfile: (token) =>
    api.get('/user/profile', {
      headers: { Authorization: `Bearer ${token}` },
    }),
};
