rm -rf *.clas

javac -cp ../framework.jar Frequency.java Word.java

jar cfm ../deploy/app2.jar manifest.mf *.class
