package agent

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/evgenytr/kpibuffer/internal/interfaces"
)

const apiUrl = "https://development.kpi-drive.ru/_api/facts/save_fact"
const bearer = "Bearer 48ab34464a5573519725deb5865cc74c"

func SendToAPI(storage interfaces.Storage) {

	client := &http.Client{}
	for {

		currFact, err := storage.PopFact()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		data := url.Values{}
		data.Set("period_start", currFact.PeriodStart)
		data.Set("period_end", currFact.PeriodEnd)
		data.Set("period_key", currFact.PeriodKey)
		data.Set("indicator_to_mo_id", strconv.Itoa(currFact.IndicatorToMoId))
		data.Set("indicator_to_mo_fact_id", strconv.Itoa(currFact.IndicatorToMoFactId))
		data.Set("value", strconv.Itoa(currFact.Value))
		data.Set("fact_time", currFact.FactTime)
		data.Set("is_plan", strconv.Itoa(currFact.IsPlan))
		data.Set("auth_user_id", strconv.Itoa(currFact.AuthUserId))
		data.Set("comment", currFact.Comment)

		fmt.Println("sending: ", data.Encode())
		r, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(data.Encode()))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Authorization", bearer)

		resp, err := client.Do(r)
		fmt.Println("response status:", resp.Status)
	}
}
