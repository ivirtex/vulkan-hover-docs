# vkCmdCopyAccelerationStructureToMemoryKHR(3) Manual Page

## Name

vkCmdCopyAccelerationStructureToMemoryKHR - Copy an acceleration structure to device memory



## [](#_c_specification)C Specification

To copy an acceleration structure to device memory call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyAccelerationStructureToMemoryKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is an a pointer to a [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

Accesses to `pInfo->src` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`. Accesses to the buffer indicated by `pInfo->dst.deviceAddress` **must** be synchronized with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an and an access type of `VK_ACCESS_TRANSFER_WRITE_BIT`.

This command produces the same results as [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html), but writes its result to a device address, and is executed on the device rather than the host. The output **may** not necessarily be bit-for-bit identical, but it can be equally used by either [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html) or [vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToAccelerationStructureKHR.html).

The defined header structure for the serialized data consists of:

- `VK_UUID_SIZE` bytes of data matching `VkPhysicalDeviceIDProperties`::`driverUUID`
- `VK_UUID_SIZE` bytes of data identifying the compatibility for comparison using [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)
- A 64-bit integer of the total size matching the value queried using `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`
- A 64-bit integer of the deserialized size to be passed in to `VkAccelerationStructureCreateInfoKHR`::`size`
- A 64-bit integer of the count of the number of acceleration structure handles following. This value matches the value queried using `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`. This will be zero for a bottom-level acceleration structure. For top-level acceleration structures this number is implementation-dependent; the number of and ordering of the handles may not match the instance descriptions which were used to build the acceleration structure.

The corresponding handles matching the values returned by [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureDeviceAddressKHR.html) or [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureHandleNV.html) are tightly packed in the buffer following the count. The application is expected to store a mapping between those handles and the original application-generated bottom-level acceleration structures to provide when deserializing. The serialized data is written to the buffer (or read from the buffer) according to the host endianness.

Valid Usage

- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-accelerationStructure-08926)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-accelerationStructure-08926  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03739)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03739  
  `pInfo->dst.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03740)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03740  
  `pInfo->dst.deviceAddress` **must** be aligned to `256` bytes
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-None-03559)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-None-03559  
  The `buffer` used to create `pInfo->src` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-parameter)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-parameter)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html) structure
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-recording)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-cmdpool)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-renderpass)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-videocoding)VUID-vkCmdCopyAccelerationStructureToMemoryKHR-videocoding  
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

vkCmdCopyAccelerationStructureToMemoryKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyAccelerationStructureToMemoryKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0