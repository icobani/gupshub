sendMessage:
	curl --request POST --url https://api.gupshup.io/sm/api/v1/msg \
         --header 'Accept: application/json' \
         --header 'Content-Type: application/x-www-form-urlencoded' \
         --header 'apikey: 2lb8bzbdbvikrv9buvhhaxb8hj65be8a' \
         --data 'message={"type":"text","text":"Merhaba Gökhan, şifren 4455"}' \
         --data src.name=ICITECH \
         --data channel=whatsapp \
         --data source=902123567077 \
         --data destination=905325401194

sendTemplate:
	curl --location --request POST 'http://api.gupshup.io/sm/api/v1/template/msg' \
	--header 'apikey: 2lb8bzbdbvikrv9buvhhaxb8hj65be8a' \
	--header 'Content-Type: application/x-www-form-urlencoded' \
	--data-urlencode 'source=902123567077' \
	--data-urlencode 'destination=905337691921' \
	--data-urlencode 'template={"id": "69b83f67-0372-4a5e-b748-0ecded143304","params": ["https://download.icibot.app"]}'