# VkPhysicalDeviceMultiDrawPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMultiDrawPropertiesEXT - Structure describing multidraw limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMultiDrawPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_multi_draw
typedef struct VkPhysicalDeviceMultiDrawPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxMultiDrawCount;
} VkPhysicalDeviceMultiDrawPropertiesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceMultiDrawPropertiesEXT` structure describe the following features:

## [](#_description)Description

- []()`maxMultiDrawCount` indicates the maximum number of draw calls which **can** be batched into a single multidraw.

If the `VkPhysicalDeviceMultiDrawPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMultiDrawPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceMultiDrawPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_multi\_draw](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multi_draw.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMultiDrawPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0