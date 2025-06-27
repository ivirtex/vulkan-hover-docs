import * as vscode from 'vscode';
import * as fs from 'fs';

export class HoverProvider implements vscode.HoverProvider {
    constructor(private symbolToManPagePath: Map<string, string>) { }

    provideHover(document: vscode.TextDocument, position: vscode.Position, token: vscode.CancellationToken): vscode.ProviderResult<vscode.Hover> {
        const range = document.getWordRangeAtPosition(position);
        const word = document.getText(range);

        if (word.toLowerCase().startsWith('vk')) {
            if (this.symbolToManPagePath.has(word)) {
                const manPagePath = this.symbolToManPagePath.get(word);
                const manPageContent = fs.readFileSync(manPagePath!, 'utf8');

                const hoverContent = new vscode.MarkdownString(manPageContent);
                return new vscode.Hover(hoverContent);
            }
        }
    }
}