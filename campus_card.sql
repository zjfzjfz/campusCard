/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : campus_card

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 01/05/2024 21:28:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account_info
-- ----------------------------
DROP TABLE IF EXISTS `account_info`;
CREATE TABLE `account_info`  (
  `c_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL,
  `balance` decimal(10, 0) NULL DEFAULT NULL,
  `validation` date NULL DEFAULT NULL,
  `limit` decimal(10, 0) NULL DEFAULT NULL,
  PRIMARY KEY (`c_id`) USING BTREE,
  INDEX `id`(`id` ASC) USING BTREE,
  CONSTRAINT `account_info_ibfk_1` FOREIGN KEY (`id`) REFERENCES `student_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `account_info_chk_1` CHECK ((`status` >= 0) and (`status` <= 4))
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_info
-- ----------------------------
INSERT INTO `account_info` VALUES ('10321540', '21122703', 0, 22, '2024-08-30', 1000);
INSERT INTO `account_info` VALUES ('15646413', '21122708', 0, 100, '2023-08-30', 200);
INSERT INTO `account_info` VALUES ('23106840', '21122702', 3, 0, '2024-08-30', 200);
INSERT INTO `account_info` VALUES ('32342344', '21122709', 0, 200, '2026-08-30', 500);
INSERT INTO `account_info` VALUES ('44664894', '21122707', 0, 301, '2025-08-30', 100);
INSERT INTO `account_info` VALUES ('49465464', '21122706', 0, 59, '2024-08-30', 200);
INSERT INTO `account_info` VALUES ('53470434', '21122700', 1, 11, '2024-08-30', 100);
INSERT INTO `account_info` VALUES ('54640403', '21122701', 2, 20, '2024-08-30', 200);
INSERT INTO `account_info` VALUES ('65046404', '21122704', 0, 55, '2026-08-30', 500);
INSERT INTO `account_info` VALUES ('96995583', '21122710', 2, 30, '2024-03-30', 1000);
INSERT INTO `account_info` VALUES ('98890264', '21122705', 0, 83, '2026-08-30', 100);

-- ----------------------------
-- Table structure for debt_repayment
-- ----------------------------
DROP TABLE IF EXISTS `debt_repayment`;
CREATE TABLE `debt_repayment`  (
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `b_debt` decimal(10, 0) NULL DEFAULT NULL,
  `l_debt` decimal(10, 0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `debt_repayment_ibfk_1` FOREIGN KEY (`id`) REFERENCES `student_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of debt_repayment
-- ----------------------------
INSERT INTO `debt_repayment` VALUES ('21122700', -2, -20);
INSERT INTO `debt_repayment` VALUES ('21122701', -30, -50);
INSERT INTO `debt_repayment` VALUES ('21122702', 0, -10);
INSERT INTO `debt_repayment` VALUES ('21122703', -20, 0);
INSERT INTO `debt_repayment` VALUES ('21122704', -4, 0);
INSERT INTO `debt_repayment` VALUES ('21122705', -3, 0);
INSERT INTO `debt_repayment` VALUES ('21122706', -1, -25);
INSERT INTO `debt_repayment` VALUES ('21122707', 0, 0);
INSERT INTO `debt_repayment` VALUES ('21122708', 0, 0);
INSERT INTO `debt_repayment` VALUES ('21122709', 0, -36);
INSERT INTO `debt_repayment` VALUES ('21122710', 0, 0);

-- ----------------------------
-- Table structure for student_info
-- ----------------------------
DROP TABLE IF EXISTS `student_info`;
CREATE TABLE `student_info`  (
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `pwd` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `i_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student_info
-- ----------------------------
INSERT INTO `student_info` VALUES ('21122700', 'b59c67bf196a4758191e42f76670ceba', '张小龙', '310113123421122700');
INSERT INTO `student_info` VALUES ('21122701', 'b59c67bf196a4758191e42f76670ceba', '李思明', '310113123421122701');
INSERT INTO `student_info` VALUES ('21122702', 'b59c67bf196a4758191e42f76670ceba', '张伟强', '310113123421122702');
INSERT INTO `student_info` VALUES ('21122703', 'b59c67bf196a4758191e42f76670ceba', '王芳华', '310113123421122703');
INSERT INTO `student_info` VALUES ('21122704', 'b59c67bf196a4758191e42f76670ceba', '赵天宇', '310113123421122704');
INSERT INTO `student_info` VALUES ('21122705', 'b59c67bf196a4758191e42f76670ceba', '刘晓丽', '310113123421122705');
INSERT INTO `student_info` VALUES ('21122706', 'b59c67bf196a4758191e42f76670ceba', '陈文斌', '310113123421122706');
INSERT INTO `student_info` VALUES ('21122707', 'b59c67bf196a4758191e42f76670ceba', '杨静怡', '310113123421122707');
INSERT INTO `student_info` VALUES ('21122708', 'b59c67bf196a4758191e42f76670ceba', '吴昊天', '310113123421122708');
INSERT INTO `student_info` VALUES ('21122709', 'b59c67bf196a4758191e42f76670ceba', '郑小燕', '310113123421122709');
INSERT INTO `student_info` VALUES ('21122710', 'b59c67bf196a4758191e42f76670ceba', '林志勇', '310113123421122710');

-- ----------------------------
-- Table structure for transaction_records
-- ----------------------------
DROP TABLE IF EXISTS `transaction_records`;
CREATE TABLE `transaction_records`  (
  `t_id` int NOT NULL AUTO_INCREMENT,
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `t_type` int NULL DEFAULT NULL,
  `t_location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `t_time` datetime NULL DEFAULT NULL,
  `t_amount` decimal(10, 0) NULL DEFAULT NULL,
  PRIMARY KEY (`t_id`) USING BTREE,
  INDEX `id`(`id` ASC) USING BTREE,
  CONSTRAINT `transaction_records_ibfk_1` FOREIGN KEY (`id`) REFERENCES `student_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `transaction_records_chk_1` CHECK ((`t_type` >= 0) and (`t_type` <= 3))
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transaction_records
-- ----------------------------
INSERT INTO `transaction_records` VALUES (1, '21122710', 0, '宝山东区食堂便利店114', '2006-01-02 15:04:05', 30);
INSERT INTO `transaction_records` VALUES (2, '21122710', 1, '一卡通', '2024-05-01 20:22:23', 20);
INSERT INTO `transaction_records` VALUES (3, '21122710', 1, '一卡通', '2024-05-01 20:23:34', 20);
INSERT INTO `transaction_records` VALUES (4, '21122710', 1, '一卡通', '2024-05-01 20:23:40', 20);
INSERT INTO `transaction_records` VALUES (5, '21122710', 1, '一卡通', '2024-05-01 20:27:37', 20);
INSERT INTO `transaction_records` VALUES (6, '21122710', 3, '一卡通', '2024-05-01 20:28:21', -20);
INSERT INTO `transaction_records` VALUES (7, '21122710', 3, '一卡通', '2024-05-01 20:28:39', -30);
INSERT INTO `transaction_records` VALUES (8, '21122710', 1, '一卡通', '2024-05-01 20:29:55', 20);
INSERT INTO `transaction_records` VALUES (9, '21122710', 3, '一卡通', '2024-05-01 20:30:00', -30);
INSERT INTO `transaction_records` VALUES (10, '21122710', 1, '一卡通', '2024-05-01 20:31:32', 10);
INSERT INTO `transaction_records` VALUES (11, '21122710', 1, '一卡通', '2024-05-01 20:31:36', 10);

-- ----------------------------
-- Triggers structure for table transaction_records
-- ----------------------------
DROP TRIGGER IF EXISTS `update_balance_after_transaction`;
delimiter ;;
CREATE TRIGGER `update_balance_after_transaction` AFTER INSERT ON `transaction_records` FOR EACH ROW BEGIN
    DECLARE current_balance NUMERIC;

    -- 获取当前id的balance
    SELECT balance INTO current_balance
    FROM account_info
    WHERE id = NEW.id
        FOR UPDATE;

    -- 更新balance
    UPDATE account_info
    SET balance = current_balance + NEW.t_amount
    WHERE id = NEW.id;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
