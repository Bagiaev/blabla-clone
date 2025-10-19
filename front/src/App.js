import React, { useState, useEffect } from 'react';
import Register from './components/Register';
import Login from './components/Login';
import ForgotPassword from './components/ForgotPassword';
import Profile from './components/Profile'; // ← ДОБАВЬТЕ ЭТОТ ИМПОРТ
import './App.css';

function App() {
  const [currentView, setCurrentView] = useState('login');
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  
  // Проверяем авторизацию при загрузке
  useEffect(() => {
    console.log("Backend URL from Android:", window.API_BASE_URL);

    const token = localStorage.getItem('token');
    if (token) {
      setIsAuthenticated(true);
      setCurrentView('profile');
    }
  }, []);

  const handleLoginSuccess = () => {
    setIsAuthenticated(true);
    setCurrentView('profile');
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
    setCurrentView('login');
  };

  const renderView = () => {
    switch (currentView) {
      case 'register':
        return <Register />;
      case 'forgot':
        return <ForgotPassword />;
      case 'profile':
        return <Profile onLogout={handleLogout} />;
      default:
        return <Login 
          onSwitchToRegister={() => setCurrentView('register')} 
          onLoginSuccess={handleLoginSuccess}
        />;
    }
  };

  return (
    <div className="App">
      <header style={styles.header}>
        <h1>🚗 BlaBla Clone Demo</h1>
        <nav style={styles.nav}>
          {!isAuthenticated ? (
            <>
              <button 
                style={currentView === 'login' ? styles.activeNavButton : styles.navButton}
                onClick={() => setCurrentView('login')}
              >
                Вход
              </button>
              <button 
                style={currentView === 'register' ? styles.activeNavButton : styles.navButton}
                onClick={() => setCurrentView('register')}
              >
                Регистрация
              </button>
              <button 
                style={currentView === 'forgot' ? styles.activeNavButton : styles.navButton}
                onClick={() => setCurrentView('forgot')}
              >
                Сброс пароля
              </button>
            </>
          ) : (
            <>
              <button 
                style={currentView === 'profile' ? styles.activeNavButton : styles.navButton}
                onClick={() => setCurrentView('profile')}
              >
                Профиль
              </button>
              <button 
                style={styles.logoutButton}
                onClick={handleLogout}
              >
                Выйти
              </button>
            </>
          )}
        </nav>
      </header>
      
      <main>
        {renderView()}
      </main>
    </div>
  );
}

// Обновите стили, добавив logoutButton
const styles = {
  // ... предыдущие стили ...
  logoutButton: {
    padding: '10px 20px',
    border: '1px solid #dc3545',
    backgroundColor: 'white',
    color: '#dc3545',
    cursor: 'pointer',
    borderRadius: '5px',
  },
};

export default App;