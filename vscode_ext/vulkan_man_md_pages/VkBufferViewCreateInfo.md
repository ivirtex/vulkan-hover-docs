# VkBufferViewCreateInfo(3) Manual Page

## Name

VkBufferViewCreateInfo - Structure specifying parameters of a newly
created buffer view



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferViewCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `buffer` is a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) on which the view will be
  created.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) describing the format of the
  data elements in the buffer.

- `offset` is an offset in bytes from the base address of the buffer.
  Accesses to the buffer view from shaders use addressing that is
  relative to this starting offset.

- `range` is a size in bytes of the buffer view. If `range` is equal to
  `VK_WHOLE_SIZE`, the range from `offset` to the end of the buffer is
  used. If `VK_WHOLE_SIZE` is used and the remaining size of the buffer
  is not a multiple of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#texel-block-size"
  target="_blank" rel="noopener">texel block size</a> of `format`, the
  nearest smaller multiple is used.

## <a href="#_description" class="anchor"></a>Description

The buffer view has a *buffer view usage* identifying which descriptor
types can be created from it. This usage **can** be defined by including
the
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
structure in the `pNext` chain, and specifying the `usage` value there.
If this structure is not included, it is equal to the
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`usage` value used to
create `buffer`.

Valid Usage

- <a href="#VUID-VkBufferViewCreateInfo-offset-00925"
  id="VUID-VkBufferViewCreateInfo-offset-00925"></a>
  VUID-VkBufferViewCreateInfo-offset-00925  
  `offset` **must** be less than the size of `buffer`

- <a href="#VUID-VkBufferViewCreateInfo-range-00928"
  id="VUID-VkBufferViewCreateInfo-range-00928"></a>
  VUID-VkBufferViewCreateInfo-range-00928  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be
  greater than `0`

- <a href="#VUID-VkBufferViewCreateInfo-range-00929"
  id="VUID-VkBufferViewCreateInfo-range-00929"></a>
  VUID-VkBufferViewCreateInfo-range-00929  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be an
  integer multiple of the texel block size of `format`

- <a href="#VUID-VkBufferViewCreateInfo-range-00930"
  id="VUID-VkBufferViewCreateInfo-range-00930"></a>
  VUID-VkBufferViewCreateInfo-range-00930  
  If `range` is not equal to `VK_WHOLE_SIZE`, the number of texel buffer
  elements given by (⌊`range` / (texel block size)⌋ × (texels per
  block)) where texel block size and texels per block are as defined in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">Compatible Formats</a> table for
  `format`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxTexelBufferElements`

- <a href="#VUID-VkBufferViewCreateInfo-offset-00931"
  id="VUID-VkBufferViewCreateInfo-offset-00931"></a>
  VUID-VkBufferViewCreateInfo-offset-00931  
  If `range` is not equal to `VK_WHOLE_SIZE`, the sum of `offset` and
  `range` **must** be less than or equal to the size of `buffer`

- <a href="#VUID-VkBufferViewCreateInfo-range-04059"
  id="VUID-VkBufferViewCreateInfo-range-04059"></a>
  VUID-VkBufferViewCreateInfo-range-04059  
  If `range` is equal to `VK_WHOLE_SIZE`, the number of texel buffer
  elements given by (⌊(size - `offset`) / (texel block size)⌋ × (texels
  per block)) where size is the size of `buffer`, and texel block size
  and texels per block are as defined in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">Compatible Formats</a> table for
  `format`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxTexelBufferElements`

- <a href="#VUID-VkBufferViewCreateInfo-buffer-00932"
  id="VUID-VkBufferViewCreateInfo-buffer-00932"></a>
  VUID-VkBufferViewCreateInfo-buffer-00932  
  `buffer` **must** have been created with a `usage` value containing at
  least one of `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT` or
  `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`

- <a href="#VUID-VkBufferViewCreateInfo-format-08778"
  id="VUID-VkBufferViewCreateInfo-format-08778"></a>
  VUID-VkBufferViewCreateInfo-format-08778  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-views-usage"
  target="_blank" rel="noopener">buffer view usage</a> contains
  `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-view-format-features"
  target="_blank" rel="noopener">format features</a> of `format`
  **must** contain `VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT`

- <a href="#VUID-VkBufferViewCreateInfo-format-08779"
  id="VUID-VkBufferViewCreateInfo-format-08779"></a>
  VUID-VkBufferViewCreateInfo-format-08779  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-views-usage"
  target="_blank" rel="noopener">buffer view usage</a> contains
  `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-view-format-features"
  target="_blank" rel="noopener">format features</a> of `format`
  **must** contain `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT`

