package site.javen.solver;

import org.huldra.math.BigInt;

import java.util.Arrays;

public class ByteDecoder {
    static byte DOCUMENT_BEGIN = '{';
    static byte TEXT_BEGIN = '"';
    static byte TOKEN_BEGIN = 0;
    static byte TOKEN_DOCUMENT = 1;
    static byte TOKEN_LOC_START = 2;
    static byte TOKEN_MAGIC_START = 3;

    private byte mToken = TOKEN_BEGIN;

    byte[] locationStrBuffer = new byte[200];
    char[] locationValueBuffer = new char[200];
    char[] magicValueBuffer = new char[200];
    int locationStrBufferLen = 0;
    int locationValueLen = 0;
    int magicValueBufferLen = -1;

    private ByteDecoderHandler handler;

    public ByteDecoder(ByteDecoderHandler handler) {
        this.handler = handler;
    }

    public int getToken() {
        return mToken;
    }

    /**
     * Decode return Offset;
     *
     * @param b
     * @return
     */
    public int decode(byte b) {

        if (mToken == TOKEN_LOC_START) {
            if (b == TEXT_BEGIN) {
                if (locationStrBufferLen == -1) {
                    locationStrBufferLen = 0;
                    locationValueLen = 0;
                    return 1;
                } else {
                    mToken = TOKEN_MAGIC_START;
                    return 10;
                }
            }
            if (b >= '0' && b <= '9') {
                locationValueBuffer[locationValueLen++] = (char) b;
            }
            locationStrBuffer[locationStrBufferLen++] = b;
            return 1;
        }
        if (mToken == TOKEN_MAGIC_START) {
            if (b == TEXT_BEGIN) {
                if (magicValueBufferLen == -1) {
                    magicValueBufferLen = 0;
                    return 1;
                } else {
                    mToken = TOKEN_BEGIN;
                    if (handler != null) {
                        handler.onFoundItem(Arrays.copyOf(locationStrBuffer, locationStrBufferLen), new BigInt(locationValueBuffer, locationValueLen), new BigInt(magicValueBuffer, magicValueBufferLen));
                    }
                    magicValueBufferLen = -1;
                    locationStrBufferLen = -1;
                    locationValueLen = -1;
                    return 3;
                }
            }
            magicValueBuffer[magicValueBufferLen++] = (char) b;
            return 1;
        }
        if (mToken == TOKEN_BEGIN) {
            if (b == DOCUMENT_BEGIN) {
                mToken = TOKEN_DOCUMENT;
                return 14;
            }
            return 1;
        }
        if (mToken == TOKEN_DOCUMENT) {
            if (b == TEXT_BEGIN) {
                mToken = TOKEN_LOC_START;
                locationStrBufferLen = -1;
                return decode(b);
            } else {
                throw new RuntimeException("error state");
            }
        }
        return -1;
    }

}
