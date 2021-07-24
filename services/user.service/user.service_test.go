package user_service_test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"

	m "../../models"
	userService "../user.service"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID: 		oid,
		Name:      "Jhon",
		Email:     "jhon@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {
		t.Error("La prueba de la persistencia a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("Se ha presentado un error en la consulta de usuarios")
		t.Fail()
	}

	if len(users) == 0{
		t.Error("La consulta no retorno datos")
		t.Fail()
	}else{
		t.Log("La prueba ha finalizado con exito")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name: "Jhon Oquendo",
		Email: "jhon.oquendo@gmail.com",
	}

	err := userService.Update(user,userId)

	if err != nil {
		t.Error("Error al tratar de actualizar usuario")
		t.Fail()
	}else {
		t.Log("La prueba de actualizacion finalizo con exito")
	}
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)

	if err != nil {
		t.Error("Error al tratar de eliminar usuario")
		t.Fail()
	}else{
		t.Log("La prueba de eliminacion finalizo con exito")
	}
}
