package common

import (
	"fmt"
)

// MessageKind is the kind of a message (that is passed between the Node and Arwen)
type MessageKind uint32

const (
	FirstKind MessageKind = iota
	Stop
	ContractDeployRequest
	ContractCallRequest
	ContractResponse
	BlockchainAccountExistsRequest
	BlockchainAccountExistsResponse
	BlockchainNewAddressRequest
	BlockchainNewAddressResponse
	BlockchainGetBalanceRequest
	BlockchainGetBalanceResponse
	BlockchainGetNonceRequest
	BlockchainGetNonceResponse
	BlockchainGetStorageDataRequest
	BlockchainGetStorageDataResponse
	BlockchainIsCodeEmptyRequest
	BlockchainIsCodeEmptyResponse
	BlockchainGetCodeRequest
	BlockchainGetCodeResponse
	BlockchainGetBlockhashRequest
	BlockchainGetBlockhashResponse
	BlockchainLastNonceRequest
	BlockchainLastNonceResponse
	BlockchainLastRoundRequest
	BlockchainLastRoundResponse
	BlockchainLastTimeStampRequest
	BlockchainLastTimeStampResponse
	BlockchainLastRandomSeedRequest
	BlockchainLastRandomSeedResponse
	BlockchainLastEpochRequest
	BlockchainLastEpochResponse
	BlockchainGetStateRootHashRequest
	BlockchainGetStateRootHashResponse
	BlockchainCurrentNonceRequest
	BlockchainCurrentNonceResponse
	BlockchainCurrentRoundRequest
	BlockchainCurrentRoundResponse
	BlockchainCurrentTimeStampRequest
	BlockchainCurrentTimeStampResponse
	BlockchainCurrentRandomSeedRequest
	BlockchainCurrentRandomSeedResponse
	BlockchainCurrentEpochRequest
	BlockchainCurrentEpochResponse
	DiagnoseWaitRequest
	DiagnoseWaitResponse
	LastKind
)

var messageKindNameByID = map[MessageKind]string{}

func init() {
	messageKindNameByID[FirstKind] = "FirstKind"
	messageKindNameByID[Stop] = "Stop"
	messageKindNameByID[ContractDeployRequest] = "ContractDeployRequest"
	messageKindNameByID[ContractCallRequest] = "ContractCallRequest"
	messageKindNameByID[ContractResponse] = "ContractResponse"
	messageKindNameByID[BlockchainAccountExistsRequest] = "BlockchainAccountExistsRequest"
	messageKindNameByID[BlockchainAccountExistsResponse] = "BlockchainAccountExistsResponse"
	messageKindNameByID[BlockchainNewAddressRequest] = "BlockchainNewAddressRequest"
	messageKindNameByID[BlockchainNewAddressResponse] = "BlockchainNewAddressResponse"
	messageKindNameByID[BlockchainGetBalanceRequest] = "BlockchainGetBalanceRequest"
	messageKindNameByID[BlockchainGetBalanceResponse] = "BlockchainGetBalanceResponse"
	messageKindNameByID[BlockchainGetNonceRequest] = "BlockchainGetNonceRequest"
	messageKindNameByID[BlockchainGetNonceResponse] = "BlockchainGetNonceResponse"
	messageKindNameByID[BlockchainGetStorageDataRequest] = "BlockchainGetStorageDataRequest"
	messageKindNameByID[BlockchainGetStorageDataResponse] = "BlockchainGetStorageDataResponse"
	messageKindNameByID[BlockchainIsCodeEmptyRequest] = "BlockchainIsCodeEmptyRequest"
	messageKindNameByID[BlockchainIsCodeEmptyResponse] = "BlockchainIsCodeEmptyResponse"
	messageKindNameByID[BlockchainGetCodeRequest] = "BlockchainGetCodeRequest"
	messageKindNameByID[BlockchainGetCodeResponse] = "BlockchainGetCodeResponse"
	messageKindNameByID[BlockchainGetBlockhashRequest] = "BlockchainGetBlockhashRequest"
	messageKindNameByID[BlockchainGetBlockhashResponse] = "BlockchainGetBlockhashResponse"
	messageKindNameByID[BlockchainLastNonceRequest] = "BlockchainLastNonceRequest"
	messageKindNameByID[BlockchainLastNonceResponse] = "BlockchainLastNonceResponse"
	messageKindNameByID[BlockchainLastRoundRequest] = "BlockchainLastRoundRequest"
	messageKindNameByID[BlockchainLastRoundResponse] = "BlockchainLastRoundResponse"
	messageKindNameByID[BlockchainLastTimeStampRequest] = "BlockchainLastTimeStampRequest"
	messageKindNameByID[BlockchainLastTimeStampResponse] = "BlockchainLastTimeStampResponse"
	messageKindNameByID[BlockchainLastRandomSeedRequest] = "BlockchainLastRandomSeedRequest"
	messageKindNameByID[BlockchainLastRandomSeedResponse] = "BlockchainLastRandomSeedResponse"
	messageKindNameByID[BlockchainLastEpochRequest] = "BlockchainLastEpochRequest"
	messageKindNameByID[BlockchainLastEpochResponse] = "BlockchainLastEpochResponse"
	messageKindNameByID[BlockchainGetStateRootHashRequest] = "BlockchainGetStateRootHashRequest"
	messageKindNameByID[BlockchainGetStateRootHashResponse] = "BlockchainGetStateRootHashResponse"
	messageKindNameByID[BlockchainCurrentNonceRequest] = "BlockchainCurrentNonceRequest"
	messageKindNameByID[BlockchainCurrentNonceResponse] = "BlockchainCurrentNonceResponse"
	messageKindNameByID[BlockchainCurrentRoundRequest] = "BlockchainCurrentRoundRequest"
	messageKindNameByID[BlockchainCurrentRoundResponse] = "BlockchainCurrentRoundResponse"
	messageKindNameByID[BlockchainCurrentTimeStampRequest] = "BlockchainCurrentTimeStampRequest"
	messageKindNameByID[BlockchainCurrentTimeStampResponse] = "BlockchainCurrentTimeStampResponse"
	messageKindNameByID[BlockchainCurrentRandomSeedRequest] = "BlockchainCurrentRandomSeedRequest"
	messageKindNameByID[BlockchainCurrentRandomSeedResponse] = "BlockchainCurrentRandomSeedResponse"
	messageKindNameByID[BlockchainCurrentEpochRequest] = "BlockchainCurrentEpochRequest"
	messageKindNameByID[BlockchainCurrentEpochResponse] = "BlockchainCurrentEpochResponse"
	messageKindNameByID[DiagnoseWaitRequest] = "DiagnoseWaitRequest"
	messageKindNameByID[DiagnoseWaitResponse] = "DiagnoseWaitResponse"
	messageKindNameByID[LastKind] = "LastKind"
}

