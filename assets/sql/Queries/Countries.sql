WITH track_to_countries_mixed AS
(
    SELECT      *
    FROM        "SHAH.S".track_to_countries
    
    UNION ALL
    
    SELECT      *
    FROM        "CARLO.QUICK".track_to_countries
    
    UNION ALL
    
    SELECT      *
    FROM        "SKULTHUM.RASHID".track_to_countries
    
    UNION ALL
    
    SELECT      *
    FROM        "DDEXTER".track_to_countries
)


// dsfsdf
SELECT      t.release_year
FROM        "SHAH.S".tracks t, track_to_countries_mixed ttc, "SHAH.S".country_to_code ctc, "SHAH.S".countries c
WHERE       t.track_id = ttc.track_id AND
            ttc.code = ctc.code AND
            ctc.country = c.name
HAVING      c.internet_use = MIN(c.internet_use)
GROUP BY    t.release_year, c.internet_use;