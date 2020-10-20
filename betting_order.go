package betting

import "time"

type PlaceInstruction struct {
	OrderType          EOrderType         `json:"orderType,omitempty"`
	SelectionID        int64              `json:"selectionId,omitempty"`
	Handicap           float64            `json:"handicap"`
	Side               ESide              `json:"side,omitempty"`
	LimitOrder         LimitOrder         `json:"limitOrder,omitempty"`
	//LimitOnCloseOrder  LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	//MarketOnCloseOrder MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string             `json:"customerOrderRef,omitempty"`
}


type PlaceExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []PlaceInstructionReport  `json:"instructionReports,omitempty"`
}


type ReplaceExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []ReplaceInstructionReport  `json:"instructionReports,omitempty"`
}


type CancelInstruction struct {
	BetID               string                      `json:"betId,omitempty"`
//	SizeReduction		float64						`json:"sizeReduction"`
}

type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport `json:"instructionReports,omitempty"`
}


type CancelInstructionReport struct {
	Status             EInstructionReportStatus    	`json:"status,omitempty"`
	ErrorCode          EInstructionReportErrorCode 	`json:"errorCode,omitempty"`
	Instruction        CancelInstruction 			`json:"instruction,omitempty"`
	SizeCancelled 	   float64 						`json:"sizeCancelled,omitempty"`
	CancelledDate 	   time.Time                   	`json:"cancelledDate,omitempty"`
}


type ReplaceInstruction struct {
	BetID               string                      `json:"betId,omitempty"`
	NewPrice float64                     `json:"newPrice,omitempty"`
}


type ReplaceInstructionReport struct {
	Status              EInstructionReportStatus    `json:"status,omitempty"`
	ErrorCode           EInstructionReportErrorCode `json:"errorCode,omitempty"`
	CancelInstructionReport CancelInstructionReport `json:"cancelInstructionReport,omitempty"`
	PlaceInstructionReport PlaceInstructionReport `json:"placeInstructionReport,omitempty"`
}

type PlaceInstructionReport struct {
	Status              EInstructionReportStatus    `json:"status,omitempty"`
	ErrorCode           EInstructionReportErrorCode `json:"errorCode,omitempty"`
	OrderStatus         EOrderStatus                `json:"orderStatus,omitempty"`
	Instruction         PlaceInstruction            `json:"instruction,omitempty"`
	BetID               string                      `json:"betId,omitempty"`
	PlacedDate          time.Time                   `json:"placedDate,omitempty"`
	AveragePriceMatched float64                     `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64                     `json:"sizeMatched,omitempty"`
}

type LimitOrder struct {
	Size            float64          `json:"size,omitempty"`
	Price           float64          `json:"price,omitempty"`
	PersistenceType EPersistenceType `json:"persistenceType,omitempty"`
	TimeInForce     ETimeInForce     `json:"timeInForce,omitempty"`
	MinFillSize     float64          `json:"minFillSize,omitempty"`
	BetTargetType   EBetTargetType   `json:"betTargetType,omitempty"`
	BetTargetSize   float64          `json:"betTargetSize,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

type MarketOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
}

// PlaceOrders to place new orders into market.
func (b *Betting) PlaceOrders(filter Filter) (placeExecutionReport PlaceExecutionReport, err error) {
	err = b.Request(&placeExecutionReport, BettingURL, "placeOrders", &filter)
	if err != nil {
		return
	}

	return
}


// PlaceOrders to place new orders into market.
func (b *Betting) ReplaceOrders(filter ReplaceFilter) (replaceExecutionReport ReplaceExecutionReport, err error) {
	err = b.ReplaceRequest(&replaceExecutionReport, BettingURL, "replaceOrders", &filter)
	if err != nil {
		return
	}

	return
}

// PlaceOrders to place new orders into market.
func (b *Betting) CancelOrders(filter CancelFilter) (cancelExecutionReport CancelExecutionReport, err error) {
	err = b.CancelRequest(&cancelExecutionReport, BettingURL, "cancelOrders", &filter)
	if err != nil {
		return
	}

	return
}
