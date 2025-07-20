# vkAllocateCommandBuffers(3) Manual Page

## Name

vkAllocateCommandBuffers - Allocate command buffers from an existing command pool



## [](#_c_specification)C Specification

To allocate command buffers, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkAllocateCommandBuffers(
    VkDevice                                    device,
    const VkCommandBufferAllocateInfo*          pAllocateInfo,
    VkCommandBuffer*                            pCommandBuffers);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the command pool.
- `pAllocateInfo` is a pointer to a [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html) structure describing parameters of the allocation. `commandPool` **may** be accessed any time one of the resulting command buffers is accessed.
- `pCommandBuffers` is a pointer to an array of [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handles in which the resulting command buffer objects are returned. The array **must** be at least the length specified by the `commandBufferCount` member of `pAllocateInfo`. Each allocated command buffer begins in the initial state.

## [](#_description)Description

`vkAllocateCommandBuffers` **can** be used to allocate multiple command buffers. If the allocation of any of those command buffers fails, the implementation **must** free all successfully allocated command buffer objects from this command, set all entries of the `pCommandBuffers` array to `NULL` and return the error.

Note

Filling `pCommandBuffers` with `NULL` values on failure is an exception to the default error behavior that output parameters will have undefined contents.

When command buffers are first allocated, they are in the [initial state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Valid Usage (Implicit)

- [](#VUID-vkAllocateCommandBuffers-device-parameter)VUID-vkAllocateCommandBuffers-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAllocateCommandBuffers-pAllocateInfo-parameter)VUID-vkAllocateCommandBuffers-pAllocateInfo-parameter  
  `pAllocateInfo` **must** be a valid pointer to a valid [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html) structure
- [](#VUID-vkAllocateCommandBuffers-pCommandBuffers-parameter)VUID-vkAllocateCommandBuffers-pCommandBuffers-parameter  
  `pCommandBuffers` **must** be a valid pointer to an array of `pAllocateInfo->commandBufferCount` [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handles
- [](#VUID-vkAllocateCommandBuffers-device-queuecount)VUID-vkAllocateCommandBuffers-device-queuecount  
  The device **must** have been created with at least `1` queue
- [](#VUID-vkAllocateCommandBuffers-pAllocateInfo::commandBufferCount-arraylength)VUID-vkAllocateCommandBuffers-pAllocateInfo::commandBufferCount-arraylength  
  `pAllocateInfo->commandBufferCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAllocateCommandBuffers)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0