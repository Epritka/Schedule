import requests
import urllib.parse

API_ENDPOINT = "https://edu.donstu.ru/api"

 requests.get(urllib.parse.urljoin(API_ENDPOINT, "raspGrouplist"))

