from telebot.services.msg import send_msg_to_chat
from telebot.handler.utils import Emoji
# from telebot.handler.daily_report import daily_report
# import schedule
# import time
import tzlocal
from apscheduler.schedulers.blocking import BlockingScheduler
sched = BlockingScheduler(timezone=str(tzlocal.get_localzone()))

class Reminder():
  def __init__(self, chatId, token):
    print('Set reminder sucessfully')
    self.chatId = chatId
    self.token = token
    self.reminder()
  def noti_daily_report(self):
    print('Noti daily report')
    header = "DAILY REPORT REMINDER"
    text = "Anh/Chị nhớ điền daily report và logwork trước khi về nhé!\nHave a good evening."
    msgId = send_msg_to_chat(token=self.token, chatId=self.chatId, text=text, header=header, emoji=Emoji.info)
    print(msgId)
  def reminder(self):
    # schedule.every().day.at("16:50").do(self.noti_daily_report)
    # while True:
    #   schedule.run_pending()
    #   time.sleep(1)
    sched.add_job(self.noti_daily_report, 'cron', day_of_week='mon-fri', hour=16, minute=50)
    sched.start()



      