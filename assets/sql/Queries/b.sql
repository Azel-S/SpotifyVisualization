SELECT  ry1, reg1, len1, reg2, len2
FROM    (
        SELECTt.release_year ry1, c.region reg1, median(length(t.title)) len1
		from "SHAH.S".tracks t
		join "SHAH.S".track_to_countries p on t.track_id = p.track_id
		join "SHAH.S".country_to_code s on p.code = s.code
		join "SHAH.S".countries c on s.country = c.name
		where c.region = :4 AND 
			  (t.release_year >= :1 AND
			  t.release_year <= :2 )
			  
		group by t.release_year, c.region
		order by t.release_year),
        
        (select t.release_year ry2, c.region reg2, median(length(t.title)) len2
		from "SHAH.S".tracks t
		join "SHAH.S".track_to_countries p on t.track_id = p.track_id
		join "SHAH.S".country_to_code s on p.code = s.code
		join "SHAH.S".countries c on s.country = c.name
		where c.region = :3 AND 
			  (t.release_year >= :1 AND
			  t.release_year <= :2 )
			  
		group by t.release_year, c.region
		order by t.release_year) WHERE ry1 = ry2;