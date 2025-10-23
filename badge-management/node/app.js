// app.js
// NOTE: This sample intentionally omits validation, JSON error shapes,
// and dependency wiring so candidates can discuss and propose improvements.

const express = require('express');
const { BadgeStore } = require('./badgestore');

class Handler {
  constructor() {
    this.badgeStore = new BadgeStore();
    this.registerBadge = this.registerBadge.bind(this);
  }

  async registerBadge(req, res) {
    try {
      const badge = await this.badgeStore.Create(req.body.serialNumber, req.body.version);
      res.status(201).json(badge);
    } catch (e) {
      res.status(500).type('text').send('failed to create badge');
    }
  }
}

// Initializes and returns an Express app (the common pattern)
function createApp() {
  const app = express();
  const handler = new Handler();

  app.use(express.json());
  app.post('/badges', handler.registerBadge);

  return app;
}

module.exports = { createApp };
