# VkPerformanceCounterDescriptionKHR(3) Manual Page

## Name

VkPerformanceCounterDescriptionKHR - Structure providing more detailed information about a counter



## [](#_c_specification)C Specification

The `VkPerformanceCounterDescriptionKHR` structure is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef struct VkPerformanceCounterDescriptionKHR {
    VkStructureType                            sType;
    void*                                      pNext;
    VkPerformanceCounterDescriptionFlagsKHR    flags;
    char                                       name[VK_MAX_DESCRIPTION_SIZE];
    char                                       category[VK_MAX_DESCRIPTION_SIZE];
    char                                       description[VK_MAX_DESCRIPTION_SIZE];
} VkPerformanceCounterDescriptionKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPerformanceCounterDescriptionFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionFlagBitsKHR.html) indicating the usage behavior for the counter.
- `name` is an array of size `VK_MAX_DESCRIPTION_SIZE`, containing a null-terminated UTF-8 string specifying the name of the counter.
- `category` is an array of size `VK_MAX_DESCRIPTION_SIZE`, containing a null-terminated UTF-8 string specifying the category of the counter.
- `description` is an array of size `VK_MAX_DESCRIPTION_SIZE`, containing a null-terminated UTF-8 string specifying the description of the counter.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerformanceCounterDescriptionKHR-sType-sType)VUID-VkPerformanceCounterDescriptionKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_DESCRIPTION_KHR`
- [](#VUID-VkPerformanceCounterDescriptionKHR-pNext-pNext)VUID-VkPerformanceCounterDescriptionKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkPerformanceCounterDescriptionFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceCounterDescriptionKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0