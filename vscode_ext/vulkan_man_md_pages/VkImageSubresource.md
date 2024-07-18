# VkImageSubresource(3) Manual Page

## Name

VkImageSubresource - Structure specifying an image subresource



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageSubresource` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageSubresource {
    VkImageAspectFlags    aspectMask;
    uint32_t              mipLevel;
    uint32_t              arrayLayer;
} VkImageSubresource;
```

## <a href="#_members" class="anchor"></a>Members

- `aspectMask` is a [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html) value
  selecting the image *aspect*.

- `mipLevel` selects the mipmap level.

- `arrayLayer` selects the array layer.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageSubresource-aspectMask-parameter"
  id="VUID-VkImageSubresource-aspectMask-parameter"></a>
  VUID-VkImageSubresource-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) values

- <a href="#VUID-VkImageSubresource-aspectMask-requiredbitmask"
  id="VUID-VkImageSubresource-aspectMask-requiredbitmask"></a>
  VUID-VkImageSubresource-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html),
[VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html),
[vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageSubresource"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
