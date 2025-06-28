# ClusterIDHUAWEI(3) Manual Page

## Name

ClusterIDHUAWEI - cluster culling shader output variable



## [](#_description)Description

`ClusterIDHUAWEI`

The `ClusterIDHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this variable will contain an integer value that specifies the id of cluster being rendered by this drawing command. When Cluster Culling Shader enable, `ClusterIDHUAWEI` will replace gl\_DrawID pass to vertex shader for cluster-related information fetching.

Valid Usage

- [](#VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07797)VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07797  
  The `ClusterIDHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07798)VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07798  
  The variable decorated with `ClusterIDHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ClusterIDHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0