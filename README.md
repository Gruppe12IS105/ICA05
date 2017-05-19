# ICA05

## Websrv    
Webserveren bruker index.html filen som ligger i mappen "static", og viser denne for brukeren.
Oppsett er relativt simpelt, clone mappen og kjør websrv.go så vil den starte en server på port 8001.   

## Webclient
Main-fila inneholder URL-ene og go-rutinene, som kjører hver sin metode i "API" pakken. 
Hver go-fil som ligger i “API” tar inn JSON som array av bytes, og konverterer det til noe man kan leseved hjelp av “Decoder” eller “Unmarshal”. 
Hvis det er en array av structs må man bruke “Decoder” og iterere over innholdet med en loop, ellers er det “Unmarshal” som brukes.
