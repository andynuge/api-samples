// server.js
const { createApp } = require('./app');

const PORT = process.env.PORT || 8080;
const app = createApp();

app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`);
});
