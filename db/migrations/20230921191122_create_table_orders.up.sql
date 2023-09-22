CREATE TABLE `orders` 
(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `customer_id` BIGINT NOT NULL,
  `order_date` TIMESTAMP NOT NULL,
  `total_amount` INT NOT NULL,
  PRIMARY KEY (`id`),FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`)
)ENGINE = InnoDB;