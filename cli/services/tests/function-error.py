def main(ctx, msg):
    msg = msg + "this line has error"
    ctx.send(msg)