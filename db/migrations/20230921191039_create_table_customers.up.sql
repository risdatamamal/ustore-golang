CREATE TABLE `customers` 
(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `user_name` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `full_name` VARCHAR(255),
  `role` VARCHAR(255) NOT NULL DEFAULT 'customer',
  PRIMARY KEY (`id`),UNIQUE INDEX (`email`)
)ENGINE = InnoDB;