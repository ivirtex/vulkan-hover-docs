# vkCmdBindTransformFeedbackBuffersEXT(3) Manual Page

## Name

vkCmdBindTransformFeedbackBuffersEXT - Bind transform feedback buffers to a command buffer



## [](#_c_specification)C Specification

To bind transform feedback buffers to a command buffer for use in subsequent drawing commands, call:

```c++
// Provided by VK_EXT_transform_feedback
void vkCmdBindTransformFeedbackBuffersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `firstBinding` is the index of the first transform feedback binding whose state is updated by the command.
- `bindingCount` is the number of transform feedback bindings whose state is updated by the command.
- `pBuffers` is a pointer to an array of buffer handles.
- `pOffsets` is a pointer to an array of buffer offsets.
- `pSizes` is `NULL` or a pointer to an array of [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) buffer sizes, specifying the maximum number of bytes to capture to the corresponding transform feedback buffer. If `pSizes` is `NULL`, or the value of the `pSizes` array element is `VK_WHOLE_SIZE`, then the maximum number of bytes captured will be the size of the corresponding buffer minus the buffer offset.

## [](#_description)Description

The values taken from elements i of `pBuffers`, `pOffsets` and `pSizes` replace the current state for the transform feedback binding `firstBinding` + i, for i in \[0, `bindingCount`). The transform feedback binding is updated to start at the offset indicated by `pOffsets`\[i] from the start of the buffer `pBuffers`\[i].

Valid Usage

- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-transformFeedback-02355)VUID-vkCmdBindTransformFeedbackBuffersEXT-transformFeedback-02355  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback` **must** be enabled
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02356)VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02356  
  `firstBinding` **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02357)VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02357  
  The sum of `firstBinding` and `bindingCount` **must** be less than or equal to `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02358)VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02358  
  All elements of `pOffsets` **must** be less than the size of the corresponding element in `pBuffers`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02359)VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02359  
  All elements of `pOffsets` **must** be a multiple of 4
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02360)VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02360  
  All elements of `pBuffers` **must** have been created with the `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT` flag
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pSize-02361)VUID-vkCmdBindTransformFeedbackBuffersEXT-pSize-02361  
  If the optional `pSize` array is specified, each element of `pSizes` **must** either be `VK_WHOLE_SIZE`, or be less than or equal to `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBufferSize`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pSizes-02362)VUID-vkCmdBindTransformFeedbackBuffersEXT-pSizes-02362  
  All elements of `pSizes` **must** be either `VK_WHOLE_SIZE`, or less than or equal to the size of the corresponding buffer in `pBuffers`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02363)VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02363  
  All elements of `pOffsets` plus `pSizes`, where the `pSizes`, element is not `VK_WHOLE_SIZE`, **must** be less than or equal to the size of the corresponding buffer in `pBuffers`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02364)VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02364  
  Each element of `pBuffers` that is non-sparse **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-None-02365)VUID-vkCmdBindTransformFeedbackBuffersEXT-None-02365  
  Transform feedback **must** not be active when the `vkCmdBindTransformFeedbackBuffersEXT` command is recorded

Valid Usage (Implicit)

- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-parameter)VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-parameter)VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-parameter  
  `pBuffers` **must** be a valid pointer to an array of `bindingCount` valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handles
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-parameter)VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-recording)VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-cmdpool)VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-videocoding)VUID-vkCmdBindTransformFeedbackBuffersEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-bindingCount-arraylength)VUID-vkCmdBindTransformFeedbackBuffersEXT-bindingCount-arraylength  
  `bindingCount` **must** be greater than `0`
- [](#VUID-vkCmdBindTransformFeedbackBuffersEXT-commonparent)VUID-vkCmdBindTransformFeedbackBuffersEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pBuffers` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics

State

Conditional Rendering

vkCmdBindTransformFeedbackBuffersEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindTransformFeedbackBuffersEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0