CREATE OR REPLACE FUNCTION init_db()
RETURNS void AS
$$
    BEGIN
        CREATE TABLE IF NOT EXISTS ad (
            ad_id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            start_at TIMESTAMP NOT NULL,
            end_at TIMESTAMP NOT NULL,
            age_start INTEGER,
            age_end INTEGER,
            gender gender_type
        );

        CREATE TABLE IF NOT EXISTS country (
            ad_id SERIAL REFERENCES ad(ad_id),
            country_name country_code,
            PRIMARY KEY (ad_id, country_name)
        );

        CREATE TABLE IF NOT EXISTS platform (
            ad_id SERIAL REFERENCES ad(ad_id),
            platform_name platform_type,
            PRIMARY KEY (ad_id, platform_name)
        );

        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 1', '2024-02-09 12:00:00', '2024-08-10 12:00:00', 18, 30, 'M');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 2', '2024-02-10 12:00:00', '2024-08-11 12:00:00', 25, 40, 'F');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end) VALUES
            ('Ad 3', '2024-02-11 12:00:00', '2024-08-12 12:00:00', 20, 35);
        INSERT INTO ad (title, start_at, end_at, age_start, age_end) VALUES
            ('Ad 4', '2024-02-12 12:00:00', '2024-08-13 12:00:00', 30, 45);
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 5', '2024-02-13 12:00:00', '2024-08-14 12:00:00', 18, 99, 'M');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 6', '2024-02-14 12:00:00', '2024-08-15 12:00:00', 22, 38, 'M');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 7', '2024-02-15 12:00:00', '2024-08-16 12:00:00', 27, 42, 'F');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end) VALUES
            ('Ad 8', '2024-02-16 12:00:00', '2024-08-17 12:00:00', 32, 47);
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 9', '2024-02-17 12:00:00', '2024-08-18 12:00:00', 19, 28, 'M');
        INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES
            ('Ad 10', '2024-02-18 12:00:00', '2024-08-19 12:00:00', 21, 36, 'F');


        INSERT INTO country (ad_id, country_name) VALUES (1, 'TW');
        INSERT INTO country (ad_id, country_name) VALUES (1, 'JP');
        INSERT INTO platform (ad_id, platform_name) VALUES (1, 'android');

        INSERT INTO country (ad_id, country_name) VALUES (2, 'TW');
        INSERT INTO platform (ad_id, platform_name) VALUES (2, 'ios');

        INSERT INTO country (ad_id, country_name) VALUES (4, 'JP');

        INSERT INTO platform (ad_id, platform_name) VALUES (5, 'ios');

        INSERT INTO country (ad_id, country_name) VALUES (6, 'TW');

        INSERT INTO country (ad_id, country_name) VALUES (6, 'JP');
        INSERT INTO platform (ad_id, platform_name) VALUES (6, 'ios');

        INSERT INTO country (ad_id, country_name) VALUES (7, 'US');
        INSERT INTO platform (ad_id, platform_name) VALUES (7, 'android');

        INSERT INTO country (ad_id, country_name) VALUES (8, 'TW');
        INSERT INTO country (ad_id, country_name) VALUES (8, 'KR');
        INSERT INTO platform (ad_id, platform_name) VALUES (8, 'ios');

        INSERT INTO country (ad_id, country_name) VALUES (9, 'US');
        INSERT INTO country (ad_id, country_name) VALUES (9, 'JP');
        INSERT INTO platform (ad_id, platform_name) VALUES (9, 'android');

        INSERT INTO country (ad_id, country_name) VALUES (10, 'TW');
        INSERT INTO platform (ad_id, platform_name) VALUES (10, 'ios');

    END;
$$
LANGUAGE plpgsql;