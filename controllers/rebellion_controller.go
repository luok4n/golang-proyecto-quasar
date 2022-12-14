package controllers

import (
	singleton "MeLi/singletons"
	"math"
	"sort"
	"strings"
)

//function to set the singleton class with the initial coordinates of each satellite
func SetCoordinates() {
	kenobi := singleton.Satellite{Name: "Kenobi", CoordinateX: -500, CoordinateY: -200}
	skywalker := singleton.Satellite{Name: "Skywalker", CoordinateX: 100, CoordinateY: -100}
	sato := singleton.Satellite{Name: "Sato", CoordinateX: 500, CoordinateY: 100}
	var satellites [3]singleton.Satellite
	satellites[0] = kenobi
	satellites[1] = skywalker
	satellites[2] = sato
	coordenate := singleton.GetInstance()
	coordenate.SetSatellites(satellites[:])
}

//function to get the location (x,y) given three distances with Trilateration
func GetLocation(distance1, distance2, distance3 float32) (x, y float32) {
	var p1 [2]float32
	var p2 [2]float32
	var p3 [2]float32
	var ex [2]float32
	var ey [2]float32
	var p3p1 [2]float32
	var jval float32 = 0
	var temp float32 = 0
	var ival float32 = 0
	var p3p1i float32 = 0
	var triptx float32 = 0
	var tripty float32 = 0
	var xval float32 = 0
	var yval float32 = 0
	var t1 float32 = 0
	var t2 float32 = 0
	var t3 float32 = 0
	var t float32 = 0
	var exx float32 = 0
	var d float32 = 0
	var eyy float32 = 0
	coordenate := singleton.GetInstance()
	p1[0] = coordenate.GetSatellites()[0].CoordinateX
	p1[1] = coordenate.GetSatellites()[0].CoordinateY
	p2[0] = coordenate.GetSatellites()[1].CoordinateX
	p2[1] = coordenate.GetSatellites()[1].CoordinateY
	p3[0] = coordenate.GetSatellites()[2].CoordinateX
	p3[1] = coordenate.GetSatellites()[2].CoordinateY

	for i := 0; i < len(p1); i++ {
		t1 = p2[i]
		t2 = p1[i]
		t = t1 - t2
		temp += (t * t)
	}
	d = float32(math.Sqrt(float64(temp)))
	for i := 0; i < len(p1); i++ {
		t1 = p2[i]
		t2 = p1[i]
		exx = (t1 - t2) / float32(math.Sqrt(float64(temp)))
		ex[i] = exx
	}
	for i := 0; i < len(p3); i++ {
		t1 = p3[i]
		t2 = p1[i]
		t3 = t1 - t2
		p3p1[i] = t3
	}
	for i := 0; i < len(ex); i++ {
		t1 = ex[i]
		t2 = p3p1[i]
		ival += (t1 * t2)
	}
	for i := 0; i < len(p3); i++ {
		t1 = p3[i]
		t2 = p1[i]
		t3 = ex[i] * ival
		t = t1 - t2 - t3
		p3p1i += (t * t)
	}
	for i := 0; i < len(p3); i++ {
		t1 = p3[i]
		t2 = p1[i]
		t3 = ex[i] * ival
		eyy = (t1 - t2 - t3) / float32(math.Sqrt(float64(p3p1i)))
		ey[i] = eyy
	}
	for i := 0; i < len(ey); i++ {
		t1 = ey[i]
		t2 = p3p1[i]
		jval += (t1 * t2)
	}
	xval = float32((math.Pow(float64(distance1), 2) - math.Pow(float64(distance2), 2) + math.Pow(float64(d), 2)) / (2 * float64(d)))
	yval = float32(((math.Pow(float64(distance1), 2) - math.Pow(float64(distance3), 2) + math.Pow(float64(ival), 2) + math.Pow(float64(jval), 2)) / (2 * float64(jval))) - ((float64(ival) / float64(jval)) * float64(xval)))
	t1 = coordenate.GetSatellites()[0].CoordinateX
	t2 = ex[0] * xval
	t3 = ey[0] * yval
	triptx = t1 + t2 + t3
	t1 = coordenate.GetSatellites()[0].CoordinateY
	t2 = ex[1] * xval
	t3 = ey[1] * yval
	tripty = t1 + t2 + t3

	return triptx, tripty
}

