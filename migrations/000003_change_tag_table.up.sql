ALTER TABLE `tags` ADD COLUMN `user_id` INT NOT NULL AFTER `ref_id`;
ALTER TABLE `user_tags` DROP FOREIGN KEY `fk_user_tags_users`;
ALTER TABLE `user_tags` DROP COLUMN `user_id`;

ALTER TABLE `tags`
ADD CONSTRAINT `fk_tags_users`
FOREIGN KEY (`user_id`)
REFERENCES `jax`.`users` (`id`)
ON DELETE CASCADE
ON UPDATE CASCADE;
