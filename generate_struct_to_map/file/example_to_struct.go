package file
	
	import "fmt"
	 
	type ExampleStruct struct{

	
	Event string
		StatusCode int
		ResponseTime time.Duration
		Method string
		Request interface {}
		URL string
		Message string
		Tag string
		OrderNumber string
		}
	
	func exampleToStruct(){

	

	data := ExampleStruct{}
		data.Event="mnogv"
		data.StatusCode=15000
		data.ResponseTime=
		data.Method="CSydD"
		data.Request=
		data.URL="XLFiN"
		data.Message="drfwU"
		data.Tag="Ozonr"
		data.OrderNumber="wpiQW"
		
	fmt.Println(data)
	}