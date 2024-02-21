package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages/go/v21"
	httpHandler "github.com/periclescesar/rinha-2024-q1-go/internal/clientes/delivery/http"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*messages.Pickle) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iMakeADebitOfToTheCustomersAccountWithIdAndDescription(amount, id int, description string) error {
	data := map[string]interface{}{
		"valor":     amount,
		"tipo":      'd',
		"descricao": description,
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return err
	}

	body := bytes.NewReader(jsonBody)
	return a.makeRequest(http.MethodPost, fmt.Sprintf("/clientes/%d/transacoes", id), body)
}

func (a *apiFeature) iWillReceiveAError(errMessage string) error {
	expected := map[string]interface{}{
		"mensagem": errMessage,
	}

	var actual interface{}
	// re-encode actual response too
	if err := json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, expected, actual)
}

type Balance struct {
	Total       int    `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int    `json:"limite"`
}

type Transaction struct {
	Valor      int    `json:"valor"`
	Tipo       string `json:"tipo"`
	Descricao  string `json:"descricao"`
	RealizedAt string `json:"realizada_em"`
}

type AccountStattment struct {
	Saldo             interface{}   `json:"saldo"`
	UltimasTransacoes []Transaction `json:"ultimas_transacoes"`
}

func (a *apiFeature) iWillSeeMyLimitOfAndBalanceOf(limit, balance int) error {
	var expected = &AccountStattment{
		Saldo: &Balance{
			Limite: limit,
			Total:  balance,
		},
	}
	var actual AccountStattment
	// re-encode actual response too
	if err := json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, expected, actual)
}

func (a *apiFeature) theLastsTransactions(nTransactions int) error {
	var actual AccountStattment
	// re-encode actual response too
	if err := json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, nTransactions, len(actual.UltimasTransacoes))
}

func (a *apiFeature) getAccountStattmentOfCustomer(id int) error {
	return a.makeRequest("GET", fmt.Sprintf("/clientes/%d/extrato", id), nil)
}

func (a *apiFeature) makeRequest(method string, path string, body io.Reader) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	r := httpHandler.SetupRouter()
	r.ServeHTTP(a.resp, req)

	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	return err
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { fmt.Println("Starting suite!") })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	ctx.Step(`^I get a account stattement of the customer\'s id (\d+)$`, api.getAccountStattmentOfCustomer)
	ctx.Step(`^I make a debit of (\d+) to the customer\'s account with id (\d+) and description "([^"]*)"$`, api.iMakeADebitOfToTheCustomersAccountWithIdAndDescription)
	ctx.Step(`^I will receive a error "([^"]*)"$`, api.iWillReceiveAError)
	ctx.Step(`^I will see my limit of (\d+) and balance of -(\d+)$`, api.iWillSeeMyLimitOfAndBalanceOf)
	ctx.Step(`^the lasts (\d+) transactions$`, api.theLastsTransactions)
}
