# vkCmdBeginTransformFeedbackEXT(3) Manual Page

## Name

vkCmdBeginTransformFeedbackEXT - Make transform feedback active in the command buffer



## [](#_c_specification)C Specification

Transform feedback for specific transform feedback buffers is made active by calling:

```c++
// Provided by VK_EXT_transform_feedback
void vkCmdBeginTransformFeedbackEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstCounterBuffer,
    uint32_t                                    counterBufferCount,
    const VkBuffer*                             pCounterBuffers,
    const VkDeviceSize*                         pCounterBufferOffsets);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `firstCounterBuffer` is the index of the first transform feedback buffer corresponding to `pCounterBuffers`\[0] and `pCounterBufferOffsets`\[0].
- `counterBufferCount` is the size of the `pCounterBuffers` and `pCounterBufferOffsets` arrays.
- `pCounterBuffers` is `NULL` or a pointer to an array of [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handles to counter buffers. Each buffer contains a 4 byte integer value representing the byte offset from the start of the corresponding transform feedback buffer from where to start capturing vertex data. If the byte offset stored to the counter buffer location was done using [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndTransformFeedbackEXT.html) it can be used to resume transform feedback from the previous location. In that case, a pipeline barrier is required between the calls to `vkCmdEndTransformFeedbackEXT` and `vkCmdBeginTransformFeedbackEXT`, with `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` as the source and destination stages, `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT` as the source access and `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT` as the destination access. If `pCounterBuffers` is `NULL`, then transform feedback will start capturing vertex data to byte offset zero in all bound transform feedback buffers. For each element of `pCounterBuffers` that is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), transform feedback will start capturing vertex data to byte zero in the corresponding bound transform feedback buffer.
- `pCounterBufferOffsets` is `NULL` or a pointer to an array of [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values specifying offsets within each of the `pCounterBuffers` where the counter values were previously written. The location in each counter buffer at these offsets **must** be large enough to contain 4 bytes of data. This data is the number of bytes captured by the previous transform feedback to this buffer. If `pCounterBufferOffsets` is `NULL`, then it is assumed the offsets are zero.

## [](#_description)Description

The active transform feedback buffers will capture primitives emitted from the corresponding `XfbBuffer` in the bound graphics pipeline. Any `XfbBuffer` emitted that does not output to an active transform feedback buffer will not be captured.

Valid Usage

- [](#VUID-vkCmdBeginTransformFeedbackEXT-transformFeedback-02366)VUID-vkCmdBeginTransformFeedbackEXT-transformFeedback-02366  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback` **must** be enabled
- [](#VUID-vkCmdBeginTransformFeedbackEXT-None-02367)VUID-vkCmdBeginTransformFeedbackEXT-None-02367  
  Transform feedback **must** not be active
- [](#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02368)VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02368  
  `firstCounterBuffer` **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02369)VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02369  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be less than or equal to `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdBeginTransformFeedbackEXT-counterBufferCount-02607)VUID-vkCmdBeginTransformFeedbackEXT-counterBufferCount-02607  
  If `counterBufferCount` is not `0`, and `pCounterBuffers` is not `NULL`, `pCounterBuffers` **must** be a valid pointer to an array of `counterBufferCount` `VkBuffer` handles that are either valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-02370)VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-02370  
  For each buffer handle in the array, if it is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) it **must** reference a buffer large enough to hold 4 bytes at the corresponding offset from the `pCounterBufferOffsets` array
- [](#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffer-02371)VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffer-02371  
  If `pCounterBuffer` is `NULL`, then `pCounterBufferOffsets` **must** also be `NULL`
- [](#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffers-02372)VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffers-02372  
  For each buffer handle in the `pCounterBuffers` array that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) it **must** have been created with a `usage` value containing `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`
- [](#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-09630)VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-09630  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be less than or equal to the number of transform feedback buffers bound by [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTransformFeedbackBuffersEXT.html)
- [](#VUID-vkCmdBeginTransformFeedbackEXT-None-06233)VUID-vkCmdBeginTransformFeedbackEXT-None-06233  
  If the [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature is not enabled, a valid graphics pipeline **must** be bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
- [](#VUID-vkCmdBeginTransformFeedbackEXT-None-04128)VUID-vkCmdBeginTransformFeedbackEXT-None-04128  
  The last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) of the bound graphics pipeline **must** have been declared with the `Xfb` execution mode
- [](#VUID-vkCmdBeginTransformFeedbackEXT-None-02373)VUID-vkCmdBeginTransformFeedbackEXT-None-02373  
  Transform feedback **must** not be made active in a render pass instance with multiview enabled
- [](#VUID-vkCmdBeginTransformFeedbackEXT-None-10656)VUID-vkCmdBeginTransformFeedbackEXT-None-10656  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-parameter)VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-parameter)VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-parameter  
  If `counterBufferCount` is not `0`, and `pCounterBufferOffsets` is not `NULL`, `pCounterBufferOffsets` **must** be a valid pointer to an array of `counterBufferCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-recording)VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-cmdpool)VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdBeginTransformFeedbackEXT-renderpass)VUID-vkCmdBeginTransformFeedbackEXT-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdBeginTransformFeedbackEXT-videocoding)VUID-vkCmdBeginTransformFeedbackEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBeginTransformFeedbackEXT-commonparent)VUID-vkCmdBeginTransformFeedbackEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pCounterBuffers` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdBeginTransformFeedbackEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginTransformFeedbackEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0