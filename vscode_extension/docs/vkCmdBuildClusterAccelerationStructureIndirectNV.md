# vkCmdBuildClusterAccelerationStructureIndirectNV(3) Manual Page

## Name

vkCmdBuildClusterAccelerationStructureIndirectNV - Build or move cluster acceleration structures



## [](#_c_specification)C Specification

To build or move a cluster acceleration structure or a cluster acceleration structure template call:

```c++
// Provided by VK_NV_cluster_acceleration_structure
void vkCmdBuildClusterAccelerationStructureIndirectNV(
    VkCommandBuffer                             commandBuffer,
    const VkClusterAccelerationStructureCommandsInfoNV* pCommandInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pCommandInfos` is a pointer to a [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html) structure containing parameters required for building or moving the cluster acceleration structure.

## [](#_description)Description

Similar to [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresKHR.html), this command **may** initiate multiple acceleration structures builds and there is no ordering or synchronization implied between any of the individual acceleration structure builds. Accesses to the acceleration structure scratch memory as identified by the [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`scratchData` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of (`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` | `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`).

Accesses to each [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstImplicitData`, [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstAddressesArray` and [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstSizesArray` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`.

Accesses to memory with input data as identified by any used values of [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`srcInfosArray`, [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`srcInfosCount` and [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`addressResolutionFlags` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_INDIRECT_COMMAND_READ_BIT`.

Valid Usage

- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-clusterAccelerationStructure-10443)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-clusterAccelerationStructure-10443  
  The [`VkPhysicalDeviceClusterAccelerationStructureFeaturesNV`::`clusterAccelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-clusterAccelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pNext-10444)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pNext-10444  
  The `pNext` chain of the bound ray tracing pipeline **must** include a [VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV.html) structure
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10445)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10445  
  `pCommandInfos->input`::`maxAccelerationStructureCount` **must** be less than or equal to the value used in `pInfo->maxAccelerationStructureCount` in [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html) to determine the memory requirements for the build operation
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-scratchData-10446)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-scratchData-10446  
  The scratch memory of the cluster acceleration structure specified in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`scratchData` **must** be larger than or equal to the scratch size queried with [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html)
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-scratchData-10447)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-scratchData-10447  
  The scratch address of the cluster acceleration structure specified in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`scratchData` **must** be aligned based on the cluster acceleration structure type and its alignment properties as queried with [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10448)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10448  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureMoveObjectsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10449)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10449  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10450)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10450  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10451)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10451  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10452)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10452  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureInstantiateClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInstantiateClusterInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10832)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10832  
  If `pCommandInfos->input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_GET_CLUSTER_TEMPLATE_INDICES_NV`, `pCommandInfos->srcInfosArray` **must** be an array of [VkClusterAccelerationStructureGetTemplateIndicesInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGetTemplateIndicesInfoNV.html) structures
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10453)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10453  
  The value in `pCommandInfos->srcInfosCount` **must** be less than or equal to `pCommandInfos->input`::`maxAccelerationStructureCount`
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10454)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10454  
  The number of inputs specified in `pCommandInfos->srcInfosArray` **must** be greater than or equal to `pCommandInfos->srcInfosCount`
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-dstAddressesArray-10455)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-dstAddressesArray-10455  
  The memory regions specified in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstAddressesArray` **must** not overlap with each other or with `pCommandInfos->scratchData`
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-dstImplicitData-10456)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-dstImplicitData-10456  
  The memory region specified in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstImplicitData` for multiple acceleration structure builds **must** not overlap with `pCommandInfos->scratchData`
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10457)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10457  
  The buffer from which the buffer device address for `pCommandInfos->scratchData` is queried **must** have been created with the `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10458)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10458  
  The buffers from which the buffer device addresses for `pCommandInfos->srcInfosArray`, `pCommandInfos->srcInfosCount` and `pCommandInfos->addressResolutionFlags` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10459)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10459  
  The buffers from which the buffer device addresses for `pCommandInfos->dstImplicitData` and `pCommandInfos->dstAddressesArray` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10460)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10460  
  If `pCommandInfos->dstImplicitData` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10461)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10461  
  If `pCommandInfos->scratchData` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10462)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10462  
  If `pCommandInfos->srcInfosCount` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10463)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10463  
  If the addresses specified in `pCommandInfos->dstAddressesArray` are the address of a non-sparse buffer then they each **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10464)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10464  
  If the addresses specified in `pCommandInfos->dstSizesArray` are the address of a non-sparse buffer then they each **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10465)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-10465  
  If the addresses specified in `pCommandInfos->srcInfosArray` are the address of a non-sparse buffer then they each **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-parameter)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-parameter)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-pCommandInfos-parameter  
  `pCommandInfos` **must** be a valid pointer to a valid [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html) structure
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-recording)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-cmdpool)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-renderpass)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-videocoding)VUID-vkCmdBuildClusterAccelerationStructureIndirectNV-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Compute

Action

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildClusterAccelerationStructureIndirectNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0