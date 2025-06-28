# vkCreateBufferView(3) Manual Page

## Name

vkCreateBufferView - Create a new buffer view object



## [](#_c_specification)C Specification

To create a buffer view, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateBufferView(
    VkDevice                                    device,
    const VkBufferViewCreateInfo*               pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBufferView*                               pView);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the buffer view.
- `pCreateInfo` is a pointer to a [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html) structure containing parameters to be used to create the buffer view.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pView` is a pointer to a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) handle in which the resulting buffer view object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateBufferView-device-09665)VUID-vkCreateBufferView-device-09665  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_COMPUTE_BIT` or `VK_QUEUE_GRAPHICS_BIT` capabilities

Valid Usage (Implicit)

- [](#VUID-vkCreateBufferView-device-parameter)VUID-vkCreateBufferView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateBufferView-pCreateInfo-parameter)VUID-vkCreateBufferView-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html) structure
- [](#VUID-vkCreateBufferView-pAllocator-parameter)VUID-vkCreateBufferView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateBufferView-pView-parameter)VUID-vkCreateBufferView-pView-parameter  
  `pView` **must** be a valid pointer to a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) handle
- [](#VUID-vkCreateBufferView-device-queuecount)VUID-vkCreateBufferView-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html), [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateBufferView)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0