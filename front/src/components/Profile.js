import React, { useState, useEffect } from 'react';

const Profile = () => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = async () => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        setError('Требуется авторизация');
        setLoading(false);
        return;
      }

      const response = await fetch('http://localhost:8000/api/user/profile', {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error('Ошибка загрузки профиля');
      }

      const data = await response.json();
      setUser(data.user);
    } catch (err) {
      setError(err.message);
      console.error('Ошибка:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.reload();
  };

  if (loading) {
    return (
      <div style={styles.container}>
        <div style={styles.loading}>Загрузка профиля...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div style={styles.container}>
        <div style={styles.error}>
          <h3>Ошибка</h3>
          <p>{error}</p>
          <button 
            style={styles.button}
            onClick={() => window.location.href = '/login'}
          >
            Перейти к входу
          </button>
        </div>
      </div>
    );
  }

  if (!user) {
    return (
      <div style={styles.container}>
        <div style={styles.error}>Профиль не найден</div>
      </div>
    );
  }

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h1>👤 Мой профиль</h1>
        <button 
          style={styles.logoutButton}
          onClick={handleLogout}
        >
          Выйти
        </button>
      </div>

      <div style={styles.profileCard}>
        <div style={styles.avatarSection}>
          <div style={styles.avatar}>
            {user.first_name?.[0]}{user.last_name?.[0]}
          </div>
          <div style={styles.rating}>
            <span style={styles.ratingStars}>⭐</span>
            <span>{user.rating || 'Нет оценок'}</span>
          </div>
        </div>

        <div style={styles.infoSection}>
          <div style={styles.field}>
            <label style={styles.label}>Имя:</label>
            <span style={styles.value}>{user.first_name}</span>
          </div>

          <div style={styles.field}>
            <label style={styles.label}>Фамилия:</label>
            <span style={styles.value}>{user.last_name}</span>
          </div>

          <div style={styles.field}>
            <label style={styles.label}>Email:</label>
            <span style={styles.value}>{user.email}</span>
          </div>

          <div style={styles.field}>
            <label style={styles.label}>Телефон:</label>
            <span style={styles.value}>{user.phone}</span>
          </div>

          {user.description && (
            <div style={styles.field}>
              <label style={styles.label}>О себе:</label>
              <span style={styles.value}>{user.description}</span>
            </div>
          )}

          <div style={styles.field}>
            <label style={styles.label}>Дата регистрации:</label>
            <span style={styles.value}>
              {new Date(user.created_at).toLocaleDateString('ru-RU')}
            </span>
          </div>

          {user.updated_at && (
            <div style={styles.field}>
              <label style={styles.label}>Обновлен:</label>
              <span style={styles.value}>
                {new Date(user.updated_at).toLocaleDateString('ru-RU')}
              </span>
            </div>
          )}
        </div>
      </div>

      <div style={styles.actions}>
        <button style={styles.editButton}>
          Редактировать профиль
        </button>
        <button style={styles.secondaryButton}>
          История поездок
        </button>
      </div>
    </div>
  );
};

const styles = {
  container: {
    maxWidth: '600px',
    margin: '0 auto',
    padding: '20px',
    fontFamily: 'Arial, sans-serif',
  },
  header: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: '30px',
    borderBottom: '2px solid #f0f0f0',
    paddingBottom: '15px',
  },
  logoutButton: {
    padding: '8px 16px',
    backgroundColor: '#dc3545',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    fontSize: '14px',
  },
  profileCard: {
    backgroundColor: 'white',
    borderRadius: '10px',
    padding: '25px',
    boxShadow: '0 2px 10px rgba(0,0,0,0.1)',
    marginBottom: '20px',
  },
  avatarSection: {
    display: 'flex',
    alignItems: 'center',
    marginBottom: '25px',
    paddingBottom: '20px',
    borderBottom: '1px solid #f0f0f0',
  },
  avatar: {
    width: '80px',
    height: '80px',
    borderRadius: '50%',
    backgroundColor: '#007bff',
    color: 'white',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    fontSize: '24px',
    fontWeight: 'bold',
    marginRight: '20px',
  },
  rating: {
    display: 'flex',
    alignItems: 'center',
    gap: '8px',
    fontSize: '16px',
  },
  ratingStars: {
    fontSize: '20px',
  },
  infoSection: {
    display: 'flex',
    flexDirection: 'column',
    gap: '15px',
  },
  field: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    padding: '10px 0',
    borderBottom: '1px solid #f8f9fa',
  },
  label: {
    fontWeight: 'bold',
    color: '#555',
    minWidth: '120px',
  },
  value: {
    color: '#333',
    textAlign: 'right',
    flex: 1,
  },
  actions: {
    display: 'flex',
    gap: '15px',
    justifyContent: 'center',
  },
  editButton: {
    padding: '12px 24px',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    fontSize: '16px',
  },
  secondaryButton: {
    padding: '12px 24px',
    backgroundColor: '#6c757d',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    fontSize: '16px',
  },
  loading: {
    textAlign: 'center',
    fontSize: '18px',
    color: '#666',
    padding: '40px',
  },
  error: {
    textAlign: 'center',
    padding: '40px',
    color: '#dc3545',
  },
  button: {
    padding: '10px 20px',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    marginTop: '15px',
  },
};

export default Profile;