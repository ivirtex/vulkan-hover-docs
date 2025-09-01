# VkBufferViewCreateInfo(3) Manual Page

## Name

VkBufferViewCreateInfo - Structure specifying parameters of a newly created buffer view



## [](#_c_specification)C Specification

The `VkBufferViewCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkBufferViewCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    VkBufferViewCreateFlags    flags;
    VkBuffer                   buffer;
    VkFormat                   format;
    VkDeviceSize               offset;
    VkDeviceSize               range;
} VkBufferViewCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `buffer` is a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) on which the view will be created.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) describing the format of the data elements in the buffer.
- `offset` is an offset in bytes from the base address of the buffer. Accesses to the buffer view from shaders use addressing that is relative to this starting offset.
- `range` is a size in bytes of the buffer view. If `range` is equal to `VK_WHOLE_SIZE`, the range from `offset` to the end of the buffer is used. If `VK_WHOLE_SIZE` is used and the remaining size of the buffer is not a multiple of the [texel block size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#texel-block-size) of `format`, the nearest smaller multiple is used.

## [](#_description)Description

The buffer view has a *buffer view usage* identifying which descriptor types can be created from it. This usage **can** be defined by including the [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure in the `pNext` chain, and specifying the `usage` value there. If this structure is not included, it is equal to the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`usage` value used to create `buffer`.

Valid Usage

- [](#VUID-VkBufferViewCreateInfo-offset-00925)VUID-VkBufferViewCreateInfo-offset-00925  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-VkBufferViewCreateInfo-range-00928)VUID-VkBufferViewCreateInfo-range-00928  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be greater than `0`
- [](#VUID-VkBufferViewCreateInfo-range-00929)VUID-VkBufferViewCreateInfo-range-00929  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be an integer multiple of the texel block size of `format`
- [](#VUID-VkBufferViewCreateInfo-range-00930)VUID-VkBufferViewCreateInfo-range-00930  
  If `range` is not equal to `VK_WHOLE_SIZE`, the number of texel buffer elements given by (⌊`range` / (texel block size)⌋ × (texels per block)) where texel block size and texels per block are as defined in the [Compatible Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility) table for `format`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxTexelBufferElements`
- [](#VUID-VkBufferViewCreateInfo-offset-00931)VUID-VkBufferViewCreateInfo-offset-00931  
  If `range` is not equal to `VK_WHOLE_SIZE`, the sum of `offset` and `range` **must** be less than or equal to the size of `buffer`
- [](#VUID-VkBufferViewCreateInfo-range-04059)VUID-VkBufferViewCreateInfo-range-04059  
  If `range` is equal to `VK_WHOLE_SIZE`, the number of texel buffer elements given by (⌊(size - `offset`) / (texel block size)⌋ × (texels per block)) where size is the size of `buffer`, and texel block size and texels per block are as defined in the [Compatible Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility) table for `format`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxTexelBufferElements`
- [](#VUID-VkBufferViewCreateInfo-buffer-00932)VUID-VkBufferViewCreateInfo-buffer-00932  
  `buffer` **must** have been created with a `usage` value containing at least one of `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT` or `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`
- [](#VUID-VkBufferViewCreateInfo-format-08778)VUID-VkBufferViewCreateInfo-format-08778  
  If the [buffer view usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-views-usage) contains `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`, then [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `format` **must** contain `VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT`
- [](#VUID-VkBufferViewCreateInfo-format-08779)VUID-VkBufferViewCreateInfo-format-08779  
  If the [buffer view usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-views-usage) contains `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`, then [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `format` **must** contain `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT`
- [](#VUID-VkBufferViewCreateInfo-buffer-00935)VUID-VkBufferViewCreateInfo-buffer-00935  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkBufferViewCreateInfo-offset-02749)VUID-VkBufferViewCreateInfo-offset-02749  
  If the [`texelBufferAlignment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-texelBufferAlignment) feature is not enabled, `offset` **must** be a multiple of `VkPhysicalDeviceLimits`::`minTexelBufferOffsetAlignment`
- [](#VUID-VkBufferViewCreateInfo-buffer-02750)VUID-VkBufferViewCreateInfo-buffer-02750  
  If the [`texelBufferAlignment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-texelBufferAlignment) feature is enabled and if `buffer` was created with `usage` containing `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`, `offset` **must** be a multiple of the lesser of [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`storageTexelBufferOffsetAlignmentBytes` or, if [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`storageTexelBufferOffsetSingleTexelAlignment` is `VK_TRUE`, the size of a texel of the requested `format`. If the size of a texel is a multiple of three bytes, then the size of a single component of `format` is used instead
- [](#VUID-VkBufferViewCreateInfo-buffer-02751)VUID-VkBufferViewCreateInfo-buffer-02751  
  If the [`texelBufferAlignment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-texelBufferAlignment) feature is enabled and if `buffer` was created with `usage` containing `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`, `offset` **must** be a multiple of the lesser of [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`uniformTexelBufferOffsetAlignmentBytes` or, if [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`uniformTexelBufferOffsetSingleTexelAlignment` is `VK_TRUE`, the size of a texel of the requested `format`. If the size of a texel is a multiple of three bytes, then the size of a single component of `format` is used instead
- [](#VUID-VkBufferViewCreateInfo-pNext-06782)VUID-VkBufferViewCreateInfo-pNext-06782  
  If the `pNext` chain includes a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure, its `exportObjectType` member **must** be `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT`
- [](#VUID-VkBufferViewCreateInfo-pNext-08780)VUID-VkBufferViewCreateInfo-pNext-08780  
  If the `pNext` chain includes a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html), its `usage` **must** not contain any other bit than `VK_BUFFER_USAGE_2_UNIFORM_TEXEL_BUFFER_BIT` or `VK_BUFFER_USAGE_2_STORAGE_TEXEL_BUFFER_BIT`
- [](#VUID-VkBufferViewCreateInfo-pNext-08781)VUID-VkBufferViewCreateInfo-pNext-08781  
  If the `pNext` chain includes a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html), its `usage` **must** be a subset of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`usage` specified or [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)::`usage` from [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`pNext` when creating `buffer`

Valid Usage (Implicit)

- [](#VUID-VkBufferViewCreateInfo-sType-sType)VUID-VkBufferViewCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO`
- [](#VUID-VkBufferViewCreateInfo-pNext-pNext)VUID-VkBufferViewCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) or [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- [](#VUID-VkBufferViewCreateInfo-sType-unique)VUID-VkBufferViewCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- [](#VUID-VkBufferViewCreateInfo-flags-zerobitmask)VUID-VkBufferViewCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkBufferViewCreateInfo-buffer-parameter)VUID-VkBufferViewCreateInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkBufferViewCreateInfo-format-parameter)VUID-VkBufferViewCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferViewCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateFlags.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferViewCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0