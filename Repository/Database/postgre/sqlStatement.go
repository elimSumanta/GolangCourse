package database

const (
	SelectCarByIDGerage = `
	SELECT 
		ownerName."ownerName", 
		cars."carName", 
		cars."idCar"  
	FROM 
		"tbl_car" cars  
	INNER JOIN 
		"tbl_Owner_Mobil" ownerName	
	on 
		cars."idCar" = ownerName."idCar" 
	where 
		ownerName."idGerage" =$1`

	SelectCarByIDCar = `
	SELECT 
		ownerName."ownerName", 
		cars."carName", 
		cars."idCar"  
	FROM 
		"tbl_car" as cars  
	INNER JOIN 
		"tbl_Owner_Mobil" as ownerName	
	on 
		cars."idCar" = ownerName."idCar" 
	where cars."idCar" =$1`
)
