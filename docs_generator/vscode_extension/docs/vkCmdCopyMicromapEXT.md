# vkCmdCopyMicromapEXT(3) Manual Page

## Name

vkCmdCopyMicromapEXT - Copy a micromap



## [](#_c_specification)C Specification

To copy a micromap call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkCmdCopyMicromapEXT(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMicromapInfoEXT*                pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` is a pointer to a [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapInfoEXT.html) structure defining the copy operation.

## [](#_description)Description

This command copies the `pInfo->src` micromap to the `pInfo->dst` micromap in the manner specified by `pInfo->mode`.

Accesses to `pInfo->src` and `pInfo->dst` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_2_MICROMAP_READ_BIT_EXT` or `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT` as appropriate.

Valid Usage

- [](#VUID-vkCmdCopyMicromapEXT-buffer-07529)VUID-vkCmdCopyMicromapEXT-buffer-07529  
  The `buffer` used to create `pInfo->src` **must** be bound to device memory
- [](#VUID-vkCmdCopyMicromapEXT-buffer-07530)VUID-vkCmdCopyMicromapEXT-buffer-07530  
  The `buffer` used to create `pInfo->dst` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMicromapEXT-commandBuffer-parameter)VUID-vkCmdCopyMicromapEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMicromapEXT-pInfo-parameter)VUID-vkCmdCopyMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapInfoEXT.html) structure
- [](#VUID-vkCmdCopyMicromapEXT-commandBuffer-recording)VUID-vkCmdCopyMicromapEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMicromapEXT-commandBuffer-cmdpool)VUID-vkCmdCopyMicromapEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyMicromapEXT-renderpass)VUID-vkCmdCopyMicromapEXT-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMicromapEXT-videocoding)VUID-vkCmdCopyMicromapEXT-videocoding  
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

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMicromapEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0