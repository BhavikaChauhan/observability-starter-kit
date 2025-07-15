// server.js

const express = require('express');
const path = require('path');

const app = express();
const PORT = 8000;

// Serve static files from the React app
app.use(express.static(path.join(__dirname, 'build')));

// Health route
app.get('/health', (req, res) => {
  res.send('Order Service is healthy!');
});

// All other routes → serve React index.html
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`Order Service running at http://localhost:${PORT}`);
});
