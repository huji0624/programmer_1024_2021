package site.javen.solver;

import java.math.BigInteger;

public interface ByteDecoderHandler {
    void onFoundItem(byte[] locationId, BigInteger locationValue, BigInteger magic);
}