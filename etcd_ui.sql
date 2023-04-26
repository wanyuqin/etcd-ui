/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : etcd_ui

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 26/04/2023 09:43:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for certificate
-- ----------------------------
DROP TABLE IF EXISTS `certificate`;
CREATE TABLE `certificate` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '自定义名称',
  `key` text COLLATE utf8mb4_general_ci,
  `ca` text COLLATE utf8mb4_general_ci,
  `cert` text COLLATE utf8mb4_general_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for connection
-- ----------------------------
DROP TABLE IF EXISTS `connection`;
CREATE TABLE `connection` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '自定义名称',
  `version` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ETCD 版本',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '连接类型 1 默认',
  `created_at` timestamp NULL DEFAULT NULL,
  `tls` tinyint DEFAULT '1' COMMENT '是否开启tls 1 关闭 2 开启',
  `active` tinyint DEFAULT NULL COMMENT '是否激活',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `endpoint` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '连接端点',
  `certificate_id` int DEFAULT '0' COMMENT '关联证书ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
