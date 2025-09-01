# vkCmdWriteAccelerationStructuresPropertiesKHR(3) Manual Page

## Name

vkCmdWriteAccelerationStructuresPropertiesKHR - Write acceleration structure result parameters to query results.



## [](#_c_specification)C Specification

To query acceleration structure size parameters call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdWriteAccelerationStructuresPropertiesKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    accelerationStructureCount,
    const VkAccelerationStructureKHR*           pAccelerationStructures,
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

Accesses to any of the acceleration structures listed in `pAccelerationStructures` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`.

- If `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, then the value written out is the number of bytes required by a compacted acceleration structure.
- If `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`, then the value written out is the number of bytes required by a serialized acceleration structure.

Valid Usage

- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructure-08924)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructure-08924  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02493)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02493  
  `queryPool` **must** have been created with a `queryType` matching `queryType`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02494)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02494  
  The queries identified by `queryPool` and `firstQuery` **must** be *unavailable*
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-buffer-03736)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-buffer-03736  
  The `buffer` used to create each acceleration structure in `pAccelerationStructures` **must** be bound to device memory
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-query-04880)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-query-04880  
  The sum of `firstQuery` plus `accelerationStructureCount` **must** be less than or equal to the number of queries in `queryPool`

<!--THE END-->

- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964  
  All acceleration structures in `pAccelerationStructures` **must** have been built prior to the execution of this command
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431  
  All acceleration structures in `pAccelerationStructures` **must** have been built with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` if `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-06742)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-06742  
  `queryType` **must** be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`, `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`, `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, or `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of `accelerationStructureCount` valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handles
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-parameter)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-recording)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-cmdpool)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-renderpass)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-videocoding)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`
- [](#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commonparent)VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commonparent  
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

Conditional Rendering

vkCmdWriteAccelerationStructuresPropertiesKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteAccelerationStructuresPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0