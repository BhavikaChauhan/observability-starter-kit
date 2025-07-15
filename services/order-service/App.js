const express = require('express');
const app = express();
const port = 8000;

app.get('/health', (req, res) => {
  res.send('Order Service is healthy!');
});

app.listen(port, () => {
  console.log(`Order Service running at http://localhost:${port}`);
});