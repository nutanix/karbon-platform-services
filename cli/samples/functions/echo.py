def main(ctx, msg):
    cfg = ctx.get_config()
    # read the given param
    param1 = cfg.get('param1', '')
    print('recieved a param1: ', param1)
    ctx.send(msg)
