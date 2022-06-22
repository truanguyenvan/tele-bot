from telebot.handler.msg import send_msg, remove_msg
from telebot.handler.utils import generate_header, Emoji
import time

def get_emoij_by_text(text):
    if text == 'error':
        return Emoji.error
    elif text == 'warning':
        return Emoji.warning
    elif text == 'success':
        return Emoji.success
    else:
        return Emoji.info 
    
def send_msg_to_chat(token, chatId, header, text, emoji):
    for i in range(5):
        data = send_msg(token=token, chat_id=chatId, header=generate_header(header), text=text, emoji=get_emoij_by_text(emoji))
        if data['ok'] == True:
            return data['result']['message_id']
        time.sleep(2)
    return None

def remove_msg_in_chat(token, chatId, msgId):
    for i in range(5):
        data = remove_msg(token=token, chat_id=chatId, msg_id=msgId)
        if data['ok'] == True:
            return True
        time.sleep(2)
    return False