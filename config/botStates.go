package config

type State struct {
	CurrentState string
	StateData    map[string]interface{}
}

var userStates = map[int64]State{}

func SetState(chatID int64, stateKey string) {
	oldState, exists := userStates[chatID]
	if !exists {
		oldState = State{StateData: make(map[string]interface{})}
	}
	userStates[chatID] = State{stateKey, oldState.StateData}
}

func GetCurrentState(chatID int64) interface{} {
	if state, exists := userStates[chatID]; exists {
		return state.CurrentState
	}
	return nil
}

func SetStateData(chatID int64, key string, value interface{}) {
	state := userStates[chatID]
	state.StateData[key] = value
	userStates[chatID] = state
}

func GetStateData(chatID int64) map[string]interface{} {
	if state, exists := userStates[chatID]; exists {
		return state.StateData
	}
	return make(map[string]interface{})
}

func ClearState(chatID int64) {
	if _, exists := userStates[chatID]; exists {
		delete(userStates, chatID)
	}
}
