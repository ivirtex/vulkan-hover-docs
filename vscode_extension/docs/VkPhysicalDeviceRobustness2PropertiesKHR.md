# VkPhysicalDeviceRobustness2PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceRobustness2PropertiesKHR - Structure describing robust buffer access properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRobustness2PropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_robustness2
typedef struct VkPhysicalDeviceRobustness2PropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       robustStorageBufferAccessSizeAlignment;
    VkDeviceSize       robustUniformBufferAccessSizeAlignment;
} VkPhysicalDeviceRobustness2PropertiesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_robustness2
typedef VkPhysicalDeviceRobustness2PropertiesKHR VkPhysicalDeviceRobustness2PropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`robustStorageBufferAccessSizeAlignment` is the number of bytes that the range of a storage buffer descriptor is rounded up to when used for bounds-checking when the [`robustBufferAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess2) feature is enabled. This value **must** be either 1 or 4.
- []()`robustUniformBufferAccessSizeAlignment` is the number of bytes that the range of a uniform buffer descriptor is rounded up to when used for bounds-checking when the [`robustBufferAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess2) feature is enabled. This value **must** be a power of two in the range \[1, 256].

## [](#_description)Description

If the `VkPhysicalDeviceRobustness2PropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRobustness2PropertiesKHR-sType-sType)VUID-VkPhysicalDeviceRobustness2PropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_robustness2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_robustness2.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRobustness2PropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0