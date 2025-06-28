# SMCountNV(3) Manual Page

## Name

SMCountNV - Number of SMs on the device



## [](#_description)Description

`SMCountNV`

Decorating a variable with the `SMCountNV` built-in decoration will make that variable contain the number of SMs on the device.

Valid Usage

- [](#VUID-SMCountNV-SMCountNV-04363)VUID-SMCountNV-SMCountNV-04363  
  The variable decorated with `SMCountNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SMCountNV-SMCountNV-04364)VUID-SMCountNV-SMCountNV-04364  
  The variable decorated with `SMCountNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SMCountNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0