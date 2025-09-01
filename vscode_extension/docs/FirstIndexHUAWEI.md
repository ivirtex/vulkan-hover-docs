# FirstIndexHUAWEI(3) Manual Page

## Name

FirstIndexHUAWEI - cluster culling shader output variable



## [](#_description)Description

`FirstIndexHUAWEI`

The `FirstIndexHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this indexed mode specific variable will contain an integer value that specifies the base index within the index buffer corresponding to a cluster.

Valid Usage

- [](#VUID-FirstIndexHUAWEI-FirstIndexHUAWEI-07799)VUID-FirstIndexHUAWEI-FirstIndexHUAWEI-07799  
  The `FirstIndexHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-FirstIndexHUAWEI-FirstIndexHUAWEI-07800)VUID-FirstIndexHUAWEI-FirstIndexHUAWEI-07800  
  The variable decorated with `FirstIndexHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FirstIndexHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0