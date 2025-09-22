-- =================================================================
--  テーブル: users
-- =================================================================
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =================================================================
--  テーブル: categories
-- =================================================================
CREATE TABLE IF NOT EXISTS categories (
    categories_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    color TEXT NOT NULL CHECK (color ~ '^#[0-9a-fA-F]{6}$'),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
    UNIQUE (user_id, name)
);

-- =================================================================
--  テーブル: diaries
-- =================================================================
CREATE TABLE IF NOT EXISTS diaries (
    diary_id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    category_id BIGINT,
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6),
    text TEXT,
    corrected_text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
        REFERENCES categories(categories_id)
        ON DELETE SET NULL
);

-- =================================================================
--  テーブル: photos
-- =================================================================
CREATE TABLE IF NOT EXISTS photos (
    photo_id BIGSERIAL PRIMARY KEY,
    diary_id BIGINT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_diary
        FOREIGN KEY(diary_id)
        REFERENCES diaries(diary_id)
        ON DELETE CASCADE
);

-- =================================================================
--  テーブル: travelogues
-- =================================================================
CREATE TABLE IF NOT EXISTS travelogues (
    travelogues_id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
);
