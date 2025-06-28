# WarpIDNV(3) Manual Page

## Name

WarpIDNV - Warp ID within an SM of a shader invocation



## [](#_description)Description

`WarpIDNV`

Decorating a variable with the `WarpIDNV` built-in decoration will make that variable contain the ID of the warp on a SM on which the current shader invocation is running. This variable is in the range \[0, `WarpsPerSMNV`-1].

Valid Usage

- [](#VUID-WarpIDNV-WarpIDNV-04420)VUID-WarpIDNV-WarpIDNV-04420  
  The variable decorated with `WarpIDNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WarpIDNV-WarpIDNV-04421)VUID-WarpIDNV-WarpIDNV-04421  
  The variable decorated with `WarpIDNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WarpIDNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0