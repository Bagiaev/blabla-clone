CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    phone VARCHAR(50) UNIQUE NOT NULL,
    avatar_url VARCHAR(500),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- определение пользователя (водитель, пассажир)
    is_driver BOOLEAN DEFAULT FALSE,

    -- Поля, которые заполняются только, если is_driver true
    car_model VARCHAR(100),
    car_image_url VARCHAR(500),
    rating INTEGER NOT NULL DEFAULT 0,
    description VARCHAR(1000) DEFAULT ''
);

CREATE TABLE IF NOT EXISTS rides (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    driver_id INTEGER,
    car_id INTEGER,

    from_location TEXT NOT NULL,
    to_location TEXT NOT NULL,

    departure_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,

    seats_available INTEGER NOT NULL,
    price_cents INTEGER NOT NULL,

    description TEXT
);

CREATE TABLE IF NOT EXISTS reports (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- При желании: внешние ключи (раскомментируйте, если нужно)
-- ALTER TABLE rides
--   ADD CONSTRAINT rides_driver_fk FOREIGN KEY (driver_id) REFERENCES users(id) ON DELETE SET NULL;

-- ALTER TABLE rides
--   ADD CONSTRAINT rides_car_fk FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE SET NULL;

-- Простейшие тестовые данные (опционально)
-- Пример заполнения тестовыми данными
INSERT INTO users (name, email, hashed_password, rating) VALUES
    ('Водитель 1', 'driver1@example.com', '$2a$10$examplehashdriver1', 5),
    ('Водитель 2', 'driver2@example.com', '$2a$10$examplehashdriver2', 4)
ON CONFLICT (email) DO NOTHING;

INSERT INTO rides (
    driver_id, car_id, from_location, to_location, departure_at, seats_available, price_cents, description
) VALUES
    (1, 1, 'Москва', 'Санкт-Петербург', NOW() + INTERVAL '1 day', 3, 1500, 'Тестовый рейс 1'),
    (2, 2, 'Москва', 'Казань', NOW() + INTERVAL '2 days', 2, 2500, 'Тестовый рейс 2')
ON CONFLICT DO NOTHING;

INSERT INTO reports (user_id, title, description) VALUES
    (1, 'Отзыв 1', 'Хороший водитель')
ON CONFLICT (email) DO NOTHING;

