CREATE TABLE IF NOT EXISTS users
(
    id               BIGINT                                 NOT NULL PRIMARY KEY,
    username         TEXT                                   NOT NULL CONSTRAINT username_check CHECK (username <> '' :: TEXT),
    password         TEXT                                   NOT NULL CONSTRAINT password_check CHECK (password <> '' :: TEXT),
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS files
(
    id               BIGINT                                 NOT NULL PRIMARY KEY,
    user_id          BIGINT                                 NOT NULL CONSTRAINT file_user_id_fkey REFERENCES users (id),
    name             TEXT                                   NOT NULL CONSTRAINT file_name_check CHECK (name <> '' :: TEXT),
    type             TEXT                                   NOT NULL CONSTRAINT file_type_check CHECK (type <> '' :: TEXT),
    size             BIGINT,
    data             TEXT                                   NOT NULL CONSTRAINT file_data_check CHECK (data <> '' :: TEXT),
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);