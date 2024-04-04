package database

import m "strore/models"

var UserDb = make(map [string] m.Person)

var ProductsDb = make(map [m.Product] int)

var AdminsDb = make(map [m.Admin] m.Admin)
