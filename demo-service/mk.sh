set -e

rm -rf logs

echo "==> Building demo-service"
./gradlew clean build
# ./gradlew clean build -x test

echo "==> Running demo-service"
java -jar build/libs/demo-*.jar