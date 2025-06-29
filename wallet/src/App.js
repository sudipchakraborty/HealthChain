import React from "react";

function App() {
  const addChainToKeplr = async () => {
    if (!window.keplr) {
      alert("❌ Keplr not installed! Please install the Keplr browser extension.");
      return;
    }

    try {
      await window.keplr.experimentalSuggestChain({
        chainId: "healthchain",
        chainName: "HealthChain",
        rpc: "http://localhost:26657",
        rest: "http://localhost:1317",
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
            gasPriceStep: {
              low: 0.00001,      // <<<< Ultra-low fees for test/dev
              average: 0.0001,
              high: 0.001,
            },
          },
        ],
        stakeCurrency: {
          coinDenom: "TOKEN",
          coinMinimalDenom: "token",
          coinDecimals: 6,
        },
      });

      await window.keplr.enable("healthchain");
      const offlineSigner = window.getOfflineSigner("healthchain");
      const accounts = await offlineSigner.getAccounts();

      alert(`✅ Chain added and wallet connected!\nYour address: ${accounts[0].address}`);
    } catch (error) {
      console.error("❌ Failed to suggest chain:", error);
      alert("❌ Failed to add chain to Keplr. Check the console for more details.");
    }
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "Arial, sans-serif" }}>
      <h1>🛠️ Add HealthChain to Keplr</h1>
      <button
        onClick={addChainToKeplr}
        style={{
          fontSize: "1rem",
          padding: "0.75rem 1.5rem",
          cursor: "pointer",
          backgroundColor: "#4caf50",
          color: "white",
          border: "none",
          borderRadius: "5px",
        }}
      >
        ➕ Add Chain to Keplr
      </button>
    </div>
  );
}

export default App;