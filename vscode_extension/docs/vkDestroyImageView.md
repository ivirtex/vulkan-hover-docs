# vkDestroyImageView(3) Manual Page

## Name

vkDestroyImageView - Destroy an image view object



## [](#_c_specification)C Specification

To destroy an image view, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyImageView(
    VkDevice                                    device,
    VkImageView                                 imageView,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the image view.
- `imageView` is the image view to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyImageView-imageView-01026)VUID-vkDestroyImageView-imageView-01026  
  All submitted commands that refer to `imageView` **must** have completed execution
- [](#VUID-vkDestroyImageView-imageView-01027)VUID-vkDestroyImageView-imageView-01027  
  If `VkAllocationCallbacks` were provided when `imageView` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyImageView-imageView-01028)VUID-vkDestroyImageView-imageView-01028  
  If no `VkAllocationCallbacks` were provided when `imageView` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyImageView-device-parameter)VUID-vkDestroyImageView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyImageView-imageView-parameter)VUID-vkDestroyImageView-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkDestroyImageView-pAllocator-parameter)VUID-vkDestroyImageView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyImageView-imageView-parent)VUID-vkDestroyImageView-imageView-parent  
  If `imageView` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `imageView` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyImageView).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0