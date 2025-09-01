# vkResetCommandPool(3) Manual Page

## Name

vkResetCommandPool - Reset a command pool



## [](#_c_specification)C Specification

To reset a command pool, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkResetCommandPool(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    VkCommandPoolResetFlags                     flags);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the command pool.
- `commandPool` is the command pool to reset.
- `flags` is a bitmask of [VkCommandPoolResetFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolResetFlagBits.html) controlling the reset operation.

## [](#_description)Description

Resetting a command pool recycles all of the resources from all of the command buffers allocated from the command pool back to the command pool. All command buffers that have been allocated from the command pool are put in the [initial state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Any primary command buffer allocated from another [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that is in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) and has a secondary command buffer allocated from `commandPool` recorded into it, becomes [invalid](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Valid Usage

- [](#VUID-vkResetCommandPool-commandPool-00040)VUID-vkResetCommandPool-commandPool-00040  
  All `VkCommandBuffer` objects allocated from `commandPool` **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)

Valid Usage (Implicit)

- [](#VUID-vkResetCommandPool-device-parameter)VUID-vkResetCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkResetCommandPool-commandPool-parameter)VUID-vkResetCommandPool-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle
- [](#VUID-vkResetCommandPool-flags-parameter)VUID-vkResetCommandPool-flags-parameter  
  `flags` **must** be a valid combination of [VkCommandPoolResetFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolResetFlagBits.html) values
- [](#VUID-vkResetCommandPool-commandPool-parent)VUID-vkResetCommandPool-commandPool-parent  
  `commandPool` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkCommandPoolResetFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolResetFlags.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetCommandPool).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0