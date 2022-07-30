package unit_tests

import (
	"MeLi/controllers"
	singleton "MeLi/singletons"
	"testing"
)

func TestGetLocation(t *testing.T) {
	kenobi := singleton.Satellite{Name: "Kenobi", CoordinateX: 10, CoordinateY: 0}
	skywalker := singleton.Satellite{Name: "Skywalker", CoordinateX: 0, CoordinateY: 10}
	sato := singleton.Satellite{Name: "Sato", CoordinateX: -10, CoordinateY: 0}
	var satellites [3]singleton.Satellite
	satellites[0] = kenobi
	satellites[1] = skywalker
	satellites[2] = sato
	coordenate := singleton.GetInstance()
	coordenate.SetSatellites(satellites[:])
	var x float32
	var y float32
	x, y = controllers.GetLocation(float32(14.14213562), float32(20), float32(14.14213562))
	if x != 0 && y != -10 {
		t.Error("El calculo no es correcto")
		t.Fail()
	} else {
		t.Log("TestGetLocation finalizado correctamente")
	}
}

func TestGetDataOrdered(t *testing.T) {
	names := []string{"sato", "kenobi", "skywalker"}
	distances := []float32{10, 5, 30}
	messages := [3][]string{{"es"}, {"el"}, {"orden"}}

	distancesOrdered := [3]float32{5, 30, 10}
	messagesOrdered := [3][]string{{"el"}, {"orden"}, {"es"}}

	distance, message := controllers.GetDataOrdered(names, distances, messages)

	if distance != distancesOrdered && message[0][0] != messagesOrdered[0][0] &&
		message[1][1] != messagesOrdered[1][1] && message[2][2] != messagesOrdered[2][2] {
		t.Error("El calculo no es correcto")
		t.Fail()
	} else {
		t.Log("TestGetDataOrdered finalizado correctamente")
	}
}

func TestGetMessage(t *testing.T) {
	mensaje1 := []string{"", "este", "es", "", "mensaje"}
	mensaje2 := []string{"", "este", "", "un", "mensaje"}
	mensaje3 := []string{"este", "", "un", "", "secreto", "muy", "importante"}

	result := controllers.GetMessage(mensaje1, mensaje2, mensaje3)

	if result != "este es un mensaje secreto muy importante" {
		t.Error("El mensaje no es correcto")
		t.Fail()
	} else {
		t.Log("TestGetMessage finalizado correctamente")
	}
}
