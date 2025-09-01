# SMIDNV(3) Manual Page

## Name

SMIDNV - SM ID on which a shader invocation is running



## [](#_description)Description

`SMIDNV`

Decorating a variable with the `SMIDNV` built-in decoration will make that variable contain the ID of the SM on which the current shader invocation is running. This variable is in the range \[0, `SMCountNV`-1].

Valid Usage

- [](#VUID-SMIDNV-SMIDNV-04365)VUID-SMIDNV-SMIDNV-04365  
  The variable decorated with `SMIDNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SMIDNV-SMIDNV-04366)VUID-SMIDNV-SMIDNV-04366  
  The variable decorated with `SMIDNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SMIDNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0