# ICA05

## Websrv    
Webserveren bruker index.html filen som ligger i mappen "static", og viser denne for brukeren.
Oppsett er relativt simpelt, clone mappen og kjør websrv.go så vil den starte en server på port 8001.   

## Webclient
Main-fila inneholder URL-ene og kjører funksjonene i de andre pakkene som go-rutiner. Hver av de fem pakkene har en go-fil som henter data fra API-en, mellomlagrer den og skriver ut dataene i terminalen, dataene blir mellomlagret som én json- og én txt-fil. Structen er lik og forklaring av hva som skjer i koden finner du her: 
[Generell Import/Eksport - Webclient](https://github.com/Gruppe12IS105/ICA05/blob/master/Webcli/generell_IE/importEksport.go) 

En liten digresjon: Hvis det er en array av structs (vil se slik ut: []struct) må man bruke “Decoder” og iterere over innholdet med en for-løkke, ellers er det “Unmarshal” som brukes. I denne webklienten anvendes kun Unmarshal-funksjonen. 
