const opentelemetry = require("@opentelemetry/sdk-node");

const sdk = new opentelemetry.NodeSDK({
  serviceName: "order-service"
});

sdk.start();
