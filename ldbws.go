package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type TrainData struct {
	XMLName xml.Name
	Body    TrainDataBody
}

type TrainDataBody struct {
	XMLName xml.Name
	Body    StationBoardResult `xml:"GetArrBoardWithDetailsResponse"`
}

type StationBoardResult struct {
	XMLName xml.Name
	Details Details `xml:"GetStationBoardResult"`
}

type Details struct {
	GeneratedAt   string        `xml:"generatedAt"`
	StationName   string        `xml:"locationName"`
	TrainServices TrainServices `xml:"trainServices"`
}

type TrainServices struct {
	Service []ServiceDetails `xml:"service"`
}

type ServiceDetails struct {
	Sta         string   `xml:"sta"`
	Eta         string   `xml:"eta"`
	Platform    string   `xml:"platform"`
	Origin      Location `xml:"origin"`
	Destination Location `xml:"destination"`
}

type Location struct {
	Location LocationName `xml:"location"`
}

type LocationName struct {
	LocationName string `xml:"locationName"`
}

func main() {
	// wsdl service url
	url := fmt.Sprintf("%s%s%s%s",
		"https://lite.realtime.nationalrail.co.uk",
		"/OpenLDBWS",
		"/ldb11",
		".asmx",
	)

	envelope := "<soapenv:Envelope xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/' xmlns:ser='http://thalesgroup.com/RTTI/2017-10-01/ldb/'>"

	tokenValue := "INSERT_TOKEN"
	header := "<soapenv:Header><AccessToken xmlns='http://thalesgroup.com/RTTI/2013-11-28/Token/types'><TokenValue>" + tokenValue + "</TokenValue></AccessToken></soapenv:Header>"

	currentStation := "SAJ"
	targetStation := "LBG"
	body := "<soapenv:Body><ser:GetArrBoardWithDetailsRequest><ser:numRows>10</ser:numRows><ser:crs>" + currentStation + "</ser:crs><ser:filterCrs>" + targetStation + "</ser:filterCrs><ser:filterType>to</ser:filterType></ser:GetArrBoardWithDetailsRequest></soapenv:Body></soapenv:Envelope>"

	payload := []byte(strings.TrimSpace(envelope + header + body))

	httpMethod := "POST"

	log.Println("-------> Preparing the request")

	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))

	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}
	// set headers
	req.Header.Set("Content-type", "text/xml")

	log.Printf("\n[*REQ*]: %+v\n", req)

	// prepare the client req
	client := &http.Client{}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)

	result := new(TrainData)
	log.Printf("result: %+v\n", result)

	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	log.Println("-> Everything is good, printing users data")

	// print the users data
	fmt.Printf("\nresult.Body.Body.Details.TrainServices.Service: %+v\n", result.Body.Body.Details.TrainServices.Service)
	fmt.Printf("\nTrain data.body %+v", result.Body)
}
