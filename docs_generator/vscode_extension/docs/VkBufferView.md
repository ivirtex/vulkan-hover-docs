# VkBufferView(3) Manual Page

## Name

VkBufferView - Opaque handle to a buffer view object



## [](#_c_specification)C Specification

A *buffer view* represents a contiguous range of a buffer and a specific format to be used to interpret the data. Buffer views are used to enable shaders to access buffer contents using [image operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures). In order to create a valid buffer view, the buffer **must** have been created with at least one of the following usage flags:

- `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`
- `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`

Buffer views are represented by `VkBufferView` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBufferView)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html), [vkCreateBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferView.html), [vkDestroyBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyBufferView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferView)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0