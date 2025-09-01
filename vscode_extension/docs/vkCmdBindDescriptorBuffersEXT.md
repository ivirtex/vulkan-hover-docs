# vkCmdBindDescriptorBuffersEXT(3) Manual Page

## Name

vkCmdBindDescriptorBuffersEXT - Binding descriptor buffers to a command buffer



## [](#_c_specification)C Specification

To bind descriptor buffers to a command buffer, call:

```c++
// Provided by VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBuffersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    bufferCount,
    const VkDescriptorBufferBindingInfoEXT*     pBindingInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptor buffers will be bound to.
- `bufferCount` is the number of elements in the `pBindingInfos` array.
- `pBindingInfos` is a pointer to an array of [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html) structures.

## [](#_description)Description

`vkCmdBindDescriptorBuffersEXT` causes any offsets previously set by [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html) that use the bindings numbered \[`0`.. `bufferCount`-1] to be no longer valid for subsequent bound pipeline commands. Any previously bound buffers at binding points greater than or equal to `bufferCount` are unbound.

Valid Usage

- [](#VUID-vkCmdBindDescriptorBuffersEXT-None-08047)VUID-vkCmdBindDescriptorBuffersEXT-None-08047  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkCmdBindDescriptorBuffersEXT-maxSamplerDescriptorBufferBindings-08048)VUID-vkCmdBindDescriptorBuffersEXT-maxSamplerDescriptorBufferBindings-08048  
  There **must** be no more than [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferBindings` elements in `pBindingInfos` with [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html)::`usage` containing `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdBindDescriptorBuffersEXT-maxResourceDescriptorBufferBindings-08049)VUID-vkCmdBindDescriptorBuffersEXT-maxResourceDescriptorBufferBindings-08049  
  There **must** be no more than [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferBindings` elements in `pBindingInfos` with [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html)::`usage` containing `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdBindDescriptorBuffersEXT-None-08050)VUID-vkCmdBindDescriptorBuffersEXT-None-08050  
  There **must** be no more than `1` element in `pBindingInfos` with [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html)::`usage` containing `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-08051)VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-08051  
  `bufferCount` **must** be less than or equal to [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxDescriptorBufferBindings`
- [](#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08053)VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08053  
  For any element of `pBindingInfos`, the buffer from which `address` was queried **must** have been created with the `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` bit set if it contains sampler descriptor data
- [](#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08054)VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08054  
  For any element of `pBindingInfos`, the buffer from which `address` was queried **must** have been created with the `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` bit set if it contains resource descriptor data
- [](#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08055)VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08055  
  For any element of `pBindingInfos`, `usage` **must** match the buffer from which `address` was queried
- [](#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-09947)VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-09947  
  For all elements of `pBindingInfos`, the buffer from which `address` was queried **must** have been created with the `VK_BUFFER_USAGE_2_DATA_GRAPH_FOREIGN_DESCRIPTOR_BIT_ARM` bit set if the command pool from which `commandBuffer` was allocated from was created with any element of [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html)::pProcessingEngines with `isForeign` set to `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-parameter)VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-parameter)VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-parameter  
  `pBindingInfos` **must** be a valid pointer to an array of `bufferCount` valid [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html) structures
- [](#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-recording)VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-cmdpool)VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or data\_graph operations
- [](#VUID-vkCmdBindDescriptorBuffersEXT-videocoding)VUID-vkCmdBindDescriptorBuffersEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-arraylength)VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-arraylength  
  `bufferCount` **must** be greater than `0`

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
Data\_Graph

State

Conditional Rendering

vkCmdBindDescriptorBuffersEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindDescriptorBuffersEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0