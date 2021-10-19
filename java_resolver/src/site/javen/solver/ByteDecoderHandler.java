package site.javen.solver;

import org.huldra.math.BigInt;



public interface ByteDecoderHandler {
    void onFoundItem(byte[] locationId, BigInt locationValue, BigInt magic);
}
