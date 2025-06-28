# WarpMaxIDARM(3) Manual Page

## Name

WarpMaxIDARM - Max ID for a warp on the core running a shader invovation



## [](#_description)Description

`WarpMaxIDARM`

Decorating a variable with the `WarpMaxIDARM` built-in decoration will make that variable contain the maximum warp ID for the core on which the current invocation is running.

Valid Usage

- [](#VUID-WarpMaxIDARM-WarpMaxIDARM-07601)VUID-WarpMaxIDARM-WarpMaxIDARM-07601  
  The variable decorated with `WarpMaxIDARM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WarpMaxIDARM-WarpMaxIDARM-07602)VUID-WarpMaxIDARM-WarpMaxIDARM-07602  
  The variable decorated with `WarpMaxIDARM` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WarpMaxIDARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0