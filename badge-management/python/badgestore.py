# badgestore.py
# This mock store simulates a database client for badges.
# It logs creation and returns a Badge object with a generated ID.

import time

class BadgeStore:
    def create(self, serial_number: str, version: str) -> dict:
        print(f"Created badge with SerialNumber: {serial_number}, Version: {version}")
        badge_id = int(time.time() * 1000)  # pseudo ID
        return {
            "id": badge_id,
            "serialNumber": serial_number,
            "version": version,
        }

def new() -> BadgeStore:
    return BadgeStore()