// MessageHandler is a message abstraction
type MessageHandler interface {
	GetNonce() uint32
	SetNonce(nonce uint32)
	GetKind() MessageKind
	SetKind(kind MessageKind)
	GetError() error
	SetError(err error)
}

// Message is the implementation of the abstraction
type Message struct {
	DialogueNonce uint32
	Kind          MessageKind
	ErrorMessage  string
}

// GetNonce gets the dialogue nonce
func (message *Message) GetNonce() uint32 {
	return message.DialogueNonce
}

// SetNonce sets the dialogue nonce
func (message *Message) SetNonce(nonce uint32) {
	message.DialogueNonce = nonce
}

// GetKind gets the message kind
func (message *Message) GetKind() MessageKind {
	return message.Kind
}

// SetKind sets the message kind
func (message *Message) SetKind(kind MessageKind) {
	message.Kind = kind
}

// GetError gets the error within the message
func (message *Message) GetError() error {
	if message.ErrorMessage == "" {
		return nil
	}

	return fmt.Errorf(message.ErrorMessage)
}

// SetError sets the error within the message
func (message *Message) SetError(err error) {
	if err != nil {
		message.ErrorMessage = err.Error()
	}
}

func (message *Message) String() string {
	kindName, _ := messageKindNameByID[message.Kind]
	return fmt.Sprintf("[kind=%s nonce=%d err=%s]", kindName, message.DialogueNonce, message.ErrorMessage)
}

// MessageStop is a message sent by Node to stop Arwen
type MessageStop struct {
	Message
}

// NewMessageStop creates a new message
func NewMessageStop() *MessageStop {
	message := &MessageStop{}
	message.Kind = Stop
	return message
}

// MessageReplier is a callback signature
type MessageReplier func(MessageHandler) MessageHandler

func noopReplier(message MessageHandler) MessageHandler {
	panic("NO-OP replier called")
}

// CreateReplySlots creates a slice of no-operation repliers, to be substituted with actual repliers (by message listeners)
func CreateReplySlots() []MessageReplier {
	slots := make([]MessageReplier, LastKind)
	for i := 0; i < len(slots); i++ {
		slots[i] = noopReplier
	}

	return slots
}

// IsHookCall returns whether a message is a hook call
func IsHookCall(message MessageHandler) bool {
	kind := message.GetKind()
	return kind >= BlockchainAccountExistsRequest && kind <= BlockchainCurrentEpochResponse
}

// IsStopRequest returns whether a message is a stop request
func IsStopRequest(message MessageHandler) bool {
	return message.GetKind() == Stop
}

// IsContractResponse returns whether a message is a contract response
func IsContractResponse(message MessageHandler) bool {
	return message.GetKind() == ContractResponse
}

// IsDiagnose returns whether a message is a diagnose request
func IsDiagnose(message MessageHandler) bool {
	kind := message.GetKind()
	return kind >= DiagnoseWaitRequest && kind <= DiagnoseWaitResponse
}
