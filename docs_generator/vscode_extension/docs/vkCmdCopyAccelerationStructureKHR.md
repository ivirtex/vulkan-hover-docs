# vkCmdCopyAccelerationStructureKHR(3) Manual Page

## Name

vkCmdCopyAccelerationStructureKHR - Copy an acceleration structure



## [](#_c_specification)C Specification

To copy an acceleration structure call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyAccelerationStructureKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyAccelerationStructureInfoKHR*   pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is a pointer to a [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

This command copies the `pInfo->src` acceleration structure to the `pInfo->dst` acceleration structure in the manner specified by `pInfo->mode`.

Accesses to `pInfo->src` and `pInfo->dst` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` as appropriate.

Valid Usage

- [](#VUID-vkCmdCopyAccelerationStructureKHR-accelerationStructure-08925)VUID-vkCmdCopyAccelerationStructureKHR-accelerationStructure-08925  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdCopyAccelerationStructureKHR-buffer-03737)VUID-vkCmdCopyAccelerationStructureKHR-buffer-03737  
  The `buffer` used to create `pInfo->src` **must** be bound to device memory
- [](#VUID-vkCmdCopyAccelerationStructureKHR-buffer-03738)VUID-vkCmdCopyAccelerationStructureKHR-buffer-03738  
  The `buffer` used to create `pInfo->dst` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-parameter)VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyAccelerationStructureKHR-pInfo-parameter)VUID-vkCmdCopyAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html) structure
- [](#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-recording)VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-cmdpool)VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyAccelerationStructureKHR-renderpass)VUID-vkCmdCopyAccelerationStructureKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyAccelerationStructureKHR-videocoding)VUID-vkCmdCopyAccelerationStructureKHR-videocoding  
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

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyAccelerationStructureKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0