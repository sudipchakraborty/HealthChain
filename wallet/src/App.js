import React, { useState } from 'react';
import './App.css';
import { StargateClient } from "@cosmjs/stargate";

const App = () => {
  const [address, setAddress] = useState('');
  const [balance, setBalance] = useState('');

  const chainInfo = {
    chainId: "healthchain",
    chainName: "HealthChain",
    rpc: "http://localhost:26657",         // Adjust if needed
    rest: "http://localhost:1317",         // Adjust if needed
    bip44: {
      coinType: 118,
    },
    bech32Config: {
      bech32PrefixAccAddr: "health",
      bech32PrefixAccPub: "healthpub",
      bech32PrefixValAddr: "healthvaloper",
      bech32PrefixValPub: "healthvaloperpub",
      bech32PrefixConsAddr: "healthvalcons",
      bech32PrefixConsPub: "healthvalconspub",
    },
    currencies: [
      {
        coinDenom: "TOKEN",
        coinMinimalDenom: "token",
        coinDecimals: 6,
      },
    ],
    feeCurrencies: [
      {
        coinDenom: "TOKEN",
        coinMinimalDenom: "token",
        coinDecimals: 6,
      },
    ],
    stakeCurrency: {
      coinDenom: "TOKEN",
      coinMinimalDenom: "token",
      coinDecimals: 6,
    },
    features: ["stargate"],
  };

  const connectWallet = async () => {
    if (!window.keplr) {
      alert("Keplr extension not installed.");
      return;
    }

    // Suggest the chain to Keplr
    await window.keplr.experimentalSuggestChain(chainInfo);

    // Enable the chain
    await window.keplr.enable(chainInfo.chainId);

    const offlineSigner = window.getOfflineSigner(chainInfo.chainId);
    const accounts = await offlineSigner.getAccounts();

    setAddress(accounts[0].address);

    const client = await StargateClient.connect(chainInfo.rpc);
    const balance = await client.getBalance(accounts[0].address, "token");
    setBalance(`${balance.amount} ${balance.denom}`);
  };

  return (
    <div className="App">
      <h2>Connect to HealthChain</h2>
      <button onClick={connectWallet}>Connect Keplr</button>
      {address && <p><b>Address:</b> {address}</p>}
      {balance && <p><b>Balance:</b> {balance}</p>}
    </div>
  );
};

export default App;
