-- MySQL dump 10.13  Distrib 5.7.15, for Linux (x86_64)
--

CREATE TABLE `bugs` (
  `id`          BIGINT AUTO_INCREMENT COMMENT '主键',
  `title`       VARCHAR(255)  NOT NULL DEFAULT '' COMMENT 'Bug标题',
  `status_id`   BIGINT        NOT NULL DEFAULT 0 COMMENT 'Bug状态ID',
  `create_id`   BIGINT        NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `content`     TEXT          COMMENT '正文/Markdown',
  `important_id`         BIGINT        NOT NULL DEFAULT 0 COMMENT '重要性ID',
  `create_time`  DATETIME    COMMENT '创建时间',
  `handle_uid`  BIGINT        NOT NULL DEFAULT 0 COMMENT '被指派人ID（多值时存最小UID）',
  `level_id`         BIGINT        NOT NULL DEFAULT 0 COMMENT '优先级别ID',
  `env_id`         BIGINT        NOT NULL DEFAULT 0 COMMENT '运行环境ID',
  `type_id`         BIGINT        NOT NULL DEFAULT 0 COMMENT '类型ID',
  `project_id`         BIGINT        NOT NULL DEFAULT 0 COMMENT '项目ID',
  `update_time`  DATETIME   COMMENT '更新时间',
  `deadline`    DATETIME   DEFAULT NULL COMMENT '截止时间',
  `dustbin`     TINYINT(1)    NOT NULL DEFAULT 0 COMMENT '软删除标识 0=正常 1=已删',
  PRIMARY KEY (`id`),
  KEY `idx_project_id`        (`project_id`),
  KEY `idx_status_id`  (`status_id`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci COMMENT='Bug主表';


CREATE TABLE `bug_user` (
  `bug_id` bigint(20) NOT NULL default 0, -- bug id
  `uid` bigint(20) NOT NULL DEFAULT 0   -- 用户id
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
  `create_time` datetime,
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

CREATE TABLE `project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `created` datetime,
  `updated` datetime ,
  `uid` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE `project_user_map` (
  `project_id` bigint(20) NOT NULL default 0,
  `uid` bigint(20) NOT NULL DEFAULT 0,
  UNIQUE KEY `name` (`project_id`, `uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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


CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 ;


CREATE TABLE `role_perm_map` (
  `rid` bigint(20) NOT NULL DEFAULT 0 comment '角色组id',
  `perm_id` bigint(20) NOT NULL DEFAULT 0 comment '页面权限的id',
  enable boolean not null default false comment '禁用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

--
-- Dumping data for table `rolegroup`
--



--
-- Table structure for table `page_perm`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `page_perm` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `info` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `page_perm`
--

LOCK TABLES `page_perm` WRITE;
/*!40000 ALTER TABLE `page_perm` DISABLE KEYS */;
INSERT INTO `page_perm` VALUES 
(1,'env','环境页面'),
(2,'important','重要性页面'),
(3,'level','优先级别页面'),
(4,'position','职位页面'),
(5,'project','项目页面'),
(6,'status','bug状态流程页面'),
(8,'version','版本页面');
/*!40000 ALTER TABLE `page_perm` ENABLE KEYS */;
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
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `create_id` bigint(20) DEFAULT 0,
  `realname` varchar(30) NOT NULL,
  `disable` tinyint(1) DEFAULT 0,
  `position_id` bigint(20) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `realname` (`realname`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ;


CREATE TABLE `user_status` (
  `id` bigint(20) NOT NULL default 0,   -- 用户id
  `status` int(8) NOT NULL default 0   -- 状态id
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','6fe722103c3fd788608fa54a531d810f97236175',
'admin@qq.com','/showimg/1594376285974981434.png',now(),now(),1,'admin',0,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usergroup`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `uid` bigint(20) DEFAULT '0',   -- 创建的uid
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user_group_map` (
  `ugid` bigint(20) NOT NULL default 0,   -- 用户组id
  `uid` bigint(20) DEFAULT '0'  -- 用户id
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;




CREATE TABLE `version` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `urlone` varchar(255) DEFAULT '',
  `urltwo` varchar(255) DEFAULT '',
  `created` datetime ,
  `updated` datetime ,
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


CREATE TABLE `position` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `level` int(4) NOT NULL DEFAULT 0 COMMENT '0: 普通用户  1：管理者',
  `hypo` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级id',
  `role_id` int(11) NOT NULL DEFAULT 0 COMMENT '所属角色组',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;