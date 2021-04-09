-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 28, 2019 at 07:47 AM
-- Server version: 10.1.37-MariaDB
-- PHP Version: 7.3.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cms_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` int(10) NOT NULL,
  `post_id` int(10) NOT NULL,
  `commentator_id` int(10) NOT NULL,
  `message` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `comments`
--

INSERT INTO `comments` (`id`, `post_id`, `commentator_id`, `message`, `created_at`, `updated_at`) VALUES
(1, 2, 14, 'khá tốt', '2019-01-23 07:46:58', '2019-01-23 07:46:58'),
(2, 3, 14, 'khá ổn', '2019-01-23 07:50:35', '2019-01-23 07:50:35'),
(3, 3, 14, 'title chưa ổn', '2019-01-23 07:55:35', '2019-01-23 07:55:35'),
(4, 4, 14, 'bài viết tốt', '2019-01-23 08:15:52', '2019-01-23 08:15:52');

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(10) NOT NULL,
  `creator` int(10) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `topic` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `creator`, `title`, `topic`, `description`, `content`, `status`, `created_at`, `updated_at`) VALUES
(1, 41, 'Man must explore, and this is exploration at its greatest', 'Problems look mighty small from 150 miles up', 'Problems look mighty small from 150 miles up', 'Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman\'s earth, if free men make it, will be truly round: a globe in practice, not in theory.', 1, '2019-01-27 17:00:00', '2019-01-27 17:00:00'),
(2, 41, 'I believe every human has a finite number of heartbeats. I don\'t intend to waste any of mine.', 'We predict too much for the next year and yet far too little for the next ten.', 'We predict too much for the next year and yet far too little for the next ten.', 'Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman\'s earth, if free men make it, will be truly round: a globe in practice, not in theory.', 1, '2019-01-27 17:00:00', '2019-01-27 17:00:00'),
(3, 40, 'I believe every human has a finite number of heartbeats', 'We predict too much for the next year and yet far too little for the next ten.', 'We predict too much for the next year and yet far too little for the next ten.', 'Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman\'s earth, if free men make it, will be truly round: a globe in practice, not in theory.', 1, '2019-01-27 17:00:00', '2019-01-27 17:00:00'),
(4, 40, 'Every human has a finite number of heartbeats', 'topic 2', 'We are challenge', 'Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman\'s earth, if free men make it, will be truly round: a globe in practice, not in theory.', 0, '2019-01-27 17:00:00', '2019-01-27 17:00:00'),
(5, 42, 'Test title', 'topic 3', 'We are challenge', 'Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman\'s earth, if free men make it, will be truly round: a globe in practice, not in theory.', 1, '2019-01-27 17:00:00', '2019-01-27 17:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `birthday` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone_number` int(10) DEFAULT NULL,
  `type` int(1) NOT NULL,
  `status` int(1) NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `confirm` int(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `email`, `name`, `gender`, `birthday`, `phone_number`, `type`, `status`, `token`, `confirm`, `created_at`, `updated_at`) VALUES
(14, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'admin@hblab.vn', 'Admin', 'Nam', '05/03/1997', 961706497, 1, 1, NULL, 1, '2019-01-18 08:46:26', '2019-01-25 02:23:55'),
(40, 'vip1', 'e10adc3949ba59abbe56e057f20f883e', 'vip1@gmail.com', 'vip1', 'Nam', '01/01/2019', 235235123, 0, 1, 'd397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467', 1, '2019-01-25 03:36:13', '2019-01-25 04:53:12'),
(41, 'anhnt1', 'e10adc3949ba59abbe56e057f20f883e', 'anhnt1@hblab.vn', 'Tuấn Anh', 'Nam', '02/01/2019', 12345678, 0, 1, '08970496b6310d4e2e84c99b4064dccc9b45195d374724f65886aa5bff97e2c5', 1, '2019-01-25 03:44:47', '2019-01-25 04:53:24'),
(42, 'anhnt2', 'e10adc3949ba59abbe56e057f20f883e', 'anhnt2@gmail.com', 'Anhnt2', 'Nam', '06/01/2019', 12345690, 0, 1, '5f4f459d62ff2f7d39e5d9429b90737b8897963d6f83bd2e6586b326235e777e', 1, '2019-01-25 03:46:13', '2019-01-25 04:53:16'),
(43, 'vip2', 'e10adc3949ba59abbe56e057f20f883e', 'vip2@gmail.com', 'vip2', 'Nam', '01/01/2019', 123124123, 0, 1, 'be4303c16602a63552f29f0ab8b44cf8e3a17972326b63844a658bdc4da3543a', 1, '2019-01-25 04:04:25', '2019-01-25 04:53:19'),
(44, 'vip3', 'e10adc3949ba59abbe56e057f20f883e', 'vip3@gmail.com', 'Vip3', 'Nam', '01/01/2019', 12345678, 0, 1, 'd879966df1d5a4d57208916aca869157148ac5d025d15ec678f4d8bbbda340e3', 1, '2019-01-25 04:12:08', '2019-01-25 04:12:08'),
(45, 'gadsg', 'e10adc3949ba59abbe56e057f20f883e', 'gsa@gmail.com', 'dsagd', 'Nam', '10/01/2019', 124123412, 0, 0, '27155cab57f6b3ad241764e4ac5b6dd1334310e17293c2260fdf09a5cf257cc4', 1, '2019-01-25 04:55:46', '2019-01-25 04:55:46');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_post_id` (`post_id`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_id` (`creator`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `comments`
--
ALTER TABLE `comments`
  ADD CONSTRAINT `fk_post_id` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

--
-- Constraints for table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `fk_user_id` FOREIGN KEY (`creator`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
