ALTER TABLE blacklist DROP FOREIGN KEY fk_blacklist_stores1;
ALTER TABLE blacklist DROP INDEX fk_blacklist_stores1_idx;
ALTER TABLE blacklist DROP COLUMN store_id;
ALTER TABLE blacklist ADD COLUMN store_url VARCHAR(255) NOT NULL AFTER id;
