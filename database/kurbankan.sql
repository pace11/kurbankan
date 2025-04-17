CREATE DATABASE IF NOT EXISTS kurbankan;

USE kurbankan;

CREATE TABLE `provinces` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(10) UNIQUE NOT NULL,
  `name` varchar(100) NOT NULL
);

CREATE TABLE `regencies` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(10) UNIQUE NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  `province_code` varchar(10)
);

CREATE TABLE `districts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(15) UNIQUE NOT NULL,
  `name` varchar(100) NOT NULL,
  `regency_code` varchar(10)
);

CREATE TABLE `villages` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(20) UNIQUE NOT NULL,
  `name` varchar(100) NOT NULL,
  `district_code` varchar(15)
);

CREATE TABLE `mosques` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `name` text NOT NULL,
  `address` text DEFAULT null,
  `photos` text DEFAULT null,
  `province_code` varchar(10),
  `regency_code` varchar(10),
  `district_code` varchar(15),
  `village_code` varchar(20),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `participants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `name` text NOT NULL,
  `address` text DEFAULT null,
  `province_code` varchar(10),
  `regency_code` varchar(10),
  `district_code` varchar(15),
  `village_code` varchar(20),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` text NOT NULL,
  `role` ENUM ('admin', 'mosque_member', 'user_member') NOT NULL DEFAULT 'user_member',
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `qurban_periods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `year` year NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `description` text DEFAULT null,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `transactions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `qurban_period_id` int,
  `mosque_id` int,
  `qurban_option_id` int,
  `is_full` boolean DEFAULT false,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `transaction_items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `transaction_id` int,
  `participant_id` int,
  `amount` decimal(12,2) NOT NULL,
  `status` ENUM ('pending', 'paid', 'cancelled') NOT NULL DEFAULT 'pending',
  `payment_type` ENUM ('VA') NOT NULL DEFAULT 'VA',
  `external_id` varchar(255) UNIQUE,
  `paid_at` timestamp DEFAULT null,
  `description` text DEFAULT null,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `qurban_options` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `qurban_period_id` int,
  `animal_type` ENUM ('cow', 'goat') NOT NULL,
  `scheme_type` ENUM ('group', 'individual') NOT NULL,
  `price` decimal(12,2) NOT NULL,
  `slots` int DEFAULT 1,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `beneficiaries` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `mosque_id` int,
  `name` text NOT NULL,
  `address` text DEFAULT null,
  `phone` varchar(20) DEFAULT null,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE `qurban_distributions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `qurban_period_id` int,
  `beneficiary_id` int,
  `mosque_id` int,
  `amount` decimal(12,2) NOT NULL,
  `note` text DEFAULT null,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

ALTER TABLE `regencies` ADD FOREIGN KEY (`province_code`) REFERENCES `provinces` (`code`);

ALTER TABLE `districts` ADD FOREIGN KEY (`regency_code`) REFERENCES `regencies` (`code`);

ALTER TABLE `villages` ADD FOREIGN KEY (`district_code`) REFERENCES `districts` (`code`);

ALTER TABLE `mosques` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `mosques` ADD FOREIGN KEY (`province_code`) REFERENCES `provinces` (`code`);

ALTER TABLE `mosques` ADD FOREIGN KEY (`regency_code`) REFERENCES `regencies` (`code`);

ALTER TABLE `mosques` ADD FOREIGN KEY (`district_code`) REFERENCES `districts` (`code`);

ALTER TABLE `mosques` ADD FOREIGN KEY (`village_code`) REFERENCES `villages` (`code`);

ALTER TABLE `participants` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `participants` ADD FOREIGN KEY (`province_code`) REFERENCES `provinces` (`code`);

ALTER TABLE `participants` ADD FOREIGN KEY (`regency_code`) REFERENCES `regencies` (`code`);

ALTER TABLE `participants` ADD FOREIGN KEY (`district_code`) REFERENCES `districts` (`code`);

ALTER TABLE `participants` ADD FOREIGN KEY (`village_code`) REFERENCES `villages` (`code`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`qurban_period_id`) REFERENCES `qurban_periods` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`mosque_id`) REFERENCES `mosques` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`qurban_option_id`) REFERENCES `qurban_options` (`id`);

ALTER TABLE `transaction_items` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);

ALTER TABLE `transaction_items` ADD FOREIGN KEY (`participant_id`) REFERENCES `participants` (`id`);

ALTER TABLE `qurban_options` ADD FOREIGN KEY (`qurban_period_id`) REFERENCES `qurban_periods` (`id`);

ALTER TABLE `beneficiaries` ADD FOREIGN KEY (`mosque_id`) REFERENCES `mosques` (`id`);

ALTER TABLE `qurban_distributions` ADD FOREIGN KEY (`qurban_period_id`) REFERENCES `qurban_periods` (`id`);

ALTER TABLE `qurban_distributions` ADD FOREIGN KEY (`beneficiary_id`) REFERENCES `beneficiaries` (`id`);

ALTER TABLE `qurban_distributions` ADD FOREIGN KEY (`mosque_id`) REFERENCES `mosques` (`id`);
