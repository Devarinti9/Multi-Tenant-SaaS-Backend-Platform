import json
import os
from urllib import request

BASE_URL = os.getenv("APP_BASE_URL", "http://localhost:8080")
TOKEN = os.getenv("JWT_TOKEN", "")


def post(path: str, payload: dict, auth: bool = False):
    req = request.Request(
        url=f"{BASE_URL}{path}",
        data=json.dumps(payload).encode("utf-8"),
        headers={
            "Content-Type": "application/json",
            **({"Authorization": f"Bearer {TOKEN}"} if auth and TOKEN else {}),
        },
        method="POST",
    )
    with request.urlopen(req) as response:
        print(path, response.status, response.read().decode())


if __name__ == "__main__":
    post("/api/v1/auth/register", {
        "name": "Demo Admin",
        "email": "admin@example.com",
        "password": "Password123!",
        "role": "admin",
        "organization_id": 1,
    })
