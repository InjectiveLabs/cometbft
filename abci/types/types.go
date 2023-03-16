package types

import (
	"bytes"
	"encoding/json"

	"github.com/cosmos/gogoproto/jsonpb"

	"github.com/cometbft/cometbft/abci/types/v1"
	"github.com/cometbft/cometbft/abci/types/v2"
	"github.com/cometbft/cometbft/abci/types/v3"
)

type Request = v3.Request
type RequestEcho = v1.RequestEcho
type RequestFlush = v1.RequestFlush
type RequestInfo = v2.RequestInfo
type RequestInitChain = v3.RequestInitChain
type RequestQuery = v1.RequestQuery
type RequestCheckTx = v1.RequestCheckTx
type RequestCommit = v1.RequestCommit
type RequestListSnapshots = v1.RequestListSnapshots
type RequestOfferSnapshot = v1.RequestOfferSnapshot
type RequestLoadSnapshotChunk = v1.RequestLoadSnapshotChunk
type RequestApplySnapshotChunk = v1.RequestApplySnapshotChunk
type RequestPrepareProposal = v2.RequestPrepareProposal
type RequestProcessProposal = v2.RequestProcessProposal
type RequestExtendVote = v3.RequestExtendVote
type RequestVerifyVoteExtension = v3.RequestVerifyVoteExtension
type RequestFinalizeBlock = v3.RequestFinalizeBlock

type Response = v3.Response
type ResponseException = v1.ResponseException
type ResponseEcho = v1.ResponseEcho
type ResponseFlush = v1.ResponseFlush
type ResponseInfo = v1.ResponseInfo
type ResponseInitChain = v3.ResponseInitChain
type ResponseQuery = v1.ResponseQuery
type ResponseCheckTx = v3.ResponseCheckTx
type ResponseCommit = v3.ResponseCommit
type ResponseListSnapshots = v1.ResponseListSnapshots
type ResponseOfferSnapshot = v1.ResponseOfferSnapshot
type ResponseLoadSnapshotChunk = v1.ResponseLoadSnapshotChunk
type ResponseApplySnapshotChunk = v1.ResponseApplySnapshotChunk
type ResponsePrepareProposal = v2.ResponsePrepareProposal
type ResponseProcessProposal = v2.ResponseProcessProposal
type ResponseExtendVote = v3.ResponseExtendVote
type ResponseVerifyVoteExtension = v3.ResponseVerifyVoteExtension
type ResponseFinalizeBlock = v3.ResponseFinalizeBlock

type ExecTxResult = v3.ExecTxResult
type EventAttribute = v2.EventAttribute

type ValidatorUpdate = v1.ValidatorUpdate

type ResponseProcessProposal_ProposalStatus = v2.ResponseProcessProposal_ProposalStatus

const (
	ResponseProcessProposal_UNKNOWN ResponseProcessProposal_ProposalStatus = v2.ResponseProcessProposal_UNKNOWN
	ResponseProcessProposal_ACCEPT  ResponseProcessProposal_ProposalStatus = v2.ResponseProcessProposal_ACCEPT
	ResponseProcessProposal_REJECT  ResponseProcessProposal_ProposalStatus = v2.ResponseProcessProposal_REJECT
)

type ResponseVerifyVoteExtension_VerifyStatus = v3.ResponseVerifyVoteExtension_VerifyStatus

const (
	ResponseVerifyVoteExtension_UNKNOWN ResponseVerifyVoteExtension_VerifyStatus = v3.ResponseVerifyVoteExtension_UNKNOWN
	ResponseVerifyVoteExtension_ACCEPT  ResponseVerifyVoteExtension_VerifyStatus = v3.ResponseVerifyVoteExtension_ACCEPT
	ResponseVerifyVoteExtension_REJECT 	ResponseVerifyVoteExtension_VerifyStatus = v3.ResponseVerifyVoteExtension_REJECT
)

const (
	CodeTypeOK uint32 = 0
)

// IsOK returns true if Code is OK.
func (r ResponseCheckTx) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseCheckTx) IsErr() bool {
	return r.Code != CodeTypeOK
}

// IsOK returns true if Code is OK.
func (r ExecTxResult) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ExecTxResult) IsErr() bool {
	return r.Code != CodeTypeOK
}

// IsOK returns true if Code is OK.
func (r ResponseQuery) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseQuery) IsErr() bool {
	return r.Code != CodeTypeOK
}

