-include .sdk/Makefile

$(if $(filter true,$(sdkloaded)),,$(error You must install bblfsh-sdk))

test-native-internal:
	cd native; \
	echo "yarn test"

build-native-internal:
	cd native; \
	yarn install; \
	cp *.js $(BUILD_PATH); \
	mv node_modules $(BUILD_PATH); \
	mv $(BUILD_PATH)/index.js $(BUILD_PATH)/native; \
	chmod +x $(BUILD_PATH)/native
