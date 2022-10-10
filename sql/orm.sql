/*
 Navicat MySQL Data Transfer

 Source Server         : tyj
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : orm

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/10/2022 18:03:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `cno` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tno` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`cno`, `tno`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course` VALUES ('c001', 'J2SE', 't002');
INSERT INTO `course` VALUES ('c002', 'Java Web', 't002');
INSERT INTO `course` VALUES ('c003', 'SSH', 't001');
INSERT INTO `course` VALUES ('c004', 'Oracle', 't001');
INSERT INTO `course` VALUES ('c005', 'SQL SERVER 2005', 't003');
INSERT INTO `course` VALUES ('c006', 'C#', 't003');
INSERT INTO `course` VALUES ('c007', 'JavaScript', 't002');
INSERT INTO `course` VALUES ('c008', 'DIV+CSS', 't001');
INSERT INTO `course` VALUES ('c009', 'PHP', 't003');
INSERT INTO `course` VALUES ('c010', 'EJB3.0', 't002');

-- ----------------------------
-- Table structure for sc
-- ----------------------------
DROP TABLE IF EXISTS `sc`;
CREATE TABLE `sc`  (
  `sno` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cno` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `score` decimal(4, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`sno`, `cno`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sc
-- ----------------------------
INSERT INTO `sc` VALUES ('s001', 'c001', 78.90);
INSERT INTO `sc` VALUES ('s001', 'c002', 82.90);
INSERT INTO `sc` VALUES ('s001', 'c003', 59.00);
INSERT INTO `sc` VALUES ('s002', 'c001', 80.90);
INSERT INTO `sc` VALUES ('s002', 'c002', 72.90);
INSERT INTO `sc` VALUES ('s003', 'c001', 81.90);
INSERT INTO `sc` VALUES ('s003', 'c002', 81.90);
INSERT INTO `sc` VALUES ('s004', 'c001', 60.90);

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`  (
  `sno` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sage` decimal(2, 0) NULL DEFAULT NULL,
  `ssex` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`sno`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student
-- ----------------------------
INSERT INTO `student` VALUES ('s001', '张三', 23, '男');
INSERT INTO `student` VALUES ('s002', '李四', 23, '男');
INSERT INTO `student` VALUES ('s003', '吴鹏', 25, '男');
INSERT INTO `student` VALUES ('s004', '琴沁', 20, '女');
INSERT INTO `student` VALUES ('s005', '王丽', 20, '女');
INSERT INTO `student` VALUES ('s006', '李波', 21, '男');
INSERT INTO `student` VALUES ('s007', '刘玉', 21, '男');
INSERT INTO `student` VALUES ('s008', '萧蓉', 21, '女');
INSERT INTO `student` VALUES ('s009', '陈萧晓', 23, '女');
INSERT INTO `student` VALUES ('s010', '陈美', 22, '女');
INSERT INTO `student` VALUES ('s011', '王丽', 21, '女');

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`  (
  `tno` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `tname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`tno`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teacher
-- ----------------------------
INSERT INTO `teacher` VALUES ('t001', '刘阳');
INSERT INTO `teacher` VALUES ('t002', '谌燕');
INSERT INTO `teacher` VALUES ('t003', '胡明星');

SET FOREIGN_KEY_CHECKS = 1;
