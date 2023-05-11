import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Frequency implements IFrequency {
    public List<Map.Entry<String, Long>> top25(List<String> wordList) {
        Map<String, Long> wordFreqs = new HashMap<>();

        for (String word : wordList) {
            wordFreqs.put(word, wordFreqs.getOrDefault(word, 0L) + 1);
        }

        List<Map.Entry<String, Long>> sortedWordFreqs = new ArrayList<>(wordFreqs.entrySet());
        sortedWordFreqs.sort(Map.Entry.<String, Long>comparingByValue().reversed());

        return sortedWordFreqs.subList(0, Math.min(sortedWordFreqs.size(), 25));
    }
}
