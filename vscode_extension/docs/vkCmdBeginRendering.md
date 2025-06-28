# vkCmdBeginRendering(3) Manual Page

## Name

vkCmdBeginRendering - Begin a dynamic render pass instance



## [](#_c_specification)C Specification

To begin a render pass instance, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdBeginRendering(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_dynamic_rendering
void vkCmdBeginRenderingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pRenderingInfo` is a pointer to a [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) structure specifying details of the render pass instance to begin.

## [](#_description)Description

After beginning a render pass instance, the command buffer is ready to record [draw commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing).

If `pRenderingInfo->flags` includes `VK_RENDERING_RESUMING_BIT` then this render pass is resumed from a render pass instance that has been suspended earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

If there is an instance of [VkTileMemorySizeInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemorySizeInfoQCOM.html) included in the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), the structure is ignored.

Valid Usage

- [](#VUID-vkCmdBeginRendering-dynamicRendering-06446)VUID-vkCmdBeginRendering-dynamicRendering-06446  
  The [`dynamicRendering`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRendering) feature **must** be enabled
- [](#VUID-vkCmdBeginRendering-commandBuffer-06068)VUID-vkCmdBeginRendering-commandBuffer-06068  
  If `commandBuffer` is a secondary command buffer, and the [`nestedCommandBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nestedCommandBuffer) feature is not enabled, `pRenderingInfo->flags` **must** not include `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09588)VUID-vkCmdBeginRendering-pRenderingInfo-09588  
  If `pRenderingInfo->pDepthAttachment` is not `NULL` and `pRenderingInfo->pDepthAttachment->imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pDepthAttachment->imageView` **must** be in the layout specified by `pRenderingInfo->pDepthAttachment->imageLayout`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09589)VUID-vkCmdBeginRendering-pRenderingInfo-09589  
  If `pRenderingInfo->pDepthAttachment` is not `NULL`, `pRenderingInfo->pDepthAttachment->imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pDepthAttachment->imageResolveMode` is not `VK_RESOLVE_MODE_NONE`, and `pRenderingInfo->pDepthAttachment->resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pDepthAttachment->resolveImageView` **must** be in the layout specified by `pRenderingInfo->pDepthAttachment->resolveImageLayout`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09590)VUID-vkCmdBeginRendering-pRenderingInfo-09590  
  If `pRenderingInfo->pStencilAttachment` is not `NULL` and `pRenderingInfo->pStencilAttachment->imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pStencilAttachment->imageView` **must** be in the layout specified by `pRenderingInfo->pStencilAttachment->imageLayout`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09591)VUID-vkCmdBeginRendering-pRenderingInfo-09591  
  If `pRenderingInfo->pStencilAttachment` is not `NULL`, `pRenderingInfo->pStencilAttachment->imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pStencilAttachment->imageResolveMode` is not `VK_RESOLVE_MODE_NONE`, and `pRenderingInfo->pStencilAttachment->resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pRenderingInfo->pStencilAttachment->resolveImageView` **must** be in the layout specified by `pRenderingInfo->pStencilAttachment->resolveImageLayout`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09592)VUID-vkCmdBeginRendering-pRenderingInfo-09592  
  For any element of `pRenderingInfo->pColorAttachments`, if `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), that image view **must** be in the layout specified by `imageLayout`
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-09593)VUID-vkCmdBeginRendering-pRenderingInfo-09593  
  For any element of `pRenderingInfo->pColorAttachments`, if either `imageResolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, or `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, and `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `resolveImageView` **must** be in the layout specified by `resolveImageLayout`
- [](#VUID-vkCmdBeginRendering-flags-10641)VUID-vkCmdBeginRendering-flags-10641  
  If `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` is included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`, `commandBuffer` **must** not have been created with `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`
- [](#VUID-vkCmdBeginRendering-flags-10642)VUID-vkCmdBeginRendering-flags-10642  
  [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags` **must** not include `VK_TILE_SHADING_RENDER_PASS_PER_TILE_EXECUTION_BIT_QCOM`

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginRendering-commandBuffer-parameter)VUID-vkCmdBeginRendering-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginRendering-pRenderingInfo-parameter)VUID-vkCmdBeginRendering-pRenderingInfo-parameter  
  `pRenderingInfo` **must** be a valid pointer to a valid [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) structure
- [](#VUID-vkCmdBeginRendering-commandBuffer-recording)VUID-vkCmdBeginRendering-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginRendering-commandBuffer-cmdpool)VUID-vkCmdBeginRendering-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginRendering-renderpass)VUID-vkCmdBeginRendering-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBeginRendering-videocoding)VUID-vkCmdBeginRendering-videocoding  
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

Graphics

Action  
State

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginRendering)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0