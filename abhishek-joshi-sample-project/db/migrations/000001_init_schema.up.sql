CREATE TABLE IF NOT EXISTS `stickers`
(
    `id` int PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `priority` int NOT NULL,
    `startTime` time,
    `endTime` time
);