//Function to get the coordinates ordered in the next order "kenobi" "skywalker" "sato"
func GetDataOrdered(names []string, distance []float32, message [3][]string) (distanceOrdered [3]float32, messageOrdered [3][]string) {
	var order [3]float32
	var orderedMessage [3][]string
	for i := 0; i < len(names); i++ {
		if names[i] == "kenobi" {
			order[0] = distance[i]
			orderedMessage[0] = message[i]
		} else if names[i] == "skywalker" {
			order[1] = distance[i]
			orderedMessage[1] = message[i]
		} else if names[i] == "sato" {
			order[2] = distance[i]
			orderedMessage[2] = message[i]
		}
	}
	return order, orderedMessage
}

//Function to get/set the message and position by one satellite
func GetTopSecretOneSatellite(satellite singleton.SatelliteBody) (coordinates singleton.ResultPosition, message string) {
	satellites := singleton.GetInstance().GetSatellites()
	message = ""
	result := singleton.ResultPosition{}
	position := singleton.Positions{}

	setSatellite(satellite, satellites)

	for i := 0; i < len(satellites); i++ {
		if satellites[i].GetDistance() == 0 {
			return result, "No hay suficiente informacion para mostrar"
		}
	}

	position.SetXY(GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance))
	result.SetMessage(GetMessage(satellites[0].Message, satellites[1].Message, satellites[2].Message))
	result.SetPositions(position)
	return result, message
}

//function to get the location (x,y) and message given an array with three satellites
func GetTopSecret(satellites []singleton.SatelliteBody) (coordinates singleton.ResultPosition) {
	position := singleton.Positions{}
	result := singleton.ResultPosition{}
	var names [3]string
	var distances [3]float32
	var messages [3][]string
	for i := 0; i < len(satellites); i++ {
		names[i] = satellites[i].Name
		distances[i] = satellites[i].Distance
		messages[i] = satellites[i].Message
	}
	distanceOrder, messagesOrdered := GetDataOrdered(names[:], distances[:], messages)
	position.SetXY(GetLocation(distanceOrder[0], distanceOrder[1], distanceOrder[2]))
	result.SetMessage(GetMessage(messagesOrdered[0], messagesOrdered[1], messagesOrdered[2]))
	result.SetPositions(position)
	return result
}

//function to set the data of one satellite
func setSatellite(satellite singleton.SatelliteBody, satellites []singleton.Satellite) {
	for i := 0; i < len(satellites); i++ {
		if satellite.Name == satellites[i].Name {
			satellites[i].Distance = satellite.Distance
			satellites[i].Message = satellite.Message
		}
	}
}

//function to get a decode message given three []string messages
func GetMessage(message1, message2, message3 []string) (message string) {
	var message1Len = len(message1)
	var message2Len = len(message2)
	var message3Len = len(message3)
	maxLen := []int{message1Len, message2Len, message3Len}
	var messageList []string

	sort.Sort(sort.Reverse(sort.IntSlice(maxLen)))

	for i := 0; i < maxLen[0]; i++ {
		if message1Len > i {
			if message1[i] != "" {
				if len(messageList) == 0 {
					messageList = append(messageList, message1[i])
				} else if len(messageList) == 1 && messageList[len(messageList)-1] != message1[i] {
					messageList = append(messageList, message1[i])
				} else if len(messageList) >= 2 && messageList[len(messageList)-1] != message1[i] && messageList[len(messageList)-2] != message1[i] {
					messageList = append(messageList, message1[i])
				}

			}
		}
		if message2Len > i {
			if message2[i] != "" {
				if len(messageList) == 0 {
					messageList = append(messageList, message2[i])
				} else if len(messageList) == 1 && messageList[len(messageList)-1] != message2[i] {
					messageList = append(messageList, message2[i])
				} else if len(messageList) >= 2 && messageList[len(messageList)-1] != message2[i] && messageList[len(messageList)-2] != message2[i] {
					messageList = append(messageList, message2[i])
				}

			}
		}
		if message3Len > i {
			if message3[i] != "" {
				if len(messageList) == 0 {
					messageList = append(messageList, message3[i])
				} else if len(messageList) == 1 && messageList[len(messageList)-1] != message3[i] {
					messageList = append(messageList, message3[i])
				} else if len(messageList) >= 2 && messageList[len(messageList)-1] != message3[i] && messageList[len(messageList)-2] != message3[i] {
					messageList = append(messageList, message3[i])
				}

			}
		}
	}

	result := strings.Join(messageList, " ")

	return result
}
