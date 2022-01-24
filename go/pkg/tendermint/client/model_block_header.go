/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BlockHeader struct for BlockHeader
type BlockHeader struct {
	Version BlockHeaderVersion `json:"version"`
	ChainId string `json:"chain_id"`
	Height string `json:"height"`
	Time string `json:"time"`
	LastBlockId BlockID `json:"last_block_id"`
	LastCommitHash string `json:"last_commit_hash"`
	DataHash string `json:"data_hash"`
	ValidatorsHash string `json:"validators_hash"`
	NextValidatorsHash string `json:"next_validators_hash"`
	ConsensusHash string `json:"consensus_hash"`
	AppHash string `json:"app_hash"`
	LastResultsHash string `json:"last_results_hash"`
	EvidenceHash string `json:"evidence_hash"`
	ProposerAddress string `json:"proposer_address"`
}

// NewBlockHeader instantiates a new BlockHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeader(version BlockHeaderVersion, chainId string, height string, time string, lastBlockId BlockID, lastCommitHash string, dataHash string, validatorsHash string, nextValidatorsHash string, consensusHash string, appHash string, lastResultsHash string, evidenceHash string, proposerAddress string) *BlockHeader {
	this := BlockHeader{}
	this.Version = version
	this.ChainId = chainId
	this.Height = height
	this.Time = time
	this.LastBlockId = lastBlockId
	this.LastCommitHash = lastCommitHash
	this.DataHash = dataHash
	this.ValidatorsHash = validatorsHash
	this.NextValidatorsHash = nextValidatorsHash
	this.ConsensusHash = consensusHash
	this.AppHash = appHash
	this.LastResultsHash = lastResultsHash
	this.EvidenceHash = evidenceHash
	this.ProposerAddress = proposerAddress
	return &this
}

// NewBlockHeaderWithDefaults instantiates a new BlockHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeaderWithDefaults() *BlockHeader {
	this := BlockHeader{}
	return &this
}

// GetVersion returns the Version field value
func (o *BlockHeader) GetVersion() BlockHeaderVersion {
	if o == nil {
		var ret BlockHeaderVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetVersionOk() (*BlockHeaderVersion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BlockHeader) SetVersion(v BlockHeaderVersion) {
	o.Version = v
}

// GetChainId returns the ChainId field value
func (o *BlockHeader) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetChainIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *BlockHeader) SetChainId(v string) {
	o.ChainId = v
}

// GetHeight returns the Height field value
func (o *BlockHeader) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockHeader) SetHeight(v string) {
	o.Height = v
}

// GetTime returns the Time field value
func (o *BlockHeader) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *BlockHeader) SetTime(v string) {
	o.Time = v
}

// GetLastBlockId returns the LastBlockId field value
func (o *BlockHeader) GetLastBlockId() BlockID {
	if o == nil {
		var ret BlockID
		return ret
	}

	return o.LastBlockId
}

// GetLastBlockIdOk returns a tuple with the LastBlockId field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetLastBlockIdOk() (*BlockID, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastBlockId, true
}

// SetLastBlockId sets field value
func (o *BlockHeader) SetLastBlockId(v BlockID) {
	o.LastBlockId = v
}

// GetLastCommitHash returns the LastCommitHash field value
func (o *BlockHeader) GetLastCommitHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastCommitHash
}

// GetLastCommitHashOk returns a tuple with the LastCommitHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetLastCommitHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastCommitHash, true
}

// SetLastCommitHash sets field value
func (o *BlockHeader) SetLastCommitHash(v string) {
	o.LastCommitHash = v
}

// GetDataHash returns the DataHash field value
func (o *BlockHeader) GetDataHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataHash
}

// GetDataHashOk returns a tuple with the DataHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetDataHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataHash, true
}

// SetDataHash sets field value
func (o *BlockHeader) SetDataHash(v string) {
	o.DataHash = v
}

