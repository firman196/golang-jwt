create table if not exist `tokens` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER UNSIGNED KEY NOT NULL,
    token TEXT,
    refresh_token TEXT
);

ALTER TABLE tokens ADD FOREIGN KEY (from_users_token) REFERENCES users (id);
ALTER TABLE tokens ADD FOREIGN KEY (to_users_token) REFERENCES users (id);