-- MySQL dump 10.13  Distrib 5.7.15, for Linux (x86_64)
--
-- Host: 172.19.0.6    Database: project
-- ------------------------------------------------------
-- Server version	5.7.27-log

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apilist`
--

LOCK TABLES `apilist` WRITE;
/*!40000 ALTER TABLE `apilist` DISABLE KEYS */;
INSERT INTO `apilist` VALUES (1,'登陆',1,'/login/login','','1,2','POST','{&#34;username&#34;,&#34;admin@qq.com&#34;,&#34;password&#34;:&#34;admin&#34;}','{&#34;username&#34;: &#34;admin&#34;, &#34;password&#34;: &#34;&#34;, &#34;token&#34;: &#34;7ccfb389542777936a64851c71a8297c06b7d6ae&#34;, &#34;statuscode&#34;: 0, &#34;avatar&#34;: &#34;&#34;}',1,0,'json'),(2,'获取用户信息',1,'/user/info','','','POST','','',1,1,'json'),(3,'获取用户总个数',1,'/dashboard/usercount','','','POST','','',1,1,'json'),(4,'获取所有bug',1,'/search/allbugs','','1718','POST','{\n        page: 1,\n        limit: 15,\n        level: &#39;&#39;,\n        project: &#39;&#39;,\n        title: &#39;&#39;,\n        status: []\n      }','',1,1,'json');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apiproject`
--

LOCK TABLES `apiproject` WRITE;
/*!40000 ALTER TABLE `apiproject` DISABLE KEYS */;
INSERT INTO `apiproject` VALUES (1,'itflow','1',0,0,0,0,0);
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
  `uid` bigint(20) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `sid` bigint(20) DEFAULT '0',
  `content` text,
  `ownerid` bigint(20) DEFAULT '0',
  `iid` bigint(20) DEFAULT '0',
  `createtime` bigint(20) DEFAULT '0',
  `vid` bigint(20) DEFAULT '0',
  `spusers` bigint(20) DEFAULT '0',
  `lid` bigint(20) DEFAULT '0',
  `eid` bigint(20) DEFAULT '0',
  `pid` bigint(20) DEFAULT '0',
  `updatetime` bigint(20) DEFAULT '0',
  `dustbin` tinyint(1) DEFAULT '0',
  `bugtitle` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bugs`
--

LOCK TABLES `bugs` WRITE;
/*!40000 ALTER TABLE `bugs` DISABLE KEYS */;
INSERT INTO `bugs` VALUES (1,1,NULL,6,'&lt;p&gt;asdfasfasdf&lt;/p&gt;',0,1,1575557996,2,1,2,1,1,0,0,'asdfasdf');
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
INSERT INTO `defaultvalue` VALUES (6,1,2);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `environment`
--

LOCK TABLES `environment` WRITE;
/*!40000 ALTER TABLE `environment` DISABLE KEYS */;
INSERT INTO `environment` VALUES (1,'1212');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `header`
--

LOCK TABLES `header` WRITE;
/*!40000 ALTER TABLE `header` DISABLE KEYS */;
INSERT INTO `header` VALUES (1,'itflowheader','1','token加密'),(2,'','','');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `headerlist`
--

LOCK TABLES `headerlist` WRITE;
/*!40000 ALTER TABLE `headerlist` DISABLE KEYS */;
INSERT INTO `headerlist` VALUES (1,'X-Token','');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `importants`
--

LOCK TABLES `importants` WRITE;
/*!40000 ALTER TABLE `importants` DISABLE KEYS */;
INSERT INTO `importants` VALUES (1,'一般');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` VALUES (1,'python开发',2,'0'),(2,'开发经理',1,'0');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `level`
--

