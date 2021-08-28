CREATE DATABASE minipos;
USE minipos;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `role` varchar(20) NOT NULL,
  `username` varchar(50) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX idx_user_username
ON users (username);

CREATE TABLE `merchants` (
  `merchant_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `merchant_name` varchar(50) NOT NULL,
  `head_office_address` varchar(100) NOT NULL,
  PRIMARY KEY (`merchant_id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX idx_merchant_name
ON merchants (merchant_name);

CREATE TABLE `customers` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `customer_name` varchar(50) NOT NULL,
  `phone` varchar(20) NOT NULL,
  PRIMARY KEY (`customer_id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `outlets` (
  `outlet_id` int(11) NOT NULL AUTO_INCREMENT,
  `merchant_id` int(11) NOT NULL,
  `outlet_name` varchar(50) NOT NULL,
  `address` varchar(100) NOT NULL,
  PRIMARY KEY (`outlet_id`),
  FOREIGN KEY (`merchant_id`) REFERENCES merchants(`merchant_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `suppliers` (
  `supplier_id` int(11) NOT NULL AUTO_INCREMENT,
  `supplier_name` varchar(50) NOT NULL,
  PRIMARY KEY (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `products` (
  `sku_id` varchar(10)  NOT NULL,
  `merchant_id` int(11) NOT NULL,
  `product_name` varchar(50) NOT NULL,
  `image` varchar(255) NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`sku_id`),
  FOREIGN KEY (`merchant_id`) REFERENCES merchants(`merchant_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX idx_product_name
ON products (product_name);

CREATE TABLE `prices` (
  `sku_id` varchar(10) NOT NULL ,
  `outlet_id` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  FOREIGN KEY (`sku_id`) REFERENCES products(`sku_id`) ON DELETE CASCADE,
  FOREIGN KEY (`outlet_id`) REFERENCES outlets(`outlet_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `purchase_transactions` (
  `transaction_id` varchar(10) NOT NULL,
  `merchant_id` int(11) NOT NULL,
  `sku_id` varchar(10)  NOT NULL,
  `supplier_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  FOREIGN KEY (`merchant_id`) REFERENCES merchants(`merchant_id`) ON DELETE CASCADE,
  FOREIGN KEY (`sku_id`) REFERENCES products(`sku_id`) ON DELETE CASCADE,
  FOREIGN KEY (`supplier_id`) REFERENCES suppliers(`supplier_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `sale_transactions` (
  `transaction_id` varchar(10) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `sku_id` varchar(10)  NOT NULL,
  `outlet_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  FOREIGN KEY (`customer_id`) REFERENCES customers(`customer_id`) ON DELETE CASCADE,
  FOREIGN KEY (`sku_id`) REFERENCES products(`sku_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `users` WRITE;

INSERT INTO `users` VALUES 
  (1,'MERCHANT','merchant1@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (2,'MERCHANT','merchant2@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (3,'CUSTOMER','customer1@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (4,'CUSTOMER','customer2@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22');

UNLOCK TABLES;

LOCK TABLES `merchants` WRITE;

INSERT INTO `merchants` VALUES 
  (1,1,'Nusantara Sport','Jl. Kaliurang 90, Yogyakarta'),
  (2,2,'Jawara Fashion','Jl. Palagan 33, Sleman');

UNLOCK TABLES;

LOCK TABLES `customers` WRITE;

INSERT INTO `customers` VALUES 
  (1,3,'Bobon Kuriniawan','085555555555'),
  (2,4,'Ira Pamungkas','085555555550');

UNLOCK TABLES;

LOCK TABLES `outlets` WRITE;

INSERT INTO `outlets` VALUES 
  (1,1,'Nusantara Sport Cabang Bantul','Jl. Bung Hatta 93, Bantul'),
  (2,1,'Nusantara Sport Cabang Kolon Progo','Jl. Bung Tomo 23,  Kolon Progo'),
  (3,2,'Jawara Fashion Cabang UGM','Jl. M.Yamin 93, Slemam'),
  (4,2,'Jawara Fashion Cabang UNY','Jl. Juanda 23,  Sleman');

UNLOCK TABLES;

LOCK TABLES `suppliers` WRITE;

INSERT INTO `suppliers` VALUES 
  (1,'PT Sepatu Jaya'),
  (2,'PT Bangun Sejahter'),
  (3,'CV Merah Merona');

UNLOCK TABLES;

LOCK TABLES `products` WRITE;

INSERT INTO `products` VALUES 
  ('PGK8R6O',1,'Sneaker Modern 15','localhost:7000/public/uploads/1630123199238701473.png', 100),
  ('PSK8RPO',1,'Flatshoes Panorama','localhost:7000/public/uploads/1630123199238701473.png', 20),
  ('PLK8R0O',2,'Casual Eterna','localhost:7000/public/uploads/1630123199238701473.png', 30),
  ('PMK8OPO',2,'Cressidi Omega','localhost:7000/public/uploads/1630123199238701473.png',90);

UNLOCK TABLES;

LOCK TABLES `prices` WRITE;

INSERT INTO `prices` VALUES 
  ('PGK8R6O',1, 800000),
  ('PGK8R6O',2, 1000000),
  ('PSK8RPO',1, 400000),
  ('PSK8RPO',2, 300000),
  ('PLK8R0O',3, 50000),
  ('PLK8R0O',4, 55000),
  ('PMK8OPO',3, 125000),
  ('PMK8OPO',4, 150000);

UNLOCK TABLES;

LOCK TABLES `purchase_transactions` WRITE;

INSERT INTO `purchase_transactions` VALUES 
  ('PTK8R6O',1,'PGK8R6O',1, 120, 500000000, '2020-08-09 10:27:22'),
  ('PTK8RPO',1,'PSK8RPO',2, 100, 160000000, '2020-08-09 10:27:22'),
  ('PTK8R0O',2,'PLK8R0O',2, 50, 2000000, '2020-08-09 10:27:22'),
  ('PTK8OPO',2,'PMK8OPO',3, 100, 10000000, '2020-08-09 10:27:22');

UNLOCK TABLES;

LOCK TABLES `sale_transactions` WRITE;

INSERT INTO `sale_transactions` VALUES 
  ('PTK8R6O',1,'PGK8R6O',1, 20, '2020-08-09 10:27:22'),
  ('PTK8RPO',2,'PSK8RPO',2, 20,'2020-08-09 10:27:22'),
  ('PTK8R0O',1,'PLK8R0O',3, 20,  '2020-08-09 10:27:22'),
  ('PTK8OPO',2,'PMK8OPO',4, 10, '2020-08-09 10:27:22');

UNLOCK TABLES;
