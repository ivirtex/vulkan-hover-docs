# FrontFacing(3) Manual Page

## Name

FrontFacing - Front face determination of a fragment



## [](#_description)Description

`FrontFacing`

Decorating a variable with the `FrontFacing` built-in decoration will make that variable contain whether the fragment is front or back facing. This variable is non-zero if the current fragment is considered to be part of a [front-facing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygons-basic) polygon primitive or of a non-polygon primitive and is zero if the fragment is considered to be part of a back-facing polygon primitive.

Valid Usage

- [](#VUID-FrontFacing-FrontFacing-04229)VUID-FrontFacing-FrontFacing-04229  
  The `FrontFacing` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FrontFacing-FrontFacing-04230)VUID-FrontFacing-FrontFacing-04230  
  The variable decorated with `FrontFacing` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FrontFacing-FrontFacing-04231)VUID-FrontFacing-FrontFacing-04231  
  The variable decorated with `FrontFacing` **must** be declared as a boolean value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FrontFacing).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0