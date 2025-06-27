# VkImageSubresourceLayers(3) Manual Page

## Name

VkImageSubresourceLayers - Structure specifying an image subresource
layers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageSubresourceLayers` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageSubresourceLayers {
    VkImageAspectFlags    aspectMask;
    uint32_t              mipLevel;
    uint32_t              baseArrayLayer;
    uint32_t              layerCount;
} VkImageSubresourceLayers;
```

## <a href="#_members" class="anchor"></a>Members

- `aspectMask` is a combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html), selecting the
  color, depth and/or stencil aspects to be copied.

- `mipLevel` is the mipmap level to copy

- `baseArrayLayer` and `layerCount` are the starting layer and number of
  layers to copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageSubresourceLayers-aspectMask-00167"
  id="VUID-VkImageSubresourceLayers-aspectMask-00167"></a>
  VUID-VkImageSubresourceLayers-aspectMask-00167  
  If `aspectMask` contains `VK_IMAGE_ASPECT_COLOR_BIT`, it **must** not
  contain either of `VK_IMAGE_ASPECT_DEPTH_BIT` or
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkImageSubresourceLayers-aspectMask-00168"
  id="VUID-VkImageSubresourceLayers-aspectMask-00168"></a>
  VUID-VkImageSubresourceLayers-aspectMask-00168  
  `aspectMask` **must** not contain `VK_IMAGE_ASPECT_METADATA_BIT`

- <a href="#VUID-VkImageSubresourceLayers-aspectMask-02247"
  id="VUID-VkImageSubresourceLayers-aspectMask-02247"></a>
  VUID-VkImageSubresourceLayers-aspectMask-02247  
  `aspectMask` **must** not include
  `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` for any index *i*

- <a href="#VUID-VkImageSubresourceLayers-layerCount-09243"
  id="VUID-VkImageSubresourceLayers-layerCount-09243"></a>
  VUID-VkImageSubresourceLayers-layerCount-09243  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> feature
  is not enabled, `layerCount` **must** not be
  `VK_REMAINING_ARRAY_LAYERS`

- <a href="#VUID-VkImageSubresourceLayers-layerCount-01700"
  id="VUID-VkImageSubresourceLayers-layerCount-01700"></a>
  VUID-VkImageSubresourceLayers-layerCount-01700  
  If `layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, it **must** be
  greater than 0

Valid Usage (Implicit)

- <a href="#VUID-VkImageSubresourceLayers-aspectMask-parameter"
  id="VUID-VkImageSubresourceLayers-aspectMask-parameter"></a>
  VUID-VkImageSubresourceLayers-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) values

- <a href="#VUID-VkImageSubresourceLayers-aspectMask-requiredbitmask"
  id="VUID-VkImageSubresourceLayers-aspectMask-requiredbitmask"></a>
  VUID-VkImageSubresourceLayers-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html),
[VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html),
[VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageIndirectCommandNV.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkImageBlit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html),
[VkImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html),
[VkImageResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve.html),
[VkImageResolve2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve2.html),
[VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html),
[VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html),
[vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToImageIndirectNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageSubresourceLayers"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
