const express = require("express");
require("./telemetry"); // Telemetry init

const app = express();

app.get("/order", (req, res) => {
  res.send("Order service called");
});

const PORT = process.env.PORT || 3003;

app.listen(PORT, () => {
  console.log(`🚀 Order service running on port ${PORT}`);
});