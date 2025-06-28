# FragSizeEXT(3) Manual Page

## Name

FragSizeEXT - Size of the screen-space area covered by the fragment



## [](#_description)Description

`FragSizeEXT`

Decorating a variable with the `FragSizeEXT` built-in decoration will make that variable contain the dimensions in pixels of the [area](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-fragment-area) that the fragment covers for that invocation.

If fragment density map is not enabled, `FragSizeEXT` will be filled with a value of (1,1).

Valid Usage

- [](#VUID-FragSizeEXT-FragSizeEXT-04220)VUID-FragSizeEXT-FragSizeEXT-04220  
  The `FragSizeEXT` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragSizeEXT-FragSizeEXT-04221)VUID-FragSizeEXT-FragSizeEXT-04221  
  The variable decorated with `FragSizeEXT` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FragSizeEXT-FragSizeEXT-04222)VUID-FragSizeEXT-FragSizeEXT-04222  
  The variable decorated with `FragSizeEXT` **must** be declared as a two-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragSizeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0