module frostwagner/todolist

go 1.13

replace frostwagner/route => ../route

replace frostwagner/dbhandle => ../dbhandle

replace frostwagner/structures => ../structures

require (
	frostwagner/dbhandle v0.0.0-00010101000000-000000000000
	frostwagner/route v0.0.0-00010101000000-000000000000
)
