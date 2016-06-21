package betting

import (
	"log"
	"testing"
	"time"
)

func TestRequestListCompetitions(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListCompetitions(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{
		//InPlayOnly: false,
		},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListCompetitions\n %v \n\n", list)
	}
}

func TestRequestListCountries(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListCountries(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{
		//InPlayOnly: false,
		},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListCountries\n %v \n\n", list)
	}
}

func TestRequestListCurrentOrders(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListCurrentOrders(Filter{})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListCurrentOrders\n %v \n\n", list)
	}
}

func TestRequestListClearedOrders(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListClearedOrders(Filter{
		SettledDateRange: DateRange{time.Now().AddDate(-1, 0, 0), time.Now()},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListClearedOrders\n %v \n\n", list)
	}
}

func TestRequestListEvents(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListEvents(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListEvents\n %v \n\n", list)
	}
}

func TestRequestListEventTypes(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListEventTypes(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListEventTypes\n %v \n\n", list)
	}
}

func TestRequestListMarketCatalogue(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListMarketCatalogue(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{},
		MaxResults:   5,
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestListMarketCatalogue\n %v \n\n", list)
	}
}