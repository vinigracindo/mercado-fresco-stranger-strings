package repository

const (
	QueryGetById = `
    select l.id as locality_id, c.country_name, p.province_name , l.locality_name  
    from countries c 
    inner join provinces p on c.id = p.country_id
    inner join localities l on p.id = l.province_id 
    where l.id = ?
    `
)
