# VkOffset3D(3) Manual Page

## Name

VkOffset3D - Structure specifying a three-dimensional offset



## <a href="#_c_specification" class="anchor"></a>C Specification

A three-dimensional offset is defined by the structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkOffset3D {
    int32_t    x;
    int32_t    y;
    int32_t    z;
} VkOffset3D;
```

## <a href="#_members" class="anchor"></a>Members

- `x` is the x offset.

- `y` is the y offset.

- `z` is the z offset.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html),
[VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html),
[VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageIndirectCommandNV.html),
[VkImageBlit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html),
[VkImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html),
[VkImageResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve.html),
[VkImageResolve2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve2.html),
[VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html),
[VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html),
[VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOffset3D"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
