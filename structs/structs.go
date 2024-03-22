package structs



type Create struct{
	Id string  `bson:"_id"`
	Name string
	Surname string
	Login string
	Password string
}
type Student struct{
	Id string  `bson:"_id"`
	StudentName string
	StudentSurname string
	StudentEmail string
	StudentPhone int
	StudentLogin string
	StudentPassword string
}
type Forproject struct{
	Id string  `bson:"_id"`
	Project_name string
	Discription string
}
type JoinStudent struct{
	Id string  `bson:"_id"`
	Project_id string 
	Student_email string
	Owner_id string 
	Student_phone int
}
type DeleteStudent struct{
	Student_id string 
}
type DeleteProject struct{
	Project_id string  
}
type OneProjectList struct{
	Project Forproject
	Join_students []JoinStudent  
}