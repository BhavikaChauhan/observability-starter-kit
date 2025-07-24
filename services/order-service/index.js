const express = require("express");
const app = express();

app.get("/", (req, res) => {
  res.send("Order Service Working!");
});

app.listen(3003, () => {
  console.log("Order service running on port 3003");
});
