// badgestore.js
// This mock store simulates a database client for badges.
// It logs creation and returns a Badge object with a generated ID.

class BadgeStore {
  async Create(serialNumber, version) {
    console.log(`Created badge with SerialNumber: ${serialNumber}, Version: ${version}`);

    // Generate a pseudo ID using current timestamp
    const id = Date.now();

    return {
      id,
      serialNumber,
      version,
    };
  }
}

module.exports = { BadgeStore };
