package model

type Admin struct {
	username string;
	password string;
	__name string;
	__surname string;
	__phone_number string;
}

func (admin Admin) RepresentAdmin () string {
	return admin.__name + " " + admin.__surname;
}
