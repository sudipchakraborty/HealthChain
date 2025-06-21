package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"healthchain/x/healthchain/types"
)

func CmdCreateRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-record [patient-id] [sys] [dia]",
		Short: "Create a new record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPatientId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argSys, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argDia, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRecord(clientCtx.GetFromAddress().String(), argPatientId, argSys, argDia)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-record [id] [patient-id] [sys] [dia]",
		Short: "Update a record",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argPatientId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argSys, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			argDia, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRecord(clientCtx.GetFromAddress().String(), id, argPatientId, argSys, argDia)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-record [id]",
		Short: "Delete a record by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRecord(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
