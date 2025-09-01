# FullyCoveredEXT(3) Manual Page

## Name

FullyCoveredEXT - Indication of whether a fragment is fully covered



## [](#_description)Description

`FullyCoveredEXT`

Decorating a variable with the `FullyCoveredEXT` built-in decoration will make that variable indicate whether the [fragment area](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-fragment-area) is fully covered by the generating primitive. This variable is non-zero if conservative rasterization is enabled and the current fragment area is fully covered by the generating primitive, and is zero if the fragment is not covered or partially covered, or conservative rasterization is disabled.

Valid Usage

- [](#VUID-FullyCoveredEXT-FullyCoveredEXT-04232)VUID-FullyCoveredEXT-FullyCoveredEXT-04232  
  The `FullyCoveredEXT` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FullyCoveredEXT-FullyCoveredEXT-04233)VUID-FullyCoveredEXT-FullyCoveredEXT-04233  
  The variable decorated with `FullyCoveredEXT` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FullyCoveredEXT-FullyCoveredEXT-04234)VUID-FullyCoveredEXT-FullyCoveredEXT-04234  
  The variable decorated with `FullyCoveredEXT` **must** be declared as a boolean value
- [](#VUID-FullyCoveredEXT-conservativeRasterizationPostDepthCoverage-04235)VUID-FullyCoveredEXT-conservativeRasterizationPostDepthCoverage-04235  
  If `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`conservativeRasterizationPostDepthCoverage` is not supported the `PostDepthCoverage` `Execution` `Mode` **must** not be declared, when a variable with the `FullyCoveredEXT` decoration is declared

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FullyCoveredEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0