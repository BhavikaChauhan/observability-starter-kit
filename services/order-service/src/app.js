const express = require("express");

const app = express();

app.get("/", (req, res) => {
  res.send("Order service running!");
});

module.exports = app;
