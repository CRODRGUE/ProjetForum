-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : mar. 31 mai 2022 à 12:38
-- Version du serveur : 5.7.36
-- Version de PHP : 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `forum`
--

-- --------------------------------------------------------

--
-- Structure de la table `append`
--

DROP TABLE IF EXISTS `append`;
CREATE TABLE IF NOT EXISTS `append` (
  `id_message` bigint(20) NOT NULL,
  `id_emoji` bigint(20) NOT NULL,
  PRIMARY KEY (`id_message`,`id_emoji`),
  KEY `id_emoji` (`id_emoji`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `categorys`
--

DROP TABLE IF EXISTS `categorys`;
CREATE TABLE IF NOT EXISTS `categorys` (
  `id_category` bigint(20) NOT NULL AUTO_INCREMENT,
  `category` varchar(50) NOT NULL,
  `picto` varchar(255) NOT NULL,
  `sub-category` varchar(255) NOT NULL,
  `path` char(255) NOT NULL,
  PRIMARY KEY (`id_category`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `categorys`
--

INSERT INTO `categorys` (`id_category`, `category`, `picto`, `sub-category`, `path`) VALUES
(14, 'Théatre', './static/img/category/la-tragedie.png', 'Classique, moderne, tragedie...', 'theatre'),
(13, 'Autobiographie', './static/img/category/sante-mentale.png', 'Ecrivain, chanteur, homme politique...', 'autobiographie'),
(12, 'Littérature', './static/img/category/livre-dhistoire.png', 'poésie, essai, discours...', 'litterature'),
(11, 'Romans', './static/img/category/dictionnaire.png', 'Classique, science fiction, policier...', 'roman'),
(10, 'Information', './static/img/category/atlas.png', 'Reglement, F.A.Q, conseil...', 'information'),
(15, 'Savoir', './static/img/category/livre-de-science.png', 'Fiancier, droits, médecine...', 'savoir'),
(16, 'Manga / Bande déssinée', './static/img/category/science-fiction.png', 'Comisc, marvel...', 'manga');

-- --------------------------------------------------------

--
-- Structure de la table `emojis`
--

DROP TABLE IF EXISTS `emojis`;
CREATE TABLE IF NOT EXISTS `emojis` (
  `id_emoji` bigint(20) NOT NULL AUTO_INCREMENT,
  `emoticon` varchar(255) NOT NULL,
  PRIMARY KEY (`id_emoji`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `follows`
--

DROP TABLE IF EXISTS `follows`;
CREATE TABLE IF NOT EXISTS `follows` (
  `id_user` bigint(20) NOT NULL,
  `id_topic` bigint(20) NOT NULL,
  PRIMARY KEY (`id_user`,`id_topic`),
  KEY `id_topic` (`id_topic`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `gets`
--

DROP TABLE IF EXISTS `gets`;
CREATE TABLE IF NOT EXISTS `gets` (
  `id_topic` bigint(20) NOT NULL,
  `id_category` bigint(20) NOT NULL,
  PRIMARY KEY (`id_topic`,`id_category`),
  KEY `id_category` (`id_category`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `likes`
--

DROP TABLE IF EXISTS `likes`;
CREATE TABLE IF NOT EXISTS `likes` (
  `id_user` bigint(20) NOT NULL,
  `id_message` bigint(20) NOT NULL,
  `id_topic` bigint(20) NOT NULL,
  PRIMARY KEY (`id_user`,`id_message`),
  KEY `id_message` (`id_message`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `likes`
--

INSERT INTO `likes` (`id_user`, `id_message`, `id_topic`) VALUES
(1, 1, 10),
(23, 59, 9),
(12, 54, 10),
(12, 43, 6),
(12, 44, 6),
(12, 45, 6),
(12, 46, 6),
(23, 55, 8),
(12, 53, 6),
(23, 62, 19),
(23, 58, 10),
(23, 57, 10),
(23, 63, 10),
(23, 60, 9),
(24, 54, 10),
(25, 67, 2),
(25, 51, 6),
(25, 46, 6);

-- --------------------------------------------------------

--
-- Structure de la table `messages`
--

DROP TABLE IF EXISTS `messages`;
CREATE TABLE IF NOT EXISTS `messages` (
  `id_message` bigint(20) NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `date_message` datetime NOT NULL,
  `id_user` bigint(20) NOT NULL,
  `id_topic` bigint(20) NOT NULL,
  PRIMARY KEY (`id_message`),
  KEY `id_user` (`id_user`),
  KEY `id_topic` (`id_topic`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=69 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `messages`
--

INSERT INTO `messages` (`id_message`, `message`, `date_message`, `id_user`, `id_topic`) VALUES
(1, 'tesfdggtgtgtgtrgtgtrtg', '2022-04-19 00:00:00', 1, 6),
(2, '74656656', '2022-04-19 00:00:00', 13, 6),
(3, 'gugjhjhvjhvjhvjh', '2022-04-19 00:00:00', 1, 6),
(4, '[value-2]', '2022-04-21 14:38:50', 10, 1),
(5, '[value-4545]', '2022-04-21 14:38:50', 12, 1),
(6, '74656656', '2022-04-19 00:00:00', 12, 6),
(37, 'écris ton message ici...', '2022-05-08 17:39:38', 13, 7),
(36, 'bite !!!!!\r\n\r\n', '2022-04-21 16:21:31', 13, 6),
(35, 'écris ton message ici...', '2022-04-21 16:10:10', 13, 6),
(33, 'bite', '2022-04-21 16:09:17', 13, 6),
(34, 'écris ton message ici...', '2022-04-21 16:10:07', 13, 6),
(31, '74656656', '2022-04-19 00:00:00', 13, 6),
(38, 'écris ton message ici... okokokokokoo', '2022-05-08 17:40:54', 13, 7),
(54, 'écris ton message ici... hihihih', '2022-05-23 14:12:20', 12, 10),
(42, 'écris ton message ici...', '2022-05-09 18:31:37', 13, 6),
(43, 'test1', '2022-05-09 18:39:17', 13, 6),
(44, 'test2\r\n', '2022-05-09 18:39:37', 13, 6),
(45, 'lolo', '2022-05-09 18:39:44', 13, 6),
(46, 'pmojmpumu', '2022-05-09 18:39:50', 13, 6),
(47, 'écris ton message ici...', '2022-05-09 18:40:04', 13, 6),
(48, 'écris ton message ici...', '2022-05-13 15:41:32', 13, 6),
(49, 'ipoolyoloil', '2022-05-13 15:41:46', 13, 6),
(50, 'ortorgtijgbgijb', '2022-05-13 15:52:23', 13, 6),
(51, 'écris ton message ici...', '2022-05-13 15:54:52', 12, 6),
(52, 'écris ton message ici...fqtbbgs', '2022-05-14 21:34:59', 12, 6),
(53, 'écris ton message ici...fvbtrdsbfdb<sgreqhghtrhatahrta', '2022-05-14 21:35:04', 12, 6),
(57, 'écris ton message ici... c\'est moi l\'aim', '2022-05-23 20:49:32', 23, 10),
(63, 'écris ton message ici...', '2022-05-24 07:48:00', 23, 10),
(59, 'écris ton message ici...', '2022-05-23 20:50:09', 23, 9),
(60, 'écris ton message iciujjtyjrng...', '2022-05-23 20:50:12', 23, 9),
(61, 'écris ton message ici...', '2022-05-23 20:51:06', 23, 19),
(64, 'écris ton message ici..k;b,gh,fhxth.', '2022-05-24 07:48:04', 23, 10),
(65, 'salut topic value-1', '2022-05-24 07:57:09', 24, 10),
(66, 'salut vous !\r\n', '2022-05-24 08:13:34', 25, 10),
(67, 'salut salut', '2022-05-24 08:14:00', 25, 2);

-- --------------------------------------------------------

--
-- Structure de la table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
CREATE TABLE IF NOT EXISTS `permissions` (
  `id_perme` bigint(20) NOT NULL AUTO_INCREMENT,
  `role` varchar(255) NOT NULL,
  PRIMARY KEY (`id_perme`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `permissions`
--

INSERT INTO `permissions` (`id_perme`, `role`) VALUES
(10, '[value-2]'),
(11, '[value-2]'),
(12, '[value-2]'),
(13, '[value-3]'),
(14, '[value-2]'),
(15, '[value-2]'),
(16, '[value-2]');

-- --------------------------------------------------------

--
-- Structure de la table `pictures`
--

DROP TABLE IF EXISTS `pictures`;
CREATE TABLE IF NOT EXISTS `pictures` (
  `id_picture` bigint(20) NOT NULL AUTO_INCREMENT,
  `image` varchar(50) DEFAULT NULL,
  `id_message` bigint(20) NOT NULL,
  PRIMARY KEY (`id_picture`),
  KEY `id_message` (`id_message`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `possess`
--

DROP TABLE IF EXISTS `possess`;
CREATE TABLE IF NOT EXISTS `possess` (
  `id_user` bigint(20) NOT NULL,
  `id_perme` bigint(20) NOT NULL,
  PRIMARY KEY (`id_user`,`id_perme`),
  KEY `id_perme` (`id_perme`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `replay`
--

DROP TABLE IF EXISTS `replay`;
CREATE TABLE IF NOT EXISTS `replay` (
  `id_message` bigint(20) NOT NULL,
  `id_message_1` bigint(20) NOT NULL,
  PRIMARY KEY (`id_message`,`id_message_1`),
  KEY `id_message_1` (`id_message_1`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `social_networks`
--

DROP TABLE IF EXISTS `social_networks`;
CREATE TABLE IF NOT EXISTS `social_networks` (
  `id_social_network` bigint(20) NOT NULL AUTO_INCREMENT,
  `insta__account` varchar(255) DEFAULT NULL,
  `meta__account` varchar(255) DEFAULT NULL,
  `snap__account` varchar(255) DEFAULT NULL,
  `twitter__account` varchar(255) DEFAULT NULL,
  `id_user` bigint(20) NOT NULL,
  PRIMARY KEY (`id_social_network`),
  UNIQUE KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `topics`
--

DROP TABLE IF EXISTS `topics`;
CREATE TABLE IF NOT EXISTS `topics` (
  `id_topic` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `subject` varchar(255) NOT NULL,
  `start_date` date NOT NULL,
  `polity` varchar(255) DEFAULT NULL,
  `description` text,
  `id_topic_picture` int(11) NOT NULL,
  `id_category` bigint(20) NOT NULL,
  `id_user` bigint(20) NOT NULL,
  PRIMARY KEY (`id_topic`),
  UNIQUE KEY `title` (`title`),
  KEY `id_topic_picture` (`id_topic_picture`),
  KEY `id_category` (`id_category`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `topics`
--

INSERT INTO `topics` (`id_topic`, `title`, `subject`, `start_date`, `polity`, `description`, `id_topic_picture`, `id_category`, `id_user`) VALUES
(1, 'ok', 'ok', '2022-04-17', '1', 'kololonlhltutgtg-l-ç', 2, 13, 5),
(2, 'ok54', 'ok', '2022-04-17', '1', 'kololonlhltutgtg-l-ç', 2, 13, 5),
(3, 'hey test', 'test test test ', '2022-04-19', '1', 'azezrezrt tbhbh hdjynghfd  regyteh t thytrhgthbdt', 4, 14, 5),
(4, 'hi hi hi hi', 'test test test ', '2022-04-19', '5', '1 1 1 1 1 1 1', 1, 14, 5),
(7, 'je suis ok', 'robert test', '2022-05-08', '1', 'salut je voudrais votre retour sur le livre de robert test !', 3, 15, 5),
(6, 'testetestetestetst', 'test test test ', '2022-04-19', '1', 'fgtgrtreazte ghhghsd trrrh te fdhtrdhstrhstrhse hst rh htsrhsth ew', 6, 14, 5),
(8, '[value-2]', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 14, 12),
(9, '[value-2lom288]', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 14, 12),
(10, '[value-1]', '[value-2]', '2022-04-17', '1', '[value-5]', 2, 14, 12),
(11, 'testetestetestetstefrvf', 'tetststettetst', '2022-05-16', '1', 'vtrhbhytbytbytbtrbr', 2, 13, 5),
(12, 'testetestetestetstefrvfmpmoi', 'tetststettetstokok', '2022-05-16', '1', 'maintenant c\'est ok ! ', 1, 13, 5),
(13, '[value-2]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(14, '[value-247]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(15, '[ve-247]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(16, '[vhtyr-èe-247]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(17, '[vhtyjujyjur-èe-247]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(18, '[èe-247]4', '[value-3]', '2022-04-19', '1', '[value-6]', 1, 13, 14),
(19, 'okooko', 'okokokook', '2022-05-23', '1', 'Je suis un nouveau topic !', 5, 13, 5),
(20, 'hjgkjgc', 'hfjhgy', '2022-05-24', '1', 'fjyjghjyfdjyftjydtydrj', 1, 14, 10);

-- --------------------------------------------------------

--
-- Structure de la table `topics_pictures`
--

DROP TABLE IF EXISTS `topics_pictures`;
CREATE TABLE IF NOT EXISTS `topics_pictures` (
  `id_topic_picture` int(11) NOT NULL AUTO_INCREMENT,
  `path` varchar(255) NOT NULL,
  PRIMARY KEY (`id_topic_picture`),
  UNIQUE KEY `path` (`path`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `topics_pictures`
--

INSERT INTO `topics_pictures` (`id_topic_picture`, `path`) VALUES
(1, '/static/img/topic/Picture_1.png'),
(2, '/static/img/topic/Picture_2.png'),
(3, '/static/img/topic/Picture_3.png'),
(4, '/static/img/topic/Picture_4.png'),
(5, '/static/img/topic/Picture_5.png'),
(6, '/static/img/topic/Picture_6.png'),
(7, '/static/img/topic/Picture_7.png'),
(8, '/static/img/topic/Picture_8.png');

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id_user` bigint(20) NOT NULL AUTO_INCREMENT,
  `lastname` varchar(255) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `brith_day` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `pseudo` varchar(50) NOT NULL,
  `description` text NOT NULL,
  `profile_picture` int(11) NOT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `pseudo` (`pseudo`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id_user`, `lastname`, `firstname`, `brith_day`, `email`, `password`, `pseudo`, `description`, `profile_picture`) VALUES
(1, 'test_1', 'test_1', '2022-04-26', 'test1@test.ts', 'aze123', 'test_1', 'hnhnhn', 1),
(2, 'test_2', 'test_2', '2022-04-26', 'test2@test.ts', 'aze123', 'test_2', 'hnghg', 1),
(13, 'Rodrigues', 'Anna-maria', '2022-04-26', 'cyril.commande@gmail.com', '{SSHA}Ni1oFnQzMZnvGdvk2nfJuVVj3lwRhOWYGw5II9Tg3vM=', 'cyril.commom', 'dfbghj;h,hngbfvdcxw', 1),
(4, 'rodrigues', 'cyril', '2022-04-26', 'test.test@outlook.fr', 'aze12398ze', '__cyril', 'Salut toi ! je suis un jeune qui kiffe les livres... =)', 3),
(5, 'test', 'test', '2022-04-26', 'cyril.opop@outlok.fr', '{SSHA}KBhymFF5daRASq9LVxES1AQn50BWHsSwKIycRqrgWr3bzvQ=', 'testcyril', 'lalalalalalalallalalala', 2),
(6, 't4545454est', 'cyril', '2022-04-26', 'testtest@oi.fr', '{SSHA}NwPJUICvn+lK0aG+an8z8aNWSRyLQBvRifAFO1RZxT/JJaw=', 'lol_lol', 'zdfzefrzfgre rezgr erezg rehthz hh', 4),
(7, 'cyril', 'cyril', '2022-04-26', 'test159753@ab.fr', '{SSHA}bZDlgo8EPAUTXVeuUhMmnYLmBdVDYWWdCw24pgkPI7xvzzM=', 'teste', '4546545487685', 4),
(8, 'test', 'test', '2022-04-26', 'ynov@test.ts', '{SSHA}ThpUyFFxquMVKEimW4xX3xJb960+/LeXDHbIv5whkVCFjgk=', 'test', 'azert', 5),
(9, 'Rodrigues', 'Anna-maria', '2022-04-26', 'cyril@gmail.com', '{SSHA}NJCh3Djz5sqh6AKv/xdL9c3639VqVZkpqfyEUKorlouumCA=', 'cyril.comm', 'rgbrjrngfr', 5),
(10, 'fvgbtbtrbtr', 'testt111', '2022-04-26', 'azerty@azert.fr', 'cyrilcyril13', 'profileEdite10', 'efvbhtnbtgrfvtrvr', 5),
(11, 'azeze', 'azeeza', '2022-04-26', 'fofo@gmail.com', '{SSHA}3aJTNTrIGwUFdrQl1KMGJjymE+fYvgFLJ/secf8BRU3q/qU=', 'cyril', 'azeza', 5),
(12, 'azrerear12', 'azrarzar12', '2022-04-26', 'testM@gmail.ts', 'cyrilcyril13', '12az12ezr12', 'zer\'rttgreggf Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas\r\n                    mauris purus, luctus vitae sapien malesuada, pellentesque\r\n                    hendrerit nisi. Pellentesque porttitor diam a leo malesuada, in\r\n                    ornare diam aliquet. Nam id dolor ullamcorper, cursus purus at,\r\n                    vulputate arcu. Integer eleifend eros id sapien consequat, ut\r\n                  cdsccs<cq  gravida erat volutpat. luctus vitae sapien malesuada, pellentesque\r\n                    hendrerit nisi. Pellentesque porttitor diam a leo malesuada', 3),
(14, 'Rodrigues', 'Anna-maria', '2022-04-26', 'test.test@test.fr', '{SSHA}r5f0s1YE8Ai78dlrL4wyMMcGpc/zzHreZsT6xH5v', 'cyrilom', '894134+6464', 5),
(15, 'roro', 'cyril', '2022-04-26', 'test.testa@gmail.com', '{SSHA}XVelrCzQkp1vvhVz7laJEq/xUlJyq0Xk5YAxchGI', 'testCyriltest', 'salut je suis un compte de test !', 3),
(16, 'cr', 'cr', '2022-04-26', 'aqw@gmail.com', '$2a$14$DtG4K5R8GsU2.3cWOA6nhuTWlKR5EGgiLzKqGEZjmc0qq0hpkQQ6q', 'je suis ok', 'je suis ok ', 3),
(17, 'areree', 'azer', '2022-04-26', 'jklm@gmail.com', '$2a$14$pSawbjBDt7G7g.MnGfLF5.vxiqLxMy4/mj4irh9sNJBYIuuEtw41i', 'lolololololololol', 'uihlè-goi_fèik i-kg(èufrjfèyh ujyvkjujbjrvv ju(ufgj', 2),
(18, 'Rodrigues', 'Anna-maria', '2022-04-26', 'bobo@outlook.fr', '$2a$14$hwLq3j/kP7JYfDAEtVAateBX7r7spHArLOMk5bN0KgJpjYUKl1KEO', '_aa_', 'fgtykigukjhbvdfvbyntyjnuy,', 4),
(19, 'Rodrigues', 'Anna-maria', '2022-04-26', 'nono@outlook.fr', '$2a$14$cx93BW3y0g8J1P8kHmmpSeQtlYcsjHCUCjApKy5IDXHqeOV9Zlsi2', 'aaaaaazzzzz', 'defrgthjkl:m!:;,nbvcxsdcrgvty-ehtvyhtvehtrhtvrht', 1),
(20, 'trhgrtgtrgtr', 'zdfrbgfbrtd', '2022-04-26', 'testrr@outlook.fr', '$2a$14$UE4G4qiVuta5.N9e0d.1YuGDY5eDFGk0mz0Ijy33h8AqNRz3t7ZIu', 'ascftyj', '\"rf(grtyjutyytjtyjtyjyt', 2),
(21, 'grezgzergze', 'grzegrezg', '2022-04-26', 'gtgtg@fjrhfr.com', '$2a$14$w5m4JI3WYNUt4x6AS8i8q.jqrv1coZ6o4QaKUutnhI.NE/k6PDP1m', 'zsefrg', 'rhgtrhyetjurkiukoililçy', 4),
(22, 'Test', 'Compte', '2022-05-13', 'comptetest00@outlook.fr', '$2a$14$6whtzno9kAnCVmGlXXPYYu1s5W82gIZTqnf8xxsPV30i7s2lm4592', 'CompteTest00', 'Salut ! je suis un compte de test l\'ami', 4),
(23, '00test', 'test00', '2022-05-19', 'test00test@outlook.fr', '$2a$14$OyETyFkVbEdAh9vKtH8WGuQGNj648LbBaGl269p3xa/h.QmmqhmPG', 'test00test', 'Salut toi j\'aime les bites bien grosses !!!', 2),
(24, 'testtest', 'test', '2015-05-15', 'test02@outlook.fr', '$2a$14$mjMIJg3EKcNdSS.5sNLAu.nh0FbEFFUoihFWxXwY1qmioAHpbIRi.', 'test02', 'Bonjours ! je suis un compte de test', 2),
(25, 'test', 'cyril', '2001-06-01', 'test04@outlook.fr', '$2a$14$4y7dWilsgFNrxBHOHdV/hubs5FETzCdcawUvRiivIXJlqXwLq5luq', 'cyrilTest0001', 'Salut ! Je suis un compte de test !!!!', 5);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
