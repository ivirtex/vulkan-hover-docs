# vkFreeCommandBuffers(3) Manual Page

## Name

vkFreeCommandBuffers - Free command buffers



## [](#_c_specification)C Specification

To free command buffers, call:

```c++
// Provided by VK_VERSION_1_0
void vkFreeCommandBuffers(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    uint32_t                                    commandBufferCount,
    const VkCommandBuffer*                      pCommandBuffers);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the command pool.
- `commandPool` is the command pool from which the command buffers were allocated.
- `commandBufferCount` is the length of the `pCommandBuffers` array.
- `pCommandBuffers` is a pointer to an array of handles of command buffers to free.

## [](#_description)Description

Any primary command buffer that is in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) and has any element of `pCommandBuffers` recorded into it, becomes [invalid](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Valid Usage

- [](#VUID-vkFreeCommandBuffers-pCommandBuffers-00047)VUID-vkFreeCommandBuffers-pCommandBuffers-00047  
  All elements of `pCommandBuffers` **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkFreeCommandBuffers-pCommandBuffers-00048)VUID-vkFreeCommandBuffers-pCommandBuffers-00048  
  `pCommandBuffers` **must** be a valid pointer to an array of `commandBufferCount` `VkCommandBuffer` handles, each element of which **must** either be a valid handle or `NULL`

Valid Usage (Implicit)

- [](#VUID-vkFreeCommandBuffers-device-parameter)VUID-vkFreeCommandBuffers-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkFreeCommandBuffers-commandPool-parameter)VUID-vkFreeCommandBuffers-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle
- [](#VUID-vkFreeCommandBuffers-commandBufferCount-arraylength)VUID-vkFreeCommandBuffers-commandBufferCount-arraylength  
  `commandBufferCount` **must** be greater than `0`
- [](#VUID-vkFreeCommandBuffers-commandPool-parent)VUID-vkFreeCommandBuffers-commandPool-parent  
  `commandPool` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkFreeCommandBuffers-pCommandBuffers-parent)VUID-vkFreeCommandBuffers-pCommandBuffers-parent  
  Each element of `pCommandBuffers` that is a valid handle **must** have been created, allocated, or retrieved from `commandPool`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized
- Host access to each member of `pCommandBuffers` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkFreeCommandBuffers).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0