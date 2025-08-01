const express = require('express');
const cors = require('cors');
const path = require('path');

const app = express();
const PORT = 4002;

app.use(cors());
app.use(express.json());

// Serve frontend
app.use(express.static(path.join(__dirname, 'build')));

app.get('/orders', (req, res) => {
  res.json([{ id: 1, product: 'Book', quantity: 2 }]);
});

app.post('/orders', (req, res) => {
  const order = req.body;
  res.status(201).json({ message: 'Order created', order });
});

// Fallback to frontend
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`Order service running on http://localhost:${PORT}`);
});