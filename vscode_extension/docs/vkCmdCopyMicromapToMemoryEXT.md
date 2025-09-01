# vkCmdCopyMicromapToMemoryEXT(3) Manual Page

## Name

vkCmdCopyMicromapToMemoryEXT - Copy a micromap to device memory



## [](#_c_specification)C Specification

To copy a micromap to device memory call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkCmdCopyMicromapToMemoryEXT(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMicromapToMemoryInfoEXT*        pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is an a pointer to a [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html) structure defining the copy operation.

## [](#_description)Description

Accesses to `pInfo->src` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`. Accesses to the buffer indicated by `pInfo->dst.deviceAddress` **must** be synchronized with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage and an access type of `VK_ACCESS_TRANSFER_WRITE_BIT`.

This command produces the same results as [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html), but writes its result to a device address, and is executed on the device rather than the host. The output **may** not necessarily be bit-for-bit identical, but it can be equally used by either [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToMicromapEXT.html) or [vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToMicromapEXT.html).

The defined header structure for the serialized data consists of:

- `VK_UUID_SIZE` bytes of data matching `VkPhysicalDeviceIDProperties`::`driverUUID`
- `VK_UUID_SIZE` bytes of data identifying the compatibility for comparison using [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMicromapCompatibilityEXT.html) The serialized data is written to the buffer (or read from the buffer) according to the host endianness.

Valid Usage

- [](#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07536)VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07536  
  `pInfo->dst.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07537)VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07537  
  `pInfo->dst.deviceAddress` **must** be aligned to `256` bytes
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07538)VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07538  
  If the buffer pointed to by `pInfo->dst.deviceAddress` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-buffer-07539)VUID-vkCmdCopyMicromapToMemoryEXT-buffer-07539  
  The `buffer` used to create `pInfo->src` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-parameter)VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-parameter)VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html) structure
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-recording)VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-cmdpool)VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-renderpass)VUID-vkCmdCopyMicromapToMemoryEXT-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMicromapToMemoryEXT-videocoding)VUID-vkCmdCopyMicromapToMemoryEXT-videocoding  
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

vkCmdCopyMicromapToMemoryEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMicromapToMemoryEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0