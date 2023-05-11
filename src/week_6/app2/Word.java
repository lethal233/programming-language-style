import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Word implements IWord {
    public  List<String> extractWords(String pathToFile) {
        String data = readFromFile(pathToFile);
        Set<String> stopWords = getStopWords();
        return Stream.of(data.split("[\\W_]+"))
                .map(String::toLowerCase)
                .filter(word -> !stopWords.contains(word))
                .collect(Collectors.toList());
    }

    private  Set<String> getStopWords() {
        String stopWordsFile = "../stop_words.txt";
        String[] stopWordsArray = readFromFile(stopWordsFile).split(",");
        Set<String> stopWords = new HashSet<>(Arrays.asList(stopWordsArray));
        for (char c = 'a'; c <= 'z'; c++) {
            stopWords.add(Character.toString(c));
        }
        return stopWords;
    }

    private  String readFromFile(String filePath) {
        try {
            return Files.lines(Paths.get(filePath))
                    .collect(Collectors.joining(" "));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
