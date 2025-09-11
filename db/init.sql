-- =================================================================
--  テーブル: users
--  ユーザー情報を格納します。
-- =================================================================
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY, --追加時に自動で割り振る
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at NOT NULL DEFAULT CURRENT_TIMESTAMP, --"TIMESTAMPTZ"にすると追加時に自動で振られる
    updated_at NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =================================================================
--  テーブル: categories
--  (新規追加) ユーザーが日記を分類するためのカテゴリです。
-- =================================================================
CREATE TABLE IF NOT EXISTS categories (
    categories_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    color TEXT NOT NULL CHECK (color ~ '^#[0-9a-fA-F]{6}$'), -- HEXカラーコード形式(#FFFFFF)に制限
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    UNIQUE (user_id, name) -- ユーザーごとに同じカテゴリ名は登録不可
);

-- =================================================================
--  テーブル: diaries
--  日記の情報を格納します。テキスト情報もこのテーブルに含まれます。
-- =================================================================
CREATE TABLE IF NOT EXISTS diaries (
    diary_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    category_id BIGINT, -- カテゴリは必須ではないためNULLを許可
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6),
    text TEXT,
    corrected_text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- ユーザーが任意の日時を指定することも可能
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
        REFERENCES categories(id)
        ON DELETE SET NULL -- 参照先のカテゴリが削除された場合、このカラムはNULLになります
);

-- =================================================================
--  テーブル: photos
--  日記に紐づく写真のURLを格納します。
-- =================================================================
CREATE TABLE IF NOT EXISTS photos (
    photo_id BIGSERIAL PRIMARY KEY,
    diary_id BIGINT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_diary
        FOREIGN KEY(diary_id)
        REFERENCES diaries(id)
        ON DELETE CASCADE
);

-- =================================================================
--  テーブル: travelogues
--  AIによってユーザーの日記から生成された旅行記を格納します。
-- =================================================================
CREATE TABLE IF NOT EXISTS travelogues (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);