package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type JsonReservationDetailParam struct {
	SubscriptionPlan string
	AddOnPackage     string
	CreditLimit      int
	HaloMsisdn       string
	FirstName        string
	LastName         string
	Pob              string
	Dob              string
	OtherMsisdn      string
	Email            string
	BillingCycle     string
}

type JsonPersonalDataFormParam struct {
	FirstName       string
	LastName        string
	Pob             string
	Dob             string
	IdentityNumber  string
	ExpiredNumber   string
	Nokk            string
	MotherName      string
	IdentityAddress string
	City            string
	Province        string
	PostalCode      string
	Country         string
	Email           string
	HaloMsisdn      string
	ItemName        string
	BillingCycle    string
	PricePlan       string
	Cls             string
	Ir              string
}

func main() {
	start := time.Now()
	rsvParam := JsonReservationDetailParam{
		SubscriptionPlan: "test",
		AddOnPackage:     "param.AddOnPackage",
		CreditLimit:      3838733,
		HaloMsisdn:       "param.HaloMsisdn",
		FirstName:        "param.FirstName",
		LastName:         "param.LastName",
		Pob:              "param.Pob",
		Dob:              "param.Dob",
		OtherMsisdn:      "",
		Email:            "param.Email",
		BillingCycle:     "param.BillingCycle",
	}

	persParam := JsonPersonalDataFormParam{
		FirstName:       "param.FirstName",
		LastName:        "param.LastName",
		Pob:             "param.Pob",
		Dob:             "param.Dob",
		IdentityNumber:  "param.IdentityNumber",
		ExpiredNumber:   "param.ExpiredNumber",
		Nokk:            "param.Nokk",
		MotherName:      "param.MotherName",
		IdentityAddress: "param.IdentityAddress",
		City:            "param.City",
		Province:        "param.Province",
		PostalCode:      "param.PostalCode",
		Country:         "param.Country",
		Email:           "param.Email",
		HaloMsisdn:      "param.HaloMsisdn",
		ItemName:        "param.ItemName",
		BillingCycle:    "param.BillingCycle",
		PricePlan:       "param.PricePlan",
		Cls:             "param.Cls",
		Ir:              "param.Ir",
	}
	fmt.Println("halo")

	// documents := []string{
	// cpuNums := runtime.NumCPU()
	// runtime.GOMAXPROCS(1)
	// fmt.Println(cpuNums)

	var wg sync.WaitGroup
	wg.Add(3)

	var detail string
	var personal string
	var personal2 string
	go func(detail *string) {
		res1 := GenerateReservationDetails(rsvParam)
		*detail = res1
		wg.Done()
		fmt.Println(runtime.NumCPU())
	}(&detail)
	go func(personal *string) {
		res2 := GeneratePersonalDataForm(persParam, "personal-web-psb.html")
		*personal = res2
		wg.Done()
		fmt.Println(runtime.NumCPU())

	}(&personal)
	go func(personal2 *string) {
		res3 := GeneratePersonalDataForm(persParam, "terms.html")
		*personal2 = res3
		wg.Done()

	}(&personal2)
	wg.Wait()
	documents := []string{detail, personal, personal2}
	fmt.Println(documents)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	// start2 := time.Now()

	// TestGoRoutine()
	// elapsed2 := time.Since(start2)
	// log.Printf("Binomial took %s", elapsed2)
	start3 := time.Now()

	res4 := GenerateReservationDetails(rsvParam)
	res5 := GeneratePersonalDataForm(persParam, "personal-web-psb.html")
	res6 := GeneratePersonalDataForm(persParam, "terms.html")
	fmt.Println(res4, res5, res6)
	elapsed3 := time.Since(start3)
	log.Printf("Binomial took %s", elapsed3)

	// }
}

func GenerateReservationDetails(param JsonReservationDetailParam) (base64string string) {
	pdfg := GeneratePDF("number-reservation.html", "/home/ademawan/go/src/generate/generate_pdf/number-reservation.html", param)
	pdfg.Create()
	filebyte := pdfg.Bytes()

	strBase64 := base64.StdEncoding.EncodeToString(filebyte)
	// log.Println(strBase64)
	return strBase64
}

