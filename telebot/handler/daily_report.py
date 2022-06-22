from telebot.handler.msg import sendNotification
from telebot.handler.utils import Emoji, generate_header
def daily_report(token, chatId):
    header = "DAILY REPORT REMINDER"
    text = "Anh/Chị nhớ điền daily report và logwork trước khi về nhé!\nHave a good evening."
    data = sendNotification(token=token, chat_id=chatId, text=text, header=generate_header(header), emoji=Emoji.info)
    return data['ok']