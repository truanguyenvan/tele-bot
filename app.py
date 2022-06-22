from requests import head
from telebot.credentials import bot_token, bot_chat_id
from telebot.services.reminder import Reminder
from telebot.services.msg import send_msg_to_chat, remove_msg_in_chat
from flask import Flask, request
from waitress import serve
import multiprocessing
import time
app = Flask(__name__)

@app.route('/telegram/noti', methods=['GET','POST'])
def push_noti():
    chatId = request.args.get('chat_id')
    if chatId == None:
        return {'code': 'failed', 'msg': 'Query param invalid'}
    emoji = request.args.get('emoji')
    text = request.args.get('text')
    header = request.args.get('header')

    msgId = send_msg_to_chat(token=bot_token, chatId=chatId, header=header, text=text, emoji=emoji)
    if msgId:
        t = Interval(1.0, lambda: remove_msg_in_chat(token=bot_token, chatId=chatId, msgId=str(msgId)))
        t.start()
        time.sleep(2)
        t.cancel()
        return {'code': 'ok', 'msg': 'send msg oke rồi nhé!'}

    return {'code': 'failed', 'msg': 'send msg tạch rồi nhé!'}

@app.route('/telegram/remove', methods=['GET','POST'])
def remove_msg():
    chatId = request.args.get('chat_id')
    msgId = request.args.get('message_id')
    if chatId == None or msgId == None:
        return {'code': 'failed', 'msg': 'Query param invalid'}

    ok = remove_msg_in_chat(token=bot_token, chatId=chatId, msgId=msgId)
    if ok:
        return {'code': 'ok', 'msg': 'remove msg oke rồi nhé!'}

    return {'code': 'failed', 'msg': 'remove msg tạch rồi nhé!'}
    

def run_web_server_service():
    print('Run app service')
    serve(app, host="0.0.0.0", port=5000)
    # app.run(host='0.0.0.0', port=5000, debug=False)

def run_reminder_service():
    print('Run remider service')
    Reminder(chatId=bot_chat_id, token=bot_token)

if __name__ == '__main__':
    # Tạo hai tiến trình process
    p1 = multiprocessing.Process(target=run_web_server_service)
    p2 = multiprocessing.Process(target=run_reminder_service )
 
    # Bắt đầu process 1
    p1.start()
    # Bắt đầu process 2
    p2.start()
 
    # Chờ tới khi process 1 hoàn thành
    p1.join()
    # Chờ tới khi process 2 hoàn thành
    p2.join()
 
    # Cả hai processes hoàn thành
    print("Done!") 
