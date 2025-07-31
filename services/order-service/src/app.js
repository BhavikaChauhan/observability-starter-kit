import React from "react";
import ReactDOM from "react-dom/client";
import "./telemetry";

function App() {
  return <h1>Order Service React App (OTEL Instrumented)</h1>;
}

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<App />);