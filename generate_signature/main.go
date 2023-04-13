package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Profiles struct {
	Balance             string `json:"balance"`
	Chargingmode        string `json:"chargingmode"`
	Gracedate           string `json:"gracedate"`
	Canceldate          string `json:"canceldate"`
	CustomerId          string `json:"customer_id"`
	CustomerName        string `json:"customer_name"`
	CustomerType        string `json:"customer_type"`
	CustomerSubtype     string `json:"customer_subtype"`
	Postcusttype        string `json:"postcusttype"`
	Postcustsubtype     string `json:"postcustsubtype"`
	SubscriberStatus    string `json:"subscriber_status"`
	BillcycleCode       string `json:"billcycle_code"`
	Billcycle           string `json:"billcycle"`
	AccountId           string `json:"account_id"`
	LastBillAmount      string `json:"last_bill_amount"`
	LastBillDue         string `json:"last_bill_due"`
	Postemail           string `json:"postmail"`
	Postaddress         string `json:"postaddress"`
	Priceplan           string `json:"priceplan"`
	CreditLimitRoaming  string `json:"credit_limit_roaming"`
	CreditLimitDomestic string `json:"credit_limit_domestic"`
	Puk1                string `json:"puk1"`
	Puk2                string `json:"puk2"`
	DspCusttype         string `json:"dsp_custtype"`
	Location            string `json:"location"`
	Imsi                string `json:"imsi"`
	Imei                string `json:"imei"`
	ScpId               string `json:"scp_id"`
	Handsettype         string `json:"handsettype"`
	Brand               string `json:"brand"`
	Mnc                 string `json:"mnc"`
	Mcc                 string `json:"mcc"`
	Lac                 string `json:"lac"`
	Ci                  string `json:"ci"`
	Skey                string `json:"skey"`
	Kecamatan           string `json:"kecamatan"`
	CardType            string `json:"card_type"`
	Lte                 string `json:"lte"`
	Regional            string `json:"regional"`
	Custtype            string `json:"custtype"`
	Countryname         string `json:"countryname"`
	Network4g           string `json:"network4g"`
	Odbroam             string `json:"odbroam"`
	SgsntplName         string `json:"sgsntpl_name"`
	Corporateprt        string `json:"corporateprt"`
	BonusSms            string `json:"bonus_sms"`
	BonusVoice          string `json:"bonus_voice"`
	BonusData           string `json:"bonus_data"`
	BonusMonetary       string `json:"bonus_monetary"`
	Pcrfprofilenames    string `json:"pcrfprofilenames"`
	Los                 string `json:"los"`
	TcashBalance        string `json:"tcash_balance"`
	Pointcategory       string `json:"pointcategory"`
	Pointlte            string `json:"pointlte"`
	Pointregular        string `json:"pointregular"`
	Pointlteexpdate     string `json:"pointlteexpdate"`
	Pointregularexpdate string `json:"pointregularexpdate"`
}

