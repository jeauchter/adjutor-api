CREATE DATABASE IF NOT EXISTS auth;
CREATE DATABASE IF NOT EXISTS inventory;
CREATE DATABASE IF NOT EXISTS product;
CREATE DATABASE IF NOT EXISTS salesorder;

GRANT ALL ON `auth`.* TO 'dev'@'%';
GRANT ALL ON `inventory`.* TO 'dev'@'%';
GRANT ALL ON `product`.* TO 'dev'@'%';
GRANT ALL ON `salesorder`.* TO 'dev'@'%';