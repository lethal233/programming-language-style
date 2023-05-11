cd ~/APL-Java/Week6/framework

javac IWord.java IFrequency.java Twenty.java

jar cfm ~/APL-Java/Week6/framework.jar manifest.mf *.class

cd ~/APL-Java/Week6/app1

chmod +x compile.sh

./compile.sh

cd ~/APL-Java/Week6/app2

chmod +x compile.sh

./compile.sh
