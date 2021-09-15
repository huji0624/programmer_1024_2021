package site.javen.solver;

import java.nio.MappedByteBuffer;

/**
 * 解析任务
 */
public class DecodeTask implements Runnable {

    final ByteDecoder mDecoder;

    final MappedByteBuffer mByteBuffer;

    final String mName;

    final long begin;

    final long end;

    public DecodeTask(String name, MappedByteBuffer byteBuffer, long begin, long length, ByteDecoderHandler handler) throws Exception {
        this.mDecoder = new ByteDecoder(handler);
        this.mName = name;
        this.mByteBuffer = byteBuffer;
        this.begin = begin;
        this.end = begin + length;
    }


    @Override
    public void run() {
        try {
            doWork();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * 具体解析
     */
    private void doWork() throws Exception {
//        Utils.measureTime(mName + ":" + begin, () -> {
            long endValue = Math.min(mByteBuffer.limit(), end);
            for (long i = begin; ; ) {
                final int offset = mDecoder.decode(mByteBuffer.get((int) i));
                if (offset == -1) {
                    throw new RuntimeException("offset <0");
                }
                i += offset;
                if (mDecoder.getToken() == ByteDecoder.TOKEN_BEGIN && i >= endValue) {
                    break;
                }
            }
//        });
    }

}