* read
    * final ChannelPipeline pipeline = pipeline();
    * pipeline.fireChannelRead(readBuf.get(i));
    * pipeline.fireChannelReadComplete();
    * pipeline.fireExceptionCaught(exception);

---
