create table if not exists `tokens` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER,
    token TEXT,
    refresh_token TEXT,
    FOREIGN KEY(user_id) REFERENCES users(id)
);