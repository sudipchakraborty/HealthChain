import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { MsgCreateRecord } from "../ts-client/healthchain.healthchain/types/healthchain/healthchain/tx";

export async function connectKeplrAndSendRecord({ patientId, sys, dia }) {
  await window.keplr.experimentalSuggestChain({
    chainId: "healthchain",
    chainName: "Healthchain Local",
    rpc: "http://localhost:26657",
    rest: "http://localhost:1317",
    bip44: { coinType: 118 },
    bech32Config: {
      bech32PrefixAccAddr: "health",
      bech32PrefixAccPub: "healthpub",
      bech32PrefixValAddr: "healthvaloper",
      bech32PrefixValPub: "healthvaloperpub",
      bech32PrefixConsAddr: "healthvalcons",
      bech32PrefixConsPub: "healthvalconspub"
    },
    currencies: [{ coinDenom: "ahealth", coinMinimalDenom: "ahealth", decimals: 6 }],
    feeCurrencies: [{ coinDenom: "ahealth", coinMinimalDenom: "ahealth", decimals: 6 }],
    stakeCurrency: { coinDenom: "ahealth", coinMinimalDenom: "ahealth", decimals: 6 },
    gasPriceStep: { low: 0.01, average: 0.025, high: 0.03 }
  });

  await window.keplr.enable("healthchain");
  const offlineSigner = window.getOfflineSigner("healthchain");
  const accounts = await offlineSigner.getAccounts();
  const sender = accounts[0].address;

  // Register the message type
  const registry = new Registry([
    ["/healthchain.healthchain.MsgCreateRecord", MsgCreateRecord],
  ]);

  const client = await SigningStargateClient.connectWithSigner(
    "http://localhost:26657",
    offlineSigner,
    { registry }
  );

  const msg = {
    typeUrl: "/healthchain.healthchain.MsgCreateRecord",
    value: {
      creator: sender,
      patientId: String(patientId),
      sys: Number(sys),
      dia: Number(dia),
    }
  };

  const fee = {
    amount: [{ denom: "ahealth", amount: "5000" }],
    gas: "200000"
  };

  const result = await client.signAndBroadcast(sender, [msg], fee, "Create record from Dapp");

  // Success/failure only
  if (result.code === 0) {
    return { success: true, message: "Record created successfully!" };
  } else {
    return { success: false, message: `Tx failed: ${result.rawLog || "Unknown error"}` };
  }
}
