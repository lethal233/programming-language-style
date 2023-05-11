import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Pattern;

public class Word implements IWord {
    public List<String> extractWords(String pathToFile) {
        String strData = null;
        try {
            strData = new String(Files.readAllBytes(Paths.get(pathToFile)));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        Pattern pattern = Pattern.compile("[\\W_]+");
        String[] wordArray = pattern.matcher(strData).replaceAll(" ").toLowerCase().split(" ");

        List<String> stopWords = getStopWords();
        List<String> wordList = new ArrayList<>(Arrays.asList(wordArray));

        wordList.removeIf(stopWords::contains);
        return wordList;
    }

    private List<String> getStopWords() {
        String stopWordsFile = "../stop_words.txt";
        String[] stopWordsArray = new String[0];
        try {
            stopWordsArray = new String(Files.readAllBytes(Paths.get(stopWordsFile))).split(",");
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        List<String> stopWords = new ArrayList<>(Arrays.asList(stopWordsArray));

        for (char c = 'a'; c <= 'z'; c++) {
            stopWords.add(Character.toString(c));
        }

        return stopWords;
    }
}
