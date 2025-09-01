# InstanceCountHUAWEI(3) Manual Page

## Name

InstanceCountHUAWEI - cluster culling shader output variable



## [](#_description)Description

`InstanceCountHUAWEI`

The `InstanceCountHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this variable will contain an integer value that specifies the number of instance to draw in a cluster.

Valid Usage

- [](#VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07807)VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07807  
  The `InstanceCountHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07808)VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07808  
  The variable decorated with `InstanceCountHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InstanceCountHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0