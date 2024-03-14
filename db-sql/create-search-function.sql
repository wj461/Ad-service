CREATE or REPLACE FUNCTION search_ads(age INTEGER, gender TEXT, country country_code, platform platform_type , offset_val INTEGER default 0, limit_val INTEGER default 5)
RETURNS TABLE (
    ad_id INTEGER,
    title TEXT,
    end_at TIMESTAMP
) AS
$$
    DECLARE
        sql_statement TEXT;
    BEGIN
        if limit_val IS NULL THEN
            limit_val := 5;
        END IF;
        if offset_val IS NULL THEN
            offset_val := 0;
        END IF;

        sql_statement := 'SELECT DISTINCT ad.ad_id, ad.title, ad.end_at FROM ad LEFT JOIN country ON ad.ad_id = country.ad_id LEFT JOIN platform ON ad.ad_id = platform.ad_id WHERE start_at <= now() AND end_at >= now()';
        IF gender = 'M' THEN
            sql_statement := sql_statement || ' AND male = true ';
        ELSIF gender = 'F' THEN
            sql_statement := sql_statement || ' AND female = true ';
        END IF;

        IF age IS NOT NULL THEN
            sql_statement := sql_statement || ' AND age_start <= ' || age || ' AND age_end >= ' || age ;
        END IF;

        IF country IS NOT NULL THEN
            sql_statement := sql_statement || ' AND (country_name = ' || quote_literal(country) || ' OR country_name IS NULL) ';
        END IF;

        IF platform IS NOT NULL THEN
            sql_statement := sql_statement || ' AND (platform_name = ' || quote_literal(platform) || ' OR platform_name IS NULL) ';
        END IF;

        sql_statement := sql_statement || ' ORDER BY ad.end_at ASC' || ' OFFSET ' || offset_val || ' LIMIT ' || limit_val;

        RETURN Query EXECUTE Sql_statement;
    END;
$$
LANGUAGE plpgsql;

