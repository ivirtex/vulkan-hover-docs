# vkCmdCopyMemoryToMicromapEXT(3) Manual Page

## Name

vkCmdCopyMemoryToMicromapEXT - Copy device memory to a micromap



## [](#_c_specification)C Specification

To copy device memory to a micromap call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkCmdCopyMemoryToMicromapEXT(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMemoryToMicromapInfoEXT*        pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is a pointer to a [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html) structure defining the copy operation.

## [](#_description)Description

Accesses to `pInfo->dst` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`. Accesses to the buffer indicated by `pInfo->src.deviceAddress` **must** be synchronized with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage and an access type of `VK_ACCESS_TRANSFER_READ_BIT`.

This command can accept micromaps produced by either [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html) or [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html).

Valid Usage

- [](#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07543)VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07543  
  `pInfo->src.deviceAddress` **must** be a valid device address for a buffer bound to device memory
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07544)VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07544  
  `pInfo->src.deviceAddress` **must** be aligned to `256` bytes
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07545)VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07545  
  If the buffer pointed to by `pInfo->src.deviceAddress` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-buffer-07546)VUID-vkCmdCopyMemoryToMicromapEXT-buffer-07546  
  The `buffer` used to create `pInfo->dst` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-parameter)VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-parameter)VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html) structure
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-recording)VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-cmdpool)VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-renderpass)VUID-vkCmdCopyMemoryToMicromapEXT-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMemoryToMicromapEXT-videocoding)VUID-vkCmdCopyMemoryToMicromapEXT-videocoding  
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

vkCmdCopyMemoryToMicromapEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMemoryToMicromapEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0