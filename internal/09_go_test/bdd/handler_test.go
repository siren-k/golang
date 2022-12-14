package bdd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

var payloads []HandlerRequest
var resps []*httptest.ResponseRecorder

func weCreateAHandlerRequestPayloadWith(arg1 *godog.Table) error {
	for _, row := range arg1.Rows {
		h := HandlerRequest{
			Name: row.Cells[0].Value,
		}
		payloads = append(payloads, h)
	}
	return nil
}

func wePOSTTheHandlerRequestToHello() error {
	for _, p := range payloads {
		v, err := json.Marshal(p)
		if err != nil {
			return err
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/hello", bytes.NewBuffer(v))

		Handler(w, r)
		resps = append(resps, w)
	}
	return nil
}

func theResponseCodeShouldBe(arg1 int) error {
	for _, r := range resps {
		if got, want := r.Code, arg1; got != want {
			return fmt.Errorf("got: %d; want %d", got, want)
		}
	}
	return nil
}

func theResponseBodyShouldBe(arg1 *godog.Table) error {
	for c, row := range arg1.Rows {
		b := bytes.Buffer{}
		b.ReadFrom(resps[c].Body)
		if got, want := b.String(), row.Cells[0].Value; got != want {
			return fmt.Errorf("got: %s; want %s", got, want)
		}
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response body should be:$`, theResponseBodyShouldBe)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^we create a HandlerRequest payload with:$`, weCreateAHandlerRequestPayloadWith)
	ctx.Step(`^we POST the HandlerRequest to \/hello$`, wePOSTTheHandlerRequestToHello)

}
