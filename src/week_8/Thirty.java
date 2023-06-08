import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.concurrent.*;

public class Thirty {

    static BlockingQueue<String> wordSpace = new LinkedBlockingQueue<>();
    static BlockingQueue<Map<String, Integer>> freqSpace = new LinkedBlockingQueue<>();
    static Set<String> stopwords = new HashSet<>();

    static class Worker implements Runnable {
        @Override
        public void run() {
            Map<String, Integer> wordFreqs = new HashMap<>();
            while (true) {
                try {
                    String word = wordSpace.poll(1, TimeUnit.SECONDS);
                    if (word == null) break;
                    if (!stopwords.contains(word)) {
                        wordFreqs.put(word, wordFreqs.getOrDefault(word, 0) + 1);
                    }
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
            try {
                freqSpace.put(wordFreqs);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }

    public static void main(String[] args) {
        try {
            List<String> lines = Files.readAllLines(Paths.get("../stop_words.txt"));
            lines.forEach(l->stopwords.addAll(List.of(l.split(","))));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        try {
            Files.lines(Paths.get(args[0]))
                    .map(line -> line.toLowerCase().split("[\\W_]+"))
                    .flatMap(Arrays::stream)
                    .filter(word -> word.length() > 1)
                    .forEach(wordSpace::add);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        Thread[] workers = new Thread[5];
        for (int i = 0; i < workers.length; i++) {
            workers[i] = new Thread(new Worker());
            workers[i].start();
        }

        for (Thread worker : workers) {
            try {
                worker.join();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }

        Map<String, Integer> wordFreqs = new HashMap<>();
        while (!freqSpace.isEmpty()) {
            Map<String, Integer> freqs = freqSpace.poll();
            if (freqs == null) continue;
            for (String word : freqs.keySet()) {
                wordFreqs.put(word, wordFreqs.getOrDefault(word, 0) + freqs.get(word));
            }
        }

        wordFreqs.entrySet().stream()
                .sorted(Collections.reverseOrder(Map.Entry.comparingByValue()))
                .limit(25)
                .forEach(entry -> System.out.println(entry.getKey() + " - " + entry.getValue()));
    }
}
