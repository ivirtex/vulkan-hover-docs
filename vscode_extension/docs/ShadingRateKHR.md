# ShadingRateKHR(3) Manual Page

## Name

ShadingRateKHR - Shading rate of a fragment



## [](#_description)Description

`ShadingRateKHR`

Decorating a variable with the `ShadingRateKHR` built-in decoration will make that variable contain the [fragment shading rate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate) for the current fragment invocation.

Valid Usage

- [](#VUID-ShadingRateKHR-ShadingRateKHR-04490)VUID-ShadingRateKHR-ShadingRateKHR-04490  
  The `ShadingRateKHR` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-ShadingRateKHR-ShadingRateKHR-04491)VUID-ShadingRateKHR-ShadingRateKHR-04491  
  The variable decorated with `ShadingRateKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ShadingRateKHR-ShadingRateKHR-04492)VUID-ShadingRateKHR-ShadingRateKHR-04492  
  The variable decorated with `ShadingRateKHR` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ShadingRateKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0