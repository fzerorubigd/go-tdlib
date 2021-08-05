// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// BotInfo Provides information about a bot and its supported commands
type BotInfo struct {
	tdCommon
	Description string       `json:"description"` // Long description shown on the user info page
	Commands    []BotCommand `json:"commands"`    // A list of commands supported by the bot
}

// MessageType return the string telegram-type of BotInfo
func (botInfo *BotInfo) MessageType() string {
	return "botInfo"
}

// NewBotInfo creates a new BotInfo
//
// @param description Long description shown on the user info page
// @param commands A list of commands supported by the bot
func NewBotInfo(description string, commands []BotCommand) *BotInfo {
	botInfoTemp := BotInfo{
		tdCommon:    tdCommon{Type: "botInfo"},
		Description: description,
		Commands:    commands,
	}

	return &botInfoTemp
}
