# VkPhysicalDeviceTexelBufferAlignmentProperties(3) Manual Page

## Name

VkPhysicalDeviceTexelBufferAlignmentProperties - Structure describing the texel buffer alignment requirements supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTexelBufferAlignmentProperties` structure is defined as:

```c++
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

```c++
// Provided by VK_EXT_texel_buffer_alignment
typedef VkPhysicalDeviceTexelBufferAlignmentProperties VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`storageTexelBufferOffsetAlignmentBytes` is a byte alignment that is sufficient for a storage texel buffer of any format. The value **must** be a power of two.
- []()`storageTexelBufferOffsetSingleTexelAlignment` indicates whether single texel alignment is sufficient for a storage texel buffer of any format.
- []()`uniformTexelBufferOffsetAlignmentBytes` is a byte alignment that is sufficient for a uniform texel buffer of any format. The value **must** be a power of two.
- []()`uniformTexelBufferOffsetSingleTexelAlignment` indicates whether single texel alignment is sufficient for a uniform texel buffer of any format.

If the `VkPhysicalDeviceTexelBufferAlignmentProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

If the single texel alignment property is `VK_FALSE`, then the buffer view’s offset **must** be aligned to the corresponding byte alignment value. If the single texel alignment property is `VK_TRUE`, then the buffer view’s offset **must** be aligned to the lesser of the corresponding byte alignment value or the size of a single texel, based on [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html)::`format`. If the size of a single texel is a multiple of three bytes, then the size of a single component of the format is used instead.

These limits **must** not advertise a larger alignment than the [required](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-required) maximum minimum value of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`minTexelBufferOffsetAlignment`, for any format that supports use as a texel buffer.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTexelBufferAlignmentProperties-sType-sType)VUID-VkPhysicalDeviceTexelBufferAlignmentProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_texel\_buffer\_alignment](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_texel_buffer_alignment.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTexelBufferAlignmentProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0