- <a href="#VUID-VkBufferViewCreateInfo-buffer-00935"
  id="VUID-VkBufferViewCreateInfo-buffer-00935"></a>
  VUID-VkBufferViewCreateInfo-buffer-00935  
  If `buffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkBufferViewCreateInfo-offset-02749"
  id="VUID-VkBufferViewCreateInfo-offset-02749"></a>
  VUID-VkBufferViewCreateInfo-offset-02749  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-texelBufferAlignment"
  target="_blank" rel="noopener"><code>texelBufferAlignment</code></a>
  feature is not enabled, `offset` **must** be a multiple of
  `VkPhysicalDeviceLimits`::`minTexelBufferOffsetAlignment`

- <a href="#VUID-VkBufferViewCreateInfo-buffer-02750"
  id="VUID-VkBufferViewCreateInfo-buffer-02750"></a>
  VUID-VkBufferViewCreateInfo-buffer-02750  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-texelBufferAlignment"
  target="_blank" rel="noopener"><code>texelBufferAlignment</code></a>
  feature is enabled and if `buffer` was created with `usage` containing
  `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`, `offset` **must** be a
  multiple of the lesser of
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`storageTexelBufferOffsetAlignmentBytes`
  or, if
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`storageTexelBufferOffsetSingleTexelAlignment`
  is `VK_TRUE`, the size of a texel of the requested `format`. If the
  size of a texel is a multiple of three bytes, then the size of a
  single component of `format` is used instead

- <a href="#VUID-VkBufferViewCreateInfo-buffer-02751"
  id="VUID-VkBufferViewCreateInfo-buffer-02751"></a>
  VUID-VkBufferViewCreateInfo-buffer-02751  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-texelBufferAlignment"
  target="_blank" rel="noopener"><code>texelBufferAlignment</code></a>
  feature is enabled and if `buffer` was created with `usage` containing
  `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`, `offset` **must** be a
  multiple of the lesser of
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`uniformTexelBufferOffsetAlignmentBytes`
  or, if
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)::`uniformTexelBufferOffsetSingleTexelAlignment`
  is `VK_TRUE`, the size of a texel of the requested `format`. If the
  size of a texel is a multiple of three bytes, then the size of a
  single component of `format` is used instead

- <a href="#VUID-VkBufferViewCreateInfo-pNext-06782"
  id="VUID-VkBufferViewCreateInfo-pNext-06782"></a>
  VUID-VkBufferViewCreateInfo-pNext-06782  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT`

- <a href="#VUID-VkBufferViewCreateInfo-pNext-08780"
  id="VUID-VkBufferViewCreateInfo-pNext-08780"></a>
  VUID-VkBufferViewCreateInfo-pNext-08780  
  If the `pNext` chain includes a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html),
  its `usage` **must** not contain any other bit than
  `VK_BUFFER_USAGE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR` or
  `VK_BUFFER_USAGE_2_STORAGE_TEXEL_BUFFER_BIT_KHR`

- <a href="#VUID-VkBufferViewCreateInfo-pNext-08781"
  id="VUID-VkBufferViewCreateInfo-pNext-08781"></a>
  VUID-VkBufferViewCreateInfo-pNext-08781  
  If the `pNext` chain includes a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html),
  its `usage` **must** be a subset of the
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`usage` specified or
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)::`usage`
  from [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`pNext` when
  creating `buffer`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferViewCreateInfo-sType-sType"
  id="VUID-VkBufferViewCreateInfo-sType-sType"></a>
  VUID-VkBufferViewCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO`

- <a href="#VUID-VkBufferViewCreateInfo-pNext-pNext"
  id="VUID-VkBufferViewCreateInfo-pNext-pNext"></a>
  VUID-VkBufferViewCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  or
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- <a href="#VUID-VkBufferViewCreateInfo-sType-unique"
  id="VUID-VkBufferViewCreateInfo-sType-unique"></a>
  VUID-VkBufferViewCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- <a href="#VUID-VkBufferViewCreateInfo-flags-zerobitmask"
  id="VUID-VkBufferViewCreateInfo-flags-zerobitmask"></a>
  VUID-VkBufferViewCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkBufferViewCreateInfo-buffer-parameter"
  id="VUID-VkBufferViewCreateInfo-buffer-parameter"></a>
  VUID-VkBufferViewCreateInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkBufferViewCreateInfo-format-parameter"
  id="VUID-VkBufferViewCreateInfo-format-parameter"></a>
  VUID-VkBufferViewCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkBufferViewCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateFlags.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferViewCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
