package site.javen.solver;


import org.apache.hc.client5.http.HttpHostConnectException;
import org.apache.hc.client5.http.classic.HttpClient;
import org.apache.hc.client5.http.classic.methods.HttpPost;
import org.apache.hc.client5.http.impl.classic.HttpClients;
import org.apache.hc.core5.http.ClassicHttpResponse;
import org.apache.hc.core5.http.ContentType;
import org.apache.hc.core5.http.io.entity.StringEntity;

import java.io.File;
import java.io.RandomAccessFile;
import java.math.BigInteger;
import java.nio.MappedByteBuffer;
import java.nio.channels.FileChannel;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;

import static site.javen.solver.Utils.*;

public class Main implements ByteDecoderHandler {
    private static final File DATA_DIR = new File("/Users/coder/go/src/programmer_1024_2021/data_generator/data");

    public static void main(String[] args) throws Exception {
        Utils.log("[io threads]:%d   [network threads]:%d", (maxThreadsCount - netThreads), netThreads);
        log("开始寻宝..... Press Enter to Start");
        System.in.read();
        measureAvgTime("计算", () -> {
            new Main().doWork();
        }, 1);
    }

    final AtomicInteger matchCount = new AtomicInteger(0);
    final AtomicInteger submitCount = new AtomicInteger(0);
    final static int maxThreadsCount = Runtime.getRuntime().availableProcessors() * 2;
    final static int netThreads = 4;
    final ExecutorService ioService;

    final ExecutorService netService;

    private Main() {
        ioService = new ThreadPoolExecutor(maxThreadsCount - netThreads, Integer.MAX_VALUE,
                0L, TimeUnit.MILLISECONDS,
                new LinkedBlockingQueue<>());
        netService = new ThreadPoolExecutor(netThreads, Integer.MAX_VALUE,
                0L, TimeUnit.MILLISECONDS,
                new LinkedBlockingQueue<>());
    }

    public void doWork() throws Exception {
        List<Future<File>> tasks = new ArrayList<>(100);
        File[] dataFiles = DATA_DIR.listFiles((dir, name) -> name.endsWith(".data"));
        if (dataFiles == null) {
            return;
        }
        int ioThreads = maxThreadsCount - netThreads;
        int totalFile = dataFiles.length;

        float perFileSplits = totalFile / (float) ioThreads;//每个文件切成多少份
        for (File dataFile : dataFiles) {
            long length = dataFile.length();
            long begin = 0;
            long perSize = Math.round(length * perFileSplits);
            MappedByteBuffer byteBuffer;
            try (RandomAccessFile raf = new RandomAccessFile(dataFile, "r"); FileChannel inChannel = raf.getChannel()) {
                byteBuffer = inChannel.map(FileChannel.MapMode.READ_ONLY, 0, raf.length());
            }
            while (begin < length) {
                tasks.add(ioService.submit(new DecodeTask(dataFile.getName(), byteBuffer, begin, perSize, this), dataFile));
                begin += perSize;
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
        netService.shutdownNow();
        log("匹配数:" + matchCount.get() + " 提交数:" + submitCount.get());
    }


    @Override
    public void onFoundItem(byte[] locationId, BigInteger locationValue, BigInteger magic) {
        if (isMatchMagic(locationValue, magic)) {
            String loc = new String(locationId);
            matchCount.incrementAndGet();
            postResultToServer(loc);
        }
    }


    private final CopyOnWriteArrayList<Future> networkQueue = new CopyOnWriteArrayList<>(new ArrayList<>(10000));


    /**
     * 提交到服务器
     *
     * @param locationId
     */
    private void postResultToServer(String locationId) {
        HttpClient httpClient = HttpClients.createMinimal();
        final String postData = String.format("{\"token\":\"%s\",\"locationid\":\"%s\"}", "test2", locationId);
        HttpPost post = new HttpPost("http://47.104.220.230/h5/");
        post.setEntity(new StringEntity(postData, ContentType.APPLICATION_JSON));
        networkQueue.add(netService.submit(() -> {
            submitCount.incrementAndGet();
            try {
                ClassicHttpResponse response = (ClassicHttpResponse) httpClient.execute(post);
                if (response.getCode() == 200) {
                    Utils.log("%s 提交成功   %s", locationId, new String(response.getEntity().getContent().readAllBytes()));
                } else {
                    Utils.log("%s 提交失败  %s", locationId, response.getCode());
                }
            } catch (HttpHostConnectException e) {
                Utils.log("%s 提交失败 网络原因", locationId);
            } catch (Exception e) {
                Utils.log("%s 提交失败 %s", locationId, e.getMessage());
            }

        }, locationId));
    }
}




