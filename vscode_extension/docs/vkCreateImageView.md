# vkCreateImageView(3) Manual Page

## Name

vkCreateImageView - Create an image view from an existing image



## [](#_c_specification)C Specification

To create an image view, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateImageView(
    VkDevice                                    device,
    const VkImageViewCreateInfo*                pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkImageView*                                pView);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the image view.
- `pCreateInfo` is a pointer to a [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html) structure containing parameters to be used to create the image view.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pView` is a pointer to a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle in which the resulting image view object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateImageView-device-09667)VUID-vkCreateImageView-device-09667  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`, `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities
- [](#VUID-vkCreateImageView-image-09179)VUID-vkCreateImageView-image-09179  
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`image` **must** have been created from `device`

Valid Usage (Implicit)

- [](#VUID-vkCreateImageView-device-parameter)VUID-vkCreateImageView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateImageView-pCreateInfo-parameter)VUID-vkCreateImageView-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html) structure
- [](#VUID-vkCreateImageView-pAllocator-parameter)VUID-vkCreateImageView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateImageView-pView-parameter)VUID-vkCreateImageView-pView-parameter  
  `pView` **must** be a valid pointer to a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkCreateImageView-device-queuecount)VUID-vkCreateImageView-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateImageView)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0