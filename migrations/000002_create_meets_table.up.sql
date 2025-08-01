BEGIN;

CREATE TABLE meets (
    id UUID DEFAULT uuid_generate_v7() PRIMARY KEY,
    owner_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    title varchar(255) NOT NULL,
    identifier varchar (8) UNIQUE NOT NULL,
    capacity smallint CHECK (capacity > 0),

    CONSTRAINT fk_meets_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES users(id)
            ON DELETE RESTRICT
            ON UPDATE CASCADE
);

CREATE INDEX idx_meets_owner_id ON meets(owner_id);
CREATE INDEX idx_meets_identifier_id ON meets(identifier);

COMMENT ON TABLE meets IS 'Таблица комнат видеоконференций';
COMMENT ON COLUMN meets.created_at IS 'Дата и время создания комнаты';
COMMENT ON COLUMN meets.title IS 'Название встречи';
COMMENT ON COLUMN meets.identifier IS 'Уникальный публичный идентификатор для входа';
COMMENT ON COLUMN meets.capacity IS 'Вместимость комнаты';


COMMIT;