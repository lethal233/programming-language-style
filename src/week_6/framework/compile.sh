rm -rf *.class

javac IWord.java IFrequency.java Twenty.java

jar cfm ../framework.jar manifest.mf *.class