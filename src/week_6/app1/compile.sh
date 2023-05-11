rm -rf *.class

javac -cp ../framework.jar Frequency.java Word.java

jar cfm ../deploy/app1.jar manifest.mf *.class
