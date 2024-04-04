package model

type Admin struct {
	Username     string
	Password     string
	Name         string
	Surname      string
	Phone_number string
}

func (admin Admin) RepresentAdmin() string {
	return admin.Name + " " + admin.Surname
}
