-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Янв 23 2022 г., 20:48
-- Версия сервера: 8.0.24
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `library`
--

-- --------------------------------------------------------

--
-- Структура таблицы `books`
--

CREATE TABLE `books` (
  `id_book` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `price_of_book` varchar(50) DEFAULT '0',
  `num_of_copies` varchar(50) DEFAULT '0',
  `authors` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default',
  `cover_photo` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default',
  `price_per_day` varchar(50) DEFAULT '0',
  `reg_date` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `books`
--

INSERT INTO `books` (`id_book`, `name`, `price_of_book`, `num_of_copies`, `authors`, `cover_photo`, `price_per_day`, `reg_date`) VALUES
(20, 'reg', '4', '4', 'reg', 'default', '4', '12-20-2021'),
(22, '100', '100', '98', '100', 'default', '100', '12-21-2021'),
(25, 'newBook', '123', '3', 'authors', 'default', '1', '12-28-2021'),
(26, 'refactor', '123', '8', '4', 'default', '2', '12-29-2021'),
(40, 'hunger games', '1000', '5', 'Collin', 'https://s00.yaplakal.com/pics/pics_original/1/2/6/11606621.jpg', '1', '01-18-2022'),
(41, 'hunger games', '1000', '5', 'Collin', 'https://s00.yaplakal.com/pics/pics_original/1/2/6/11606621.jpg', '1', '01-19-2022'),
(46, 'ippppppa', '1000', '0', 'default', 'default', '0', 'default'),
(47, 'ippppppa', '1000', '5', 'Collin', 'https://mobimg.b-cdn.net/v3/fetch/c4/c493aac67877288476b0fc52d55f55cf.jpeg', '1', '01-20-2022'),
(50, 'harry potter', '1200', '4', 'Collin', 'https://image.tmdb.org/t/p/original/8f9dnOtpArDrOMEylpSN9Sc6fuz.jpg', '1', '01-20-2022'),
(58, 'Harry Potter', '1000', '5', 'Rowling', 'https://image.tmdb.org/t/p/original/8f9dnOtpArDrOMEylpSN9Sc6fuz.jpg', '1', '01-20-2022'),
(59, 'newOne', '12', '5', 'newOne', 'https://image.tmdb.org/t/p/original/8f9dnOtpArDrOMEylpSN9Sc6fuz.jpg', '1', '01-23-2022'),
(60, 'newOne', '12', '5', 'newOne', 'https://image.tmdb.org/t/p/original/8f9dnOtpArDrOMEylpSN9Sc6fuz.jpg', '1', '01-23-2022'),
(71, 'qwrty12', '12', '5', 'newOne', 'https://image.tmdb.org/t/p/original/8f9dnOtpArDrOMEylpSN9Sc6fuz.jpg', '1', '01-23-2022');

-- --------------------------------------------------------

--
-- Структура таблицы `book_genre`
--

CREATE TABLE `book_genre` (
  `id` int NOT NULL,
  `id_book` int NOT NULL,
  `id_genre` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `book_genre`
--

INSERT INTO `book_genre` (`id`, `id_book`, `id_genre`) VALUES
(10, 46, 1),
(11, 46, 2),
(12, 47, 1),
(13, 47, 2),
(24, 58, 2),
(25, 58, 4),
(30, 60, 3),
(31, 60, 6),
(33, 71, 5),
(34, 71, 6);

-- --------------------------------------------------------

--
-- Структура таблицы `genres`
--

CREATE TABLE `genres` (
  `id_genre` int NOT NULL,
  `book_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `id_book` int DEFAULT NULL,
  `genre_of_book` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Структура таблицы `genres2`
--

CREATE TABLE `genres2` (
  `id_genre` int NOT NULL,
  `name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `genres2`
--

INSERT INTO `genres2` (`id_genre`, `name`) VALUES
(1, 'horror'),
(2, 'fantasy'),
(3, 'advanture'),
(4, 'detective'),
(5, 'thriller'),
(6, 'historical'),
(7, 'biography');

-- --------------------------------------------------------

--
-- Структура таблицы `reeders`
--

CREATE TABLE `reeders` (
  `id_reeder` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `surname` varchar(50) NOT NULL,
  `date_of_birth` varchar(20) NOT NULL DEFAULT '0',
  `address` varchar(50) NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `reeders`
--

INSERT INTO `reeders` (`id_reeder`, `name`, `surname`, `date_of_birth`, `address`, `email`) VALUES
(12, 'tttt', 'tttt', 'tttt', 'tttt', 'tttt'),
(13, 'Polina', 'Nalivayko', '28.04.2001', 'Minsk', 'nalivayko0428@gmail.com'),
(17, 'Ivan', 'Ivanov', '12.12.2012', 'minsk', 'ivanov@gmail.com'),
(19, 'Van', 'Ivanov', '12.12.2012', 'minsk', 'ivanov11@gmail.com'),
(20, 'Vanya', '11', '12.12.2012', 'minsk', 'ivanov111@gmail.com');

-- --------------------------------------------------------

--
-- Структура таблицы `rent`
--

CREATE TABLE `rent` (
  `id_rent` int NOT NULL,
  `id_book` int NOT NULL,
  `id_reeder` int NOT NULL,
  `first_date` varchar(20) NOT NULL,
  `last_date` varchar(20) NOT NULL,
  `fine` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `completed` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `rent`
--

INSERT INTO `rent` (`id_rent`, `id_book`, `id_reeder`, `first_date`, `last_date`, `fine`, `completed`) VALUES
(6, 20, 13, '12-29-2021', '01-28-2022', '0', 'completed'),
(7, 20, 13, '12-31-2021', '01-30-2022', '0', 'completed'),
(9, 26, 13, '01-13-2022', '02-12-2022', '0', '0'),
(10, 22, 13, '01-13-2021', '02-12-2021', '0', '0'),
(12, 26, 13, '01-19-2022', '02-18-2022', '0', 'completed'),
(14, 50, 13, '01-20-2022', '02-19-2022', '0', 'completed');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id_book`),
  ADD UNIQUE KEY `id` (`id_book`),
  ADD KEY `id_book` (`id_book`);

--
-- Индексы таблицы `book_genre`
--
ALTER TABLE `book_genre`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_book` (`id_book`),
  ADD KEY `id_genre` (`id_genre`);

--
-- Индексы таблицы `genres`
--
ALTER TABLE `genres`
  ADD PRIMARY KEY (`id_genre`),
  ADD KEY `id_book` (`id_book`) USING BTREE;
ALTER TABLE `genres` ADD FULLTEXT KEY `book_name` (`book_name`);

--
-- Индексы таблицы `genres2`
--
ALTER TABLE `genres2`
  ADD PRIMARY KEY (`id_genre`);

--
-- Индексы таблицы `reeders`
--
ALTER TABLE `reeders`
  ADD PRIMARY KEY (`id_reeder`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `id_reeder` (`id_reeder`);

--
-- Индексы таблицы `rent`
--
ALTER TABLE `rent`
  ADD PRIMARY KEY (`id_rent`),
  ADD KEY `id_book` (`id_book`),
  ADD KEY `id_reeder` (`id_reeder`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `books`
--
ALTER TABLE `books`
  MODIFY `id_book` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=72;

--
-- AUTO_INCREMENT для таблицы `book_genre`
--
ALTER TABLE `book_genre`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;

--
-- AUTO_INCREMENT для таблицы `genres`
--
ALTER TABLE `genres`
  MODIFY `id_genre` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT для таблицы `genres2`
--
ALTER TABLE `genres2`
  MODIFY `id_genre` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT для таблицы `reeders`
--
ALTER TABLE `reeders`
  MODIFY `id_reeder` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT для таблицы `rent`
--
ALTER TABLE `rent`
  MODIFY `id_rent` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `book_genre`
--
ALTER TABLE `book_genre`
  ADD CONSTRAINT `book_genre_ibfk_1` FOREIGN KEY (`id_book`) REFERENCES `books` (`id_book`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `book_genre_ibfk_2` FOREIGN KEY (`id_genre`) REFERENCES `genres2` (`id_genre`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Ограничения внешнего ключа таблицы `genres`
--
ALTER TABLE `genres`
  ADD CONSTRAINT `genres_ibfk_1` FOREIGN KEY (`id_book`) REFERENCES `books` (`id_book`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `rent`
--
ALTER TABLE `rent`
  ADD CONSTRAINT `rent_ibfk_1` FOREIGN KEY (`id_book`) REFERENCES `books` (`id_book`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `rent_ibfk_2` FOREIGN KEY (`id_reeder`) REFERENCES `reeders` (`id_reeder`) ON DELETE RESTRICT ON UPDATE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
