# vkResetCommandBuffer(3) Manual Page

## Name

vkResetCommandBuffer - Reset a command buffer to the initial state



## [](#_c_specification)C Specification

To reset a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkResetCommandBuffer(
    VkCommandBuffer                             commandBuffer,
    VkCommandBufferResetFlags                   flags);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer to reset. The command buffer **can** be in any state other than [pending](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), and is moved into the [initial state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).
- `flags` is a bitmask of [VkCommandBufferResetFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferResetFlagBits.html) controlling the reset operation.

## [](#_description)Description

Any primary command buffer that is in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) and has `commandBuffer` recorded into it, becomes [invalid](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

After a command buffer is reset, any objects or memory specified by commands recorded into the command buffer **must** no longer be accessed when the command buffer is accessed by the implementation.

Valid Usage

- [](#VUID-vkResetCommandBuffer-commandBuffer-00045)VUID-vkResetCommandBuffer-commandBuffer-00045  
  `commandBuffer` **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkResetCommandBuffer-commandBuffer-00046)VUID-vkResetCommandBuffer-commandBuffer-00046  
  `commandBuffer` **must** have been allocated from a pool that was created with the `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT`

Valid Usage (Implicit)

- [](#VUID-vkResetCommandBuffer-commandBuffer-parameter)VUID-vkResetCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkResetCommandBuffer-flags-parameter)VUID-vkResetCommandBuffer-flags-parameter  
  `flags` **must** be a valid combination of [VkCommandBufferResetFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferResetFlagBits.html) values

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCommandBufferResetFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferResetFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetCommandBuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0