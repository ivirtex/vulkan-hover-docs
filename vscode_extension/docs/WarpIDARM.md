# WarpIDARM(3) Manual Page

## Name

WarpIDARM - Warp ID within a core of a shader invocation



## [](#_description)Description

`WarpIDARM`

Decorating a variable with the `WarpIDARM` built-in decoration will make that variable contain the ID of the warp on a core on which the current shader invocation is running. This variable is in the range \[0, `WarpMaxIDARM`].

Valid Usage

- [](#VUID-WarpIDARM-WarpIDARM-07603)VUID-WarpIDARM-WarpIDARM-07603  
  The variable decorated with `WarpIDARM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WarpIDARM-WarpIDARM-07604)VUID-WarpIDARM-WarpIDARM-07604  
  The variable decorated with `WarpIDARM` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WarpIDARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0