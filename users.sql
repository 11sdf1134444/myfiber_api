/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3309
 Source Server Type    : MySQL
 Source Server Version : 100035 (10.0.35-MariaDB)
 Source Host           : localhost:3309
 Source Schema         : gin_db

 Target Server Type    : MySQL
 Target Server Version : 100035 (10.0.35-MariaDB)
 File Encoding         : 65001

 Date: 18/02/2025 00:51:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'admin', '123456');
INSERT INTO `users` VALUES (2, 'user', 'password');

SET FOREIGN_KEY_CHECKS = 1;
