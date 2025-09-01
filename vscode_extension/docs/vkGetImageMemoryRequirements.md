# vkGetImageMemoryRequirements(3) Manual Page

## Name

vkGetImageMemoryRequirements - Returns the memory requirements for specified Vulkan object



## [](#_c_specification)C Specification

To determine the memory requirements for an image resource which is not created with the `VK_IMAGE_CREATE_DISJOINT_BIT` flag set, call:

```c++
// Provided by VK_VERSION_1_0
void vkGetImageMemoryRequirements(
    VkDevice                                    device,
    VkImage                                     image,
    VkMemoryRequirements*                       pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `image` is the image to query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure in which the memory requirements of the image object are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetImageMemoryRequirements-image-01588)VUID-vkGetImageMemoryRequirements-image-01588  
  `image` **must** not have been created with the `VK_IMAGE_CREATE_DISJOINT_BIT` flag set
- [](#VUID-vkGetImageMemoryRequirements-image-04004)VUID-vkGetImageMemoryRequirements-image-04004  
  If `image` was created with the `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` external memory handle type, then `image` **must** be bound to memory
- [](#VUID-vkGetImageMemoryRequirements-image-08960)VUID-vkGetImageMemoryRequirements-image-08960  
  If `image` was created with the `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX` external memory handle type, then `image` **must** be bound to memory

Valid Usage (Implicit)

- [](#VUID-vkGetImageMemoryRequirements-device-parameter)VUID-vkGetImageMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageMemoryRequirements-image-parameter)VUID-vkGetImageMemoryRequirements-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkGetImageMemoryRequirements-pMemoryRequirements-parameter)VUID-vkGetImageMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure
- [](#VUID-vkGetImageMemoryRequirements-image-parent)VUID-vkGetImageMemoryRequirements-image-parent  
  `image` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageMemoryRequirements).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0