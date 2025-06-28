# ClusterShadingRateHUAWEI(3) Manual Page

## Name

ClusterShadingRateHUAWEI - cluster culling shader output variable



## [](#_description)Description

`ClusterShadingRateHUAWEI`

The `ClusterShadingRateHUAWEI` decoration can be used to decorate a cluster culling shader output variable. This variable will contain an integer value specifying the shading rate of a rendering cluster.

Valid Usage

- [](#VUID-ClusterShadingRateHUAWEI-ClusterShadingRateHUAWEI-09448)VUID-ClusterShadingRateHUAWEI-ClusterShadingRateHUAWEI-09448  
  The `ClusterShadingRateHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-ClusterShadingRateHUAWEI-ClusterShadingRateHUAWEI-09449)VUID-ClusterShadingRateHUAWEI-ClusterShadingRateHUAWEI-09449  
  The variable decorated with `ClusterShadingRateHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ClusterShadingRateHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0