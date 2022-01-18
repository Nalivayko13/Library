-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Янв 18 2022 г., 16:41
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
  `genre` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default',
  `price_of_book` varchar(50) DEFAULT '0',
  `num_of_copies` varchar(50) DEFAULT '0',
  `authors` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default',
  `cover_photo` varchar(50) DEFAULT 'default',
  `price_per_day` varchar(50) DEFAULT '0',
  `reg_date` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `books`
--

INSERT INTO `books` (`id_book`, `name`, `genre`, `price_of_book`, `num_of_copies`, `authors`, `cover_photo`, `price_per_day`, `reg_date`) VALUES
(20, 'reg', 'reg', '4', '4', 'reg', 'default', '4', '12-20-2021'),
(22, '100', '100', '100', '99', '100', 'default', '100', '12-21-2021'),
(25, 'newBook', 'genre', '123', '3', 'authors', 'default', '1', '12-28-2021'),
(26, 'refactor', 'refactor', '123', '8', '4', 'default', '2', '12-29-2021'),
(39, 'dog3', 'default', '1000000', '121', 'testinggenre', 'https://vistapointe.net/images/cane-corso-2.jpg', '1', '01-18-2022');

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

--
-- Дамп данных таблицы `genres`
--

INSERT INTO `genres` (`id_genre`, `book_name`, `id_book`, `genre_of_book`) VALUES
(28, 'dog3', NULL, 'testgenreaaaa');

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
(13, 'Polina', 'Nalivayko', '28.04.2001', 'Minsk', 'nalivayko0428@gmail.com');

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
(10, 22, 13, '01-13-2021', '02-12-2021', '0', '0');

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
-- Индексы таблицы `genres`
--
ALTER TABLE `genres`
  ADD PRIMARY KEY (`id_genre`),
  ADD KEY `id_book` (`id_book`) USING BTREE;
ALTER TABLE `genres` ADD FULLTEXT KEY `book_name` (`book_name`);

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
  MODIFY `id_book` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- AUTO_INCREMENT для таблицы `genres`
--
ALTER TABLE `genres`
  MODIFY `id_genre` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT для таблицы `reeders`
--
ALTER TABLE `reeders`
  MODIFY `id_reeder` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT для таблицы `rent`
--
ALTER TABLE `rent`
  MODIFY `id_rent` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

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
