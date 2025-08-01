BEGIN;

CREATE TABLE users (
   id UUID DEFAULT uuid_generate_v7() PRIMARY KEY,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL  DEFAULT NOW(),
   updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
   email VARCHAR(255) NOT NULL UNIQUE,
   name VARCHAR(64) NOT NULL,
   surname VARCHAR(64) NOT NULL,
   patronymic VARCHAR(64)
);

COMMENT ON TABLE users IS 'Таблица пользователей приложения';
COMMENT ON COLUMN users.created_at IS 'Дата и время создания пользователя';
COMMENT ON COLUMN users.updated_at IS 'Дата и время последнего обновления пользователя';

COMMIT;