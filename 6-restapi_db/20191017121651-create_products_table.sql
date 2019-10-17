
-- +migrate Up
CREATE TABLE IF NOT EXISTS `products` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
LOCK TABLES `products` WRITE;
INSERT INTO `products` VALUES ('111','Sony Playstation 4','Home video game console','500$'),('222','Nintendo Switch','Home video game console','320$'),('333','Xbox One S','Home video game console','400$');
UNLOCK TABLES;


-- +migrate Down
DROP TABLE products;
