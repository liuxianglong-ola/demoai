package code

const (
	RelayBuildRequestErr   = "relay.build_request_err"
	RelayRequestErr        = "relay.request_err"
	RelayTokenNoExistsErr  = "relay.token_no_exists_err"
	RelayModelCannotUseErr = "relay.model_cannot_use"
	RelayBadChannel        = "relay.bad_channel"
	RelayBadModel          = "relay.bad_model"
)

var relayMap = map[string]int{
	RelayBuildRequestErr:   1,
	RelayRequestErr:        2,
	RelayTokenNoExistsErr:  3,
	RelayModelCannotUseErr: 4,
	RelayBadChannel:        5,
	RelayBadModel:          6,
}
