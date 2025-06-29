import React, { useState } from "react";
import { connectKeplrAndSendRecord } from "./chain";

function App() {
  const [patientId, setPatientId] = useState("");
  const [sys, setSys] = useState("");
  const [dia, setDia] = useState("");
  const [status, setStatus] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setStatus("Sending...");
    try {
      const res = await connectKeplrAndSendRecord({
        patientId,
        sys,
        dia,
      });
      setStatus(res.message);
    } catch (err) {
      setStatus("Error: " + err.message);
    }
  };

  return (
    <div style={{ padding: 32 }}>
      <h2>Healthchain DApp</h2>
      <form onSubmit={handleSubmit}>
        <input
          placeholder="Patient ID"
          value={patientId}
          onChange={e => setPatientId(e.target.value)}
        />
        <input
          placeholder="Sys"
          value={sys}
          type="number"
          onChange={e => setSys(e.target.value)}
        />
        <input
          placeholder="Dia"
          value={dia}
          type="number"
          onChange={e => setDia(e.target.value)}
        />
        <button type="submit">Send Record</button>
      </form>
      <div style={{ marginTop: 16 }}>{status}</div>
    </div>
  );
}

export default App;
