-- MySQL dump 10.13  Distrib 5.7.26, for Linux (x86_64)
--
-- Host: localhost    Database: bug
-- ------------------------------------------------------
-- Server version	5.7.26-log

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
-- Table structure for table `apilist`
--

DROP TABLE IF EXISTS `apilist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `apilist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `pid` bigint(20) DEFAULT NULL,
  `url` varchar(200) DEFAULT NULL,
  `information` varchar(255) DEFAULT NULL,
  `opts` varchar(100) DEFAULT NULL,
  `methods` varchar(100) DEFAULT NULL,
  `resp` text,
  `result` text,
  `uid` bigint(20) DEFAULT NULL,
  `hid` bigint(20) DEFAULT NULL,
  `calltype` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apilist`
--

LOCK TABLES `apilist` WRITE;
/*!40000 ALTER TABLE `apilist` DISABLE KEYS */;
/*!40000 ALTER TABLE `apilist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `apiproject`
--

DROP TABLE IF EXISTS `apiproject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `apiproject` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `ownerid` varchar(100) DEFAULT '0',
  `auth` tinyint(1) DEFAULT '0',
  `readuser` tinyint(1) DEFAULT '0',
  `edituser` tinyint(1) DEFAULT '0',
  `rid` tinyint(1) DEFAULT '0',
  `eid` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apiproject`
--

