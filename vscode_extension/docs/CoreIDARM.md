# CoreIDARM(3) Manual Page

## Name

CoreIDARM - Core ID on which a shader invocation is running



## [](#_description)Description

`CoreIDARM`

Decorating a variable with the `CoreIDARM` built-in decoration will make that variable contain the ID of the core on which the current shader invocation is running. This variable is in the range \[0, `CoreMaxIDARM`].

Valid Usage

- [](#VUID-CoreIDARM-CoreIDARM-07599)VUID-CoreIDARM-CoreIDARM-07599  
  The variable decorated with `CoreIDARM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-CoreIDARM-CoreIDARM-07600)VUID-CoreIDARM-CoreIDARM-07600  
  The variable decorated with `CoreIDARM` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CoreIDARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0