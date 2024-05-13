// You can import and use all API from the 'vscode' module
// as well as import your extension to test it
import * as vscode from 'vscode';
import * as assert from 'assert';
import * as vulkanInlineDocs from '../extension';

import path from 'path';

const test_file_dir = 'test_files';

suite('Vulkan Hover Docs Test Suite', () => {
	vscode.window.showInformationMessage('Start all tests.');
	let test_root = path.resolve(__dirname, '../../src/test');
	let test_files_path = path.join(test_root, test_file_dir);

	test("createSymbolToManPagePathMap returns correct map", () => {
		let symbolToManPagePath = vulkanInlineDocs.createSymbolToManPagePathMap(test_files_path);

		let expectedMap = new Map<string, string>();
		expectedMap.set('PFN_vkDebugUtilsMessengerCallbackEXT', path.join(test_files_path, 'PFN_vkDebugUtilsMessengerCallbackEXT.md'));
		expectedMap.set('VK_SUBPASS_EXTERNAL', path.join(test_files_path, 'VK_SUBPASS_EXTERNAL.md'));
		expectedMap.set('VkAccelerationStructureInstanceKHR', path.join(test_files_path, 'VkAccelerationStructureInstanceKHR.md'));
		expectedMap.set('VkBufferCopy', path.join(test_files_path, 'VkBufferCopy.md'));

		assert.deepEqual(symbolToManPagePath, expectedMap);
	});
});
