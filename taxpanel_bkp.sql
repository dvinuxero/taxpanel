-- MySQL dump 10.13  Distrib 5.5.46, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: taxpanel
-- ------------------------------------------------------
-- Server version	5.5.46-0ubuntu0.14.04.2

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
-- Table structure for table `expensive`
--

DROP TABLE IF EXISTS `expensive`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `expensive` (
  `id` varchar(50) NOT NULL DEFAULT '',
  `description` varchar(120) DEFAULT NULL,
  `status` varchar(20) DEFAULT 'ACTIVE',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `expensive`
--

LOCK TABLES `expensive` WRITE;
/*!40000 ALTER TABLE `expensive` DISABLE KEYS */;
INSERT INTO `expensive` VALUES ('ABL_SANTAFE','alumbrado barrido y limpieza','ACTIVE'),('ALQUILER','alquiler de departamento','ACTIVE'),('AROMATIZADORES','aromatizadores','ACTIVE'),('AYSA_LANUS','agua de lanus','ACTIVE'),('AYSA_SANTAFE','agua de santa fe','ACTIVE'),('CABLEVISION_LANUS','cablevision de lanus','ACTIVE'),('CARGA_SUBE','carga de sube mensual','ACTIVE'),('COMIDA_MENSUAL','comida mensual','ACTIVE'),('CORTE_PELO_ALVI','corte de pelo de alvi','ACTIVE'),('CUOTA_MAMA','cuota mama mensual','ACTIVE'),('DAVINCI','facultad eve davinci','ACTIVE'),('DESPARACITANTE_ABRIL','desparacitante abril','ACTIVE'),('DOPING_LEO','analisis de doping de leo','INACTIVE'),('EDENOR_CARRANZA','edenor de carranza','INACTIVE'),('EDENOR_SANTAFE','luz santa fe','ACTIVE'),('EDESUR_LANUS','edesur de lanus','ACTIVE'),('EXPENSAS_CARRANZA','expensas de carranza','ACTIVE'),('EXPENSAS_SANTAFE','expensas de santa fe','ACTIVE'),('EXTRA_HOGAR_ABUELA','pago de cosas extras del hogar de la abuela','ACTIVE'),('GALENO_LEO','obra social galeno de leo','ACTIVE'),('GYM_EVE','gimnasio para eve','INACTIVE'),('INMOBILIARIO_LANUS','inmobiliario de lanus','ACTIVE'),('INSTITUTO_LEO','instituto de leo','INACTIVE'),('LIMPIEZA_CARRANZA','chica limpieza casa mama','ACTIVE'),('METROGAS_CARRANZA','metrogas de carranza','INACTIVE'),('METROGAS_LANUS','metrogas de lanus','ACTIVE'),('METROGAS_SANTAFE','gas de santa fe','ACTIVE'),('MOVISTAR','celular personal','INACTIVE'),('MUNICIPALIDAD_LANUS','municipalidad de lanus','ACTIVE'),('PANIALES_ABUELA','paniales de la abuela','ACTIVE'),('PASTILLAS_EVE','pastillas anticonseptivas eve','ACTIVE'),('PERSONAL_LEO','personal de leo','INACTIVE'),('PIEDRITAS_ABRIL','piedritas sanitarias abril','ACTIVE'),('PIPETA_ABRIL','pipeta abril','ACTIVE'),('PSICOLOGO','sesiones de psicologo','INACTIVE'),('REMEDIOS_ABUELA','remedios para la abuela','ACTIVE'),('TARJETA_AMEX','tarjeta de credito amex santander','ACTIVE'),('TARJETA_VISA','tarjeta de credito visa santander','ACTIVE'),('TELECOM_CARRANZA','telecom de carranza','ACTIVE'),('TELECOM_SANTAFE','telefono y arnet de santa fe','ACTIVE'),('TIRAS_MAMA','tiras para medir glucemia mama','INACTIVE'),('UAI','cuota facultad','ACTIVE'),('VENTOLIN','ventolin anti agitacion','ACTIVE'),('VIAJE_ALVI','cuota viaje de egresados de alvi','INACTIVE');
/*!40000 ALTER TABLE `expensive` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `expensive_id` varchar(50) DEFAULT NULL,
  `tag_id` varchar(50) DEFAULT NULL,
  `mount` int(10) DEFAULT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (2,'ABL_SANTAFE','mandatory',145,'2016-01-23 22:57:45'),(3,'ALQUILER','mandatory',4200,'2016-01-23 23:02:05'),(4,'AROMATIZADORES','mandatory',120,'2016-01-23 23:02:17'),(5,'AYSA_LANUS','extra',67,'2016-01-23 23:03:04'),(6,'CABLEVISION_LANUS','extra',423,'2016-01-23 23:03:28'),(7,'CARGA_SUBE','mandatory',300,'2016-01-23 23:03:39'),(9,'COMIDA_MENSUAL','mandatory',2000,'2016-01-23 23:04:38'),(10,'CORTE_PELO_ALVI','extra',100,'2016-01-23 23:04:53'),(11,'CUOTA_MAMA','extra',2400,'2016-01-23 23:05:07'),(12,'DAVINCI','mandatory',1900,'2016-01-23 23:05:23'),(13,'DESPARACITANTE_ABRIL','mandatory',100,'2016-01-23 23:05:36'),(14,'DOPING_LEO','extra',800,'2016-01-23 23:05:49'),(15,'EDENOR_SANTAFE','mandatory',40,'2016-01-23 23:06:34'),(16,'EDESUR_LANUS','extra',69,'2016-01-23 23:07:07'),(17,'EXPENSAS_CARRANZA','extra',590,'2016-01-23 23:07:43'),(18,'EXPENSAS_SANTAFE','mandatory',1349,'2016-01-23 23:08:11'),(19,'EXTRA_HOGAR_ABUELA','mandatory',0,'2016-01-23 23:08:38'),(20,'GALENO_LEO','extra',1000,'2016-01-23 23:08:50'),(21,'INMOBILIARIO_LANUS','extra',137,'2016-01-23 23:09:18'),(22,'INSTITUTO_LEO','extra',1000,'2016-01-23 23:09:41'),(23,'LIMPIEZA_CARRANZA','mandatory',680,'2016-01-23 23:10:16'),(24,'METROGAS_LANUS','extra',43,'2016-01-23 23:11:03'),(25,'METROGAS_SANTAFE','mandatory',71,'2016-01-23 23:11:09'),(26,'MUNICIPALIDAD_LANUS','extra',263,'2016-01-23 23:11:35'),(27,'PANIALES_ABUELA','mandatory',450,'2016-01-23 23:11:49'),(28,'PASTILLAS_EVE','mandatory',50,'2016-01-23 23:11:58'),(29,'PIEDRITAS_ABRIL','mandatory',40,'2016-01-23 23:12:08'),(30,'PIPETA_ABRIL','mandatory',50,'2016-01-23 23:12:20'),(31,'REMEDIOS_ABUELA','mandatory',500,'2016-01-23 23:12:39'),(32,'TARJETA_AMEX','mandatory',1000,'2016-01-23 23:12:51'),(33,'TELECOM_CARRANZA','extra',347,'2016-01-23 23:13:08'),(34,'TELECOM_SANTAFE','mandatory',401,'2016-01-23 23:13:31'),(35,'TIRAS_MAMA','extra',440,'2016-01-23 23:13:42'),(36,'UAI','mandatory',3040,'2016-01-23 23:13:52'),(37,'VENTOLIN','mandatory',100,'2016-01-23 23:14:00'),(38,'ABL_SANTAFE','mandatory',146,'2016-01-23 23:29:22');
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(120) DEFAULT NULL,
  `salary` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'dario villanustre',21000);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-01-24 12:54:43
