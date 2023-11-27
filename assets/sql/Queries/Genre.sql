SELECT      t.release_year, SUM(followers) followers
FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
WHERE       t.track_id = att.track_id AND
            att.artist_id = a.artist_id AND
            a.artist_id = atg.artist_id AND
            atg.genre = 'rock'
GROUP BY    t.release_year;