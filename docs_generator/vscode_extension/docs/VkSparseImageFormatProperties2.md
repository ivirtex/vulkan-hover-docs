# VkSparseImageFormatProperties2(3) Manual Page

## Name

VkSparseImageFormatProperties2 - Structure specifying sparse image format properties



## [](#_c_specification)C Specification

The `VkSparseImageFormatProperties2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkSparseImageFormatProperties2 {
    VkStructureType                  sType;
    void*                            pNext;
    VkSparseImageFormatProperties    properties;
} VkSparseImageFormatProperties2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_physical_device_properties2
typedef VkSparseImageFormatProperties2 VkSparseImageFormatProperties2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `properties` is a [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties.html) structure which is populated with the same values as in [vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSparseImageFormatProperties2-sType-sType)VUID-VkSparseImageFormatProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2`
- [](#VUID-VkSparseImageFormatProperties2-pNext-pNext)VUID-VkSparseImageFormatProperties2-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html), [vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageFormatProperties2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0