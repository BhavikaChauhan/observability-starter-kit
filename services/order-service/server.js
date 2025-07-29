const express = require("express");
require("./src/telemetry");
const app = require("./src/app");

const server = express();

server.use(app);

server.listen(3002, () => {
  console.log("Order service listening on port 3002");
});