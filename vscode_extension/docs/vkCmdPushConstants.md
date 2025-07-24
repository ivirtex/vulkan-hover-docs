# vkCmdPushConstants(3) Manual Page

## Name

vkCmdPushConstants - Update the values of push constants



## [](#_c_specification)C Specification

To update push constants, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdPushConstants(
    VkCommandBuffer                             commandBuffer,
    VkPipelineLayout                            layout,
    VkShaderStageFlags                          stageFlags,
    uint32_t                                    offset,
    uint32_t                                    size,
    const void*                                 pValues);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which the push constant update will be recorded.
- `layout` is the pipeline layout used to program the push constant updates.
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages that will use the push constants in the updated range.
- `offset` is the start offset of the push constant range to update, in units of bytes.
- `size` is the size of the push constant range to update, in units of bytes.
- `pValues` is a pointer to an array of `size` bytes containing the new push constant values.

## [](#_description)Description

When a command buffer begins recording, all push constant values are undefined. Reads of undefined push constant values by the executing shader return undefined values.

Push constant values **can** be updated incrementally, causing shader stages in `stageFlags` to read the new data from `pValues` for push constants modified by this command, while still reading the previous data for push constants not modified by this command. When a [bound pipeline command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-bindpoint-commands) is issued, the bound pipeline’s layout **must** be compatible with the layouts used to set the values of all push constants in the pipeline layout’s push constant ranges, as described in [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility). Binding a pipeline with a layout that is not compatible with the push constant layout does not disturb the push constant values.

Note

As `stageFlags` needs to include all flags the relevant push constant ranges were created with, any flags that are not supported by the queue family that the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) used to allocate `commandBuffer` was created on are ignored.

Valid Usage

- [](#VUID-vkCmdPushConstants-offset-01795)VUID-vkCmdPushConstants-offset-01795  
  For each byte in the range specified by `offset` and `size` and for each shader stage in `stageFlags`, there **must** be a push constant range in `layout` that includes that byte and that stage
- [](#VUID-vkCmdPushConstants-offset-01796)VUID-vkCmdPushConstants-offset-01796  
  For each byte in the range specified by `offset` and `size` and for each push constant range that overlaps that byte, `stageFlags` **must** include all stages in that push constant range’s [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html)::`stageFlags`
- [](#VUID-vkCmdPushConstants-offset-00368)VUID-vkCmdPushConstants-offset-00368  
  `offset` **must** be a multiple of `4`
- [](#VUID-vkCmdPushConstants-size-00369)VUID-vkCmdPushConstants-size-00369  
  `size` **must** be a multiple of `4`
- [](#VUID-vkCmdPushConstants-offset-00370)VUID-vkCmdPushConstants-offset-00370  
  `offset` **must** be less than `VkPhysicalDeviceLimits`::`maxPushConstantsSize`
- [](#VUID-vkCmdPushConstants-size-00371)VUID-vkCmdPushConstants-size-00371  
  `size` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

Valid Usage (Implicit)

- [](#VUID-vkCmdPushConstants-commandBuffer-parameter)VUID-vkCmdPushConstants-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushConstants-layout-parameter)VUID-vkCmdPushConstants-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkCmdPushConstants-stageFlags-parameter)VUID-vkCmdPushConstants-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-vkCmdPushConstants-stageFlags-requiredbitmask)VUID-vkCmdPushConstants-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-vkCmdPushConstants-pValues-parameter)VUID-vkCmdPushConstants-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `size` bytes
- [](#VUID-vkCmdPushConstants-commandBuffer-recording)VUID-vkCmdPushConstants-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushConstants-commandBuffer-cmdpool)VUID-vkCmdPushConstants-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushConstants-videocoding)VUID-vkCmdPushConstants-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdPushConstants-size-arraylength)VUID-vkCmdPushConstants-size-arraylength  
  `size` **must** be greater than `0`
- [](#VUID-vkCmdPushConstants-commonparent)VUID-vkCmdPushConstants-commonparent  
  Both of `commandBuffer`, and `layout` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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
Compute

State

Conditional Rendering

vkCmdPushConstants is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushConstants)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0