// IsAccepted returns true if Code is ACCEPT
func (r ResponseProcessProposal) IsAccepted() bool {
	return r.Status == ResponseProcessProposal_ACCEPT
}

// IsStatusUnknown returns true if Code is UNKNOWN
func (r ResponseProcessProposal) IsStatusUnknown() bool {
	return r.Status == ResponseProcessProposal_UNKNOWN
}

func (r ResponseVerifyVoteExtension) IsAccepted() bool {
	return r.Status == ResponseVerifyVoteExtension_ACCEPT
}

// IsStatusUnknown returns true if Code is Unknown
func (r ResponseVerifyVoteExtension) IsStatusUnknown() bool {
	return r.Status == ResponseVerifyVoteExtension_UNKNOWN
}

//---------------------------------------------------------------------------
// override JSON marshaling so we emit defaults (ie. disable omitempty)

var (
	jsonpbMarshaller = jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	jsonpbUnmarshaller = jsonpb.Unmarshaler{}
)

func (r *ResponseCheckTx) MarshalJSON() ([]byte, error) {
	s, err := jsonpbMarshaller.MarshalToString(r)
	return []byte(s), err
}

func (r *ResponseCheckTx) UnmarshalJSON(b []byte) error {
	reader := bytes.NewBuffer(b)
	return jsonpbUnmarshaller.Unmarshal(reader, r)
}

func (r *ExecTxResult) MarshalJSON() ([]byte, error) {
	s, err := jsonpbMarshaller.MarshalToString(r)
	return []byte(s), err
}

func (r *ExecTxResult) UnmarshalJSON(b []byte) error {
	reader := bytes.NewBuffer(b)
	return jsonpbUnmarshaller.Unmarshal(reader, r)
}

func (r *ResponseQuery) MarshalJSON() ([]byte, error) {
	s, err := jsonpbMarshaller.MarshalToString(r)
	return []byte(s), err
}

func (r *ResponseQuery) UnmarshalJSON(b []byte) error {
	reader := bytes.NewBuffer(b)
	return jsonpbUnmarshaller.Unmarshal(reader, r)
}

func (r *ResponseCommit) MarshalJSON() ([]byte, error) {
	s, err := jsonpbMarshaller.MarshalToString(r)
	return []byte(s), err
}

func (r *ResponseCommit) UnmarshalJSON(b []byte) error {
	reader := bytes.NewBuffer(b)
	return jsonpbUnmarshaller.Unmarshal(reader, r)
}

func (r *EventAttribute) MarshalJSON() ([]byte, error) {
	s, err := jsonpbMarshaller.MarshalToString(r)
	return []byte(s), err
}

func (r *EventAttribute) UnmarshalJSON(b []byte) error {
	reader := bytes.NewBuffer(b)
	return jsonpbUnmarshaller.Unmarshal(reader, r)
}

// Some compile time assertions to ensure we don't
// have accidental runtime surprises later on.

// jsonEncodingRoundTripper ensures that asserted
// interfaces implement both MarshalJSON and UnmarshalJSON
type jsonRoundTripper interface {
	json.Marshaler
	json.Unmarshaler
}

var _ jsonRoundTripper = (*ResponseCommit)(nil)
var _ jsonRoundTripper = (*ResponseQuery)(nil)
var _ jsonRoundTripper = (*ExecTxResult)(nil)
var _ jsonRoundTripper = (*ResponseCheckTx)(nil)

var _ jsonRoundTripper = (*EventAttribute)(nil)

// deterministicExecTxResult constructs a copy of response that omits
// non-deterministic fields. The input response is not modified.
func deterministicExecTxResult(response *ExecTxResult) *ExecTxResult {
	return &ExecTxResult{
		Code:      response.Code,
		Data:      response.Data,
		GasWanted: response.GasWanted,
		GasUsed:   response.GasUsed,
	}
}

// MarshalTxResults encodes the the TxResults as a list of byte
// slices. It strips off the non-deterministic pieces of the TxResults
// so that the resulting data can be used for hash comparisons and used
// in Merkle proofs.
func MarshalTxResults(r []*ExecTxResult) ([][]byte, error) {
	s := make([][]byte, len(r))
	for i, e := range r {
		d := deterministicExecTxResult(e)
		b, err := d.Marshal()
		if err != nil {
			return nil, err
		}
		s[i] = b
	}
	return s, nil
}

// -----------------------------------------------
// construct Result data
