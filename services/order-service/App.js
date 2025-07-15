const express = require('express');
const app = express();
const PORT = 8000;

app.get('/health', (req, res) => {
  res.send('Order Service is healthy!');
});

app.listen(PORT, () => {
  console.log(`Order Service running on port ${PORT}`);
});
