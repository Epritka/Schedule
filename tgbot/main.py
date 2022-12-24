from aiogram.dispatcher.filters.state import State, StatesGroup
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram import Bot, Dispatcher, executor, types
from aiogram.dispatcher import FSMContext
from dotenv import load_dotenv
from datetime import datetime, timedelta
import logging
import asyncio
import requests
import pytz
import json
import os

import keyboard
import inline_keyboard

DAYS = {
    1: "Понедельник",
    2: "Вторник",
    3: "Среда",
    4: "Четверг",
    5: "Пятница",
    6: "Суббота",
    7: "Воскресенье"
}

load_dotenv()

ENGINE_URL = str(os.getenv("ENGINE_HOST"))
USER_MANAGER_URL = str(os.getenv("USER_MANAGER_HOST"))
TG_TOKEN = str(os.getenv("TG_TOKEN"))

ENGINE_URL = f"http://{ENGINE_URL}/api/v1/"
USER_MANAGER_URL = f"http://{USER_MANAGER_URL}/api/v1/"

USERS = {}
USERS_GROUP = {}

logging.basicConfig(level=logging.INFO)

bot = Bot(token=TG_TOKEN)

storage = MemoryStorage()

dp = Dispatcher(bot, storage=storage)


class Form(StatesGroup):
    name = State()


def add_group(tgId, groupId):
    global USERS_GROUP
    USERS_GROUP[tgId] = groupId


# def addStudent(tgId=None, userId=None, groupId=None, studentId=None):
#     global USERS
#     if tgId != None:
#         if userId != None:
#             USERS[tgId] = {
#                 "userId": userId,
#             }
#         else:
#             USERS[tgId] = {
#                 "userId": 0,
#             }

#         if groupId != None:
#             USERS[tgId] = {
#                 "groupId": groupId,
#             }
#         else:
#             USERS[tgId] = {
#                 "groupId": 0,
#             }

#         if studentId != None:
#             USERS[tgId] = {
#                 "studentId": studentId,
#             }
#         else:
#             USERS[tgId] = {
#                 "studentId": 0,
#             }


@dp.message_handler(commands="help")
async def show_help_message(message: types.Message):
    await message.answer(text=f"Этот бот позволяет смотреть расписание вашего любимого университета ♥")


@dp.message_handler(commands=["start"])
async def show_start_message(message: types.Message):
    await message.answer(f"Привет, {message.from_user.full_name}!",
                         reply_markup=inline_keyboard.DEFAULT)


@dp.callback_query_handler(text="login")
async def process_callback_login(callback_query: types.CallbackQuery):
    await bot.answer_callback_query(callback_query.id)
    await Form.name.set()
    await bot.send_message(
        callback_query.from_user.id,
        text="Введите группу"
    )


@dp.message_handler(state=Form.name)
async def process_text_group_name(message: types.Message, state: FSMContext):
    await state.finish()

    response = requests.get(f"{ENGINE_URL}group/{message.text}")

    if response.status_code == 200:
        data = response.json()["data"]

        tgId = message.from_user.id
        add_group(tgId=tgId, groupId=data["id"])
        global USERS
        # print(USERS[message.from_user.id])
        # response = requests.post(f"{ENGINE_URL}student/", json={
        #     "id": USERS[message.from_user.id]["studentId"],
        #     "userId": USERS[message.from_user.id]["userId"],
        #     "groupId": data["id"],
        # })

        await bot.send_message(
            message.from_user.id,
            text="Принято. Можете узнать расписание.",
            reply_markup=keyboard.SCHEDULE,
        )
    elif response.status_code == 404:
        await message.reply("Проверьте данные и попробуйте еще раз.")
    else:
        await message.reply("Извините...Что-то пошло не так.")


async def get_day(moscow_time, tgId):
    global USERS_GROUP
    json_data = {
        "groupId": USERS_GROUP[tgId],
        "date": str(int(moscow_time.timestamp())),
    }

    response = requests.get(f"{ENGINE_URL}day/", params=json_data)

    if response.status_code == 200:
        data = response.json()["data"]

        text = f"{'Нечётная' if data['weekType'] == 'odd' else 'Чётная'} неделя\n{DAYS[data['number']]}\n\n"
        text += f""

        for l in data["lessons"]:
            text += f"{l['startTime']}-{l['endTime']}\n"

            if "subGroup" in l:
                text += f"{l['subGroup']}\n"
            text += f"{l['name']}\n"
            text += f"{l['teacher']}\n"
            text += f"{l['auditorium']}\n"
            text += f"{l['type']}\n\n"

        return text
    elif response.status_code == 404:
        return "Пар нет, можно отдохнуть."
    else:
        return ""


@dp.callback_query_handler(text="today")
async def process_callback_today(callback_query: types.CallbackQuery):
    await bot.answer_callback_query(callback_query.id)

    tgId = callback_query.from_user.id
    moscow_time = datetime.now(pytz.timezone('Europe/Moscow'))

    text = await get_day(moscow_time=moscow_time, tgId=tgId)
    if text != "":
        await bot.send_message(
            callback_query.from_user.id,
            text=text,
            reply_markup=keyboard.SCHEDULE,
        )
    else:
        await callback_query.message.reply("Извините...Что-то пошло не так.")

@dp.callback_query_handler(text="tomorrow")
async def process_callback_today(callback_query: types.CallbackQuery):
    await bot.answer_callback_query(callback_query.id)

    tgId = callback_query.from_user.id

    moscow_time = datetime.now(pytz.timezone('Europe/Moscow'))

    text = await get_day(moscow_time=moscow_time + timedelta(1), tgId=tgId)
    if text != "":
        await bot.send_message(
            callback_query.from_user.id,
            text=text,
            reply_markup=keyboard.SCHEDULE,
        )
    else:
        await callback_query.message.reply("Извините...Что-то пошло не так.")

@dp.message_handler(content_types="text")
async def process_text_show_schedule(message: types.Message):
    if message.text == "Посмотреть расписание":
        await bot.send_message(
            message.from_user.id,
            text="Выберите на какой день хотите посмотреть расписание",
            reply_markup=inline_keyboard.SELECT_DAY,
        )


@dp.message_handler()
async def echo_message(msg: types.Message):
    await bot.send_message(msg.from_user.id, "неизвестная команда.")

if __name__ == "__main__":
    add_group(tgId=403918258, groupId=1)
    # response = requests.get(f"{USER_MANAGER_URL}user/")

    # users = response.json()["data"]

    # for u in users:
    #     addStudent(tgId=u["telegramUserId"], userId=u["id"])

    # response = requests.get(f"{ENGINE_URL}student/")

    # students = response.json()["data"]

    # for s in students:
    #     print
    #     for k, v in USERS:
    #         if v["id"] == s["userId"]:
    #             USERS[k]["groupId"] = s["groupId"]
    #             USERS[k]["studentId"] = s["id"]

    # print(USERS)
    # print(USERS_GROUP)
    executor.start_polling(dp, skip_updates=True)
