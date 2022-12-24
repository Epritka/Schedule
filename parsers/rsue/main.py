from time import sleep
from dotenv import load_dotenv
import requests
import proccessor
import json
import os

if __name__ == "__main__":
    print("start parser")

    load_dotenv()
    server_host = os.getenv("SERVER_HOST")
    schedules = proccessor.parse_schedule_page()

    file = open("rsue.json", "w+", encoding="utf-8")
    file.write(json.dumps(schedules, ensure_ascii=False))
    file.close()

    print("sending schedules...")
    while True:
        try:
            response = requests.post(
                f"http://{server_host}/api/v1/schedule/", json=schedules)
        except Exception as e:
            print("failed to send:", e)
            sleep(5)
            continue

        code = response.status_code
        if code < 300 and code >= 200:
            break
        else:
            print("failed to send:", code, response.text)
        exit(1)
    print("completed...")
