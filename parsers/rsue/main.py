import requests
from bs4 import BeautifulSoup
from urllib3.exceptions import InsecureRequestWarning

requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

try:
    response = requests.post("https://rsue.ru/raspisanie/", data={
        "f": 3,
        "k": 4,
        "g": 2,
    }, verify=False)
except Exception as e:
    print(e)

file = open("rsue.html", "w+")
file.write(response.text)
file.close()

soup = BeautifulSoup(response.text, "lxml")

soup = soup.find("div", {"id": "content"}).find_all("div", {"class": "container"})[1]

days = {
    "Понедельник": 1,
    "Вторник": 2,
    "Среда": 3,
    "Четверг": 4,
    "Пятница": 5,
    "Суббота": 6,
}

for elem in soup:
    print(elem.find())

# soup = soup.body.div.prettify()

# print(soup)
