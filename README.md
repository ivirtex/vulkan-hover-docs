# Vulkan Hover Docs

This repo consists of 2 projects:

- **docs_generator**, which is a small CLI written in Go for scraping, converting and post-processing HTML pages from Khronos registry.
- **vscode_extension**, which is an VS Code extension that matches Vulkan symbols in the editor to Markdown pages generated by `docs_generator`.

### Example

![preview](https://github.com/ivirtex/vulkan-hover-docs/blob/master/vscode_extension/images/screenshot.png?raw=true)

## Motivation

If you have ever wrote something in Vulkan, you probably know how verbose and extensive this API is.  
For me, going back and forth to the documentation was pretty annoying and time consuming, that's why I created this project to help others with learning and using Vulkan.
