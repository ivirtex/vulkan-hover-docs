# Vulkan Hover Docs

## Overview 
This repo consists of 2 projects:
- **vk_docs_generator**, which is a small CLI written in Rust for downloading, converting and post-processing Vulkan manual pages from Khronos registry.
- **vscode_ext**, which is an VS Code extension that matches Vulkan symbols in your workspace to generated Markdown files from **vk_docs_generator**.

### Example:
![preview](https://github.com/ivirtex/vulkan-hover-docs/blob/master/vscode_ext/images/screenshot.png?raw=true)

## Motivation
If you have ever wrote something in Vulkan, you probably know how verbose and extensive this API is.  
For me, going back and forth to the documentation was pretty annoying and time consuming, that's why I created this project to help others with learning and using Vulkan.
