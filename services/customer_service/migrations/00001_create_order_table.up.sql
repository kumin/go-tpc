CREATE TABLE IF NOT EXISTS `order` (
  id INT PRIMARY KEY AUTO_INCREMENT,
  num_products INT NOT NULL,
  money VARCHAR(255) NOT NULL,
  shipping_address TEXT NOT NULL,
  status INT NOT NULL
);