type DetailTransaction struct {
	Month       string `json:"month"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
type CheckTierMessage struct {
	TierStatus           string              `json:"tier_status"`
	TierTransactionPrice string              `json:"tier_transaction_price"`
	TransactionPriceLeft string              `json:"transaction_price_left"`
	TierTransactionDate  string              `json:"tier_transaction_date"`
	PurchaseBefore       string              `json:"purchase_before"`
	NextTier             string              `json:"next_tier,omitempty"`
	Percentage           string              `json:"percentage"`
	DetailTransactions   []DetailTransaction `json:"detail_transaction"`
	IsSilver             bool                `json:"is_silver"`
	IsGold               bool                `json:"is_gold"`
	IsPlatinum           bool                `json:"is_platinum"`
	IsDiamond            bool                `json:"is_diamond"`
	IsCorporate          bool                `json:"is_corporate"`
	NextMonth            string              `json:"next_month"`
	IsDown               bool                `json:"is_down"`
}

func main() {
	signature := CreateSignature()
	fmt.Println("signature:", signature)

	signatureprod := CreateSignatureProd()
	fmt.Println("signature prod:", signatureprod)

	// sp := strings.Split("https://api.digitalcore.telkomsel.co.id/preprod-web/scrt/subscriber/profiles/6281234567890?profileitems=customer_type", "?")
	// sp2 := strings.Split(sp[0], "/")
	// fmt.Println(sp2[len(sp2)-1])

	sp := strings.Split("https://api.digitalcore.telkomsel.co.id/preprod-web/scrt/subscriber/profiles/6281234567890?profileitems=customer_type", "?")
	sp2 := strings.Split(sp[0], "/")
	msisdn := sp2[len(sp2)-1]
	fmt.Println(msisdn)

	i := 2
	detailTransactions := make([]DetailTransaction, 0)
	detailTransaction := DetailTransaction{}
	detailTransaction.Description = ""
	detailTransaction.Value = "5000"
	month, _ := TimeTransaction(2)
	detailTransaction.Month = month
	detailTransactions = append(detailTransactions, detailTransaction)
	detailTransaction2 := DetailTransaction{}
	detailTransaction2.Description = ""
	detailTransaction2.Value = "122000"
	month2, _ := TimeTransaction(1)

	detailTransaction2.Month = month2
	detailTransactions = append(detailTransactions, detailTransaction2)
	detailTransactions = append(detailTransactions, detailTransaction2)
	detailTransaction3 := DetailTransaction{}
	detailTransaction3.Description = ""
	detailTransaction3.Value = "0"
	month3, _ := TimeTransaction(0)
	detailTransaction3.Month = month3
	detailTransactions = append(detailTransactions, detailTransaction3)

	detailTransactions[i] = detailTransactions[len(detailTransactions)-1] // Copy last element to index i.
	detailTransactions = detailTransactions[:len(detailTransactions)-1]
	fmt.Println(detailTransactions)
	month2, date := TimeTransaction(0)
	// fmt.Println(month2, date)

	// for i := len(detailTransactions); i > 0; i-- {
	// 	fmt.Println(i)
	// }

	// ==============================================================
	//===============================================================
	var REV_MTD_VALUE, REV_M2_VALUE, REV_M1_VALUE = "90600", "159700", "230000"
	TIER := "GOLD"
	result := CheckTierMessage{}
	rev_mtd_int, _ := strconv.ParseFloat(REV_MTD_VALUE, 64)
	rev_mt2_int, _ := strconv.ParseFloat(REV_M2_VALUE, 64)
	rev_mt1_int, _ := strconv.ParseFloat(REV_M1_VALUE, 64)

	sum := rev_mt1_int + rev_mt2_int + rev_mtd_int

	sumAVG := int(sum)

	threshold_gold, _ := strconv.Atoi("300000")
	threshold_platinum, _ := strconv.Atoi("900000")
	threshold_diamond, _ := strconv.Atoi("3000000")
	countPercentage := 0.0
	formatPercentage := ""

	var price_left_tostring string

	timeFull, date := TimeTransaction(-1)
	fmt.Println(date)
	var nextMonthYear = timeFull
	var nextMonth string
	if timeFull != "" {
		monthYearSplit := strings.Split(timeFull, " ")
		nextMonth = monthYearSplit[0]
	}
	// nextTierDate, _ := helper.TimeTransaction(-2)
	var down bool
	fmt.Println(down)

	if strings.EqualFold(TIER, strings.ToLower("SILVER")) || strings.EqualFold(TIER, strings.ToLower("Silver")) {
		countPercentage = (float64(sumAVG) / float64(threshold_gold)) * 100
		formatPercentage = fmt.Sprintf("%.2f", countPercentage)

		price_left_tostring = strconv.Itoa(threshold_gold - sumAVG)
		result.TierStatus = TIER
		result.NextTier = "GOLD"
		result.PurchaseBefore = "01 " + timeFull
		result.Percentage = formatPercentage
		result.IsSilver = true
		result.TierTransactionDate = nextMonthYear
		result.NextMonth = nextMonth
		if sumAVG > threshold_gold {
			result.TransactionPriceLeft = "0"
		} else {
			result.TransactionPriceLeft = price_left_tostring

		}
	}

	if strings.EqualFold(TIER, "GOLD") || strings.EqualFold(TIER, "Gold") {
		countPercentage = (float64(sumAVG) / float64(threshold_platinum)) * 100
		formatPercentage = fmt.Sprintf("%.2f", countPercentage)
		price_left_tostring = strconv.Itoa(threshold_platinum - sumAVG)
		result.TierStatus = TIER
		result.NextTier = "PLATINUM"
		result.PurchaseBefore = "01 " + timeFull
		result.TierTransactionDate = nextMonthYear
		result.Percentage = formatPercentage
		result.IsGold = true
		result.NextMonth = nextMonth

		if sumAVG > threshold_platinum {
			result.TransactionPriceLeft = "0"
		} else {
			result.TransactionPriceLeft = price_left_tostring
		}
		if sumAVG < threshold_gold {

			countPercentage = (float64(sumAVG) / float64(threshold_gold)) * 100
			formatPercentage = fmt.Sprintf("%.2f", countPercentage)
			price_left_tostring = strconv.Itoa(threshold_gold - sumAVG)
			result.TierStatus = TIER
			result.NextTier = "GOLD"
			result.PurchaseBefore = "01 " + timeFull
			nextMonthYear, _ = TimeTransaction(-1)
			result.TierTransactionDate = nextMonthYear
			result.Percentage = formatPercentage
			result.IsGold = true
			result.NextMonth = nextMonth
			result.IsDown = true

			if sumAVG > threshold_platinum {
				result.TransactionPriceLeft = "0"
			} else {
				result.TransactionPriceLeft = price_left_tostring
			}
		}
	}

	if strings.EqualFold(TIER, "PLATINUM") || strings.EqualFold(TIER, "Platinum") {
		countPercentage = (float64(sumAVG) / float64(threshold_diamond)) * 100
		formatPercentage = fmt.Sprintf("%.2f", countPercentage)
		price_left_tostring = strconv.Itoa(threshold_diamond - sumAVG)
		result.TierStatus = TIER
		result.NextTier = "DIAMOND"
		result.PurchaseBefore = "01 " + timeFull
		result.IsPlatinum = true
		result.Percentage = formatPercentage
		result.TierTransactionDate = nextMonthYear
		result.NextMonth = nextMonth

		if sumAVG > threshold_diamond {
			result.TransactionPriceLeft = "0"
		} else {
			result.TransactionPriceLeft = price_left_tostring
		}
		if sumAVG < threshold_diamond && strings.ToLower(TIER) == "diamond" {
			down = true
		}
	}

	if strings.EqualFold(TIER, "DIAMOND") || strings.EqualFold(TIER, "Diamond") {
		price_left_tostring = strconv.Itoa(sumAVG)
		result.TierStatus = TIER
		result.NextTier = ""
		result.PurchaseBefore = ""
		result.TierTransactionDate = ""
		result.TransactionPriceLeft = price_left_tostring
		result.Percentage = os.Getenv("PERCENTAGE_DIAMOND")
		result.IsDiamond = true
		result.NextMonth = nextMonth
		if sumAVG < threshold_gold && strings.ToLower(TIER) == "gold" {
			down = true
		}
		if sumAVG < threshold_platinum && strings.ToLower(TIER) == "platinum" {
			down = true
		}

	}
	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

}

func CreateSignature() string {

	api_key := "eedfevw8q7rym4zhdsa64cma"
	fmt.Println(api_key)

	createSignature := api_key + "CSkjSURNZA" + strconv.FormatInt(time.Now().Unix(), 10)

	hash := md5.Sum([]byte(createSignature))

	xSignature := hex.EncodeToString(hash[:])

	return xSignature
}
func CreateSignatureProd() string {

	api_key := "4wj7qbrzr8yv57qjby4c9wg7"
	fmt.Println(api_key)

	createSignature := api_key + "2JZQ3h3mBS" + strconv.FormatInt(time.Now().Unix(), 10)

	hash := md5.Sum([]byte(createSignature))

	xSignature := hex.EncodeToString(hash[:])

	return xSignature
}

func TimeTransaction(montMinus int) (string, string) {

	format := "02-January-2006"

	timeNow := time.Now().AddDate(0, int(math.Abs(float64(montMinus))), 0)
	if montMinus >= 0 {
		timeNow = time.Now().AddDate(0, -montMinus, 0)
	}

	times := timeNow.Format(format)

	sp := strings.Split(times, "-")

	join := fmt.Sprintf("%s %s", sp[1], sp[2])
	return join, sp[0]

}

func TierTrxDateAndPurchaseBefore() (string, string) {

	format := "02-January-2006"

	timeNow := time.Now().AddDate(1, 0, 0)

	times := timeNow.Format(format)

	sp := strings.Split(times, "-")

	t, _, _ := time.Now().Date()

	timeFull := fmt.Sprintf("01 desember %d", t)

	return timeFull, sp[2]

}
