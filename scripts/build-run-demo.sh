set -e

cd ../demo-service

echo "==> Building demo-service..."
./gradlew clean build

echo "==> Starting demo-service..."
java -jar build/libs/demo-*.jar

cd -

exit
