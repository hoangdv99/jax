ALTER TABLE `tags` DROP FOREIGN KEY `fk_tags_users`;
ALTER TABLE `user_tags` ADD COLUMN `user_id` INT;

ALTER TABLE `user_tags`
ADD CONSTRAINT `fk_user_tags_users`
FOREIGN KEY (`user_id`)
REFERENCES `jax`.`users` (`id`)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE `tags` DROP COLUMN `user_id`;
