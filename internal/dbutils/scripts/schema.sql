CREATE TABLE users (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   username VARCHAR(50) NOT NULL,
   realname VARCHAR(100),
   password TEXT NOT NULL,
   create_datetime DATETIME NOT NULL
);

CREATE UNIQUE INDEX idx_username ON users (username);

CREATE TABLE campaigns (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   description TEXT,
   key_field VARCHAR(36) NOT NULL,
   user_id INTEGER NOT NULL,
   create_datetime DATETIME NOT NULL,
   FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_campaign_key ON campaigns (key_field);

CREATE TABLE characters (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   player_name VARCHAR(100) NOT NULL,
   xp_bonus INTEGER NOT NULL,
   campaign_id INTEGER NOT NULL,
   create_datetime DATETIME NOT NULL,
   FOREIGN KEY(campaign_id) REFERENCES campaigns(id)
);

CREATE TABLE xp_bonus_reasons (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   xp_bonus INTEGER NOT NULL,
   reason TEXT NOT NULL,
   character_id INTEGER NOT NULL,
   create_datetime DATETIME NOT NULL,
   FOREIGN KEY(character_id) REFERENCES characters(id)
);

CREATE TABLE xp_awards (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   xp_award INTEGER NOT NULL,
   xp_award_with_bonus INTEGER NOT NULL,
   reason TEXT NOT NULL,
   character_id INTEGER NOT NULL,
   create_datetime DATETIME NOT NULL,
   FOREIGN KEY(character_id) REFERENCES characters(id)
);
