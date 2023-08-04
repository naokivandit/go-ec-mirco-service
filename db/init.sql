CREATE DATABASE IF NOT EXISTS `order`;
USE `order`;
SOURCE /docker-entrypoint-initdb.d/migrations/order/000001_create_orders_table.up.sql;
SOURCE /docker-entrypoint-initdb.d/migrations/order/000002_create_order_items_table.up.sql;

CREATE DATABASE IF NOT EXISTS `product`;
use `product`;
SOURCE /docker-entrypoint-initdb.d/migrations/product/000001_create_products_table.up.sql;
SOURCE /docker-entrypoint-initdb.d/migrations/product/000002_insert_products.up.sql;

CREATE DATABASE IF NOT EXISTS `setting`;
use `setting`;
SOURCE /docker-entrypoint-initdb.d/migrations/setting/000001_create_settings_table.up.sql;
SOURCE /docker-entrypoint-initdb.d/migrations/setting/000002_insert_settings.up.sql;