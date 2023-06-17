package file
	
	import "fmt"
	 
	type ExampleStruct struct{

	
	Id string
		TransactionNumber string
		OrderId string
		PaymentId string
		Msisdn string
		PaymentMethod string
		CustomerName string
		Periode string
		TotalDiscount int
		TotalBill int
		TotalTax int
		Qty int
		Cashback int
		Tag string
		Name string
		UserId string
		DeviceId string
		Email string
		StatusTransaction string
		CreatedAt string
		UpdatedAt string
		OrderNumber string
		URLPayload string
		TransactionPackageId string
		}
	
	func exampleToStruct(){

	

	data := ExampleStruct{}
		data.Id="bBFtq"
		data.TransactionNumber="JRnSS"
		data.OrderId="DPtgO"
		data.PaymentId="gmhuS"
		data.Msisdn="qIBsy"
		data.PaymentMethod="Vacec"
		data.CustomerName="Jkuqc Dhfr"
		data.Periode="gMhqW"
		data.TotalDiscount=100000
		data.TotalBill=85000
		data.TotalTax=30000
		data.Qty=55000
		data.Cashback=85000
		data.Tag="XIXCY"
		data.Name="Gvwdi Imck"
		data.UserId="gVWMd"
		data.DeviceId="rTyTZ"
		data.Email="lisybq@gmail.com"
		data.StatusTransaction="fbUMO"
		data.CreatedAt="yrwEX"
		data.UpdatedAt="CCGlO"
		data.OrderNumber="wKLjg"
		data.URLPayload="QLgKw"
		data.TransactionPackageId="SvONk"
		
	fmt.Println(data)
	}