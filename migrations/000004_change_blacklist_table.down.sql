-- Rollback: Remove `store_url` column
ALTER TABLE blacklist DROP COLUMN store_url;

-- Rollback: Add back `store_id` column
ALTER TABLE blacklist ADD COLUMN store_id INT;

-- Rollback: Add back index (assuming it was on store_id)
ALTER TABLE blacklist ADD INDEX fk_blacklist_stores1_idx (store_id);

-- Rollback: Add back foreign key constraint
ALTER TABLE blacklist ADD CONSTRAINT fk_blacklist_stores1
  FOREIGN KEY (store_id) REFERENCES stores(id);
  