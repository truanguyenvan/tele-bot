

class Emoji:
    error = '❌'
    info = 'ℹ️'
    success = '✅'
    warning = '⚠️'


def generate_header(text):
    return '''\t**{}**\n\n'''.format(text).replace('-', " ")