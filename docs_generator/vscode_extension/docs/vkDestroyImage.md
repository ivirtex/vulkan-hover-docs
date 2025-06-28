# vkDestroyImage(3) Manual Page

## Name

vkDestroyImage - Destroy an image object



## [](#_c_specification)C Specification

To destroy an image, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyImage(
    VkDevice                                    device,
    VkImage                                     image,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the image.
- `image` is the image to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyImage-image-01000)VUID-vkDestroyImage-image-01000  
  All submitted commands that refer to `image`, either directly or via a `VkImageView`, **must** have completed execution
- [](#VUID-vkDestroyImage-image-01001)VUID-vkDestroyImage-image-01001  
  If `VkAllocationCallbacks` were provided when `image` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyImage-image-01002)VUID-vkDestroyImage-image-01002  
  If no `VkAllocationCallbacks` were provided when `image` was created, `pAllocator` **must** be `NULL`
- [](#VUID-vkDestroyImage-image-04882)VUID-vkDestroyImage-image-04882  
  `image` **must** not have been acquired from [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainImagesKHR.html)

Valid Usage (Implicit)

- [](#VUID-vkDestroyImage-device-parameter)VUID-vkDestroyImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyImage-image-parameter)VUID-vkDestroyImage-image-parameter  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkDestroyImage-pAllocator-parameter)VUID-vkDestroyImage-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyImage-image-parent)VUID-vkDestroyImage-image-parent  
  If `image` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `image` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyImage)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0