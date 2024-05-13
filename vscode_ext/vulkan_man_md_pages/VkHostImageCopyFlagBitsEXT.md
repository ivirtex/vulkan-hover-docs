# VkHostImageCopyFlagBitsEXT(3) Manual Page

## Name

VkHostImageCopyFlagBitsEXT - Bitmask specifying additional copy
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)::`flags`,
[VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html)::`flags`,
and
[VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html)::`flags`,
specifying additional copy parameters are:

``` c
// Provided by VK_EXT_host_image_copy
typedef enum VkHostImageCopyFlagBitsEXT {
    VK_HOST_IMAGE_COPY_MEMCPY_EXT = 0x00000001,
} VkHostImageCopyFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_HOST_IMAGE_COPY_MEMCPY_EXT` specifies that no memory layout
  swizzling is to be applied during data copy. For copies between memory
  and images, this flag indicates that image data in host memory is
  swizzled in exactly the same way as the image data on the device.
  Using this flag indicates that the implementations **may** use a
  simple memory copy to transfer the data between the host memory and
  the device memory. The format of the swizzled data in host memory is
  platform dependent and is not defined in this specification.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkHostImageCopyFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
