// Get Popularity
SELECT      t.release_year, MAX(t.popularity)
FROM        tracks t
WHERE       t.release_year >= :1 AND // :1 is the first passed in paramater
            t.release_year <= :2     // :2 is the second passed in parameter
GROUP BY    t.release_year
ORDER BY    t.release_year ASC