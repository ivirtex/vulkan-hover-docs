# vkCmdCopyMemoryToAccelerationStructureKHR(3) Manual Page

## Name

vkCmdCopyMemoryToAccelerationStructureKHR - Copy device memory to an acceleration structure



## [](#_c_specification)C Specification

To copy device memory to an acceleration structure call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyMemoryToAccelerationStructureKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is a pointer to a [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

Accesses to `pInfo->dst` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`. Accesses to the buffer indicated by `pInfo->src.deviceAddress` **must** be synchronized with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an access type of `VK_ACCESS_TRANSFER_READ_BIT`.

This command can accept acceleration structures produced by either [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html) or [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html).

The structure provided as input to deserialize is as described in [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html), with any acceleration structure handles filled in with the newly-queried handles to bottom level acceleration structures created before deserialization. These do not need to be built at deserialize time, but **must** be created.

Valid Usage

- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-accelerationStructure-08927)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-accelerationStructure-08927  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03742)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03742  
  `pInfo->src.deviceAddress` **must** be a valid device address for a buffer bound to device memory
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03743)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03743  
  `pInfo->src.deviceAddress` **must** be aligned to `256` bytes
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03744)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03744  
  If the buffer pointed to by `pInfo->src.deviceAddress` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-buffer-03745)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-buffer-03745  
  The `buffer` used to create `pInfo->dst` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-parameter)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-parameter)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html) structure
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-recording)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-cmdpool)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-renderpass)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-videocoding)VUID-vkCmdCopyMemoryToAccelerationStructureKHR-videocoding  
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

vkCmdCopyMemoryToAccelerationStructureKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMemoryToAccelerationStructureKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0