-- -----------------------------------------------------
-- Schema test
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `test` DEFAULT CHARACTER SET utf8 ;
USE `test` ;

-- -----------------------------------------------------
-- Table `test`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `test`.`user` (
  `user_id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `email` VARCHAR(128) UNIQUE NOT NULL COMMENT 'メールアドレス',
  `password` VARCHAR(128) NOT NULL COMMENT 'パスワードダイジェスト',
  `admin` BOOLEAN NOT NULL COMMENT '管理者',
  PRIMARY KEY (`user_id`),
  INDEX `idx_user_id` (`user_id` ASC))
ENGINE = InnoDB
COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- Table `test`.`post`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `test`.`post` (
  `post_id` INT AUTO_INCREMENT NOT NULL COMMENT '投稿ID',
  `content` VARCHAR(64) NOT NULL COMMENT '内容',
  `user_id` VARCHAR(128) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`post_id`))
ENGINE = InnoDB
COMMENT = '投稿';