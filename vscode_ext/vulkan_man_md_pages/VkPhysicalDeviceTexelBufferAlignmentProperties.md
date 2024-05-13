# VkPhysicalDeviceTexelBufferAlignmentProperties(3) Manual Page

## Name

VkPhysicalDeviceTexelBufferAlignmentProperties - Structure describing
the texel buffer alignment requirements supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceTexelBufferAlignmentProperties` structure is
defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceTexelBufferAlignmentProperties {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       storageTexelBufferOffsetAlignmentBytes;
    VkBool32           storageTexelBufferOffsetSingleTexelAlignment;
    VkDeviceSize       uniformTexelBufferOffsetAlignmentBytes;
    VkBool32           uniformTexelBufferOffsetSingleTexelAlignment;
} VkPhysicalDeviceTexelBufferAlignmentProperties;
```

or the equivalent

``` c
// Provided by VK_EXT_texel_buffer_alignment
typedef VkPhysicalDeviceTexelBufferAlignmentProperties VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-storageTexelBufferOffsetAlignmentBytes"></span>
  `storageTexelBufferOffsetAlignmentBytes` is a byte alignment that is
  sufficient for a storage texel buffer of any format. The value
  **must** be a power of two.

- <span id="extension-limits-storageTexelBufferOffsetSingleTexelAlignment"></span>
  `storageTexelBufferOffsetSingleTexelAlignment` indicates whether
  single texel alignment is sufficient for a storage texel buffer of any
  format.

- <span id="extension-limits-uniformTexelBufferOffsetAlignmentBytes"></span>
  `uniformTexelBufferOffsetAlignmentBytes` is a byte alignment that is
  sufficient for a uniform texel buffer of any format. The value
  **must** be a power of two.

- <span id="extension-limits-uniformTexelBufferOffsetSingleTexelAlignment"></span>
  `uniformTexelBufferOffsetSingleTexelAlignment` indicates whether
  single texel alignment is sufficient for a uniform texel buffer of any
  format.

If the `VkPhysicalDeviceTexelBufferAlignmentProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

If the single texel alignment property is `VK_FALSE`, then the buffer
view’s offset **must** be aligned to the corresponding byte alignment
value. If the single texel alignment property is `VK_TRUE`, then the
buffer view’s offset **must** be aligned to the lesser of the
corresponding byte alignment value or the size of a single texel, based
on [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html)::`format`. If
the size of a single texel is a multiple of three bytes, then the size
of a single component of the format is used instead.

These limits **must** not advertise a larger alignment than the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-required"
target="_blank" rel="noopener">required</a> maximum minimum value of
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`minTexelBufferOffsetAlignment`,
for any format that supports use as a texel buffer.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceTexelBufferAlignmentProperties-sType-sType"
  id="VUID-VkPhysicalDeviceTexelBufferAlignmentProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceTexelBufferAlignmentProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_texel_buffer_alignment](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_texel_buffer_alignment.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceTexelBufferAlignmentProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
