from time import sleep
from dotenv import load_dotenv
import requests
import parser
import os

if __name__ == "__main__":
    load_dotenv()
    server_host = os.getenv("SERVER_HOST")
    schedules = parser.parse_schedule_page()

    print("sending schedules...")
    while True:
        try:
            response = requests.post(
                f"http://{server_host}/upload", json=schedules)
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
