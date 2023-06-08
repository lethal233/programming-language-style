import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

public class ThirtyTwo {

    private static List<String> stopwords = new ArrayList<>();

    private static List<String> partition(String data, int nLines) {
        List<String> lines = Arrays.asList(data.split("\n"));
        List<String> chunks = new ArrayList<>();
        for (int i = 0; i < lines.size(); i += nLines) {
            chunks.add(String.join("\n", lines.subList(i, Math.min(i + nLines, lines.size()))));
        }
        return chunks;
    }

    private static List<Map.Entry<String, Integer>> splitWords(String data) {
        String[] words = data.replaceAll("[\\W_]+", " ").toLowerCase().split("\\s+");
        List<Map.Entry<String, Integer>> result = new ArrayList<>();
        for (String w : words) {
            if (!stopwords.contains(w) && !w.isBlank()) {
                result.add(new AbstractMap.SimpleEntry<>(w, 1));
            }
        }
        return result;
    }

    private static Map<String, List<Map.Entry<String, Integer>>> regroup(Stream<List<Map.Entry<String, Integer>>> pairsList) {
        Map<String, List<Map.Entry<String, Integer>>> map = new HashMap<>();
        pairsList.forEach(list -> list.forEach(pair -> map.computeIfAbsent(pair.getKey(), k -> new ArrayList<>()).add(pair)));
        return map;
    }

    private static Map.Entry<String, Integer> countWords(Map.Entry<String, List<Map.Entry<String, Integer>>> entry) {
        int sum = entry.getValue().stream().map(Map.Entry::getValue).reduce(0, Integer::sum);
        return new AbstractMap.SimpleEntry<>(entry.getKey(), sum);
    }

    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Paths.get("../stop_words.txt"));
        lines.forEach(l -> stopwords.addAll(List.of(l.split(","))));
        stopwords.addAll(Arrays.asList("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"));

        String data = Files.readString(Paths.get(args[0]));
        List<String> chunks = partition(data, 200);
        Stream<List<Map.Entry<String, Integer>>> splits = chunks.stream().map(ThirtyTwo::splitWords);
        Map<String, List<Map.Entry<String, Integer>>> groups = regroup(splits);
        List<Map.Entry<String, Integer>> wordFreqs = groups.entrySet().stream()
                .map(ThirtyTwo::countWords)
                .sorted(Map.Entry.<String, Integer>comparingByValue().reversed())
                .toList();

        wordFreqs.subList(0, 25).forEach(e -> System.out.println(e.getKey() + " - " + e.getValue()));
    }
}
