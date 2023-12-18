--
-- Create database `url_shortener`
--
CREATE DATABASE IF NOT EXISTS url_shortener;

USE url_shortener;

--
-- Table structure for table `urls`
--

CREATE TABLE IF NOT EXISTS `urls` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(10) DEFAULT NULL,
  `url` varchar(400) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `ip` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `urls_pk` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Table structure for table `visitors`
--

CREATE TABLE IF NOT EXISTS `visitors` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) NOT NULL,
  `visited_at` timestamp NOT NULL,
  `url_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `visitors_urls_id_fk` (`url_id`),
  CONSTRAINT `visitors_urls_id_fk` FOREIGN KEY (`url_id`) REFERENCES `urls` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
