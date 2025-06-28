# vkCmdWriteAccelerationStructuresPropertiesNV(3) Manual Page

## Name

vkCmdWriteAccelerationStructuresPropertiesNV - Write acceleration structure result parameters to query results.



## [](#_c_specification)C Specification

To query acceleration structure size parameters call:

```c++
// Provided by VK_NV_ray_tracing
void vkCmdWriteAccelerationStructuresPropertiesNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    accelerationStructureCount,
    const VkAccelerationStructureNV*            pAccelerationStructures,
    VkQueryType                                 queryType,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `accelerationStructureCount` is the count of acceleration structures for which to query the property.
- `pAccelerationStructures` is a pointer to an array of existing previously built acceleration structures.
- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value specifying the type of queries managed by the pool.
- `queryPool` is the query pool that will manage the results of the query.
- `firstQuery` is the first query index within the query pool that will contain the `accelerationStructureCount` number of results.

## [](#_description)Description

Accesses to any of the acceleration structures listed in `pAccelerationStructures` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`.

Valid Usage

- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03755)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03755  
  `queryPool` **must** have been created with a `queryType` matching `queryType`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03756)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03756  
  The queries identified by `queryPool` and `firstQuery` **must** be *unavailable*
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructure-03757)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructure-03757  
  `accelerationStructure` **must** be bound completely and contiguously to a single `VkDeviceMemory` object via [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-04958)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-04958  
  All acceleration structures in `pAccelerationStructures` **must** have been built prior to the execution of this command
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-06215)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-06215  
  All acceleration structures in `pAccelerationStructures` **must** have been built with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` if `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-06216)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-06216  
  `queryType` **must** be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of `accelerationStructureCount` valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handles
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-recording)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-cmdpool)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-renderpass)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-videocoding)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructureCount-arraylength)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commonparent)VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commonparent  
  Each of `commandBuffer`, `queryPool`, and the elements of `pAccelerationStructures` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteAccelerationStructuresPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0