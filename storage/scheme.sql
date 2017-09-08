SET PASSWORD FOR 'root'@'localhost' = PASSWORD('123');

-- --------------------------------------------------------
-- Хост:                         127.0.0.1
-- Версия сервера:               10.1.23-MariaDB-1~jessie - mariadb.org binary distribution
-- Операционная система:         debian-linux-gnu
-- HeidiSQL Версия:              9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

drop database if exists t4rest;

-- Дамп структуры базы данных t4rest
CREATE DATABASE IF NOT EXISTS `t4rest` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `t4rest`;


-- Экспортируемые данные не выделены.
-- Дамп структуры для таблица t4rest.visits
CREATE TABLE `visits` (
  `id` INT(10) UNSIGNED NOT NULL,
  `location` INT(10) UNSIGNED NOT NULL,
  `user` INT(10) UNSIGNED NOT NULL,
  `visited_at` INT(11) NOT NULL,
  `mark` TINYINT(2) UNSIGNED NOT NULL,
  `gender` VARCHAR(10) NOT NULL,
  `birth_date` INT(11) NOT NULL,
  `country` VARCHAR(255) NOT NULL,
  `distance` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `user` (`user`),
  INDEX `location` (`location`)
);

-- Экспортируемые данные не выделены.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;


set global max_prepared_stmt_count=32764*2;