CREATE TABLE stickers
(
    id int PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    priority int NOT NULL,
    packid int,
    startTime time,
    endTime time
);

INSERT INTO stickers
    (id, name, priority, startTime, endTime, packid)
VALUES
    ('1', 'A1', '1', '08:00:00', '11:59:59', '1'),
    ('2', 'A2', '1', '12:00:00', '03:59:59', '2'),
    ('3', 'A3', '1', '04:00:00', '23:00:00', '3'),
    ('4', 'B1', '2', '08:00:00', '11:59:59', '1'),
    ('5', 'B2', '2', '12:00:00', '03:59:59', '2'),
    ('6', 'B3', '2', '04:00:00', '23:00:00', '3'),
    ('7', 'C1', '3', '08:00:00', '11:59:59', '1'),
    ('8', 'C2', '3', '12:00:00', '03:59:59', '2'),
    ('9', 'C3', '3', '04:00:00', '23:00:00', '3');