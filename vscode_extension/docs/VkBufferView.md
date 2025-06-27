# VkBufferView(3) Manual Page

## Name

VkBufferView - Opaque handle to a buffer view object



## <a href="#_c_specification" class="anchor"></a>C Specification

A *buffer view* represents a contiguous range of a buffer and a specific
format to be used to interpret the data. Buffer views are used to enable
shaders to access buffer contents using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures"
target="_blank" rel="noopener">image operations</a>. In order to create
a valid buffer view, the buffer **must** have been created with at least
one of the following usage flags:

- `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`

- `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`

Buffer views are represented by `VkBufferView` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBufferView)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html),
[vkCreateBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html),
[vkDestroyBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferView"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
