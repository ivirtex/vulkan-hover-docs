# vkBeginCommandBuffer(3) Manual Page

## Name

vkBeginCommandBuffer - Start recording a command buffer



## [](#_c_specification)C Specification

To begin recording a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkBeginCommandBuffer(
    VkCommandBuffer                             commandBuffer,
    const VkCommandBufferBeginInfo*             pBeginInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the handle of the command buffer which is to be put in the recording state.
- `pBeginInfo` is a pointer to a [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html) structure defining additional information about how the command buffer begins recording.

## [](#_description)Description

Valid Usage

- [](#VUID-vkBeginCommandBuffer-commandBuffer-00049)VUID-vkBeginCommandBuffer-commandBuffer-00049  
  `commandBuffer` **must** not be in the [recording or pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkBeginCommandBuffer-commandBuffer-00050)VUID-vkBeginCommandBuffer-commandBuffer-00050  
  If `commandBuffer` was allocated from a [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which did not have the `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT` flag set, `commandBuffer` **must** be in the [initial state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkBeginCommandBuffer-commandBuffer-00051)VUID-vkBeginCommandBuffer-commandBuffer-00051  
  If `commandBuffer` is a secondary command buffer, the `pInheritanceInfo` member of `pBeginInfo` **must** be a valid `VkCommandBufferInheritanceInfo` structure
- [](#VUID-vkBeginCommandBuffer-commandBuffer-00052)VUID-vkBeginCommandBuffer-commandBuffer-00052  
  If `commandBuffer` is a secondary command buffer and either the `occlusionQueryEnable` member of the `pInheritanceInfo` member of `pBeginInfo` is `VK_FALSE`, or the [`occlusionQueryPrecise`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-occlusionQueryPrecise) feature is not enabled, then `pBeginInfo->pInheritanceInfo->queryFlags` **must** not contain `VK_QUERY_CONTROL_PRECISE_BIT`
- [](#VUID-vkBeginCommandBuffer-commandBuffer-02840)VUID-vkBeginCommandBuffer-commandBuffer-02840  
  If `commandBuffer` is a primary command buffer, then `pBeginInfo->flags` **must** not set both the `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` and the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flags

Valid Usage (Implicit)

- [](#VUID-vkBeginCommandBuffer-commandBuffer-parameter)VUID-vkBeginCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkBeginCommandBuffer-pBeginInfo-parameter)VUID-vkBeginCommandBuffer-pBeginInfo-parameter  
  `pBeginInfo` **must** be a valid pointer to a valid [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html) structure

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBeginCommandBuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0