// GetValidatorsHash returns the ValidatorsHash field value
func (o *BlockHeader) GetValidatorsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorsHash
}

// GetValidatorsHashOk returns a tuple with the ValidatorsHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetValidatorsHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidatorsHash, true
}

// SetValidatorsHash sets field value
func (o *BlockHeader) SetValidatorsHash(v string) {
	o.ValidatorsHash = v
}

// GetNextValidatorsHash returns the NextValidatorsHash field value
func (o *BlockHeader) GetNextValidatorsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextValidatorsHash
}

// GetNextValidatorsHashOk returns a tuple with the NextValidatorsHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetNextValidatorsHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextValidatorsHash, true
}

// SetNextValidatorsHash sets field value
func (o *BlockHeader) SetNextValidatorsHash(v string) {
	o.NextValidatorsHash = v
}

// GetConsensusHash returns the ConsensusHash field value
func (o *BlockHeader) GetConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsensusHash
}

// GetConsensusHashOk returns a tuple with the ConsensusHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetConsensusHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsensusHash, true
}

// SetConsensusHash sets field value
func (o *BlockHeader) SetConsensusHash(v string) {
	o.ConsensusHash = v
}

// GetAppHash returns the AppHash field value
func (o *BlockHeader) GetAppHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppHash
}

// GetAppHashOk returns a tuple with the AppHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetAppHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppHash, true
}

// SetAppHash sets field value
func (o *BlockHeader) SetAppHash(v string) {
	o.AppHash = v
}

// GetLastResultsHash returns the LastResultsHash field value
func (o *BlockHeader) GetLastResultsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastResultsHash
}

// GetLastResultsHashOk returns a tuple with the LastResultsHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetLastResultsHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastResultsHash, true
}

// SetLastResultsHash sets field value
func (o *BlockHeader) SetLastResultsHash(v string) {
	o.LastResultsHash = v
}

// GetEvidenceHash returns the EvidenceHash field value
func (o *BlockHeader) GetEvidenceHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvidenceHash
}

// GetEvidenceHashOk returns a tuple with the EvidenceHash field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetEvidenceHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EvidenceHash, true
}

// SetEvidenceHash sets field value
func (o *BlockHeader) SetEvidenceHash(v string) {
	o.EvidenceHash = v
}

// GetProposerAddress returns the ProposerAddress field value
func (o *BlockHeader) GetProposerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProposerAddress
}

// GetProposerAddressOk returns a tuple with the ProposerAddress field value
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetProposerAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProposerAddress, true
}

// SetProposerAddress sets field value
func (o *BlockHeader) SetProposerAddress(v string) {
	o.ProposerAddress = v
}

func (o BlockHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["chain_id"] = o.ChainId
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["last_block_id"] = o.LastBlockId
	}
	if true {
		toSerialize["last_commit_hash"] = o.LastCommitHash
	}
	if true {
		toSerialize["data_hash"] = o.DataHash
	}
	if true {
		toSerialize["validators_hash"] = o.ValidatorsHash
	}
	if true {
		toSerialize["next_validators_hash"] = o.NextValidatorsHash
	}
	if true {
		toSerialize["consensus_hash"] = o.ConsensusHash
	}
	if true {
		toSerialize["app_hash"] = o.AppHash
	}
	if true {
		toSerialize["last_results_hash"] = o.LastResultsHash
	}
	if true {
		toSerialize["evidence_hash"] = o.EvidenceHash
	}
	if true {
		toSerialize["proposer_address"] = o.ProposerAddress
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeader struct {
	value *BlockHeader
	isSet bool
}

func (v NullableBlockHeader) Get() *BlockHeader {
	return v.value
}

func (v *NullableBlockHeader) Set(val *BlockHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeader(val *BlockHeader) *NullableBlockHeader {
	return &NullableBlockHeader{value: val, isSet: true}
}

func (v NullableBlockHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

