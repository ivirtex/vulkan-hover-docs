# FirstInstanceHUAWEI(3) Manual Page

## Name

FirstInstanceHUAWEI - cluster culling shader output variable



## [](#_description)Description

`FirstInstanceHUAWEI`

The `FirstInstanceHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this variable will contain an integer value that specifies the instance ID of the first instance to draw.

Valid Usage

- [](#VUID-FirstInstanceHUAWEI-FirstInstanceHUAWEI-07801)VUID-FirstInstanceHUAWEI-FirstInstanceHUAWEI-07801  
  The `FirstInstanceHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-FirstInstanceHUAWEI-FirstInstanceHUAWEI-07802)VUID-FirstInstanceHUAWEI-FirstInstanceHUAWEI-07802  
  The variable decorated with `FirstInstanceHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FirstInstanceHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0