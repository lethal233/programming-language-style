import java.util.*;

public interface IFrequency {
  public List<Map.Entry<String, Long>> top25(List<String> wordList);
}