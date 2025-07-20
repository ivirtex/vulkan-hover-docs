# vkCmdBuildPartitionedAccelerationStructuresNV(3) Manual Page

## Name

vkCmdBuildPartitionedAccelerationStructuresNV - Command for building a PTLAS



## [](#_c_specification)C Specification

To build a partitioned top level acceleration structure, call:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
void vkCmdBuildPartitionedAccelerationStructuresNV(
    VkCommandBuffer                             commandBuffer,
    const VkBuildPartitionedAccelerationStructureInfoNV* pBuildInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pBuildInfo` is a pointer to a [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html) structure containing parameters required for building a PTLAS.

## [](#_description)Description

Accesses to the acceleration structure scratch memory as identified by the [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`scratchData` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of (`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` | `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`).

Accesses to each [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`srcAccelerationStructureData` and [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`dstAccelerationStructureData` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, as appropriate.

Accesses to memory with input data as identified by any used values of [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`srcInfos` and [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`srcInfosCount` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_INDIRECT_COMMAND_READ_BIT`.

Valid Usage

- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-partitionedAccelerationStructure-10536)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-partitionedAccelerationStructure-10536  
  The [`VkPhysicalDevicePartitionedAccelerationStructureFeaturesNV`::`partitionedAccelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-partitionedAccelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10537)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10537  
  The count specified in `pBuildInfo->input`::`instanceCount` for the build operation **must** not exceed the value provided in `pInfo->instanceCount` when calling [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) to determine the memory size
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10538)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10538  
  The count specified in `pBuildInfo->input`::`maxInstancePerPartitionCount` for the build operation **must** not exceed the value provided in `pInfo->maxInstancePerPartitionCount` when calling [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) to determine the memory size
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10539)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10539  
  The count specified in `pBuildInfo->input`::`partitionCount` for the build operation **must** not exceed the value provided in `pInfo->partitionCount` when calling [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) to determine the memory size
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10540)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10540  
  The count specified in `pBuildInfo->input`::`maxInstanceInGlobalPartitionCount` for the build operation **must** not exceed the value provided in `pInfo->maxInstanceInGlobalPartitionCount` when calling [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) to determine the memory size
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10541)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10541  
  The scratch memory for the partitioned acceleration structure build specified in `pBuildInfo->scratchData` **must** be larger than or equal to the scratch size queried with [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html)
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10542)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10542  
  `pBuildInfo->scratchData` **must** be aligned to `256` bytes
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10543)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10543  
  The destination memory of the partitioned acceleration structure build specified in `pBuildInfo->dstAccelerationStructureData` **must** be larger than or equal to the size queried with [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html)
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10544)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10544  
  `pBuildInfo->srcAccelerationStructureData` **must** be aligned to `256` bytes
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10545)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10545  
  `pBuildInfo->dstAccelerationStructureData` **must** be aligned to `256` bytes
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10546)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10546  
  The number of inputs specified in `pBuildInfo->srcInfos` **must** be greater than or equal to `pBuildInfo->srcInfosCount`
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10547)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10547  
  The memory region containing the acceleration structure at address `pBuildInfo->srcAccelerationStructureData` **must** not overlap with scratch memory region at address `pBuildInfo->scratchData`
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10548)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10548  
  The memory region containing the acceleration structure at address `pBuildInfo->dstAccelerationStructureData` **must** not overlap with scratch memory region at address `pBuildInfo->scratchData`
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10549)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10549  
  The memory regions containing the acceleration structures at addresses `pBuildInfo->srcAccelerationStructureData` and `pBuildInfo->dstAccelerationStructureData` **must** not overlap with each other
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10550)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10550  
  The buffer from which the buffer device address for `pBuildInfo->scratchData` is queried **must** have been created with the `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10551)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10551  
  The buffers from which the buffer device addresses for `pBuildInfo->srcInfos` and `pBuildInfo->srcInfosCount` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10552)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10552  
  The buffers from which the buffer device addresses for `pBuildInfo->srcAccelerationStructureData` and `pBuildInfo->dstAccelerationStructureData` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10553)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10553  
  If `pBuildInfo->srcAccelerationStructureData` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10554)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10554  
  If `pBuildInfo->dstAccelerationStructureData` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10555)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10555  
  If `pBuildInfo->scratchData` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10556)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10556  
  If `pBuildInfo->srcInfos` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10557)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-10557  
  If `pBuildInfo->srcInfosCount` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-parameter)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-parameter)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-pBuildInfo-parameter  
  `pBuildInfo` **must** be a valid pointer to a valid [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html) structure
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-recording)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-cmdpool)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-renderpass)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildPartitionedAccelerationStructuresNV-videocoding)VUID-vkCmdBuildPartitionedAccelerationStructuresNV-videocoding  
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

Conditional Rendering

vkCmdBuildPartitionedAccelerationStructuresNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildPartitionedAccelerationStructuresNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0