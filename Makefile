include .sdk/Makefile

test-native:
	cd native; \
	echo "yarn test"

build-native:
	cd native; \
	yarn install; \
	cp *.js $(BUILD_PATH); \
	mv node_modules $(BUILD_PATH); \
	mv $(BUILD_PATH)/index.js $(BUILD_PATH)/native; \
	chmod +x $(BUILD_PATH)/native
