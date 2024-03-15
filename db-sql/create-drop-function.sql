CREATE OR REPLACE FUNCTION reset_db()
RETURNS void AS 
$$
    BEGIN
        DROP TABLE IF EXISTS country;
        DROP TABLE IF EXISTS platform;
        DROP TABLE IF EXISTS ad;
    END;
$$ 
LANGUAGE plpgsql;
