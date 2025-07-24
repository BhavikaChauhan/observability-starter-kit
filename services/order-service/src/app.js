const express = require("express");
require("./telemetry"); // Telemetry init
const app = express();

app.get("/order", (req, res) => {
  res.send("Order service called");
});

app.listen(5002, () => {
  console.log("Order service running on port 5002");
});