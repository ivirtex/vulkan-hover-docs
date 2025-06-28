# IndexCountHUAWEI(3) Manual Page

## Name

IndexCountHUAWEI - cluster culling shader output variable



## [](#_description)Description

`IndexCountHUAWEI`

The `IndexCountHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this indexed mode specific variable will contain an integer value that specifies the number of indexed vertices in a cluster to draw.

Valid Usage

- [](#VUID-IndexCountHUAWEI-IndexCountHUAWEI-07805)VUID-IndexCountHUAWEI-IndexCountHUAWEI-07805  
  The `IndexCountHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-IndexCountHUAWEI-IndexCountHUAWEI-07806)VUID-IndexCountHUAWEI-IndexCountHUAWEI-07806  
  The variable decorated with `IndexCountHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#IndexCountHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0