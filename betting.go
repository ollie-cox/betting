package betting

import (
	"bytes"
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

type Betting struct {
	*Client
}

// Request function for send requests to betfair via REST JSON
func (b *Betting) Request(reqStruct interface{}, url BetfairRestURL, method string, filter *Filter) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	urlBuild := bytes.NewBuffer([]byte{})
	urlBuild.WriteString(string(url))
	urlBuild.WriteString("/")
	urlBuild.WriteString(method)
	urlBuild.WriteString("/")

	req.SetRequestURI(urlBuild.String())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	if filter != nil {
		filterBody, err := ffjson.Marshal(&filter)
		if err != nil {
			return err
		}
		fmt.Print(urlBuild.String() + " " )
		fmt.Println(string(filterBody))

		req.SetBody(filterBody)
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		err = ffjson.Unmarshal(resp.Body(), &bettingError)
		if err != nil {
			return err
		}

		return fmt.Errorf("Error with code - %s and string - %s", bettingError.Faultcode, bettingError.Faultstring)
	}
	fmt.Println("resp: " + string(resp.Body()))
	err = ffjson.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}

// Request function for send requests to betfair via REST JSON
func (b *Betting) CancelRequest(reqStruct interface{}, url BetfairRestURL, method string, filter *CancelFilter) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	urlBuild := bytes.NewBuffer([]byte{})
	urlBuild.WriteString(string(url))
	urlBuild.WriteString("/")
	urlBuild.WriteString(method)
	urlBuild.WriteString("/")

	req.SetRequestURI(urlBuild.String())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	if filter != nil {
		filterBody, err := ffjson.Marshal(&filter)
		if err != nil {
			return err
		}
		fmt.Print(urlBuild.String() + " " )
		fmt.Println(string(filterBody))

		req.SetBody(filterBody)
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		err = ffjson.Unmarshal(resp.Body(), &bettingError)
		if err != nil {
			return err
		}

		return fmt.Errorf("Error with code - %s and string - %s", bettingError.Faultcode, bettingError.Faultstring)
	}

	fmt.Println("resp: " + string(resp.Body()))

	err = ffjson.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}




// Request function for send requests to betfair via REST JSON
func (b *Betting) ReplaceRequest(reqStruct interface{}, url BetfairRestURL, method string, filter *ReplaceFilter) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	urlBuild := bytes.NewBuffer([]byte{})
	urlBuild.WriteString(string(url))
	urlBuild.WriteString("/")
	urlBuild.WriteString(method)
	urlBuild.WriteString("/")

	req.SetRequestURI(urlBuild.String())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	if filter != nil {
		filterBody, err := ffjson.Marshal(&filter)
		if err != nil {
			return err
		}
		fmt.Print(urlBuild.String() + " " )
		fmt.Println(string(filterBody))

		req.SetBody(filterBody)
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		err = ffjson.Unmarshal(resp.Body(), &bettingError)
		if err != nil {
			return err
		}

		return fmt.Errorf("Error with code - %s and string - %s", bettingError.Faultcode, bettingError.Faultstring)
	}

	fmt.Println("resp: " + string(resp.Body()))

	err = ffjson.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}




// Request function for send requests to betfair via REST JSON
func (b *Betting) RunnerRequest(reqStruct interface{}, url BetfairRestURL, method string, filter *RunnerFilter) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	urlBuild := bytes.NewBuffer([]byte{})
	urlBuild.WriteString(string(url))
	urlBuild.WriteString("/")
	urlBuild.WriteString(method)
	urlBuild.WriteString("/")

	req.SetRequestURI(urlBuild.String())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	if filter != nil {
		filterBody, err := ffjson.Marshal(&filter)
		if err != nil {
			return err
		}
		fmt.Print(urlBuild.String() + " " )
		fmt.Println(string(filterBody))

		req.SetBody(filterBody)
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		err = ffjson.Unmarshal(resp.Body(), &bettingError)
		if err != nil {
			return err
		}

		return fmt.Errorf("Error with code - %s and string - %s", bettingError.Faultcode, bettingError.Faultstring)
	}

	fmt.Println("resp: " + string(resp.Body()))

	err = ffjson.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}
