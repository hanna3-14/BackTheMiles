DROP TABLE IF EXISTS distances;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS eventResults;
DROP TABLE IF EXISTS goals;
DROP TABLE IF EXISTS results;

CREATE TABLE IF NOT EXISTS distances (
	name TEXT,
	distanceInMeters INT
);

CREATE TABLE IF NOT EXISTS events (
	name TEXT,
	location TEXT
);

CREATE TABLE IF NOT EXISTS eventResults (
	event_id INT,
	result_id INT
);

CREATE TABLE IF NOT EXISTS goals (
	distance TEXT,
	time TEXT
);

CREATE TABLE IF NOT EXISTS results (
	event_id INT,
	date TEXT,
	distance_id INT,
	time_gross_hours INT,
	time_gross_minutes INT,
	time_gross_seconds INT,
	time_net_hours INT,
	time_net_minutes INT,
	time_net_seconds INT,
	category TEXT,
	agegroup TEXT,
	place_total INT,
	place_category INT,
	place_agegroup INT,
	finisher_total INT,
	finisher_category INT,
	finisher_agegroup INT
);

INSERT INTO distances (name, distanceInMeters) VALUES ('Marathon', 42195);
INSERT INTO distances (name, distanceInMeters) VALUES ('Half Marathon', 21097);
INSERT INTO distances (name, distanceInMeters) VALUES ('10K', 10000);
INSERT INTO distances (name, distanceInMeters) VALUES ('5K', 5000);
INSERT INTO distances (name, distanceInMeters) VALUES ('Badische Meile', 8889);

INSERT INTO events (name, location) VALUES ('Baden-Marathon', 'Karlsruhe');
INSERT INTO events (name, location) VALUES ('Schwarzwald-Marathon', 'Bräunlingen');
INSERT INTO events (name, location) VALUES ('Bienwald-Marathon', 'Kandel');
INSERT INTO events (name, location) VALUES ('Freiburg Marathon', 'Freiburg');

INSERT INTO goals (distance, time) VALUES ('Marathon', '4:59:59');
INSERT INTO goals (distance, time) VALUES ('Half Marathon', '1:59:59');

INSERT INTO results (event_id, date, distance_id, time_gross_hours, time_gross_minutes, time_gross_seconds, time_net_hours, time_net_minutes, time_net_seconds, category, agegroup, place_total, place_category, place_agegroup, finisher_total, finisher_category, finisher_agegroup) VALUES (1, '2022-09-18', 2, 2, 27, 10, 2, 21, 40, 'W', 'W', 2336, 626, 195, 2664, 809, 230);
INSERT INTO results (event_id, date, distance_id, time_gross_hours, time_gross_minutes, time_gross_seconds, time_net_hours, time_net_minutes, time_net_seconds, category, agegroup, place_total, place_category, place_agegroup, finisher_total, finisher_category, finisher_agegroup) VALUES (2, '2022-10-09', 2, 2, 10, 50, 2, 9, 45, 'W', 'W', 535, 151, 25, 664, 216, 39);
INSERT INTO results (event_id, date, distance_id, time_gross_hours, time_gross_minutes, time_gross_seconds, time_net_hours, time_net_minutes, time_net_seconds, category, agegroup, place_total, place_category, place_agegroup, finisher_total, finisher_category, finisher_agegroup) VALUES (3, '2023-03-12', 2, 2, 11, 17, 2, 9, 14, 'W', 'W', 928, 226, 74, 1148, 340, 109);
INSERT INTO results (event_id, date, distance_id, time_gross_hours, time_gross_minutes, time_gross_seconds, time_net_hours, time_net_minutes, time_net_seconds, category, agegroup, place_total, place_category, place_agegroup, finisher_total, finisher_category, finisher_agegroup) VALUES (4, '2023-03-26', 1, 5, 45, 57, 5, 29, 9, 'W', 'W', 916, 164, 60, 930, 166, 62);

INSERT INTO eventResults (event_id, result_id) VALUES (1, 1);
INSERT INTO eventResults (event_id, result_id) VALUES (2, 2);
INSERT INTO eventResults (event_id, result_id) VALUES (3, 3);
INSERT INTO eventResults (event_id, result_id) VALUES (4, 4);
