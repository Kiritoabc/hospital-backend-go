/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : hospital

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 02/07/2023 11:16:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for call
-- ----------------------------
DROP TABLE IF EXISTS `call`;
CREATE TABLE `call`  (
  `call_id` int NOT NULL COMMENT '出诊ID',
  `doctor_id` int NULL DEFAULT NULL COMMENT '医生ID',
  `call_date` date NULL DEFAULT NULL COMMENT '出诊日期',
  `from_time` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `to_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `doctor_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '医生姓名',
  PRIMARY KEY (`call_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of call
-- ----------------------------

-- ----------------------------
-- Table structure for depart
-- ----------------------------
DROP TABLE IF EXISTS `depart`;
CREATE TABLE `depart`  (
  `depart_id` int NOT NULL AUTO_INCREMENT COMMENT '科室ID',
  `depart_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '科室名称',
  `depart_introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '科室介绍',
  PRIMARY KEY (`depart_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of depart
-- ----------------------------
INSERT INTO `depart` VALUES (1, '呼吸内科', '呼吸类');
INSERT INTO `depart` VALUES (2, '消化内科', '消化系统');
INSERT INTO `depart` VALUES (3, '神经内科', '神经类');
INSERT INTO `depart` VALUES (4, '心血管内科', '心血管内科');
INSERT INTO `depart` VALUES (5, '肾内科', '肾内科');
INSERT INTO `depart` VALUES (6, '血液内科', '血液内科');
INSERT INTO `depart` VALUES (7, '免疫科', '免疫科');
INSERT INTO `depart` VALUES (8, '内分泌科', '内分泌科');
INSERT INTO `depart` VALUES (9, '妇科', '包治百病');

-- ----------------------------
-- Table structure for doctor
-- ----------------------------
DROP TABLE IF EXISTS `doctor`;
CREATE TABLE `doctor`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '唯一识别ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '医生姓名',
  `sex` int NOT NULL COMMENT '性别 0:女，1:男，2：其他',
  `birthday` date NULL DEFAULT NULL COMMENT '出生日期',
  `id_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '身份证号',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号码',
  `fee` double NULL DEFAULT NULL COMMENT '挂号费',
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '基本介绍',
  `depart_id` int NULL DEFAULT NULL COMMENT '科室ID',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '生成的ID',
  `role` int NOT NULL COMMENT '角色 0：医生，1：管理员，2：医生加管理员',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of doctor
-- ----------------------------
INSERT INTO `doctor` VALUES (1, '张三', 1, NULL, NULL, '15065211470', NULL, NULL, 1, 'null', '123455', 0);
INSERT INTO `doctor` VALUES (2, '李四', 1, NULL, NULL, '15065211111', NULL, NULL, 1, 'null', '123455', 1);

-- ----------------------------
-- Table structure for room
-- ----------------------------
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room`  (
  `room_id` int NOT NULL AUTO_INCREMENT COMMENT '诊室ID',
  `room_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '诊室名称',
  `room_introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '诊室介绍',
  `depart_id` int NOT NULL COMMENT '隶属科室ID',
  PRIMARY KEY (`room_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of room
-- ----------------------------
INSERT INTO `room` VALUES (2, 'nihao', '包治百病', 1);
INSERT INTO `room` VALUES (3, '147', '包治百病', 1);
INSERT INTO `room` VALUES (4, '1请问是47', '包治百病', 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '15065211111', '123456');
INSERT INTO `user` VALUES (2, '15065211470', '123456');

SET FOREIGN_KEY_CHECKS = 1;
