* read `AbstractNioByteChannel.NioByteUnsafe`
    * final ChannelPipeline pipeline = pipeline();
    * pipeline.fireChannelRead(byteBuf);
    * pipeline.fireChannelReadComplete();
    * closeOnRead(pipeline); (if close)

---
