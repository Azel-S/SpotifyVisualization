WITH genre_1 AS
(
    SELECT      t.release_year, SUM(followers) followers_1
    FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
    WHERE       t.track_id = att.track_id AND
                att.artist_id = a.artist_id AND
                a.artist_id = atg.artist_id AND
                atg.genre = 'pop' AND
                t.release_year >= 2000 AND
                t.release_year <= 2005
    GROUP BY    t.release_year
    ORDER BY    t.release_year ASC
    ),
genre_2 AS
(
    SELECT      t.release_year, SUM(followers) followers_2
    FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
    WHERE       t.track_id = att.track_id AND
                att.artist_id = a.artist_id AND
                a.artist_id = atg.artist_id AND
                atg.genre = 'rock' AND
                t.release_year >= 2000 AND
                t.release_year <= 2005
    GROUP BY    t.release_year
)

SELECT      *
FROM        genre_1 NATURAL JOIN genre_2;