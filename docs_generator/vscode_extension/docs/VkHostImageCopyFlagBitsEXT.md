# VkHostImageCopyFlagBits(3) Manual Page

## Name

VkHostImageCopyFlagBits - Bitmask specifying additional copy parameters



## [](#_c_specification)C Specification

Bits which **can** be set in [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html)::`flags`, [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html)::`flags`, and [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html)::`flags`, specifying additional copy parameters are:

```c++
// Provided by VK_VERSION_1_4
typedef enum VkHostImageCopyFlagBits {
    VK_HOST_IMAGE_COPY_MEMCPY_BIT = 0x00000001,
  // VK_HOST_IMAGE_COPY_MEMCPY is a deprecated alias
    VK_HOST_IMAGE_COPY_MEMCPY = VK_HOST_IMAGE_COPY_MEMCPY_BIT,
  // Provided by VK_EXT_host_image_copy
    VK_HOST_IMAGE_COPY_MEMCPY_BIT_EXT = VK_HOST_IMAGE_COPY_MEMCPY_BIT,
  // Provided by VK_EXT_host_image_copy
  // VK_HOST_IMAGE_COPY_MEMCPY_EXT is a deprecated alias
    VK_HOST_IMAGE_COPY_MEMCPY_EXT = VK_HOST_IMAGE_COPY_MEMCPY_BIT,
} VkHostImageCopyFlagBits;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkHostImageCopyFlagBits VkHostImageCopyFlagBitsEXT;
```

## [](#_description)Description

- `VK_HOST_IMAGE_COPY_MEMCPY_BIT` specifies that no memory layout swizzling is to be applied during data copy. For copies between memory and images, this flag indicates that image data in host memory is swizzled in exactly the same way as the image data on the device. Using this flag indicates that the implementations **may** use a simple memory copy to transfer the data between the host memory and the device memory. The format of the swizzled data in host memory is platform dependent and is not defined in this specification.

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkHostImageCopyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkHostImageCopyFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0