import java.io.*;
import java.util.*;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.atomic.AtomicBoolean;

abstract class ActiveWFObject extends Thread {
    protected AtomicBoolean stopMe = new AtomicBoolean(false);
    protected BlockingQueue<Object[]> queue = new LinkedBlockingQueue<>();

    abstract void dispatch(Object[] message);

    @Override
    public void run() {
        while (!stopMe.get()) {
            try {
                Object[] message = queue.take();
                dispatch(message);
                if (message[0].equals("die")) {
                    stopMe.set(true);
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

class DataStorageManager extends ActiveWFObject {
    private String data = "";
    private StopWordManager stopWordManager;

    @Override
    void dispatch(Object[] message) {
        if (message[0].equals("init")) {
            init(Arrays.copyOfRange(message, 1, message.length));
        } else if (message[0].equals("send_word_freqs")) {
            processWords(Arrays.copyOfRange(message, 1, message.length));
        } else {
            stopWordManager.queue.add(message);
        }
    }

    private void init(Object[] message) {
        String pathToFile = (String) message[0];
        stopWordManager = (StopWordManager) message[1];
        try (BufferedReader reader = new BufferedReader(new FileReader(pathToFile))) {
            String line;
            StringBuilder sb = new StringBuilder();
            while ((line = reader.readLine()) != null) {
                sb.append(line).append("\n");
            }
            data = sb.toString().replaceAll("[\\W_]+", " ").toLowerCase();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void processWords(Object[] message) {
        WordFrequencyController recipient = (WordFrequencyController) message[0];
        String[] words = data.split(" ");
        for (String word : words) {
            stopWordManager.queue.add(new Object[]{"filter", word});
        }
        stopWordManager.queue.add(new Object[]{"top25", recipient});
    }
}

public class TwentyNine {
    public static void main(String[] args) {
        WordFrequencyManager wordFreqManager = new WordFrequencyManager();
        wordFreqManager.start();

        StopWordManager stopWordManager = new StopWordManager();
        stopWordManager.queue.add(new Object[]{"init", wordFreqManager});
        stopWordManager.start();

        DataStorageManager storageManager = new DataStorageManager();
        storageManager.queue.add(new Object[]{"init", args[0], stopWordManager});
        storageManager.start();

        WordFrequencyController wfController = new WordFrequencyController();
        wfController.queue.add(new Object[]{"run", storageManager});
        wfController.start();

        try {
            wordFreqManager.join();
            stopWordManager.join();
            storageManager.join();
            wfController.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}


class StopWordManager extends ActiveWFObject {
    private List<String> stopWords = new ArrayList<>();
    private WordFrequencyManager wordFreqsManager;

    @Override
    void dispatch(Object[] message) {
        if (message[0].equals("init")) {
            init((Object[]) Arrays.copyOfRange(message, 1, message.length));
        } else if (message[0].equals("filter")) {
            filter((Object[]) Arrays.copyOfRange(message, 1, message.length));
        } else {
            wordFreqsManager.queue.add(message);
        }
    }

    private void init(Object[] message) {
        wordFreqsManager = (WordFrequencyManager) message[0];
        try (BufferedReader reader = new BufferedReader(new FileReader("../stop_words.txt"))) {
            stopWords.addAll(Arrays.asList(reader.readLine().split(",")));
            stopWords.addAll(Arrays.asList("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z".split(",")));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void filter(Object[] message) {
        String word = (String) message[0];
        if (!stopWords.contains(word)) {
            wordFreqsManager.queue.add(new Object[]{"word", word});
        }
    }
}

class WordFrequencyManager extends ActiveWFObject {
    private Map<String, Integer> wordFreqs = new HashMap<>();

    @Override
    void dispatch(Object[] message) {
        if (message[0].equals("word")) {
            incrementCount((String) message[1]);
        } else if (message[0].equals("top25")) {
            top25((WordFrequencyController) message[1]);
        }
    }

    private void incrementCount(String word) {
        wordFreqs.put(word, wordFreqs.getOrDefault(word, 0) + 1);
    }

    private void top25(WordFrequencyController recipient) {
        List<Map.Entry<String, Integer>> sortedFreqs = new ArrayList<>(wordFreqs.entrySet());
        sortedFreqs.sort((a, b) -> b.getValue().compareTo(a.getValue()));
        recipient.queue.add(new Object[]{"top25", sortedFreqs.subList(0, 25)});
    }
}

class WordFrequencyController extends ActiveWFObject {
    private DataStorageManager storageManager;

    @Override
    void dispatch(Object[] message) {
        if (message[0].equals("run")) {
            run((DataStorageManager) message[1]);
        } else if (message[0].equals("top25")) {
            display((List<Map.Entry<String, Integer>>) message[1]);
        } else {
            throw new IllegalArgumentException("Message not understood " + message[0]);
        }
    }

    private void run(DataStorageManager storageManager) {
        this.storageManager = storageManager;
        storageManager.queue.add(new Object[]{"send_word_freqs", this});
    }

    private void display(List<Map.Entry<String, Integer>> wordFreqs) {
        for (Map.Entry<String, Integer> entry : wordFreqs) {
            System.out.println(entry.getKey() + " - " + entry.getValue());
        }
        storageManager.queue.add(new Object[]{"die"});
        stopMe.set(true);
    }
}