func GeneratePersonalDataForm(param JsonPersonalDataFormParam, fileName string) (base64string string) {
	pdfg := GeneratePDF(fileName, "/home/ademawan/go/src/generate/generate_pdf/"+fileName, param)
	pdfg.Create()
	filebyte := pdfg.Bytes()

	strBase64 := base64.StdEncoding.EncodeToString(filebyte)
	// log.Println(strBase64)
	return strBase64
}

func GeneratePDF(fileName, filePath string, param interface{}) *wkhtmltopdf.PDFGenerator {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(err.Error())

		log.Fatal(err)
	}

	t := template.New(fileName)

	t, err = t.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, fileName, param); err != nil {

		log.Println(err)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(&tpl))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)
	pdfg.NoCollate.Set(false)
	// pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	return pdfg
}

type DocFile struct {
	DocTransactionID string `json:"doc_transaction_id"`
	DocType          string `json:"doc_type"`
}

func TestGoRoutine() {

	rsvParam := JsonReservationDetailParam{
		SubscriptionPlan: "test",
		AddOnPackage:     "param.AddOnPackage",
		CreditLimit:      3838733,
		HaloMsisdn:       "param.HaloMsisdn",
		FirstName:        "param.FirstName",
		LastName:         "param.LastName",
		Pob:              "param.Pob",
		Dob:              "param.Dob",
		OtherMsisdn:      "",
		Email:            "param.Email",
		BillingCycle:     "param.BillingCycle",
	}

	persParam := JsonPersonalDataFormParam{
		FirstName:       "param.FirstName",
		LastName:        "param.LastName",
		Pob:             "param.Pob",
		Dob:             "param.Dob",
		IdentityNumber:  "param.IdentityNumber",
		ExpiredNumber:   "param.ExpiredNumber",
		Nokk:            "param.Nokk",
		MotherName:      "param.MotherName",
		IdentityAddress: "param.IdentityAddress",
		City:            "param.City",
		Province:        "param.Province",
		PostalCode:      "param.PostalCode",
		Country:         "param.Country",
		Email:           "param.Email",
		HaloMsisdn:      "param.HaloMsisdn",
		ItemName:        "param.ItemName",
		BillingCycle:    "param.BillingCycle",
		PricePlan:       "param.PricePlan",
		Cls:             "param.Cls",
		Ir:              "param.Ir",
	}

	var docFile1 = DocFile{}
	var docFile2 = DocFile{}
	var docFile3 = DocFile{}
	var paket = []interface{}{&docFile1, &docFile2, &docFile3}
	docType := "DOCTYPE_SNK_DOCMAN"

	var wg sync.WaitGroup
	wg.Add(3)

	var detail string
	var personalWebPsb string
	var personalTerms string

	go func(detail *string) {
		res1 := GenerateReservationDetails(rsvParam)
		*detail = res1
		wg.Done()
	}(&detail)
	go func(personalWebPsb *string) {
		res2 := GeneratePersonalDataForm(persParam, "personal-web-psb.html")
		*personalWebPsb = res2
		wg.Done()

	}(&personalWebPsb)
	go func(personalTerms *string) {
		res3 := GeneratePersonalDataForm(persParam, "terms.html")
		*personalTerms = res3
		wg.Done()

	}(&personalTerms)
	wg.Wait()

	documents := []string{detail, personalWebPsb, personalTerms}
	wg.Add(3)
	for i, document := range documents {

		// trxID, _ := FileUploadDocman(document, docType, rsvParam.HaloMsisdn)

		// paket[i].(*DocFile).DocTransactionID = trxID
		// paket[i].(*DocFile).DocType = docType
		go func(doc *DocFile, docType, document, haloMsisdn *string) {
			trxID, _ := FileUploadDocman(*document, *docType, *haloMsisdn)

			doc.DocTransactionID = trxID
			doc.DocType = *docType
			wg.Done()

		}(paket[i].(*DocFile), &docType, &document, &rsvParam.HaloMsisdn)

	}
	wg.Wait()
	var docFiles = []DocFile{docFile1, docFile2, docFile3}
	fmt.Println(docFiles)

}
func FileUploadDocman(document, docType, haloMsisdn string) (string, error) {
	time.Sleep(time.Second * 5)
	timeInt := time.Now().UnixNano()
	timeString := strconv.Itoa(int(timeInt))
	return timeString, nil
}
