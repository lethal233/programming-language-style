import java.io.FileInputStream;
import java.io.IOException;
import java.lang.reflect.Method;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.List;
import java.util.Map;
import java.util.Properties;

public class Twenty {
    public static void main(String[] args) throws Exception {
        Properties properties = new Properties();
        try {
            properties.load(new FileInputStream("config.properties"));
        } catch (IOException e) {
            e.printStackTrace();
            return;
        }
        try {
          String wordPlugin = properties.getProperty("word");
          String freqPlugin = properties.getProperty("freq");
          Class cls1 = loadClass(wordPlugin, "Word");
          Class cls2 = loadClass(freqPlugin, "Frequency");

          IWord word = (IWord) cls1.newInstance();
          IFrequency freq = (IFrequency) cls2.newInstance();

          List<Map.Entry<String, Long>> res = freq.top25(word.extractWords(args[0]));

          for (Map.Entry<String, Long> re : res) {
              System.out.println(re.getKey() + "  -  " + re.getValue());
          }
        } catch (Exception e) {
          e.printStackTrace();
        }
        
    }

    private static Class loadClass(String jarPath, String className) throws Exception {
        URL jarURL = new URL("file:" + jarPath);
        URLClassLoader classLoader = new URLClassLoader(new URL[]{jarURL});
        return classLoader.loadClass(className);
    }
}