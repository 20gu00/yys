-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: gotravel_admin
-- ------------------------------------------------------
-- Server version	5.7.38

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `gotravel_admin`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `gotravel_admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gotravel_admin`;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `introduction` varchar(50) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `user_id` int(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='admin';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'2022-09-11 19:40:29','2022-09-11 19:40:29','admin','admin','adminManager','2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3',1);
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Current Database: `gotravel_usercenter`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `gotravel_usercenter` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gotravel_usercenter`;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  `mobile` char(11) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `info` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'2022-09-10 20:56:49','2022-09-10 20:56:49','1970-01-01 08:00:00',0,0,'13019991092','76a2173be6393254e72ffa4d6df1030a','m3o90woC',0,'','');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_auth`
--

DROP TABLE IF EXISTS `user_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_auth` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  `user_id` bigint(20) NOT NULL DEFAULT '0',
  `auth_key` varchar(64) NOT NULL DEFAULT '' COMMENT '平台唯一id',
  `auth_type` varchar(12) NOT NULL DEFAULT '' COMMENT '平台类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE,
  UNIQUE KEY `idx_userId_key` (`user_id`,`auth_type`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户授权表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_auth`
--

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES (1,'2022-09-10 20:56:49','2022-09-10 20:56:49','1970-01-01 08:00:00',0,0,1,'13019991092','system'),(2,'2022-09-10 21:03:36','2022-09-10 21:03:36','1970-01-01 08:00:00',0,0,2,'13019991091','system');
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Current Database: `gotravel_travel`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `gotravel_travel` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gotravel_travel`;

--
-- Table structure for table `homestay`
--

DROP TABLE IF EXISTS `homestay`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homestay` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '标题',
  `sub_title` varchar(128) NOT NULL DEFAULT '' COMMENT '副标题',
  `banner` varchar(4096) NOT NULL DEFAULT '' COMMENT '轮播图，第一张封面',
  `info` varchar(4069) NOT NULL DEFAULT '' COMMENT '介绍',
  `people_num` tinyint(1) NOT NULL DEFAULT '0' COMMENT '容纳人的数量',
  `homestay_business_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿店铺id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房东id，冗余字段',
  `row_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:下架 1:上架',
  `row_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '售卖类型0：按房间出售 1:按人次出售',
  `food_info` varchar(2048) NOT NULL COMMENT '餐食标准',
  `food_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '餐食价格（分）',
  `homestay_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿价格（分）',
  `market_homestay_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿市场价格（分）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='每一间民宿';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homestay`
--

LOCK TABLES `homestay` WRITE;
/*!40000 ALTER TABLE `homestay` DISABLE KEYS */;
INSERT INTO `homestay` VALUES (1,'2022-09-11 18:04:04','2022-09-11 21:49:30','1970-01-01 08:00:00',0,100,'cheap','LOFT复式','小橙公寓','广东省江门市江海区又一居',100,999,1000,1,0,'满汉全席+kfc',1009,2009,2509),(12,'2022-09-11 21:48:32','2022-09-11 21:48:32','1970-01-01 08:00:00',0,10,'浪琴湾','LOFT复式,北欧商务大床房,近万达广场','小橙公寓','广东省江门市江海区又一居',10,99,100,1,0,'满汉全席',100,200,250);
/*!40000 ALTER TABLE `homestay` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `homestay_activity`
--

DROP TABLE IF EXISTS `homestay_activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homestay_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `row_type` varchar(32) NOT NULL DEFAULT '' COMMENT '活动类型',
  `data_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '业务表id（id跟随活动类型走）',
  `row_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:下架 1:上架',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  `activity_id` int(100) NOT NULL DEFAULT '0' COMMENT 'activity_id',
  PRIMARY KEY (`id`),
  KEY `idx_rowType` (`row_type`,`row_status`,`del_state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='每一间民宿';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homestay_activity`
--

LOCK TABLES `homestay_activity` WRITE;
/*!40000 ALTER TABLE `homestay_activity` DISABLE KEYS */;
INSERT INTO `homestay_activity` VALUES (12,'2022-09-11 21:50:13','2022-09-11 21:50:13','1970-01-01 08:00:00',0,'情人节促销',1,1,11,0),(13,'2022-09-11 21:50:21','2022-09-11 21:50:21','1970-01-01 08:00:00',0,'情人节促销',1,1,99,0);
/*!40000 ALTER TABLE `homestay_activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `homestay_business`
--

DROP TABLE IF EXISTS `homestay_business`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homestay_business` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '店铺名称',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联的用户id',
  `info` varchar(128) NOT NULL DEFAULT '' COMMENT '店铺介绍',
  `boss_info` varchar(128) NOT NULL DEFAULT '' COMMENT '房东介绍',
  `license_fron` varchar(256) NOT NULL DEFAULT '' COMMENT '营业执照正面',
  `license_back` varchar(256) NOT NULL DEFAULT '' COMMENT '营业执照背面',
  `row_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:禁止营业 1:正常营业',
  `star` double(2,1) NOT NULL DEFAULT '0.0' COMMENT '店铺整体评价，冗余',
  `tags` varchar(32) NOT NULL DEFAULT '' COMMENT '每个店家一个标签，自己编辑',
  `cover` varchar(1024) NOT NULL DEFAULT '' COMMENT '封面图',
  `header_img` varchar(1024) NOT NULL DEFAULT '' COMMENT '店招门头图片',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_userId` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='民宿店铺';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homestay_business`
--

LOCK TABLES `homestay_business` WRITE;
/*!40000 ALTER TABLE `homestay_business` DISABLE KEYS */;
INSERT INTO `homestay_business` VALUES (4,'2022-09-11 21:54:42','2022-09-11 21:54:42','1970-01-01 08:00:00',0,'恒大泉都',1,'豪华度假别墅','人好说话','licensefron.png','lcenseback.png',1,0.0,'good','cover.png','',1);
/*!40000 ALTER TABLE `homestay_business` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `homestay_comment`
--

DROP TABLE IF EXISTS `homestay_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homestay_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `homestay_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `content` varchar(64) NOT NULL DEFAULT '' COMMENT '评论内容',
  `star` float NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='民宿评价';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homestay_comment`
--

LOCK TABLES `homestay_comment` WRITE;
/*!40000 ALTER TABLE `homestay_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `homestay_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Current Database: `gotravel_order`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `gotravel_order` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gotravel_order`;

--
-- Table structure for table `homestay_order`
--

DROP TABLE IF EXISTS `homestay_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homestay_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(4) NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本号',
  `sn` char(25) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '下单用户id',
  `homestay_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿id',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '标题',
  `sub_title` varchar(128) NOT NULL DEFAULT '' COMMENT '副标题',
  `cover` varchar(1024) NOT NULL DEFAULT '' COMMENT '封面',
  `info` varchar(4069) NOT NULL DEFAULT '' COMMENT '介绍',
  `people_num` tinyint(1) NOT NULL DEFAULT '0' COMMENT '容纳人的数量',
  `row_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '售卖类型0：按房间出售 1:按人次出售',
  `need_food` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:不需要餐食 1:需要参数',
  `food_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '餐食标准',
  `food_price` bigint(20) NOT NULL COMMENT '餐食价格(分)',
  `homestay_price` bigint(20) NOT NULL COMMENT '民宿价格(分)',
  `market_homestay_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿市场价格(分)',
  `homestay_business_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺id',
  `homestay_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺房东id',
  `live_start_date` date NOT NULL COMMENT '开始入住日期',
  `live_end_date` date NOT NULL COMMENT '结束入住日期',
  `live_people_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '实际入住人数',
  `trade_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '-1: 已取消 0:待支付 1:未使用 2:已使用  3:已退款 4:已过期',
  `trade_code` char(8) NOT NULL DEFAULT '' COMMENT '确认码',
  `remark` varchar(64) NOT NULL DEFAULT '' COMMENT '用户下单备注',
  `order_total_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单总价格（餐食总价格+民宿总价格）(分)',
  `food_total_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '餐食总价格(分)',
  `homestay_total_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '民宿总价格(分)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sn` (`sn`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='每一间民宿';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homestay_order`
--

LOCK TABLES `homestay_order` WRITE;
/*!40000 ALTER TABLE `homestay_order` DISABLE KEYS */;
INSERT INTO `homestay_order` VALUES (2,'2022-09-11 20:41:38','2022-09-11 20:41:38','1970-01-01 08:00:00',0,1,'991100778899',1,999,'海景房下单','good','cover.png','海景房',10,0,1,'满汉全席',100,200,250,999,1,'1970-01-01','1970-01-01',10,0,'999','gogogo',300,100,500),(3,'2022-09-11 20:41:50','2022-09-11 20:41:50','1970-01-01 08:00:00',0,1,'129991100778899',1,999,'海景房下单','good','cover.png','海景房',10,0,1,'满汉全席',100,200,250,999,1,'1970-01-01','1970-01-01',10,0,'999','gogogo',300,100,500);
/*!40000 ALTER TABLE `homestay_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Current Database: `gotravel_payment`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `gotravel_payment` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gotravel_payment`;

--
-- Table structure for table `third_payment`
--

DROP TABLE IF EXISTS `third_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `third_payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sn` char(25) NOT NULL DEFAULT '' COMMENT '流水单号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint(1) NOT NULL DEFAULT '0',
  `version` bigint(20) NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `pay_mode` varchar(20) NOT NULL DEFAULT '' COMMENT '支付方式 1:微信支付',
  `trade_type` varchar(20) NOT NULL DEFAULT '' COMMENT '第三方支付类型',
  `trade_state` varchar(20) NOT NULL DEFAULT '' COMMENT '第三方交易状态',
  `pay_total` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付总金额(分)',
  `transaction_id` char(32) NOT NULL DEFAULT '' COMMENT '第三方支付单号',
  `trade_state_desc` varchar(256) NOT NULL DEFAULT '' COMMENT '支付状态描述',
  `order_sn` char(25) NOT NULL DEFAULT '' COMMENT '业务单号',
  `service_type` varchar(32) NOT NULL DEFAULT '' COMMENT '业务类型 ',
  `pay_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '平台内交易状态   -1:支付失败 0:未支付 1:支付成功 2:已退款',
  `pay_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:00' COMMENT '支付成功时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sn` (`sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方支付流水记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `third_payment`
--

LOCK TABLES `third_payment` WRITE;
/*!40000 ALTER TABLE `third_payment` DISABLE KEYS */;
/*!40000 ALTER TABLE `third_payment` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-12 10:51:03
