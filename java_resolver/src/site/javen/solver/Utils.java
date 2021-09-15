package site.javen.solver;


import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.math.BigInteger;
import java.util.List;

public class Utils {
    static final BigInteger B1024 = BigInteger.valueOf(1024);

    interface IgnoreExceptionRunnable {
        void run() throws Exception;
    }

    public static void measureTime(String tag, IgnoreExceptionRunnable runnable) throws Exception {
        long begin = System.nanoTime();
        try {
            runnable.run();
        } finally {
            log("[%d] %s 耗时:%d ms", System.currentTimeMillis(), tag, (System.nanoTime() - begin) / 1000000);
        }
    }

    public static void measureAvgTime(String tag, IgnoreExceptionRunnable runnable, int times) throws Exception {
        times = Math.max(times, 1);
        long begin = System.nanoTime();
        try {
            for (int i = 0; i < times; i++) {
                measureTime("第" + (i + 1) + "次", runnable);
            }

        } finally {
            log("[%d] %s %d次平均耗时:%d ms", System.currentTimeMillis(), tag, times, (System.nanoTime() - begin) / 1000000 / times);
        }
    }

    /**
     * 是否和 Magic 匹配
     *
     * @param locationId
     * @param magicId
     * @return
     */
    public static boolean isMatchMagic(BigInteger locationId, BigInteger magicId) {
        int cmpValue = locationId.compareTo(magicId);
        if (cmpValue < 0) {//locationId<magic
            if (locationId.add(B1024).compareTo(magicId) == 0) {
                return true;
            }
            return locationId.multiply(B1024).compareTo(magicId) == 0;
        } else if (cmpValue > 0) {
            if (locationId.subtract(B1024).compareTo(magicId) == 0) {
                return true;
            }
            return magicId.compareTo(B1024) < 0 && locationId.mod(B1024).compareTo(magicId) == 0;
        }
        return false;
    }


//    /**
//     * 检查是否全匹配
//     *
//     * @param refJsonFile
//     * @param matchIds
//     * @return
//     * @throws IOException
//     */
//    public static boolean isAllMatch(File refJsonFile, List<String> matchIds) throws IOException {
//        //检查 matchIds
//        byte[] bytes = new FileInputStream(refJsonFile).readAllBytes();
//        JSONArray jsonObject = JSONArray.parseArray(new String(bytes, 0, bytes.length));
//        boolean result = true;
//        for (String matchId : matchIds) {
//            if (!jsonObject.contains(matchId)) {
//                System.out.println(matchId);
//                result = false;
//            }
//        }
//        return result;
//    }

    public static void log(String fmt, Object... args) {
        System.out.println(String.format(fmt, args));
    }
}
