# Network Rail 

Service for making SOAP requests to the Network Rail API in Go

### Resources
[Documentation for Network Rail Live Departure Board Web Service (LDBWS)](http://lite.realtime.nationalrail.co.uk/openldbws/)

[Register for Access / Token](https://opendata.nationalrail.co.uk/)

#### Example SOAP request
```
    <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://thalesgroup.com/RTTI/2017-10-01/ldb/">
        <soapenv:Header>
            <AccessToken xmlns="http://thalesgroup.com/RTTI/2013-11-28/Token/types">
                <TokenValue>[INSERTTOKEN]</TokenValue>
            </AccessToken>
        </soapenv:Header>
        <soapenv:Body>
            <ser:GetArrBoardWithDetailsRequest>
                <ser:numRows>10</ser:numRows>
                <ser:crs>SAJ</ser:crs>
                <ser:filterCrs>LBG</ser:filterCrs>
                <ser:filterType>to</ser:filterType>
            </ser:GetArrBoardWithDetailsRequest>
        </soapenv:Body>
    </soapenv:Envelope>
```


