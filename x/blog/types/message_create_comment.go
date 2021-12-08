package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateComment{}

func NewMsgCreateComment(creator string, postID uint64, title string, body string) *MsgCreateComment {
	return &MsgCreateComment{
		Creator: creator,
		PostID:  postID,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreateComment) Route() string {
	return RouterKey
}

func (msg *MsgCreateComment) Type() string {
	return "CreateComment"
}

func (msg *MsgCreateComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
