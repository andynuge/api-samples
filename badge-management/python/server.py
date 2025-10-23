# app.py
# NOTE: This sample intentionally omits validation, JSON error shapes,
# and deeper dependency wiring so candidates can discuss and propose improvements.

from flask import Flask, request, jsonify
from badgestore import BadgeStore, new as new_badgestore


class Handler:
    def __init__(self, badge_store: BadgeStore):
        self.badge_store = badge_store

    def register_badge(self):
        if request.method != "POST":
            return "method not allowed", 405

        # Flask parses JSON via request.get_json(); we keep errors plain-text intentionally
        data = request.get_json(silent=True)
        if data is None:
            return "invalid JSON body", 400

        serial_number = data.get("serialNumber")
        version = data.get("version")

        try:
            badge = self.badge_store.create(serial_number, version)
            # Intentionally using numeric status code; candidates can discuss Location header, etc.
            return jsonify(badge), 201
        except Exception:
            return "failed to create badge", 500

def create_app() -> Flask:
    app = Flask(__name__)

    store = new_badgestore()
    handler = Handler(store)

    # Route definition: mirrors POST /badges
    app.add_url_rule("/badges", view_func=handler.register_badge, methods=["POST"])

    return app

if __name__ == "__main__":
    app = create_app()
    # Default port 8080 to match your other samples
    print("Listening on :8080")
    app.run(host="0.0.0.0", port=8080)
