# VkPhysicalDeviceGlobalPriorityQueryFeatures(3) Manual Page

## Name

VkPhysicalDeviceGlobalPriorityQueryFeatures - Structure describing whether global priority query can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceGlobalPriorityQueryFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceGlobalPriorityQueryFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           globalPriorityQuery;
} VkPhysicalDeviceGlobalPriorityQueryFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_global_priority
typedef VkPhysicalDeviceGlobalPriorityQueryFeatures VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_global_priority_query
typedef VkPhysicalDeviceGlobalPriorityQueryFeatures VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`globalPriorityQuery` indicates whether the implementation supports the ability to query global queue priorities.

If the `VkPhysicalDeviceGlobalPriorityQueryFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceGlobalPriorityQueryFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceGlobalPriorityQueryFeatures-sType-sType)VUID-VkPhysicalDeviceGlobalPriorityQueryFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES`

## [](#_see_also)See Also

[VK\_EXT\_global\_priority\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_global_priority_query.html), [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceGlobalPriorityQueryFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0