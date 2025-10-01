# vkCmdPipelineBarrier2(3) Manual Page

## Name

vkCmdPipelineBarrier2 - Insert a memory dependency



## [](#_c_specification)C Specification

To record a pipeline barrier, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdPipelineBarrier2(
    VkCommandBuffer                             commandBuffer,
    const VkDependencyInfo*                     pDependencyInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
void vkCmdPipelineBarrier2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkDependencyInfo*                     pDependencyInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pDependencyInfo` is a pointer to a [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structure defining the scopes of this operation.

## [](#_description)Description

When [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2.html) is submitted to a queue, it defines memory dependencies between commands that were submitted to the same queue before it, and those submitted to the same queue after it.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) of each memory dependency defined by `pDependencyInfo` are applied to operations that occurred earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) of each memory dependency defined by `pDependencyInfo` are applied to operations that occurred later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

If `vkCmdPipelineBarrier2` is recorded within a render pass instance, the synchronization scopes are limited to a subset of operations within the same subpass or render pass instance.

Valid Usage

- [](#VUID-vkCmdPipelineBarrier2-None-07889)VUID-vkCmdPipelineBarrier2-None-07889  
  If `vkCmdPipelineBarrier2` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, the render pass **must** have been created with at least one subpass dependency that expresses a dependency from the current subpass to itself, does not include `VK_DEPENDENCY_BY_REGION_BIT` if this command does not, does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT` if this command does not, and has [synchronization scopes](#synchronization-dependencies-scopes) and [access scopes](#synchronization-dependencies-access-scopes) that are all supersets of the scopes defined in this command
- [](#VUID-vkCmdPipelineBarrier2-bufferMemoryBarrierCount-01178)VUID-vkCmdPipelineBarrier2-bufferMemoryBarrierCount-01178  
  If `vkCmdPipelineBarrier2` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, it **must** not include any buffer memory barriers
- [](#VUID-vkCmdPipelineBarrier2-image-04073)VUID-vkCmdPipelineBarrier2-image-04073  
  If `vkCmdPipelineBarrier2` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, the `image` member of any image memory barrier included in this command **must** be an attachment used in the current subpass both as an input attachment, and as either a color, color resolve, or depth/stencil attachment
- [](#VUID-vkCmdPipelineBarrier2-image-09373)VUID-vkCmdPipelineBarrier2-image-09373  
  If `vkCmdPipelineBarrier2` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, and the `image` member of any image memory barrier is a color resolve attachment, the corresponding color attachment **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-vkCmdPipelineBarrier2-image-09374)VUID-vkCmdPipelineBarrier2-image-09374  
  If `vkCmdPipelineBarrier2` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, and the `image` member of any image memory barrier is a color resolve attachment, it **must** have been created with a non-zero [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` value
- [](#VUID-vkCmdPipelineBarrier2-oldLayout-01181)VUID-vkCmdPipelineBarrier2-oldLayout-01181  
  If `vkCmdPipelineBarrier2` is called within a render pass instance, the `oldLayout` and `newLayout` members of any image memory barrier included in this command **must** be equal
- [](#VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-01182)VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-01182  
  If `vkCmdPipelineBarrier2` is called within a render pass instance, the `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members of any memory barrier included in this command **must** be equal
- [](#VUID-vkCmdPipelineBarrier2-None-07890)VUID-vkCmdPipelineBarrier2-None-07890  
  If `vkCmdPipelineBarrier2` is called within a render pass instance, and the source stage masks of any memory barriers include [framebuffer-space stages](#synchronization-framebuffer-regions), destination stage masks of all memory barriers **must** only include [framebuffer-space stages](#synchronization-framebuffer-regions)
- [](#VUID-vkCmdPipelineBarrier2-dependencyFlags-07891)VUID-vkCmdPipelineBarrier2-dependencyFlags-07891  
  If `vkCmdPipelineBarrier2` is called within a render pass instance, and the source stage masks of any memory barriers include [framebuffer-space stages](#synchronization-framebuffer-regions), then `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`
- [](#VUID-vkCmdPipelineBarrier2-None-07892)VUID-vkCmdPipelineBarrier2-None-07892  
  If `vkCmdPipelineBarrier2` is called within a render pass instance, the source and destination stage masks of any memory barriers **must** only include graphics pipeline stages
- [](#VUID-vkCmdPipelineBarrier2-dependencyFlags-01186)VUID-vkCmdPipelineBarrier2-dependencyFlags-01186  
  If `vkCmdPipelineBarrier2` is called outside of a render pass instance, the dependency flags **must** not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-vkCmdPipelineBarrier2-None-07893)VUID-vkCmdPipelineBarrier2-None-07893  
  If `vkCmdPipelineBarrier2` is called inside a render pass instance, and there is more than one view in the current subpass, dependency flags **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-vkCmdPipelineBarrier2-None-09553)VUID-vkCmdPipelineBarrier2-None-09553  
  If none of the [`shaderTileImageColorReadAccess`](#features-shaderTileImageColorReadAccess), [`shaderTileImageStencilReadAccess`](#features-shaderTileImageStencilReadAccess), or [`shaderTileImageDepthReadAccess`](#features-shaderTileImageDepthReadAccess) features are enabled, and the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `vkCmdPipelineBarrier2` **must** not be called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdPipelineBarrier2-None-09554)VUID-vkCmdPipelineBarrier2-None-09554  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, and `vkCmdPipelineBarrier2` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), there **must** be no buffer or image memory barriers specified by this command
- [](#VUID-vkCmdPipelineBarrier2-None-09586)VUID-vkCmdPipelineBarrier2-None-09586  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, and `vkCmdPipelineBarrier2` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), memory barriers specified by this command **must** only include `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, or `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` in their access masks
- [](#VUID-vkCmdPipelineBarrier2-image-09555)VUID-vkCmdPipelineBarrier2-image-09555  
  If `vkCmdPipelineBarrier2` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), and the `image` member of any image memory barrier is used as an attachment in the current render pass instance, it **must** be in the `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` or `VK_IMAGE_LAYOUT_GENERAL` layout
- [](#VUID-vkCmdPipelineBarrier2-srcStageMask-09556)VUID-vkCmdPipelineBarrier2-srcStageMask-09556  
  If `vkCmdPipelineBarrier2` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), this command **must** only specify [framebuffer-space stages](#synchronization-framebuffer-regions) in `srcStageMask` and `dstStageMask`
- [](#VUID-vkCmdPipelineBarrier2-synchronization2-03848)VUID-vkCmdPipelineBarrier2-synchronization2-03848  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdPipelineBarrier2-srcStageMask-09673)VUID-vkCmdPipelineBarrier2-srcStageMask-09673  
  The `srcStageMask` member of any element of the `pMemoryBarriers` member of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdPipelineBarrier2-dstStageMask-09674)VUID-vkCmdPipelineBarrier2-dstStageMask-09674  
  The `dstStageMask` member of any element of the `pMemoryBarriers` member of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdPipelineBarrier2-srcStageMask-09675)VUID-vkCmdPipelineBarrier2-srcStageMask-09675  
  If a buffer or image memory barrier does not specify an [acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire), or if it does but `pDependencyInfo->dependencyFlags` includes `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the respective `srcStageMask` member of the element of the `pBufferMemoryBarriers` or `pImageMemoryBarriers` members of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdPipelineBarrier2-dstStageMask-09676)VUID-vkCmdPipelineBarrier2-dstStageMask-09676  
  If a buffer or image memory barrier does not specify an [release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release), or if it does but `pDependencyInfo->dependencyFlags` includes `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the respective `dstStageMask` member of the element of the `pBufferMemoryBarriers` or `pImageMemoryBarriers` members of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-10387)VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-10387  
  If a buffer or image memory barrier specifies a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers), either the `srcQueueFamilyIndex` or `dstQueueFamilyIndex` member of the element of the `pBufferMemoryBarriers` or `pImageMemoryBarriers` members of `pDependencyInfo` and the queue family index that was used to create the command pool that `commandBuffer` was allocated from **must** be equal

Valid Usage (Implicit)

- [](#VUID-vkCmdPipelineBarrier2-commandBuffer-parameter)VUID-vkCmdPipelineBarrier2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPipelineBarrier2-pDependencyInfo-parameter)VUID-vkCmdPipelineBarrier2-pDependencyInfo-parameter  
  `pDependencyInfo` **must** be a valid pointer to a valid [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structure
- [](#VUID-vkCmdPipelineBarrier2-commandBuffer-recording)VUID-vkCmdPipelineBarrier2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPipelineBarrier2-commandBuffer-cmdpool)VUID-vkCmdPipelineBarrier2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_TRANSFER\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT  
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Synchronization

Conditional Rendering

vkCmdPipelineBarrier2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPipelineBarrier2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0