# FragmentSizeNV(3) Manual Page

## Name

FragmentSizeNV - Size of the screen-space area covered by the fragment



## [](#_description)Description

`FragmentSizeNV`

Decorating a variable with the `FragmentSizeNV` built-in decoration will make that variable contain the width and height of the fragment.

Valid Usage

- [](#VUID-FragmentSizeNV-FragmentSizeNV-04226)VUID-FragmentSizeNV-FragmentSizeNV-04226  
  The `FragmentSizeNV` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragmentSizeNV-FragmentSizeNV-04227)VUID-FragmentSizeNV-FragmentSizeNV-04227  
  The variable decorated with `FragmentSizeNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FragmentSizeNV-FragmentSizeNV-04228)VUID-FragmentSizeNV-FragmentSizeNV-04228  
  The variable decorated with `FragmentSizeNV` **must** be declared as a two-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragmentSizeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0