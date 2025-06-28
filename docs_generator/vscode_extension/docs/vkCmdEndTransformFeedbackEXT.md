# vkCmdEndTransformFeedbackEXT(3) Manual Page

## Name

vkCmdEndTransformFeedbackEXT - Make transform feedback inactive in the command buffer



## [](#_c_specification)C Specification

Transform feedback for specific transform feedback buffers is made inactive by calling:

```c++
// Provided by VK_EXT_transform_feedback
void vkCmdEndTransformFeedbackEXT(
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
- `pCounterBuffers` is `NULL` or a pointer to an array of [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handles to counter buffers. The counter buffers are used to record the current byte positions of each transform feedback buffer where the next vertex output data would be captured. This **can** be used by a subsequent [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginTransformFeedbackEXT.html) call to resume transform feedback capture from this position. It can also be used by [vkCmdDrawIndirectByteCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectByteCountEXT.html) to determine the vertex count of the draw call.
- `pCounterBufferOffsets` is `NULL` or a pointer to an array of [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values specifying offsets within each of the `pCounterBuffers` where the counter values can be written. The location in each counter buffer at these offsets **must** be large enough to contain 4 bytes of data. The data stored at this location is the byte offset from the start of the transform feedback buffer binding where the next vertex data would be written. If `pCounterBufferOffsets` is `NULL`, then it is assumed the offsets are zero.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdEndTransformFeedbackEXT-transformFeedback-02374)VUID-vkCmdEndTransformFeedbackEXT-transformFeedback-02374  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback` **must** be enabled
- [](#VUID-vkCmdEndTransformFeedbackEXT-None-02375)VUID-vkCmdEndTransformFeedbackEXT-None-02375  
  Transform feedback **must** be active
- [](#VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02376)VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02376  
  `firstCounterBuffer` **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02377)VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02377  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be less than or equal to `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdEndTransformFeedbackEXT-counterBufferCount-02608)VUID-vkCmdEndTransformFeedbackEXT-counterBufferCount-02608  
  If `counterBufferCount` is not `0`, and `pCounterBuffers` is not `NULL`, `pCounterBuffers` **must** be a valid pointer to an array of `counterBufferCount` `VkBuffer` handles that are either valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-02378)VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-02378  
  For each buffer handle in the array, if it is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) it **must** reference a buffer large enough to hold 4 bytes at the corresponding offset from the `pCounterBufferOffsets` array
- [](#VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffer-02379)VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffer-02379  
  If `pCounterBuffer` is `NULL`, then `pCounterBufferOffsets` **must** also be `NULL`
- [](#VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffers-02380)VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffers-02380  
  For each buffer handle in the `pCounterBuffers` array that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) it **must** have been created with a `usage` value containing `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`
- [](#VUID-vkCmdEndTransformFeedbackEXT-None-10657)VUID-vkCmdEndTransformFeedbackEXT-None-10657  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-parameter)VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-parameter)VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-parameter  
  If `counterBufferCount` is not `0`, and `pCounterBufferOffsets` is not `NULL`, `pCounterBufferOffsets` **must** be a valid pointer to an array of `counterBufferCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-recording)VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-cmdpool)VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdEndTransformFeedbackEXT-renderpass)VUID-vkCmdEndTransformFeedbackEXT-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndTransformFeedbackEXT-videocoding)VUID-vkCmdEndTransformFeedbackEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdEndTransformFeedbackEXT-commonparent)VUID-vkCmdEndTransformFeedbackEXT-commonparent  
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

Graphics

State

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndTransformFeedbackEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0