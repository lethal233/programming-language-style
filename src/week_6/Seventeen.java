import java.io.File;
import java.io.IOException;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.*;

public class Seventeen {
    public static void main(String[] args) throws IOException {
        new WordFrequencyController(args[0]).run();
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter a class to inspect:");
        String name = sc.nextLine();
        printClassInfo(name);
    }
    private static void printClassInfo(String className) {
        Class clazz = null;
        try {
            clazz = Class.forName(className);
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        if (clazz == null) {
            return;
        }
        System.out.println("Getting information about " + className);

        System.out.println("Fields:");
        for (Field field : clazz.getDeclaredFields()) {
            System.out.println(" - " + field.getType().getName() + " " + field.getName());
        }

        System.out.println("Methods:");
        for (Method method : clazz.getMethods()) {
            System.out.println(" - " + method.getName());
        }

        System.out.println("Superclasses:");
        for (Class<?> superclass = clazz.getSuperclass(); superclass != null; superclass = superclass.getSuperclass()) {
            System.out.println(" - " + superclass.getName());
        }

        System.out.println("Implemented Interfaces:");
        for (Class<?> iface : clazz.getInterfaces()) {
            System.out.println(" - " + iface.getName());
        }

    }
}
abstract class TFExercise {
    public String getInfo() {
        return this.getClass().getName();
    }
}

class WordFrequencyController extends TFExercise {
    private DataStorageManager storageManager;
    private StopWordManager stopWordManager;
    private WordFrequencyManager wordFreqManager;

    public WordFrequencyController(String pathToFile) throws IOException {
        this.storageManager = new DataStorageManager(pathToFile);
        this.stopWordManager = new StopWordManager();
        this.wordFreqManager = new WordFrequencyManager();
    }

    public void run() {
        try {
            Method getWords = DataStorageManager.class.getMethod("getWords");
            Method isStopWords = StopWordManager.class.getMethod("isStopWord", String.class);
            Method incrementCount = WordFrequencyManager.class.getMethod("incrementCount", String.class);
            List<String> res = (List<String>) getWords.invoke(this.storageManager);
            for (String word :  res) {
                if (!(boolean) isStopWords.invoke(this.stopWordManager, word)) {
                    incrementCount.invoke(this.wordFreqManager, word);
                }
            }
            Method sorted = WordFrequencyManager.class.getMethod("sorted");
            List<WordFrequencyPair> pairs = (List<WordFrequencyPair>) sorted.invoke(this.wordFreqManager);
            int numWordsPrinted = 0;
            for (WordFrequencyPair pair : pairs) {
                System.out.println(pair.getWord() + " - " + pair.getFrequency());

                numWordsPrinted++;
                if (numWordsPrinted >= 25) {
                    break;
                }
            }
        } catch (NoSuchMethodException | IllegalAccessException | InvocationTargetException e) {
            e.printStackTrace();
        }
    }
}
/** Models the contents of the file. */
class DataStorageManager extends TFExercise {
    private List<String> words;

    public DataStorageManager(String pathToFile) throws IOException {
        this.words = new ArrayList<String>();

        Scanner f = new Scanner(new File(pathToFile), "UTF-8");
        try {
            f.useDelimiter("[\\W_]+");
            while (f.hasNext()) {
                this.words.add(f.next().toLowerCase());
            }
        } finally {
            f.close();
        }
    }

    public List<String> getWords() {
        return this.words;
    }

    public String getInfo() {
        return super.getInfo() + ": My major data structure is a " + this.words.getClass().getName();
    }
}
class StopWordManager extends TFExercise {
    private Set<String> stopWords;

    public StopWordManager() throws IOException {
        this.stopWords = new HashSet<String>();

        Scanner f = new Scanner(new File("../stop_words.txt"), "UTF-8");
        try {
            f.useDelimiter(",");
            while (f.hasNext()) {
                this.stopWords.add(f.next());
            }
        } finally {
            f.close();
        }

        // Add single-letter words
        for (char c = 'a'; c <= 'z'; c++) {
            this.stopWords.add("" + c);
        }
    }

    public boolean isStopWord(String word) {
        return this.stopWords.contains(word);
    }

    public String getInfo() {
        return super.getInfo() + ": My major data structure is a " + this.stopWords.getClass().getName();
    }
}

/** Keeps the word frequency data. */
class WordFrequencyManager extends TFExercise {
    private Map<String, MutableInteger> wordFreqs;

    public WordFrequencyManager() {
        this.wordFreqs = new HashMap<String, MutableInteger>();
    }

    public void incrementCount(String word) {
        MutableInteger count = this.wordFreqs.get(word);
        if (count == null) {
            this.wordFreqs.put(word, new MutableInteger(1));
        } else {
            count.setValue(count.getValue() + 1);
        }
    }

    public List<WordFrequencyPair> sorted() {
        List<WordFrequencyPair> pairs = new ArrayList<WordFrequencyPair>();
        for (Map.Entry<String, MutableInteger> entry : wordFreqs.entrySet()) {
            pairs.add(new WordFrequencyPair(entry.getKey(), entry.getValue().getValue()));
        }
        Collections.sort(pairs);
        Collections.reverse(pairs);
        return pairs;
    }

    public String getInfo() {
        return super.getInfo() + ": My major data structure is a " + this.wordFreqs.getClass().getName();
    }
}

class MutableInteger {
    private int value;

    public MutableInteger(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }

    public void setValue(int value) {
        this.value = value;
    }
}

class WordFrequencyPair implements Comparable<WordFrequencyPair> {
    private String word;
    private int frequency;

    public WordFrequencyPair(String word, int frequency) {
        this.word = word;
        this.frequency = frequency;
    }

    public String getWord() {
        return word;
    }

    public int getFrequency() {
        return frequency;
    }

    public int compareTo(WordFrequencyPair other) {
        return this.frequency - other.frequency;
    }
}

