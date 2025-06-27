# vkBeginCommandBuffer(3) Manual Page

## Name

vkBeginCommandBuffer - Start recording a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin recording a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkBeginCommandBuffer(
    VkCommandBuffer                             commandBuffer,
    const VkCommandBufferBeginInfo*             pBeginInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the handle of the command buffer which is to be put
  in the recording state.

- `pBeginInfo` is a pointer to a
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html) structure
  defining additional information about how the command buffer begins
  recording.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-00049"
  id="VUID-vkBeginCommandBuffer-commandBuffer-00049"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-00049  
  `commandBuffer` **must** not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">recording or pending state</a>

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-00050"
  id="VUID-vkBeginCommandBuffer-commandBuffer-00050"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-00050  
  If `commandBuffer` was allocated from a
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which did not have the
  `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT` flag set,
  `commandBuffer` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">initial state</a>

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-00051"
  id="VUID-vkBeginCommandBuffer-commandBuffer-00051"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-00051  
  If `commandBuffer` is a secondary command buffer, the
  `pInheritanceInfo` member of `pBeginInfo` **must** be a valid
  `VkCommandBufferInheritanceInfo` structure

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-00052"
  id="VUID-vkBeginCommandBuffer-commandBuffer-00052"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-00052  
  If `commandBuffer` is a secondary command buffer and either the
  `occlusionQueryEnable` member of the `pInheritanceInfo` member of
  `pBeginInfo` is `VK_FALSE`, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-occlusionQueryPrecise"
  target="_blank" rel="noopener"><code>occlusionQueryPrecise</code></a>
  feature is not enabled, then
  `pBeginInfo->pInheritanceInfo->queryFlags` **must** not contain
  `VK_QUERY_CONTROL_PRECISE_BIT`

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-02840"
  id="VUID-vkBeginCommandBuffer-commandBuffer-02840"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-02840  
  If `commandBuffer` is a primary command buffer, then
  `pBeginInfo->flags` **must** not set both the
  `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` and the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flags

Valid Usage (Implicit)

- <a href="#VUID-vkBeginCommandBuffer-commandBuffer-parameter"
  id="VUID-vkBeginCommandBuffer-commandBuffer-parameter"></a>
  VUID-vkBeginCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkBeginCommandBuffer-pBeginInfo-parameter"
  id="VUID-vkBeginCommandBuffer-pBeginInfo-parameter"></a>
  VUID-vkBeginCommandBuffer-pBeginInfo-parameter  
  `pBeginInfo` **must** be a valid pointer to a valid
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html) structure

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBeginCommandBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
