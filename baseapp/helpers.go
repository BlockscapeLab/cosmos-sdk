package baseapp

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (app *BaseApp) Check(tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	return app.runTx(runTxModeCheck, nil, tx)
}

func (app *BaseApp) Simulate(txBytes []byte, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	return app.runTx(runTxModeSimulate, txBytes, tx)
}

func (app *BaseApp) Deliver(tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	return app.runTx(runTxModeDeliver, nil, tx)
}

// Context with current {check, deliver}State of the app used by tests.
func (app *BaseApp) NewContext(isCheckTx bool, header abci.Header) sdk.Context {
	if header.Height < 0 {
		return sdk.NewContext(app.cms, header, true, app.logger)
	}
	if isCheckTx {
		return sdk.NewContext(app.checkState.ms, header, true, app.logger).
			WithMinGasPrices(app.minGasPrices)
	}

	return sdk.NewContext(app.deliverState.ms, header, false, app.logger)
}

func (app *BaseApp) NewUncachedContext(isCheckTx bool, header abci.Header) sdk.Context {
	return sdk.NewContext(app.cms, header, isCheckTx, app.logger)
}
