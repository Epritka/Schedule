from urllib3.exceptions import InsecureRequestWarning
from bs4 import BeautifulSoup as BS
import requests
import json
import dto as dto

requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

RSUE_URL = "https://rsue.ru/raspisanie/"


def send_request(url: str, data):
    try:
        return requests.post(url, data=data, verify=False)
    except Exception as e:
        print(e)


def fill_schedule_from_page(page, schedule: dto.ScheduleDTO):
    container = BS(page, "lxml").find("div", {"id": "content"}).find_all(
        "div", {"class": "container"})[1]

    day_number = 1
    week_type = 1
    for child in container.children:
        if child.name == "h1":
            if child.text == "Четная неделя":
                week_type = 0

        if child.name == "div":
            for d in child.children:
                lessons = d.find_all("div", {"class": "day"})

                day = dto.DayDTO(day_number)

                for l in lessons:
                    info = l.find_all("span")

                    time = str(info[0].contents[0]).replace(" ", "").split("—")

                    lesson = dto.LessonDTO(
                        str(time[0]),
                        str(time[1]),
                        str(info[2].text),
                        str(info[3].text),
                        str(info[4].text),
                        str(info[5].text),
                    )

                    sub_group_number = str(info[1].text)

                    if sub_group_number != "":
                        sub_group_number = sub_group_number.replace(
                            " ", "").split(":")[1]
                        lesson.SubGroup = sub_group_number

                    day.Lessons.append(lesson)

                if day_number == 6:
                    day_number = 1
                    if week_type == 0:
                        schedule.EvenWeek.append(day)
                        week_type = 1
                    else:
                        schedule.OddWeek.append(day)
                        week_type = 0
                else:
                    day_number += 1


def parse_schedule_page():
    schedules = []
    schedule = dto.ScheduleDTO(dto.EducationalInstitutionDTO("РГЭУ (РИНХ)"))

    response = send_request(RSUE_URL, None)

    faculty_options = BS(response.text, "lxml").find(
        "select", {"id": "type"}).find_all("option")

    for option in faculty_options:
        faculty_id = int(option["value"])

        if faculty_id == 0:
            continue

        faculty = dto.FacultyDTO(faculty_id, str(option.text))

        response = send_request(
            RSUE_URL+"query.php",
            {
                "query": "getKinds",
                "type_id": faculty_id,
            }
        )

        for y in json.loads(response.text):
            year_id = y["kind_id"]
            year_name = y["kind"]

            year = dto.YearDTO(year_id, year_name)

            response = send_request(
                RSUE_URL + "query.php",
                {
                    "query": "getCategories",
                    "type_id": faculty_id,
                    "kind_id": year_id,
                }
            )

            for g in json.loads(response.text):

                group_id = g["category_id"]
                group_name = g["category"]

                response = send_request(
                    RSUE_URL, {"f": faculty_id, "k": year_id, "g": group_id})

                schedule.Faculty = faculty
                schedule.Year = year
                schedule.Group = dto.GroupDTO(group_id, group_name)

                fill_schedule_from_page(response.text, schedule)
                schedules.append(schedule.to_dict())

    return schedules
