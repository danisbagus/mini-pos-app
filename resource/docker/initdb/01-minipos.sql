CREATE DATABASE minipos;
USE minipos;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `role` varchar(20) NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX idx_user_username
ON users (username);

LOCK TABLES `users` WRITE;

INSERT INTO `users` VALUES 
  (1,'MERCHANT','merchant1@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (2,'MERCHANT','merchant2@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (3,'CUSTOMER','customer1@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22'),
  (4,'CUSTOMER','customer2@live.com','$2a$14$w0cQq0zj1eAuy45xtKIif.3ZUoJbuWClEwK35Xhh3U1pLNOtOf/4m','2020-08-09 10:27:22');


UNLOCK TABLES;
