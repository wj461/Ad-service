INSERT INTO ad (title, start_at, end_at, age_start, age_end, male, female) VALUES
    ('Ad 1', '2024-03-09 12:00:00', '2024-08-10 12:00:00', 18, 30, true, false),
    ('Ad 2', '2024-03-10 12:00:00', '2024-08-11 12:00:00', 25, 40, false, true),
    ('Ad 3', '2024-03-11 12:00:00', '2024-08-12 12:00:00', 20, 35, true, true),
    ('Ad 4', '2024-03-12 12:00:00', '2024-08-13 12:00:00', 30, 45, true, true),
    ('Ad 5', '2024-03-13 12:00:00', '2024-08-14 12:00:00', 18, 25, true, false);

INSERT INTO country (ad_id, country_name) VALUES (1, 'TW');
INSERT INTO country (ad_id, country_name) VALUES (1, 'JP');
INSERT INTO platform (ad_id, platform_name) VALUES (1, 'android');

INSERT INTO country (ad_id, country_name) VALUES (2, 'TW');
INSERT INTO platform (ad_id, platform_name) VALUES (2, 'ios');

INSERT INTO country (ad_id, country_name) VALUES (4, 'JP');

INSERT INTO platform (ad_id, platform_name) VALUES (5, 'ios');