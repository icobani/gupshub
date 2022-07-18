# gupshub

[markdown gupshub package](https://docs.gupshup.io/reference/msg)

# Walet Balance

	var wa gupshub.WhatsApp
	wa.ApiKey = os.Getenv("WAApiKey")
	if balance, err := wa.Balance(); err == nil {
		fmt.Printf("Balance : %+v\n", balance)
	} else {
		fmt.Printf("Err : %+v\n", err.Error())
	}

# Template List

    var wa gupshub.WhatsApp
    wa.AppName = os.Getenv("WAAppName")
    wa.ApiKey = os.Getenv("WAApiKey")
    if templates, err := wa.GetTemplateList(); err == nil {
        fmt.Printf("Templates : %+v\n", templates)
    } else {
        fmt.Printf("Err : %+v\n", err.Error())
    }

# Contact List

    var wa gupshub.WhatsApp
    wa.AppName = os.Getenv("WAAppName")
    wa.ApiKey = os.Getenv("WAApiKey")
    if contacts, err := wa.Contacts(); err == nil {
        fmt.Printf("Contacts : %+v\n", contacts)
    } else {
        fmt.Printf("Err : %+v\n", err.Error())
    }

# Contact Opt-In

    toPhoneNumber := "905325401194"
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if err := wa.OptIn(toPhoneNumber); err == nil {
		fmt.Println(err.Error())		
	}

# Contact Opt-Out

    toPhoneNumber := "905325401194"
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if err := wa.OptOut(toPhoneNumber); err == nil {
		fmt.Println(err.Error())		
	}

# Send Template Message

    toPhoneNumber := "905325401194"
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	_ = wa.OptIn(toPhoneNumber)
	templateMessageRequest := wa.NewTemplateMessageRequest()
	templateMessageRequest.Source = os.Getenv("WaSourcePhoneNumber")
	templateMessageRequest.Destination = toPhoneNumber
	templateMessageRequest.Template = &gupshub.TemplateData{
		ID:     "b2aee92d-bf3a-461a-b389-28098feb8fb7",
		Params: []string{"Ibrahim", "2234"},
	}
	templateMessageRequest.Send()