LOCK TABLES `apiproject` WRITE;
/*!40000 ALTER TABLE `apiproject` DISABLE KEYS */;
/*!40000 ALTER TABLE `apiproject` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bugs`
--

DROP TABLE IF EXISTS `bugs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bugs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL default '0',
  `title` varchar(50) NOT NULL DEFAULT '',
  `sid` bigint(20) DEFAULT '0',
  `content` text,
  `ownerid` bigint(20) DEFAULT '0',
  `iid` bigint(20) DEFAULT '0',
  `createtime` bigint(20) DEFAULT '0',
  `vid` bigint(20) DEFAULT '0',
  `spusers` varchar(255) DEFAULT '',
  `lid` bigint(20) DEFAULT '0',
  `eid` bigint(20) DEFAULT '0',
  `pid` bigint(20) DEFAULT '0',
  `updatetime` bigint(20) DEFAULT '0',
  `dustbin` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bugs`
--

LOCK TABLES `bugs` WRITE;
/*!40000 ALTER TABLE `bugs` DISABLE KEYS */;
/*!40000 ALTER TABLE `bugs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `defaultvalue`
--

DROP TABLE IF EXISTS `defaultvalue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `defaultvalue` (
  `status` bigint(20) NOT NULL DEFAULT '0',
  `important` bigint(20) NOT NULL DEFAULT '0',
  `level` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `defaultvalue`
--

LOCK TABLES `defaultvalue` WRITE;
/*!40000 ALTER TABLE `defaultvalue` DISABLE KEYS */;
INSERT INTO `defaultvalue` VALUES (0,0,0);
/*!40000 ALTER TABLE `defaultvalue` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `email`
--

DROP TABLE IF EXISTS `email`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `password` varchar(50) DEFAULT '',
  `port` int(11) DEFAULT '25',
  `createuser` tinyint(1) DEFAULT '0',
  `createbug` tinyint(1) DEFAULT '0',
  `passbug` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `email`
--

LOCK TABLES `email` WRITE;
/*!40000 ALTER TABLE `email` DISABLE KEYS */;
/*!40000 ALTER TABLE `email` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `environment`
--

DROP TABLE IF EXISTS `environment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `environment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `envname` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `envname` (`envname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `environment`
--

LOCK TABLES `environment` WRITE;
/*!40000 ALTER TABLE `environment` DISABLE KEYS */;
/*!40000 ALTER TABLE `environment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `header`
--

DROP TABLE IF EXISTS `header`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `header` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `hhids` varchar(100) DEFAULT '',
  `remark` varchar(30) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `header`
--

LOCK TABLES `header` WRITE;
/*!40000 ALTER TABLE `header` DISABLE KEYS */;
/*!40000 ALTER TABLE `header` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `headerlist`
--

DROP TABLE IF EXISTS `headerlist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `headerlist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `k` varchar(200) NOT NULL,
  `v` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `headerlist`
--

LOCK TABLES `headerlist` WRITE;
/*!40000 ALTER TABLE `headerlist` DISABLE KEYS */;
/*!40000 ALTER TABLE `headerlist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `importants`
--

DROP TABLE IF EXISTS `importants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `importants` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `importants`
--

LOCK TABLES `importants` WRITE;
/*!40000 ALTER TABLE `importants` DISABLE KEYS */;
/*!40000 ALTER TABLE `importants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `informations`
--

DROP TABLE IF EXISTS `informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `informations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `bid` bigint(20) NOT NULL DEFAULT '0',
  `info` varchar(200) NOT NULL DEFAULT '',
  `time` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `informations`
--

LOCK TABLES `informations` WRITE;
/*!40000 ALTER TABLE `informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jobs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `level` bigint(20) NOT NULL DEFAULT '2',
  `hypo` varchar(30) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `level`
--

DROP TABLE IF EXISTS `level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `level`
--

LOCK TABLES `level` WRITE;
/*!40000 ALTER TABLE `level` DISABLE KEYS */;
/*!40000 ALTER TABLE `level` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `exectime` bigint(20) DEFAULT '0',
  `classify` varchar(30) NOT NULL DEFAULT '',
  `content` text,
  `ip` varchar(40) DEFAULT '',
  `username` varchar(50) DEFAULT '',
  `action` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (1,1588911474,'setting','','127.0.0.1','admin','deluser');
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `options`
--

DROP TABLE IF EXISTS `options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `options` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `info` varchar(100) DEFAULT '',
  `tid` bigint(20) DEFAULT '0',
  `df` varchar(10) DEFAULT '',
  `need` varchar(10) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `options`
--

LOCK TABLES `options` WRITE;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
/*!40000 ALTER TABLE `options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;

CREATE TABLE `project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `participant` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 
--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restfulname`
--

DROP TABLE IF EXISTS `restfulname`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `restfulname` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restfulname`
--

LOCK TABLES `restfulname` WRITE;
/*!40000 ALTER TABLE `restfulname` DISABLE KEYS */;
/*!40000 ALTER TABLE `restfulname` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rolegroup`
--

DROP TABLE IF EXISTS `rolegroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rolegroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `permids` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `perm`;
CREATE TABLE `perm` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `find` tinyint(1) NOT NULL DEFAULT '0',
  `remove` tinyint(1) NOT NULL DEFAULT '0',
  `revise` tinyint(1) NOT NULL DEFAULT '0',
  `increase` tinyint(1) NOT NULL DEFAULT '0',
  `rid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 
--
-- Dumping data for table `rolegroup`
--

LOCK TABLES `rolegroup` WRITE;
/*!40000 ALTER TABLE `rolegroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `rolegroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role` varchar(30) NOT NULL DEFAULT '',
  `info` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role` (`role`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 

INSERT INTO roles(role,info) VALUES('env', "环境页面");
INSERT INTO roles(role,info) VALUES('important', "重要性页面");
INSERT INTO roles(role,info) VALUES('level', "重要性页面");
INSERT INTO roles(role,info) VALUES('log', "查看日志");
INSERT INTO roles(role,info) VALUES('position', "职位页面");
INSERT INTO roles(role,info) VALUES('project', "项目页面");
INSERT INTO roles(role,info) VALUES('status', "bug状态流程页面");
INSERT INTO roles(role,info) VALUES('statusgroup', "状态组页面");
INSERT INTO roles(role,info) VALUES('version', "版本页面");
INSERT INTO roles(role,info) VALUES('allbug', "所有bug的页面");
--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sharefile`
--

DROP TABLE IF EXISTS `sharefile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sharefile` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `filepath` varchar(200) NOT NULL,
  `readuser` tinyint(1) DEFAULT '0',
  `rid` bigint(20) DEFAULT '0',
  `isfile` tinyint(1) DEFAULT '0',
  `ownerid` bigint(20) DEFAULT '0',
  `wid` bigint(20) DEFAULT '0',
  `writeuser` tinyint(1) DEFAULT '0',
  `size` bigint(20) DEFAULT '0',
  `updatetime` bigint(20) DEFAULT '0',
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sharefile`
--

LOCK TABLES `sharefile` WRITE;
/*!40000 ALTER TABLE `sharefile` DISABLE KEYS */;
/*!40000 ALTER TABLE `sharefile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `status`
--

DROP TABLE IF EXISTS `status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status`
--

LOCK TABLES `status` WRITE;
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
/*!40000 ALTER TABLE `status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `statusgroup`
--

DROP TABLE IF EXISTS `statusgroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `statusgroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `sids` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `statusgroup`
--

LOCK TABLES `statusgroup` WRITE;
/*!40000 ALTER TABLE `statusgroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `statusgroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `types`
--

DROP TABLE IF EXISTS `types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `type` bigint(20) NOT NULL DEFAULT '0',
  `opts` varchar(200) DEFAULT '',
  `tid` bigint(20) DEFAULT '0',
  `default` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `types`
--

LOCK TABLES `types` WRITE;
/*!40000 ALTER TABLE `types` DISABLE KEYS */;
/*!40000 ALTER TABLE `types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) NOT NULL,
  `password` varchar(40) NOT NULL,
  `email` varchar(50) NOT NULL,
  `headimg` varchar(100) DEFAULT '',
  `createtime` bigint(20) DEFAULT '0',
  `createuid` bigint(20) DEFAULT '0',
  `realname` varchar(30) NOT NULL,
  `showstatus` varchar(200) DEFAULT '',
  `disable` tinyint(1) DEFAULT '0',
  `bugsid` bigint(20) DEFAULT '0',
  `level` bigint(20) DEFAULT '2',
  `rid` bigint(20) DEFAULT '0',
  `jid` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `realname` (`realname`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','69ad5117e7553ecfa7f918a223426dd8da08a57f','admin@qq.com','http://127.0.0.1:10001/showimg/1588816996826858100.png',1557131883,0,'admin','',0,0,0,0,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usergroup`
--

DROP TABLE IF EXISTS `usergroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usergroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `ids` varchar(200) DEFAULT '',
  `cuid` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usergroup`
--

LOCK TABLES `usergroup` WRITE;
/*!40000 ALTER TABLE `usergroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `usergroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `version`
--

DROP TABLE IF EXISTS `version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `version` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `urlone` varchar(30) DEFAULT '',
  `urltwo` varchar(30) DEFAULT '',
  `createtime` varchar(30) DEFAULT '0',
  `createuid` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `version`
--

LOCK TABLES `version` WRITE;
/*!40000 ALTER TABLE `version` DISABLE KEYS */;
/*!40000 ALTER TABLE `version` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-08 12:18:38
