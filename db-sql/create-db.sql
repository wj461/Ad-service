
CREATE TABLE IF NOT EXISTS ad (
    ad_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    age_start INTEGER,
    age_end INTEGER,
    male BOOLEAN,
    female BOOLEAN
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

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'country_code') THEN
        CREATE TYPE country_code AS ENUM (
            'AF', 'AX', 'AL', 'DZ', 'AS', 'AD', 'AO', 'AI', 'AQ', 'AG', 
            'AR', 'AM', 'AW', 'AU', 'AT', 'AZ', 'BS', 'BH', 'BD', 'BB', 
            'BY', 'BE', 'BZ', 'BJ', 'BM', 'BT', 'BO', 'BQ', 'BA', 'BW', 
            'BV', 'BR', 'IO', 'BN', 'BG', 'BF', 'BI', 'CV', 'KH', 'CM', 
            'CA', 'KY', 'CF', 'TD', 'CL', 'CN', 'CX', 'CC', 'CO', 'KM', 
            'CD', 'CG', 'CK', 'CR', 'CI', 'HR', 'CU', 'CW', 'CY', 'CZ', 
            'DK', 'DJ', 'DM', 'DO', 'EC', 'EG', 'SV', 'GQ', 'ER', 'EE', 
            'SZ', 'ET', 'FK', 'FO', 'FJ', 'FI', 'FR', 'GF', 'PF', 'TF', 
            'GA', 'GM', 'GE', 'DE', 'GH', 'GI', 'GR', 'GL', 'GD', 'GP', 
            'GU', 'GT', 'GG', 'GN', 'GW', 'GY', 'HT', 'HM', 'VA', 'HN', 
            'HK', 'HU', 'IS', 'IN', 'ID', 'IR', 'IQ', 'IE', 'IM', 'IL', 
            'IT', 'JM', 'JP', 'JE', 'JO', 'KZ', 'KE', 'KI', 'KP', 'KR', 
            'KW', 'KG', 'LA', 'LV', 'LB', 'LS', 'LR', 'LY', 'LI', 'LT', 
            'LU', 'MO', 'MG', 'MW', 'MY', 'MV', 'ML', 'MT', 'MH', 'MQ', 
            'MR', 'MU', 'YT', 'MX', 'FM', 'MD', 'MC', 'MN', 'ME', 'MS', 
            'MA', 'MZ', 'MM', 'NA', 'NR', 'NP', 'NL', 'NC', 'NZ', 'NI', 
            'NE', 'NG', 'NU', 'NF', 'MK', 'MP', 'NO', 'OM', 'PK', 'PW', 
            'PS', 'PA', 'PG', 'PY', 'PE', 'PH', 'PN', 'PL', 'PT', 'PR', 
            'QA', 'RO', 'RU', 'RW', 'BL', 'SH', 'KN', 'LC', 'MF', 'PM', 
            'VC', 'WS', 'SM', 'ST', 'SA', 'SN', 'RS', 'SC', 'SL', 'SG', 
            'SX', 'SK', 'SI', 'SB', 'SO', 'ZA', 'GS', 'SS', 'ES', 'LK', 
            'SD', 'SR', 'SJ', 'SE', 'CH', 'SY', 'TW', 'TJ', 'TZ', 'TH', 
            'TL', 'TG', 'TK', 'TO', 'TT', 'TN', 'TR', 'TM', 'TC', 'TV', 
            'UG', 'UA', 'AE', 'GB', 'UM', 'US', 'UY', 'UZ', 'VU', 'VE', 
            'VN', 'VG', 'VI', 'WF', 'EH', 'YE', 'ZM', 'ZW', ''
        );
    END IF;
END$$;

CREATE TYPE platform_type AS ENUM (
    'android', 'ios', 'web', ''
);

DROP TABLE country;
DROP TABLE platform;
DROP TABLE ad;


DROP FUNCTION search_ads;

DROP TYPE country_code;
DROP TYPE platform_type;