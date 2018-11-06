from aiocqhttp import CQHttp

bot = CQHttp(api_root='http://127.0.0.1:5700/',
             access_token='',
             secret='')


@bot.on_message()
async def handle_msg(context):
    await bot.send(context, '你好呀，下面一条是你刚刚发的：')
    return {'reply': context['message']}


@bot.on_notice('group_increase')
async def handle_group_increase(context):
    await bot.send(context, message='欢迎新人～', auto_escape=True)


@bot.on_request('group', 'friend')
async def handle_request(context):
    return {'approve': True}


bot.run(host='127.0.0.1', port=8080)