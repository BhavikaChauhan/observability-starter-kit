const express = require('express');
const path = require('path');
const app = express();
const PORT = process.env.PORT || 3003;

app.use(express.static(path.join(__dirname, '../public')));

// Health check
app.get('/health', (req, res) => res.send('Order service is healthy'));

// All other routes → index.html
app.get('*', (req, res) => {
  res.sendFile(path.resolve(__dirname, '../public', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`Order service running on port ${PORT}`);
});
