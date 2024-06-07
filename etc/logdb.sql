-- MariaDB dump 10.19-11.3.2-MariaDB, for osx10.19 (x86_64)
--
-- Host: localhost    Database: logdb
-- ------------------------------------------------------
-- Server version	11.3.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

create database logdb;

--
-- Table structure for table `kv_high_frequencies`
--

DROP TABLE IF EXISTS `kv_high_frequencies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kv_high_frequencies` (
  `data_no` bigint(20) NOT NULL,
  `data_date` date NOT NULL,
  `data_time` varchar(256) NOT NULL,
  `interval` bigint(20) DEFAULT NULL,
  `current_now` bigint(20) DEFAULT NULL,
  `tuning_now` bigint(20) DEFAULT NULL,
  `d00108` bigint(20) DEFAULT NULL,
  `temp_up` bigint(20) DEFAULT NULL,
  `temp_down` bigint(20) DEFAULT NULL,
  `welding_time` bigint(20) DEFAULT NULL,
  `cooling_time` bigint(20) DEFAULT NULL,
  `tuning_speed` bigint(20) DEFAULT NULL,
  `current_value1` bigint(20) DEFAULT NULL,
  `tuning_value1` bigint(20) DEFAULT NULL,
  `welding_time1` bigint(20) DEFAULT NULL,
  `cooling_time1` bigint(20) DEFAULT NULL,
  `welding_time2` bigint(20) DEFAULT NULL,
  `tuning_value2` bigint(20) DEFAULT NULL,
  `current_value2` bigint(20) DEFAULT NULL,
  `welding_sw2` bigint(20) DEFAULT NULL,
  `set_temp_up` bigint(20) DEFAULT NULL,
  `high_voltage_memory` bigint(20) DEFAULT NULL,
  `set_temp_down` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`data_no`,`data_date`,`data_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `kv_microwaves`
--

DROP TABLE IF EXISTS `kv_microwaves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kv_microwaves` (
  `data_no` bigint(20) NOT NULL,
  `data_date` date NOT NULL,
  `data_time` varchar(256) NOT NULL,
  `mw_pin1` bigint(20) DEFAULT NULL,
  `mw_pin2` bigint(20) DEFAULT NULL,
  `mw_pin3` bigint(20) DEFAULT NULL,
  `mw_pin4` bigint(20) DEFAULT NULL,
  `mw_pf1` bigint(20) DEFAULT NULL,
  `mw_pf2` bigint(20) DEFAULT NULL,
  `mw_pf3` bigint(20) DEFAULT NULL,
  `mw_pf4` bigint(20) DEFAULT NULL,
  `mw_pr1` bigint(20) DEFAULT NULL,
  `mw_pr2` bigint(20) DEFAULT NULL,
  `mw_pr3` bigint(20) DEFAULT NULL,
  `mw_pr4` bigint(20) DEFAULT NULL,
  `stop_cv` bigint(20) DEFAULT NULL,
  `water_level` bigint(20) DEFAULT NULL,
  `far_infrared_heater` bigint(20) DEFAULT NULL,
  `mw_output` bigint(20) DEFAULT NULL,
  `mw_total_output` bigint(20) DEFAULT NULL,
  `temp_right` bigint(20) DEFAULT NULL,
  `temp_left` bigint(20) DEFAULT NULL,
  `temp_top` bigint(20) DEFAULT NULL,
  `temp_mirror` bigint(20) DEFAULT NULL,
  `temp_floor` bigint(20) DEFAULT NULL,
  `temp_door` bigint(20) DEFAULT NULL,
  `in_room_pressure` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`data_no`,`data_date`,`data_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `login_users`
--

DROP TABLE IF EXISTS `login_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `login_users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `department` varchar(256) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-07 22:16:34
