package site.javen.solver;


import org.apache.hc.client5.http.async.methods.SimpleHttpRequest;
import org.apache.hc.client5.http.async.methods.SimpleHttpResponse;
import org.apache.hc.client5.http.impl.async.CloseableHttpAsyncClient;
import org.apache.hc.client5.http.impl.async.HttpAsyncClients;
import org.apache.hc.core5.concurrent.FutureCallback;
import org.apache.hc.core5.http.ContentType;

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
import java.util.concurrent.atomic.AtomicInteger;

import static site.javen.solver.Utils.*;

public class Main implements ByteDecoderHandler {
    private static final File DATA_DIR = new File("/Users/coder/go/src/programmer_1024_2021/data_generator/data");

    public static void main(String[] args) throws Exception {
        log("开始寻宝.....");
        measureAvgTime("计算", () -> {
            new Main().doWork();
        }, 10);
    }

    final AtomicInteger matchCount = new AtomicInteger(0);

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
        log("匹配数:" + matchCount.get());
        for (Future future : networkQueue) {
            future.get();
        }
        ioService.shutdownNow();
        http2Default.initiateShutdown();
    }


    @Override
    public void onFoundItem(byte[] locationId, BigInteger locationValue, BigInteger magic) {
        if (isMatchMagic(locationValue, magic)) {
            String loc = new String(locationId);
            matchCount.incrementAndGet();
            postResultToServer(loc);
        }
    }


    private CopyOnWriteArrayList<Future> networkQueue = new CopyOnWriteArrayList<>(new ArrayList<>(10000));

    CloseableHttpAsyncClient http2Default = HttpAsyncClients.createDefault();

    {
        http2Default.start();
    }

    /**
     * 提交到服务器
     *
     * @param locationId
     */
    private void postResultToServer(String locationId) {
        SimpleHttpRequest post = SimpleHttpRequest.create("POST", "http://localhost/dig");
        post.setBody(String.format("{\n" +
                "                \"token\":\"%s\",\n" +
                "                \"locationid\":\"%s\"\n" +
                "            }", "test2", locationId), ContentType.APPLICATION_JSON);
        Future<SimpleHttpResponse> execute = http2Default.execute(post, new FutureCallback<>() {
            @Override
            public void completed(SimpleHttpResponse simpleHttpResponse) {
            }

            @Override
            public void failed(Exception e) {
                e.printStackTrace();
            }


            @Override
            public void cancelled() {

            }
        });
        networkQueue.add(execute);
    }
}




