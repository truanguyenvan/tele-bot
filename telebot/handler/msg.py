import requests


def send_msg(token, chat_id, text, header, emoji):
    msg = str(emoji) + header + text.replace ('NT', '\n').replace ('TT', '\t')
    send_text = 'https://api.telegram.org/bot' + token + '/sendMessage?chat_id=' + chat_id + '&parse_mode=Markdown&text=' + msg
    response = requests.get(send_text)
    return response.json()

def remove_msg (token, chat_id, msg_id):
    send_text = 'https://api.telegram.org/bot' + token + '/deleteMessage?chat_id=' + chat_id + '&message_id=' + msg_id
    response = requests.get(send_text)
    return response.json()