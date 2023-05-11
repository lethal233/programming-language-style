import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Frequency implements IFrequency {
    public List<Map.Entry<String, Long>> top25(List<String> wordList) {
        Map<String, Long> counts = wordList.stream()
                .collect(Collectors.groupingBy(e -> e, Collectors.counting()));

        return counts.entrySet().stream()
                .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
                .limit(25)
                .collect(Collectors.toList());
    }
}
