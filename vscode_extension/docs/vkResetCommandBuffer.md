# vkResetCommandBuffer(3) Manual Page

## Name

vkResetCommandBuffer - Reset a command buffer to the initial state



## <a href="#_c_specification" class="anchor"></a>C Specification

To reset a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkResetCommandBuffer(
    VkCommandBuffer                             commandBuffer,
    VkCommandBufferResetFlags                   flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer to reset. The command buffer
  **can** be in any state other than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending</a>, and is moved into the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">initial state</a>.

- `flags` is a bitmask of
  [VkCommandBufferResetFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlagBits.html)
  controlling the reset operation.

## <a href="#_description" class="anchor"></a>Description

Any primary command buffer that is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a> and has
`commandBuffer` recorded into it, becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

Valid Usage

- <a href="#VUID-vkResetCommandBuffer-commandBuffer-00045"
  id="VUID-vkResetCommandBuffer-commandBuffer-00045"></a>
  VUID-vkResetCommandBuffer-commandBuffer-00045  
  `commandBuffer` **must** not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkResetCommandBuffer-commandBuffer-00046"
  id="VUID-vkResetCommandBuffer-commandBuffer-00046"></a>
  VUID-vkResetCommandBuffer-commandBuffer-00046  
  `commandBuffer` **must** have been allocated from a pool that was
  created with the `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkResetCommandBuffer-commandBuffer-parameter"
  id="VUID-vkResetCommandBuffer-commandBuffer-parameter"></a>
  VUID-vkResetCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkResetCommandBuffer-flags-parameter"
  id="VUID-vkResetCommandBuffer-flags-parameter"></a>
  VUID-vkResetCommandBuffer-flags-parameter  
  `flags` **must** be a valid combination of
  [VkCommandBufferResetFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlagBits.html)
  values

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCommandBufferResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetCommandBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