LOCK TABLES `level` WRITE;
/*!40000 ALTER TABLE `level` DISABLE KEYS */;
INSERT INTO `level` VALUES (1,'1'),(2,'2');
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
) ENGINE=InnoDB AUTO_INCREMENT=601 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (1,1557134129,'login','nickname: admin has login','45.124.112.138'),(2,1557134832,'header','addheader:operator:admin, id: 1, name: itflowheader','45.124.112.138'),(3,1557134861,'api','addapi:operator:admin, id: 2, name: 获取用户信息','45.124.112.138'),(4,1557135024,'api','addapi:operator:admin, id: 3, name: 获取用户总个数','45.124.112.138'),(5,1557135281,'login','nickname: admin has login','183.11.128.112'),(6,1557135303,'login','nickname: admin has login','58.57.156.161'),(7,1557135573,'login','nickname: admin has login','45.124.112.138'),(8,1557135712,'api','addapi:operator:admin, id: 4, name: 获取所有bug','116.25.146.205'),(9,1557136134,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(10,1557136249,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(11,1557136437,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(12,1557136468,'login','nickname: admin has login','223.104.63.18'),(13,1557136469,'login','nickname: admin has login','223.104.63.18'),(14,1557136787,'login','nickname: admin has login','111.207.194.227'),(15,1557136957,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(16,1557137176,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(17,1557137401,'api','updateapi: operator:admin, id: 4, change to 获取所有bug','116.25.146.205'),(18,1557138842,'login','nickname: admin has login','171.210.152.39'),(19,1557143155,'login','nickname: admin has login','116.25.146.205'),(20,1557143157,'login','nickname: admin has login','116.25.146.205'),(21,1557143176,'login','nickname: admin has login','116.25.146.205'),(22,1557143183,'login','nickname: admin has login','116.25.146.205'),(23,1557143205,'login','nickname: admin has login','116.25.146.205'),(24,1557143895,'login','nickname: admin has login','112.231.178.187'),(25,1557144273,'login','nickname: admin has login','223.99.166.194'),(26,1557144274,'login','nickname: admin has login','223.99.166.194'),(27,1557144275,'login','nickname: admin has login','223.99.166.194'),(28,1557144275,'login','nickname: admin has login','223.99.166.194'),(29,1557145199,'login','nickname: admin has login','223.104.246.136'),(30,1557147449,'login','nickname: admin has login','114.221.22.208'),(31,1557160663,'login','nickname: admin has login','113.97.34.21'),(32,1557190254,'login','nickname: admin has login','119.137.55.133'),(33,1557190910,'login','nickname: admin has login','223.104.215.81'),(34,1557281393,'login','nickname: admin has login','112.25.234.118'),(35,1557282212,'login','nickname: admin has login','1.58.197.55'),(36,1557300451,'login','nickname: admin has login','171.83.99.177'),(37,1557302337,'login','nickname: admin has login','101.206.169.57'),(38,1557302355,'login','nickname: admin has login','101.206.169.57'),(39,1557302408,'login','nickname: admin has login','60.176.110.57'),(40,1557306026,'login','nickname: admin has login','116.21.133.9'),(41,1557347155,'login','nickname: admin has login','117.136.40.227'),(42,1557372653,'login','nickname: admin has login','113.109.198.218'),(43,1557372662,'login','nickname: admin has login','113.109.198.218'),(44,1557372666,'login','nickname: admin has login','113.109.198.218'),(45,1557372672,'login','nickname: admin has login','113.109.198.218'),(46,1557372672,'login','nickname: admin has login','113.109.198.218'),(47,1557372673,'login','nickname: admin has login','113.109.198.218'),(48,1557372674,'login','nickname: admin has login','113.109.198.218'),(49,1557372674,'login','nickname: admin has login','113.109.198.218'),(50,1557372674,'login','nickname: admin has login','113.109.198.218'),(51,1557372674,'login','nickname: admin has login','113.109.198.218'),(52,1557372674,'login','nickname: admin has login','113.109.198.218'),(53,1557372686,'login','nickname: admin has login','113.109.198.218'),(54,1557378882,'login','nickname: admin has login','171.216.222.20'),(55,1557384762,'login','nickname: admin has login','113.109.198.218'),(56,1557388325,'login','nickname: admin has login','223.104.63.109'),(57,1557388589,'login','nickname: admin has login','106.6.189.131'),(58,1557450605,'login','nickname: admin has login','61.150.43.35'),(59,1557450916,'login','nickname: admin has login','119.4.252.222'),(60,1557471066,'login','nickname: admin has login','117.136.11.14'),(61,1557473956,'login','nickname: admin has login','112.91.82.50'),(62,1557474477,'login','nickname: admin has login','117.80.100.207'),(63,1557474583,'login','nickname: admin has login','220.202.152.94'),(64,1557474584,'login','nickname: admin has login','220.202.152.94'),(65,1557474584,'login','nickname: admin has login','220.202.152.94'),(66,1557474584,'login','nickname: admin has login','220.202.152.94'),(67,1557474584,'login','nickname: admin has login','220.202.152.94'),(68,1557474588,'login','nickname: admin has login','58.250.17.93'),(69,1557474791,'project','addproject:operator:admin, id: 1, name: 123','117.80.100.207'),(70,1557474814,'login','nickname: admin has login','113.91.149.215'),(71,1557474834,'login','nickname: admin has login','163.179.125.138'),(72,1557474979,'login','nickname: admin has login','58.243.250.48'),(73,1557475259,'login','nickname: admin has login','220.181.102.176'),(74,1557475446,'login','nickname: admin has login','180.175.177.64'),(75,1557475756,'login','nickname: admin has login','113.102.165.244'),(76,1557475802,'restproject','addrestproject:operator:admin, id: 2, name: ','113.102.165.244'),(77,1557475805,'restproject','deleterestproject: operator:admin, id: 2','113.102.165.244'),(78,1557476335,'login','nickname: admin has login','115.230.122.117'),(79,1557476689,'login','nickname: admin has login','61.183.117.38'),(80,1557477057,'login','nickname: admin has login','115.60.191.38'),(81,1557477103,'login','nickname: admin has login','1.119.40.66'),(82,1557477213,'login','nickname: admin has login','101.84.151.116'),(83,1557477262,'login','nickname: admin has login','112.65.61.64'),(84,1557477439,'login','nickname: admin has login','222.222.62.180'),(85,1557540913,'login','nickname: admin has login','114.221.18.94'),(86,1557557313,'login','nickname: admin has login','183.17.57.232'),(87,1557592360,'login','nickname: admin has login','117.136.0.198'),(88,1557656111,'login','nickname: admin has login','219.143.13.45'),(89,1557923493,'login','nickname: admin has login','45.124.112.138'),(90,1558001609,'login','nickname: admin has login','116.21.135.243'),(91,1558012714,'login','nickname: admin has login','113.97.33.146'),(92,1558075736,'login','nickname: admin has login','113.89.239.253'),(93,1558079669,'login','nickname: admin has login','58.243.254.80'),(94,1558082949,'login','nickname: admin has login','113.89.239.253'),(95,1558082990,'login','nickname: admin has login','113.89.239.253'),(96,1558083019,'login','nickname: admin has login','113.89.239.253'),(97,1558083019,'login','nickname: admin has login','113.89.239.253'),(98,1558083020,'login','nickname: admin has login','113.89.239.253'),(99,1558083023,'login','nickname: admin has login','113.89.239.253'),(100,1558083023,'login','nickname: admin has login','113.89.239.253'),(101,1558083023,'login','nickname: admin has login','113.89.239.253'),(102,1558083024,'login','nickname: admin has login','113.89.239.253'),(103,1558083024,'login','nickname: admin has login','113.89.239.253'),(104,1558083024,'login','nickname: admin has login','113.89.239.253'),(105,1558083024,'login','nickname: admin has login','113.89.239.253'),(106,1558083034,'login','nickname: admin has login','113.89.239.253'),(107,1558083034,'login','nickname: admin has login','113.89.239.253'),(108,1558083034,'login','nickname: admin has login','113.89.239.253'),(109,1558083036,'login','nickname: admin has login','113.89.239.253'),(110,1558083036,'login','nickname: admin has login','113.89.239.253'),(111,1558083036,'login','nickname: admin has login','113.89.239.253'),(112,1558083037,'login','nickname: admin has login','113.89.239.253'),(113,1558083257,'login','nickname: admin has login','113.89.239.253'),(114,1558083293,'login','nickname: admin has login','113.89.239.253'),(115,1558083293,'login','nickname: admin has login','113.89.239.253'),(116,1558083293,'login','nickname: admin has login','113.89.239.253'),(117,1558083296,'login','nickname: admin has login','113.89.239.253'),(118,1558083297,'login','nickname: admin has login','113.89.239.253'),(119,1558083297,'login','nickname: admin has login','113.89.239.253'),(120,1558083323,'login','nickname: admin has login','113.89.239.253'),(121,1558083336,'login','nickname: admin has login','113.89.239.253'),(122,1558083349,'login','nickname: admin has login','113.89.239.253'),(123,1558083349,'login','nickname: admin has login','113.89.239.253'),(124,1558083350,'login','nickname: admin has login','113.89.239.253'),(125,1558083350,'login','nickname: admin has login','113.89.239.253'),(126,1558083350,'login','nickname: admin has login','113.89.239.253'),(127,1558083375,'login','nickname: admin has login','113.89.239.253'),(128,1558083511,'login','nickname: admin has login','113.89.239.253'),(129,1558083514,'login','nickname: admin has login','113.89.239.253'),(130,1558085426,'login','nickname: admin has login','113.89.239.253'),(131,1558085462,'login','nickname: admin has login','113.89.239.253'),(132,1558085520,'login','nickname: admin has login','113.89.239.253'),(133,1558149391,'login','nickname: admin has login','113.89.239.253'),(134,1558149420,'login','nickname: admin has login','113.89.239.253'),(135,1558149443,'login','nickname: admin has login','113.89.239.253'),(136,1558156123,'login','nickname: admin has login','112.24.207.201'),(137,1558156245,'login','nickname: admin has login','139.227.11.238'),(138,1558156364,'login','nickname: admin has login','58.62.93.253'),(139,1558157415,'login','nickname: admin has login','114.249.0.29'),(140,1558159448,'login','nickname: admin has login','113.89.239.253'),(141,1558160537,'login','nickname: admin has login','115.239.237.166'),(142,1558161003,'login','nickname: admin has login','222.172.244.213'),(143,1558163634,'login','nickname: admin has login','116.25.146.247'),(144,1558163660,'login','nickname: admin has login','116.25.146.247'),(145,1558163668,'login','nickname: admin has login','116.25.146.247'),(146,1558163774,'login','nickname: admin has login','116.25.146.247'),(147,1558163809,'login','nickname: admin has login','116.25.146.247'),(148,1558164014,'login','nickname: admin has login','116.25.146.247'),(149,1558164017,'login','nickname: admin has login','116.25.146.247'),(150,1558164232,'login','nickname: admin has login','116.25.146.247'),(151,1558164248,'login','nickname: admin has login','116.25.146.247'),(152,1558164252,'login','nickname: admin has login','116.25.146.247'),(153,1558164290,'login','nickname: admin has login','116.25.146.247'),(154,1558165675,'login','nickname: admin has login','116.25.146.247'),(155,1558165678,'login','nickname: admin has login','116.25.146.247'),(156,1558165807,'login','nickname: admin has login','116.25.146.247'),(157,1558165808,'login','nickname: admin has login','116.25.146.247'),(158,1558165808,'login','nickname: admin has login','116.25.146.247'),(159,1558165808,'login','nickname: admin has login','116.25.146.247'),(160,1558165809,'login','nickname: admin has login','116.25.146.247'),(161,1558165809,'login','nickname: admin has login','116.25.146.247'),(162,1558165814,'login','nickname: admin has login','116.25.146.247'),(163,1558165863,'login','nickname: admin has login','116.25.146.247'),(164,1558165864,'login','nickname: admin has login','116.25.146.247'),(165,1558165865,'login','nickname: admin has login','116.25.146.247'),(166,1558165957,'login','nickname: admin has login','116.25.146.247'),(167,1558165965,'login','nickname: admin has login','116.25.146.247'),(168,1558165977,'login','nickname: admin has login','116.25.146.247'),(169,1558165978,'login','nickname: admin has login','116.25.146.247'),(170,1558165979,'login','nickname: admin has login','116.25.146.247'),(171,1558165980,'login','nickname: admin has login','116.25.146.247'),(172,1558166000,'login','nickname: admin has login','116.25.146.247'),(173,1558166016,'login','nickname: admin has login','116.25.146.247'),(174,1558166039,'login','nickname: admin has login','116.25.146.247'),(175,1558166283,'login','nickname: admin has login','116.25.146.247'),(176,1558166283,'login','nickname: admin has login','116.25.146.247'),(177,1558166283,'login','nickname: admin has login','116.25.146.247'),(178,1558166283,'login','nickname: admin has login','116.25.146.247'),(179,1558166283,'login','nickname: admin has login','116.25.146.247'),(180,1558166283,'login','nickname: admin has login','116.25.146.247'),(181,1558166283,'login','nickname: admin has login','116.25.146.247'),(182,1558166283,'login','nickname: admin has login','116.25.146.247'),(183,1558166283,'login','nickname: admin has login','116.25.146.247'),(184,1558166283,'login','nickname: admin has login','116.25.146.247'),(185,1558166283,'login','nickname: admin has login','116.25.146.247'),(186,1558166283,'login','nickname: admin has login','116.25.146.247'),(187,1558166283,'login','nickname: admin has login','116.25.146.247'),(188,1558166283,'login','nickname: admin has login','116.25.146.247'),(189,1558166283,'login','nickname: admin has login','116.25.146.247'),(190,1558166283,'login','nickname: admin has login','116.25.146.247'),(191,1558166283,'login','nickname: admin has login','116.25.146.247'),(192,1558166283,'login','nickname: admin has login','116.25.146.247'),(193,1558166283,'login','nickname: admin has login','116.25.146.247'),(194,1558166283,'login','nickname: admin has login','116.25.146.247'),(195,1558166283,'login','nickname: admin has login','116.25.146.247'),(196,1558166283,'login','nickname: admin has login','116.25.146.247'),(197,1558166283,'login','nickname: admin has login','116.25.146.247'),(198,1558166283,'login','nickname: admin has login','116.25.146.247'),(199,1558166283,'login','nickname: admin has login','116.25.146.247'),(200,1558166283,'login','nickname: admin has login','116.25.146.247'),(201,1558166283,'login','nickname: admin has login','116.25.146.247'),(202,1558166283,'login','nickname: admin has login','116.25.146.247'),(203,1558166283,'login','nickname: admin has login','116.25.146.247'),(204,1558166283,'login','nickname: admin has login','116.25.146.247'),(205,1558166283,'login','nickname: admin has login','116.25.146.247'),(206,1558166283,'login','nickname: admin has login','116.25.146.247'),(207,1558166283,'login','nickname: admin has login','116.25.146.247'),(208,1558166283,'login','nickname: admin has login','116.25.146.247'),(209,1558166283,'login','nickname: admin has login','116.25.146.247'),(210,1558166283,'login','nickname: admin has login','116.25.146.247'),(211,1558166283,'login','nickname: admin has login','116.25.146.247'),(212,1558166283,'login','nickname: admin has login','116.25.146.247'),(213,1558166283,'login','nickname: admin has login','116.25.146.247'),(214,1558166283,'login','nickname: admin has login','116.25.146.247'),(215,1558166283,'login','nickname: admin has login','116.25.146.247'),(216,1558166283,'login','nickname: admin has login','116.25.146.247'),(217,1558166283,'login','nickname: admin has login','116.25.146.247'),(218,1558166283,'login','nickname: admin has login','116.25.146.247'),(219,1558166283,'login','nickname: admin has login','116.25.146.247'),(220,1558166283,'login','nickname: admin has login','116.25.146.247'),(221,1558166283,'login','nickname: admin has login','116.25.146.247'),(222,1558166283,'login','nickname: admin has login','116.25.146.247'),(223,1558166283,'login','nickname: admin has login','116.25.146.247'),(224,1558166284,'login','nickname: admin has login','116.25.146.247'),(225,1558166284,'login','nickname: admin has login','116.25.146.247'),(226,1558166284,'login','nickname: admin has login','116.25.146.247'),(227,1558166284,'login','nickname: admin has login','116.25.146.247'),(228,1558166304,'login','nickname: admin has login','116.25.146.247'),(229,1558166307,'login','nickname: admin has login','116.25.146.247'),(230,1558166438,'login','nickname: admin has login','116.25.146.247'),(231,1558166438,'login','nickname: admin has login','116.25.146.247'),(232,1558166533,'login','nickname: admin has login','116.25.146.247'),(233,1558166533,'login','nickname: admin has login','116.25.146.247'),(234,1558166562,'login','nickname: admin has login','116.25.146.247'),(235,1558166562,'login','nickname: admin has login','116.25.146.247'),(236,1558166573,'login','nickname: admin has login','116.25.146.247'),(237,1558166583,'login','nickname: admin has login','116.25.146.247'),(238,1558166609,'login','nickname: admin has login','116.25.146.247'),(239,1558166665,'login','nickname: admin has login','116.25.146.247'),(240,1558166676,'login','nickname: admin has login','116.25.146.247'),(241,1558166678,'login','nickname: admin has login','116.25.146.247'),(242,1558167032,'login','nickname: admin has login','171.212.125.223'),(243,1558167101,'login','nickname: admin has login','116.25.146.247'),(244,1558167126,'login','nickname: admin has login','116.25.146.247'),(245,1558167440,'login','nickname: admin has login','116.25.146.247'),(246,1558167667,'login','nickname: admin has login','116.25.146.247'),(247,1558167837,'login','nickname: admin has login','183.92.251.197'),(248,1558168456,'login','nickname: admin has login','116.25.146.247'),(249,1558168470,'login','nickname: admin has login','116.25.146.247'),(250,1558168512,'login','nickname: admin has login','116.25.146.247'),(251,1558168545,'login','nickname: admin has login','116.25.146.247'),(252,1558168571,'login','nickname: admin has login','116.25.146.247'),(253,1558168631,'login','nickname: admin has login','116.25.146.247'),(254,1558168646,'login','nickname: admin has login','116.25.146.247'),(255,1558173059,'login','nickname: admin has login','117.88.47.182'),(256,1558178117,'login','nickname: admin has login','117.143.160.198'),(257,1558197677,'login','nickname: admin has login','112.96.192.144'),(258,1558225016,'login','nickname: admin has login','223.73.60.7'),(259,1558227034,'login','nickname: admin has login','183.205.17.66'),(260,1558232973,'login','nickname: admin has login','116.21.65.92'),(261,1558232988,'project','updateproject: operator:admin, id: 1, change to 123','116.21.65.92'),(262,1558252477,'login','nickname: admin has login','113.97.35.9'),(263,1558275532,'login','nickname: admin has login','223.88.189.120'),(264,1558321901,'login','nickname: admin has login','116.25.146.247'),(265,1558323525,'login','nickname: admin has login','39.186.194.45'),(266,1558328900,'login','nickname: admin has login','116.25.146.247'),(267,1558328906,'login','nickname: admin has login','116.25.146.247'),(268,1558328914,'login','nickname: admin has login','116.25.146.247'),(269,1558332479,'login','nickname: admin has login','116.25.146.247'),(270,1558334417,'login','nickname: admin has login','113.87.182.23'),(271,1558334436,'login','nickname: admin has login','113.87.182.23'),(272,1558334894,'login','nickname: admin has login','113.87.182.23'),(273,1558336630,'login','nickname: admin has login','113.87.182.23'),(274,1558343294,'login','nickname: admin has login','113.87.182.23'),(275,1558343333,'login','nickname: admin has login','113.87.182.23'),(276,1558344920,'login','nickname: admin has login','113.87.182.23'),(277,1558345045,'login','nickname: admin has login','113.87.182.23'),(278,1558345096,'login','nickname: admin has login','113.87.182.23'),(279,1558345114,'login','nickname: admin has login','113.87.182.23'),(280,1558345134,'login','nickname: admin has login','113.87.182.23'),(281,1558345230,'login','nickname: admin has login','113.87.182.23'),(282,1558345279,'login','nickname: admin has login','113.87.182.23'),(283,1558345337,'login','nickname: admin has login','113.87.182.23'),(284,1558364092,'login','nickname: admin has login','221.239.255.151'),(285,1558402104,'login','nickname: admin has login','120.28.22.186'),(286,1558422448,'login','nickname: admin has login','113.87.182.23'),(287,1558426020,'login','nickname: admin has login','113.87.182.23'),(288,1558426032,'login','nickname: admin has login','113.87.182.23'),(289,1558426059,'login','nickname: admin has login','113.87.182.23'),(290,1558426099,'login','nickname: admin has login','113.87.182.23'),(291,1558428100,'login','nickname: admin has login','113.87.182.23'),(292,1558430625,'login','nickname: admin has login','113.87.182.23'),(293,1558430742,'login','nickname: admin has login','113.87.182.23'),(294,1558435630,'login','nickname: admin has login','127.0.0.1'),(295,1558435697,'login','nickname: admin has login','127.0.0.1'),(296,1558435746,'login','nickname: admin has login','127.0.0.1'),(297,1558435922,'login','nickname: admin has login','192.168.1.33'),(298,1558435991,'login','nickname: admin has login','127.0.0.1'),(299,1558436067,'login','nickname: admin has login','192.168.1.126'),(300,1558436112,'login','nickname: admin has login','192.168.1.126'),(301,1558436191,'login','nickname: admin has login','192.168.1.126'),(302,1558436266,'login','nickname: admin has login','192.168.1.126'),(303,1558436341,'login','nickname: admin has login','192.168.1.126'),(304,1558436517,'login','nickname: admin has login','192.168.1.126'),(305,1558436615,'login','nickname: admin has login','192.168.1.126'),(306,1558436627,'login','nickname: admin has login','192.168.1.126'),(307,1558436647,'login','nickname: admin has login','192.168.1.126'),(308,1558489797,'login','nickname: admin has login','124.160.90.30'),(309,1558494502,'login','nickname: admin has login','192.168.1.126'),(310,1558506553,'login','nickname: admin has login','192.168.1.126'),(311,1558508005,'login','nickname: admin has login','192.168.1.126'),(312,1558509431,'login','nickname: admin has login','192.168.1.126'),(313,1558511267,'login','nickname: admin has login','192.168.1.126'),(314,1558511393,'login','nickname: admin has login','192.168.1.126'),(315,1558511835,'login','nickname: admin has login','192.168.1.126'),(316,1558512678,'login','nickname: admin has login','113.87.180.87'),(317,1558512681,'login','nickname: admin has login','113.87.180.87'),(318,1558512682,'login','nickname: admin has login','113.87.180.87'),(319,1558512682,'login','nickname: admin has login','113.87.180.87'),(320,1558512685,'login','nickname: admin has login','113.87.180.87'),(321,1558519193,'login','nickname: admin has login','221.192.178.187'),(322,1558524560,'login','nickname: admin has login','192.168.1.126'),(323,1558527089,'login','nickname: admin has login','222.93.167.100'),(324,1558579494,'login','nickname: admin has login','192.168.1.126'),(325,1558579979,'login','nickname: admin has login','113.87.180.87'),(326,1558594091,'login','nickname: admin has login','192.168.1.126'),(327,1558594416,'login','nickname: admin has login','113.87.180.87'),(328,1558594419,'login','nickname: admin has login','113.87.180.87'),(329,1558594419,'login','nickname: admin has login','113.87.180.87'),(330,1558594419,'login','nickname: admin has login','113.87.180.87'),(331,1558594419,'login','nickname: admin has login','113.87.180.87'),(332,1558594420,'login','nickname: admin has login','113.87.180.87'),(333,1558594420,'login','nickname: admin has login','113.87.180.87'),(334,1558594420,'login','nickname: admin has login','113.87.180.87'),(335,1558594420,'login','nickname: admin has login','113.87.180.87'),(336,1558594421,'login','nickname: admin has login','113.87.180.87'),(337,1558594426,'login','nickname: admin has login','113.87.180.87'),(338,1558600749,'login','nickname: admin has login','192.168.1.126'),(339,1558603495,'login','nickname: admin has login','113.87.180.87'),(340,1558608541,'login','nickname: admin has login','192.168.1.126'),(341,1558670499,'login','nickname: admin has login','113.87.180.87'),(342,1558670529,'login','nickname: admin has login','113.87.180.87'),(343,1558677338,'login','nickname: admin has login','192.168.1.126'),(344,1558679431,'login','nickname: admin has login','113.87.15.166'),(345,1558679433,'login','nickname: admin has login','113.87.15.166'),(346,1558679433,'login','nickname: admin has login','113.87.15.166'),(347,1558679434,'login','nickname: admin has login','113.87.15.166'),(348,1558679434,'login','nickname: admin has login','113.87.15.166'),(349,1558679434,'login','nickname: admin has login','113.87.15.166'),(350,1558679434,'login','nickname: admin has login','113.87.15.166'),(351,1558679437,'login','nickname: admin has login','113.87.15.166'),(352,1558685474,'login','nickname: admin has login','192.168.1.126'),(353,1558686008,'login','nickname: admin has login','127.0.0.1'),(354,1558688383,'login','nickname: admin has login','113.87.15.166'),(355,1558927545,'login','nickname: admin has login','192.168.1.126'),(356,1558940977,'login','nickname: admin has login','113.89.236.148'),(357,1558940998,'status','addstatus:operator:admin, id: 1, name: test','113.89.236.148'),(358,1558941003,'status','addstatus:operator:admin, id: 2, name: biao+','113.89.236.148'),(359,1558941031,'status','addstatus:operator:admin, id: 3, name: ToDoList','113.89.236.148'),(360,1558943153,'login','nickname: admin has login','192.168.1.126'),(361,1558950403,'login','nickname: admin has login','192.168.1.126'),(362,1559007702,'login','nickname: admin has login','113.215.181.238'),(363,1559007756,'version','addversion:operator:admin, id: 1, name: ','113.215.181.238'),(364,1559109903,'login','nickname: admin has login','192.168.1.126'),(365,1559110561,'login','nickname: admin has login','113.87.181.175'),(366,1559114292,'status','addstatus:operator:admin, id: 4, name: 嘻嘻嘻','113.87.181.175'),(367,1559181464,'login','nickname: admin has login','192.168.1.126'),(368,1559181465,'login','nickname: admin has login','192.168.1.126'),(369,1559200518,'login','nickname: admin has login','192.168.1.126'),(370,1559203277,'status','addstatus:operator:admin, id: 5, name: react','192.168.1.126'),(371,1559203331,'login','nickname: admin has login','113.87.12.5'),(372,1559283222,'login','nickname: admin has login','192.168.1.126'),(373,1559283222,'login','nickname: admin has login','192.168.1.126'),(374,1559283222,'login','nickname: admin has login','192.168.1.126'),(375,1559283222,'login','nickname: admin has login','192.168.1.126'),(376,1559283222,'login','nickname: admin has login','192.168.1.126'),(377,1559283247,'login','nickname: admin has login','192.168.1.126'),(378,1559283433,'status','addstatus:operator:admin, id: 6, name: 测试','192.168.1.126'),(379,1559285890,'login','nickname: admin has login','113.87.12.5'),(380,1559286715,'status','updatestatus: operator:admin, id: 4, change to 嘻嘻嘻1','113.87.12.5'),(381,1559287145,'status','updatestatus: operator:admin, id: 4, change to 哈哈哈','192.168.1.126'),(382,1559287218,'status','deletestatus: operator:admin, id: 4','113.87.12.5'),(383,1559288584,'status','deletestatus: operator:admin, id: 4','192.168.1.126'),(384,1559288777,'status','addstatus:operator:admin, id: 7, name: 等待删除','192.168.1.126'),(385,1559288781,'status','deletestatus: operator:admin, id: 7','192.168.1.126'),(386,1559355297,'login','nickname: admin has login','192.168.1.126'),(387,1559371305,'login','nickname: admin has login','1.198.215.199'),(388,1559372489,'login','nickname: admin has login','192.168.1.126'),(389,1559372548,'login','nickname: admin has login','113.89.238.81'),(390,1559372550,'login','nickname: admin has login','113.89.238.81'),(391,1559372550,'login','nickname: admin has login','113.89.238.81'),(392,1559372550,'login','nickname: admin has login','113.89.238.81'),(393,1559372551,'login','nickname: admin has login','113.89.238.81'),(394,1559372551,'login','nickname: admin has login','113.89.238.81'),(395,1559372552,'login','nickname: admin has login','113.89.238.81'),(396,1559372624,'statusgroup','addstatusgroup:operator:admin, id: 1, name: Saturday','113.89.238.81'),(397,1559444331,'login','nickname: admin has login','118.247.112.197'),(398,1559532523,'login','nickname: admin has login','113.87.15.36'),(399,1559546154,'login','nickname: admin has login','113.87.15.36'),(400,1559613229,'login','nickname: admin has login','113.118.186.89'),(401,1560154124,'login','nickname: admin has login','101.64.140.84'),(402,1560327391,'login','nickname: admin has login','58.215.108.75'),(403,1560352639,'login','nickname: admin has login','58.101.88.139'),(404,1560424818,'login','nickname: admin has login','116.231.95.170'),(405,1560429695,'login','nickname: admin has login','113.87.183.87'),(406,1560429724,'level','addlevel:operator:admin, id: 1, name: 1','113.87.183.87'),(407,1560429728,'level','addlevel:operator:admin, id: 2, name: 2','113.87.183.87'),(408,1560862127,'login','nickname: admin has login','113.97.34.15'),(409,1561010706,'login','nickname: admin has login','120.198.47.70'),(410,1561032561,'login','nickname: admin has login','183.159.193.207'),(411,1561032597,'status','deletestatus: operator:admin, id: 1','183.159.193.207'),(412,1561032834,'header','addheader:operator:admin, id: 2, name: ','183.159.193.207'),(413,1561072087,'login','nickname: admin has login','171.43.249.118'),(414,1561191605,'login','nickname: admin has login','113.97.35.156'),(415,1561218262,'login','nickname: admin has login','127.0.0.1'),(416,1561222687,'login','nickname: admin has login','127.0.0.1'),(417,1561223032,'login','nickname: admin has login','127.0.0.1'),(418,1561224272,'login','nickname: admin has login','127.0.0.1'),(419,1561272506,'login','nickname: admin has login','113.97.35.156'),(420,1561286868,'login','nickname: admin has login','113.97.35.156'),(421,1561287706,'login','nickname: admin has login','172.17.0.1'),(422,1561352066,'login','nickname: admin has login','116.25.146.203'),(423,1561446749,'login','nickname: admin has login','113.87.183.128'),(424,1561446792,'version','addversion:operator:admin, id: 2, name: v1.2','113.87.183.128'),(425,1561446797,'version','deleteversion: operator:admin, id: 1','113.87.183.128'),(426,1561448000,'login','nickname: admin has login','211.171.245.213'),(427,1561448006,'login','nickname: admin has login','211.171.245.213'),(428,1561972777,'login','nickname: admin has login','119.39.130.171'),(429,1562147855,'login','nickname: admin has login','120.236.160.152'),(430,1562206080,'login','nickname: admin has login','112.208.248.104'),(431,1562206104,'login','nickname: admin has login','112.208.248.104'),(432,1562206106,'login','nickname: admin has login','112.208.248.104'),(433,1562206234,'login','nickname: admin has login','112.208.248.104'),(434,1562206289,'login','nickname: admin has login','120.236.160.152'),(435,1562227329,'login','nickname: admin has login','120.236.160.152'),(436,1562227331,'login','nickname: admin has login','120.236.160.152'),(437,1562227332,'login','nickname: admin has login','120.236.160.152'),(438,1562227332,'login','nickname: admin has login','120.236.160.152'),(439,1562227332,'login','nickname: admin has login','120.236.160.152'),(440,1562227332,'login','nickname: admin has login','120.236.160.152'),(441,1562227333,'login','nickname: admin has login','120.236.160.152'),(442,1562227333,'login','nickname: admin has login','120.236.160.152'),(443,1562227345,'login','nickname: admin has login','120.236.160.152'),(444,1562350781,'login','nickname: admin has login','121.127.226.158'),(445,1562350782,'login','nickname: admin has login','121.127.226.158'),(446,1562807269,'login','nickname: admin has login','123.160.225.189'),(447,1562979802,'login','nickname: admin has login','113.97.33.1'),(448,1563244365,'login','nickname: admin has login','180.136.144.155'),(449,1563244388,'rolegroup','addrolegroup:operator:admin, id: 1, name: cess','180.136.144.155'),(450,1563244402,'rolegroup','addrolegroup:operator:admin, id: 2, name: cewww','180.136.144.155'),(451,1563244418,'usergroup','addusergroup:operator:admin, id: 1, name: cew','180.136.144.155'),(452,1563244539,'rolegroup','updaterolegroup: operator:admin, id: 1, change to cess','180.136.144.155'),(453,1563244548,'rolegroup','updaterolegroup: operator:admin, id: 1, change to cess','180.136.144.155'),(454,1563245499,'rolegroup','deleterolegroup: operator:admin, id: 2','180.136.144.155'),(455,1563255763,'login','nickname: admin has login','180.136.144.155'),(456,1563255769,'login','nickname: admin has login','180.136.144.155'),(457,1563255779,'login','nickname: admin has login','180.136.144.155'),(458,1563412383,'login','nickname: admin has login','144.48.211.43'),(459,1563418349,'login','nickname: admin has login','111.201.60.143'),(460,1563435934,'login','nickname: admin has login','180.168.109.226'),(461,1563458115,'login','nickname: admin has login','111.162.116.76'),(462,1563858835,'login','nickname: admin has login','49.87.60.159'),(463,1564069951,'login','nickname: admin has login','110.80.164.69'),(464,1564122938,'login','nickname: admin has login','112.17.240.122'),(465,1564392072,'login','nickname: admin has login','210.108.146.8'),(466,1564501221,'login','nickname: admin has login','180.102.145.23'),(467,1564539739,'login','nickname: admin has login','43.226.16.199'),(468,1564627519,'login','nickname: admin has login','61.164.47.201'),(469,1564639793,'login','nickname: admin has login','117.186.225.182'),(470,1564749061,'login','nickname: admin has login','117.30.209.109'),(471,1564759962,'login','nickname: admin has login','117.30.209.109'),(472,1564800797,'login','nickname: admin has login','117.31.212.36'),(473,1565224611,'login','nickname: admin has login','219.159.73.73'),(474,1565234888,'login','nickname: admin has login','111.193.120.253'),(475,1565252455,'login','nickname: admin has login','171.120.89.194'),(476,1565338265,'login','nickname: admin has login','114.251.146.133'),(477,1565570215,'login','nickname: admin has login','180.109.182.88'),(478,1565939482,'login','nickname: admin has login','182.150.40.59'),(479,1566186037,'login','nickname: admin has login','112.224.70.235'),(480,1566265556,'login','nickname: admin has login','110.185.174.60'),(481,1566271857,'login','nickname: admin has login','120.202.28.53'),(482,1567059780,'login','nickname: admin has login','175.11.208.63'),(483,1567586344,'login','nickname: admin has login','119.61.21.147'),(484,1568031583,'login','nickname: admin has login','59.52.215.38'),(485,1568096960,'login','nickname: admin has login','111.121.117.61'),(486,1568104747,'login','nickname: admin has login','124.64.17.104'),(487,1568104898,'login','nickname: admin has login','36.7.84.10'),(488,1568104985,'login','nickname: admin has login','36.7.84.10'),(489,1568105168,'login','nickname: admin has login','124.64.17.104'),(490,1568249400,'login','nickname: admin has login','202.96.98.106'),(491,1568700928,'login','nickname: admin has login','124.160.89.162'),(492,1568785734,'login','nickname: admin has login','113.235.117.214'),(493,1568785870,'important','addimportant:operator:admin, id: 1, name: 一般','113.235.117.214'),(494,1568785930,'env','addenv:operator:admin, id: 1, name: 1212','113.235.117.214'),(495,1568879583,'login','nickname: admin has login','14.131.250.25'),(496,1569120686,'login','nickname: admin has login','120.234.130.205'),(497,1569215372,'login','nickname: admin has login','113.97.33.88'),(498,1569224352,'login','nickname: admin has login','113.97.33.88'),(499,1569240758,'login','nickname: admin has login','113.97.33.88'),(500,1569240760,'login','nickname: admin has login','113.97.33.88'),(501,1569240764,'login','nickname: admin has login','113.97.33.88'),(502,1569240776,'login','nickname: admin has login','113.97.33.88'),(503,1569240788,'login','nickname: admin has login','113.97.33.88'),(504,1569240793,'login','nickname: admin has login','113.97.33.88'),(505,1569254057,'login','nickname: admin has login','113.97.33.88'),(506,1569310143,'login','nickname: admin has login','112.30.41.87'),(507,1569376102,'login','nickname: admin has login','115.204.143.21'),(508,1569376104,'login','nickname: admin has login','115.204.143.21'),(509,1569380325,'login','nickname: admin has login','210.42.32.4'),(510,1569402670,'login','nickname: admin has login','113.97.35.21'),(511,1569496098,'login','nickname: admin has login','223.11.9.243'),(512,1569508196,'login','nickname: admin has login','113.97.35.21'),(513,1569721729,'login','nickname: admin has login','112.112.102.51'),(514,1569825384,'login','nickname: admin has login','183.12.239.252'),(515,1570067343,'login','nickname: admin has login','171.112.1.81'),(516,1570602353,'login','nickname: admin has login','115.175.68.31'),(517,1570602354,'login','nickname: admin has login','115.175.68.31'),(518,1570773678,'login','nickname: admin has login','58.246.229.238'),(519,1570892119,'login','nickname: admin has login','218.255.253.21'),(520,1571030196,'login','nickname: admin has login','124.93.196.25'),(521,1571030197,'login','nickname: admin has login','124.93.196.25'),(522,1571030198,'login','nickname: admin has login','124.93.196.25'),(523,1571030253,'login','nickname: admin has login','223.100.134.46'),(524,1571104999,'login','nickname: admin has login','119.61.21.147'),(525,1571119563,'login','nickname: admin has login','113.97.33.101'),(526,1571119695,'position','addposition:operator:admin, id: 1, name: python开发','113.97.33.101'),(527,1571119715,'position','addposition:operator:admin, id: 2, name: 开发经理','113.97.33.101'),(528,1571207900,'login','nickname: admin has login','182.50.124.85'),(529,1571239364,'login','nickname: admin has login','113.97.32.97'),(530,1571569868,'login','nickname: admin has login','113.97.33.58'),(531,1571591179,'login','nickname: admin has login','58.35.54.88'),(532,1571727345,'login','nickname: admin has login','127.0.0.1'),(533,1571727491,'login','nickname: admin has login','127.0.0.1'),(534,1571727496,'login','nickname: admin has login','127.0.0.1'),(535,1571727746,'login','nickname: admin has login','127.0.0.1'),(536,1571727917,'login','nickname: admin has login','127.0.0.1'),(537,1571728018,'login','nickname: admin has login','127.0.0.1'),(538,1571728022,'login','nickname: admin has login','127.0.0.1'),(539,1571728025,'login','nickname: admin has login','127.0.0.1'),(540,1571728028,'login','nickname: admin has login','127.0.0.1'),(541,1571728030,'login','nickname: admin has login','127.0.0.1'),(542,1571732442,'login','nickname: admin has login','127.0.0.1'),(543,1571732457,'login','nickname: admin has login','127.0.0.1'),(544,1571732633,'login','nickname: admin has login','127.0.0.1'),(545,1571732634,'login','nickname: admin has login','127.0.0.1'),(546,1571732639,'login','nickname: admin has login','127.0.0.1'),(547,1571732647,'login','nickname: admin has login','127.0.0.1'),(548,1571732652,'login','nickname: admin has login','127.0.0.1'),(549,1571732711,'login','nickname: admin has login','127.0.0.1'),(550,1571732792,'login','nickname: admin has login','127.0.0.1'),(551,1571732798,'login','nickname: admin has login','127.0.0.1'),(552,1571732803,'login','nickname: admin has login','127.0.0.1'),(553,1571732835,'login','nickname: admin has login','127.0.0.1'),(554,1571732836,'login','nickname: admin has login','127.0.0.1'),(555,1571732874,'login','nickname: admin has login','127.0.0.1'),(556,1571733175,'login','nickname: admin has login','127.0.0.1'),(557,1571733269,'login','nickname: admin has login','127.0.0.1'),(558,1571733293,'login','nickname: admin has login','127.0.0.1'),(559,1571733365,'login','nickname: admin has login','127.0.0.1'),(560,1571733381,'login','nickname: admin has login','127.0.0.1'),(561,1571733415,'login','nickname: admin has login','127.0.0.1'),(562,1571733490,'login','nickname: admin has login','127.0.0.1'),(563,1571734371,'login','nickname: admin has login','127.0.0.1'),(564,1571734400,'login','nickname: admin has login','127.0.0.1'),(565,1571734469,'login','nickname: admin has login','127.0.0.1'),(566,1571734495,'login','nickname: admin has login','127.0.0.1'),(567,1571734663,'login','nickname: admin has login','127.0.0.1'),(568,1571734693,'login','nickname: admin has login','127.0.0.1'),(569,1571734905,'login','nickname: admin has login','127.0.0.1'),(570,1571734923,'login','nickname: admin has login','127.0.0.1'),(571,1571735034,'login','nickname: admin has login','127.0.0.1'),(572,1571887003,'login','nickname: admin has login','114.244.130.117'),(573,1571930349,'login','nickname: admin has login','113.97.33.177'),(574,1571992563,'login','nickname: admin has login','183.28.31.197'),(575,1572234771,'login','nickname: admin has login','120.198.22.24'),(576,1572239933,'login','nickname: admin has login','218.17.157.251'),(577,1572671271,'login','nickname: admin has login','219.138.75.72'),(578,1572671502,'login','nickname: admin has login','219.138.75.72'),(579,1572875707,'login','nickname: admin has login','47.75.132.136'),(580,1572928192,'login','nickname: admin has login','61.241.197.27'),(581,1573016778,'login','nickname: admin has login','58.34.170.217'),(582,1573349550,'login','nickname: admin has login','45.120.156.156'),(583,1573363280,'login','nickname: admin has login','120.230.131.86'),(584,1573375680,'login','nickname: admin has login','110.184.161.132'),(585,1573473508,'login','nickname: admin has login','171.217.4.111'),(586,1574149728,'login','nickname: admin has login','172.19.0.1'),(587,1574149763,'project','addproject:operator:admin, id: 2, name: aaa','172.19.0.1'),(588,1575181484,'login','nickname: admin has login','127.0.0.1'),(589,1575181502,'login','nickname: admin has login','127.0.0.1'),(590,1575181583,'login','nickname: admin has login','127.0.0.1'),(591,1575181608,'login','nickname: admin has login','127.0.0.1'),(592,1575181669,'login','nickname: admin has login','127.0.0.1'),(593,1575181710,'login','nickname: admin has login','127.0.0.1'),(594,1575182073,'login','nickname: admin has login','127.0.0.1'),(595,1575271394,'login','nickname: admin has login','123.147.250.168'),(596,1575557910,'login','nickname: admin has login','127.0.0.1'),(597,1575557996,'bug','addbug:operator:admin, id: 1, name: asdfasdf','127.0.0.1'),(598,1575558387,'login','nickname: admin has login','127.0.0.1'),(599,1575597249,'login','nickname: admin has login','127.0.0.1'),(600,1575599126,'status','addstatus:operator:admin, id: 7, name: need','127.0.0.1');
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `options`
--

LOCK TABLES `options` WRITE;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
INSERT INTO `options` VALUES (1,'username','',2,'','必须'),(2,'password','',2,'','必须'),(3,'page','',10,'','必须'),(4,'limit','',10,'','必须'),(5,'level','',2,'','必须'),(6,'project','',2,'','必须'),(7,'title','',2,'','必须'),(8,'status','',14,'','必须'),(9,'page','',1,'','必须'),(10,'limit','',1,'','必须'),(11,'level','',2,'','可选'),(12,'project','',2,'','可选'),(13,'title','',2,'','可选'),(14,'status','',14,'','可选'),(15,'page','',1,'','必须'),(16,'limit','',1,'','必须'),(17,'page','',1,'','必须'),(18,'limit','',1,'','必须');
/*!40000 ALTER TABLE `options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projectname`
--

DROP TABLE IF EXISTS `projectname`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `projectname` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projectname`
--

LOCK TABLES `projectname` WRITE;
/*!40000 ALTER TABLE `projectname` DISABLE KEYS */;
INSERT INTO `projectname` VALUES (1,'123'),(2,'aaa');
/*!40000 ALTER TABLE `projectname` ENABLE KEYS */;
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
  `rolelist` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rolegroup`
--

LOCK TABLES `rolegroup` WRITE;
/*!40000 ALTER TABLE `rolegroup` DISABLE KEYS */;
INSERT INTO `rolegroup` VALUES (1,'cess','5,10,2,3,6,1');
/*!40000 ALTER TABLE `rolegroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role` (`role`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (3,'env'),(8,'important'),(9,'level'),(5,'log'),(10,'position'),(2,'project'),(7,'rolegroup'),(4,'status'),(6,'statusgroup'),(11,'usergroup'),(1,'version');
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sharefile`
--

LOCK TABLES `sharefile` WRITE;
/*!40000 ALTER TABLE `sharefile` DISABLE KEYS */;
INSERT INTO `sharefile` VALUES (13,'.',0,0,1,1,0,0,7877,1562350791,'bd_logo1.png'),(14,'.',0,0,1,1,0,0,13808248,1564069998,'AirDroid.exe'),(15,'.',1,1,0,1,1,1,0,1570773860,'7777');
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status`
--

LOCK TABLES `status` WRITE;
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
INSERT INTO `status` VALUES (2,'biao+'),(7,'need'),(5,'react'),(3,'ToDoList'),(6,'测试');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `statusgroup`
--

LOCK TABLES `statusgroup` WRITE;
/*!40000 ALTER TABLE `statusgroup` DISABLE KEYS */;
INSERT INTO `statusgroup` VALUES (1,'Saturday','3');
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `types`
--

LOCK TABLES `types` WRITE;
/*!40000 ALTER TABLE `types` DISABLE KEYS */;
INSERT INTO `types` VALUES (1,'int',0,'',0,'0'),(2,'string',0,'',0,''),(3,'bool',0,'',0,'false'),(4,'object',0,'',0,''),(5,'list',0,'',0,''),(6,'float64',0,'',0,'0');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','69ad5117e7553ecfa7f918a223426dd8da08a57f','admin@qq.com','http://120.26.164.125:10001/showimg/1569120737844973861.png',1557131883,0,'admin','',0,0,0,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usergroup`
--

LOCK TABLES `usergroup` WRITE;
/*!40000 ALTER TABLE `usergroup` DISABLE KEYS */;
INSERT INTO `usergroup` VALUES (1,'cew','1',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `version`
--

LOCK TABLES `version` WRITE;
/*!40000 ALTER TABLE `version` DISABLE KEYS */;
INSERT INTO `version` VALUES (2,'v1.2','aaa','bbb','1561446792',1);
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

-- Dump completed on 2019-12-06 10:30:07
