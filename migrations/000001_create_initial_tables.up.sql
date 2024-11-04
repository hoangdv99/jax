SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema jax
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema jax
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `jax` DEFAULT CHARACTER SET utf8 ;
USE `jax` ;

-- -----------------------------------------------------
-- Table `jax`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uid` VARCHAR(255) NULL,
  `username` VARCHAR(50) NULL DEFAULT NULL,
  `email` VARCHAR(255) NOT NULL,
  `hashed_password` VARCHAR(255) NULL,
  `role` VARCHAR(10) NOT NULL DEFAULT 'user' COMMENT 'user | admin | supervisor',
  `status` INT NOT NULL COMMENT '-10: firestore user\n0: inactive\n10: waiting\n999: active',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`stores`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`stores` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ref_id` VARCHAR(255) NULL COMMENT 'Firestore ref ID',
  `url` VARCHAR(255) NOT NULL,
  `platform` VARCHAR(45) NOT NULL COMMENT '‘woocommerce’ | ‘shopify’ | ‘shopbase’',
  `is_active` TINYINT(1) NOT NULL DEFAULT 1,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `url_UNIQUE` (`url` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`user_stores`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`user_stores` (
  `user_id` INT NOT NULL,
  `store_id` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `store_id`),
  INDEX `fk_users_stores_idx` (`user_id` ASC) VISIBLE,
  INDEX `fk_user_stores_stores_idx` (`store_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_stores_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `jax`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_user_stores_stores`
    FOREIGN KEY (`store_id`)
    REFERENCES `jax`.`stores` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`tags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`tags` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ref_id` VARCHAR(255) NULL COMMENT 'Firestore ref ID',
  `name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`user_tags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`user_tags` (
  `user_id` INT NOT NULL,
  `store_id` INT NOT NULL,
  `tag_id` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `store_id`, `tag_id`),
  INDEX `fk_user_tags_stores_idx` (`store_id` ASC) VISIBLE,
  INDEX `fk_user_tags_tags_idx` (`tag_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_tags_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `jax`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_user_tags_stores`
    FOREIGN KEY (`store_id`)
    REFERENCES `jax`.`stores` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_user_tags_tags`
    FOREIGN KEY (`tag_id`)
    REFERENCES `jax`.`tags` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`tokens`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`tokens` (
  `hash` VARCHAR(255) NOT NULL,
  `user_id` INT NOT NULL,
  `expiry` DATETIME NOT NULL,
  `scope` VARCHAR(45) NOT NULL COMMENT 'authentication | activation',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`hash`),
  INDEX `fk_tokens_users_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_tokens_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `jax`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `jax`.`blacklist`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jax`.`blacklist` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `store_id` INT NOT NULL,
  `added_by` VARCHAR(255) NOT NULL,
  `scope` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `edited_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_blacklist_stores1_idx` (`store_id` ASC) VISIBLE,
  CONSTRAINT `fk_blacklist_stores1`
    FOREIGN KEY (`store_id`)
    REFERENCES `jax`.`stores` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
