-- MySQL dump 10.13  Distrib 5.7.15, for Linux (x86_64)
--

CREATE TABLE `bugs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `title` varchar(50) DEFAULT NULL,
  `sid` bigint(20) DEFAULT '0',
  `content` text,
  `ownerid` bigint(20) DEFAULT '0',
  `iid` bigint(20) DEFAULT '0',
  `createtime` bigint(20) DEFAULT '0',
  `vid` bigint(20) DEFAULT '0',
  `spusers` varchar(255) DEFAULT '',
  `lid` bigint(20) DEFAULT '0',
  `eid` bigint(20) DEFAULT '0',
  `tid` bigint(20) DEFAULT '0',
  `pid` bigint(20) DEFAULT '0',
  `updatetime` bigint(20) DEFAULT '0',
  `dustbin` tinyint(1) DEFAULT '0',
  `deadline` int not null default 0,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `sid` (`sid`),
  KEY `iid` (`iid`),
  KEY `ownerid` (`ownerid`),
  KEY `vid` (`vid`),
  KEY `lid` (`lid`),
  KEY `eid` (`eid`),
  KEY `pid` (`pid`),
  KEY `tid` (`tid`),
  KEY `dustbin` (`dustbin`),
  KEY `updatetime` (`updatetime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `bugs`
--


CREATE TABLE `defaultvalue` (
  `created` int(11) NOT NULL DEFAULT '0',
  `completed` int(11) NOT NULL DEFAULT '0',
  `pass` int(11) NOT NULL DEFAULT '0',
  `receive` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--

insert into defaultvalue values(1, 2, 3, 4);

--
-- Table structure for table `email`
--

--
CREATE TABLE `typ` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- Dumping data for table `typ`

LOCK TABLES `typ` WRITE;
insert into typ(name) values('bug'), ('需求');
UNLOCK TABLES;

CREATE TABLE `email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `password` varchar(50) DEFAULT '',
  `port` int(11) DEFAULT '25',
  `enable` tinyint(1) NOT NULL DEFAULT '0',
  `host` varchar(30) NOT NULL DEFAULT '',
  `nickname` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

CREATE TABLE `environment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

CREATE TABLE `header` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `hhids` varchar(100) DEFAULT '',
  `remark` varchar(30) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `header`
--

LOCK TABLES `header` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `headerlist`
--

CREATE TABLE `headerlist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `k` varchar(200) NOT NULL,
  `v` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `headerlist`
--

LOCK TABLES `headerlist` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `importants`
--

CREATE TABLE `importants` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

CREATE TABLE `informations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `bid` bigint(20) NOT NULL DEFAULT '0',
  `info` varchar(200) NOT NULL DEFAULT '',
  `time` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `informations`
--

LOCK TABLES `informations` WRITE;
/*!40000 ALTER TABLE `informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `informations` ENABLE KEYS */;
UNLOCK TABLES;



--
-- Table structure for table `level`
--

CREATE TABLE `level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

CREATE TABLE `log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `exectime` bigint(20) DEFAULT '0',
  `classify` varchar(30) NOT NULL DEFAULT '',
  `ip` varchar(40) DEFAULT '',
  `uid` bigint(20) DEFAULT '0',
  `action` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `options`
--

CREATE TABLE `options` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `info` varchar(100) DEFAULT '',
  `tid` bigint(20) DEFAULT '0',
  `df` varchar(10) DEFAULT '',
  `need` varchar(10) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `options`
--

LOCK TABLES `options` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `perm`
--

CREATE TABLE `perm` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pv` int(8) NOT NULL DEFAULT '0',
  `rid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `perm`
--

LOCK TABLES `perm` WRITE;
/*!40000 ALTER TABLE `perm` DISABLE KEYS */;
/*!40000 ALTER TABLE `perm` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `ugid` bigint(20) NOT NULL DEFAULT '0',
  `uid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `restfulname` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rolegroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `permids` json,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `info` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'env','环境页面'),
(2,'important','重要性页面'),
(3,'level','优先级别页面'),
(4,'position','职位页面'),
(5,'project','项目页面'),
(6,'status','bug状态流程页面'),
(8,'version','版本页面');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sharefile`
--

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `status`
--

LOCK TABLES `status` WRITE;
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
insert into status values(1, '新建'), (2, '已完成'),(3, '领取中'),(4, '处理中');
/*!40000 ALTER TABLE `status` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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
  `showstatus` json,
  `disable` tinyint(1) DEFAULT '0',
  `jid` bigint(20) DEFAULT '0',

  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `realname` (`realname`),
  KEY `jid` (`jid`),
  KEY `createuid` (`createuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','69ad5117e7553ecfa7f918a223426dd8da08a57f',
'admin@qq.com','http://120.26.164.125:10001/showimg/1594376285974981434.png',unix_timestamp(now()),0,'admin','[]',0,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usergroup`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usergroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `uids` json,
  `uid` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `usergroup`
--

LOCK TABLES `usergroup` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `version`
--

CREATE TABLE `version` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `urlone` varchar(255) DEFAULT '',
  `urltwo` varchar(255) DEFAULT '',
  `createtime` varchar(30) DEFAULT '0',
  `createuid` bigint(20) NOT NULL,
  `pid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_version` (`pid`,`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `version`
--

LOCK TABLES `version` WRITE;
UNLOCK TABLES;


CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `typ` int(11) NOT NULL DEFAULT '0',
  `pri` text COLLATE utf8mb4_unicode_ci,
  `pub` text COLLATE utf8mb4_unicode_ci,
  `user` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `uid` int(11) NOT NULL DEFAULT '0',
  `created` int(11) NOT NULL DEFAULT '0',
  `uptime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



CREATE TABLE `doc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `uid` int(11) DEFAULT '0',
  `created` int(11) DEFAULT NULL,
  `giturl` varchar(100) DEFAULT '',
  `uptime` int(11) NOT NULL DEFAULT '0',
  `dir` varchar(20) NOT NULL DEFAULT '',
  `port` int(11) NOT NULL DEFAULT '0',
  `kid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `jobs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `level` bigint(20) NOT NULL DEFAULT '2',
  `hypo` varchar(30) NOT NULL DEFAULT '0',
  `rgid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;