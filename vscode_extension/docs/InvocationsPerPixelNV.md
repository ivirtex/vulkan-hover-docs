# InvocationsPerPixelNV(3) Manual Page

## Name

InvocationsPerPixelNV - Number of fragment shader invocations for the current pixel



## [](#_description)Description

`InvocationsPerPixelNV`

Decorating a variable with the `InvocationsPerPixelNV` built-in decoration will make that variable contain the maximum number of fragment shader invocations per pixel, as derived from the effective shading rate for the fragment. If a primitive does not fully cover a pixel, the number of fragment shader invocations for that pixel **may** be less than the value of `InvocationsPerPixelNV`. If the shading rate indicates a fragment covering multiple pixels, then `InvocationsPerPixelNV` will be one.

Valid Usage

- [](#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04260)VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04260  
  The `InvocationsPerPixelNV` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04261)VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04261  
  The variable decorated with `InvocationsPerPixelNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04262)VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04262  
  The variable decorated with `InvocationsPerPixelNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InvocationsPerPixelNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0