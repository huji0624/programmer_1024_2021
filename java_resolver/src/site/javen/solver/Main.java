package site.javen.solver;


import java.io.File;
import java.io.RandomAccessFile;
import java.math.BigInteger;
import java.nio.MappedByteBuffer;
import java.nio.channels.FileChannel;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

import static site.javen.solver.Utils.*;

public class Main implements ByteDecoderHandler {
    private static final File DATA_DIR = new File("/Users/coder/go/src/programmer_1024_2021/data");

    public static void main(String[] args) throws Exception {
        log("开始寻宝.....");
        measureAvgTime("计算", () -> {
            new Main().doWork();
        }, 10);
    }

    private final CopyOnWriteArrayList<String> matchMagics = new CopyOnWriteArrayList<>(new ArrayList<>(10000));

    final ExecutorService ioService;

    private Main() {
        ioService = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 2);
    }

    public void doWork() throws Exception {
        List<Future<File>> tasks = new ArrayList<>(100);
        File[] dataFiles = DATA_DIR.listFiles((dir, name) -> name.endsWith(".data"));
        if (dataFiles == null) {
            return;
        }
        for (File dataFile : dataFiles) {
            long length = dataFile.length();
            long begin = 0;
            MappedByteBuffer byteBuffer;
            try (RandomAccessFile raf = new RandomAccessFile(dataFile, "r"); FileChannel inChannel = raf.getChannel()) {
                byteBuffer = inChannel.map(FileChannel.MapMode.READ_ONLY, 0, raf.length());
            }
            while (begin < length) {
                tasks.add(ioService.submit(new DecodeTask(dataFile.getName(), byteBuffer, begin, Constants.PER_DECODE_LENGTH, this), dataFile));
                begin += Constants.PER_DECODE_LENGTH;
            }
        }
        for (Future<File> task : tasks) {
            task.get();
        }
        ioService.shutdownNow();
        log("匹配数:" + matchMagics.size());
    }


    @Override
    public void onFoundItem(byte[] locationId, BigInteger locationValue, BigInteger magic) {
        if (isMatchMagic(locationValue, magic)) {
            matchMagics.add(new String(locationId));
        }
    }


}




