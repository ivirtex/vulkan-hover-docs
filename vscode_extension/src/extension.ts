// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';

import * as fs from 'fs';
import path from 'path';

import { HoverProvider } from './hover_provider';

const manPagesDir = "docs";

export function activate(context: vscode.ExtensionContext) {
	let manPagesPath = path.join(context.extensionPath, manPagesDir);
	let symbolToManPagePath = createSymbolToManPagePathMap(manPagesPath);

	let hoverProvider = vscode.languages.registerHoverProvider(['cpp', 'c'], new HoverProvider(symbolToManPagePath));

	context.subscriptions.push(hoverProvider);

	console.log('[Vulkan Hover Docs] Extension activated.');
}

export function createSymbolToManPagePathMap(dirPath: string): Map<string, string> {
	let symbolToPagePathMap = new Map<string, string>();

	// Save paths to all symbols
	fs.readdirSync(dirPath).forEach(symbol => {
		let symbolPath = path.join(dirPath, symbol);

		// Remove the file extension
		symbol = symbol.replace('.md', '');
		symbolToPagePathMap.set(symbol, symbolPath);
	});

	return symbolToPagePathMap;
}

export function deactivate() { }
