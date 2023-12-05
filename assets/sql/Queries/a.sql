SELECT      *
FROM        (
            SELECT      t.release_year ry1, c.region, median(length(t.title))
			FROM        "SHAH.S".tracks t
			JOIN        "SHAH.S".track_to_countries p on t.track_id = p.track_id
			JOIN        "SHAH.S".country_to_code s on p.code = s.code
			JOIN        "SHAH.S".countries c on s.country = c.name
			WHERE       c.region = 'Asia' AND
                        t.release_year >= '1920' AND
                        t.release_year <= '1940'
				  
			GROUP BY    t.release_year, c.region
			ORDER BY    t.release_year), (select t.release_year ry2, c.region, median(length(t.title))),
            (
            SELECT      t.release_year ry2, c.region, median(length(t.title)) ml2
            FROM        "SHAH.S".tracks t
			JOIN        "SHAH.S".track_to_countries p on t.track_id = p.track_id
			JOIN        "SHAH.S".country_to_code s on p.code = s.code
			JOIN        "SHAH.S".countries c on s.country = c.name
			WHERE       c.region = 'Africa' AND 
                        t.release_year >= '1920' AND
                        t.release_year <= '1940'
            )
			GROUP BY    t.release_year, c.region
			ORDER BY    t.release_year WHERE ry1 = ry2