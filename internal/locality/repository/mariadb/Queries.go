package repository

const (
	QueryGetById = `
        select l.id as locality_id, c.country_name, p.province_name , l.locality_name  
        from countries c 
        inner join provinces p on c.id = p.country_id
        inner join localities l on p.id = l.province_id 
        where l.id = ?
    `

	QueryCarryReport = `
    select 
    c.locality_id,
    l.locality_name,
    count(*) as total_locality 
    from carriers c
    left join localities l on l.id = c.locality_id
    left join provinces p on p.id = l.province_id 
    left join countries ct on ct.id = p.country_id 
    where c.locality_id = case when ? = 0 then locality_id else ? end
    group by c.locality_id 
    `

	QuerryCreateLocality = `
    INSERT INTO localities (locality_name, province_name, country_name, province_id)
    VALUES (?, ?, ?, ?)`

	QueryGetAllLocality = `
    SELECT l.id, l.locality_name, count(s.id) as sellers_count
	FROM localities as l
	LEFT JOIN sellers as s on l.id = s.locality_id
	GROUP BY l.id
	`
)
