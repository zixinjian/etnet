/*
Navicat SQLite Data Transfer

Source Server         : main
Source Server Version : 30802
Source Host           : :0

Target Server Type    : SQLite
Target Server Version : 30802
File Encoding         : 65001

Date: 2015-12-31 11:53:41
*/

PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "main"."user";
CREATE TABLE "user" (
"id"  integer PRIMARY KEY AUTOINCREMENT NOT NULL,
"sn"  varchar NOT NULL COLLATE BINARY ,
"name"  varchar NOT NULL COLLATE BINARY ,
"username"  varchar NOT NULL COLLATE BINARY ,
"password"  varchar NOT NULL,
"company"  varchar NOT NULL,
"department"  varchar NOT NULL,
"role"  varchar NOT NULL DEFAULT role_user,
"flag"  varchar NOT NULL DEFAULT flag_available,
"createtime"  varchar NOT NULL,
"creater"  varchar NOT NULL,
UNIQUE ("name" ASC, "sn" ASC)
);
