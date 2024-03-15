CREATE OR REPLACE FUNCTION check_ad(title_i TEXT, start_at_i TIMESTAMP, end_at_i TIMESTAMP, age_start_i INTEGER, age_end_i INTEGER, gender_i gender_type)
RETURNS TABLE (
    ad_id INTEGER,
    title TEXT,
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    age_start INTEGER,
    age_end INTEGER,
    gender gender_type,
    country_name country_code[],
    platform_name platform_type[]
) AS
$$
    DECLARE
        sql_statement TEXT;
    BEGIN
        sql_statement := 'SELECT DISTINCT 
            ad.ad_id, ad.title, 
            ad.start_at, ad.end_at, 
            ad.age_start, ad.age_end, 
            ad.gender,
            ARRAY_AGG(DISTINCT country_name) AS country_name, 
            ARRAY_AGG(DISTINCT platform_name) AS platform_name 
        FROM ad LEFT JOIN country ON ad.ad_id = country.ad_id 
        LEFT JOIN platform ON ad.ad_id = platform.ad_id '; 

        sql_statement := sql_statement || 
        ' WHERE ad.title = ' || quote_literal(title_i) ||
        ' AND ad.start_at = ' || quote_literal(start_at_i) ||
        ' AND ad.end_at = ' || quote_literal(end_at_i) ||
        ' AND ad.age_start = ' || age_start_i ||
        ' AND ad.age_end = ' || age_end_i ||
        ' AND ad.gender = ' || quote_literal(gender_i) ||
        ' GROUP BY ad.ad_id ORDER BY ad.ad_id ';

        RETURN Query EXECUTE sql_statement;
    END;
$$
LANGUAGE plpgsql;

DROP Function check_ad