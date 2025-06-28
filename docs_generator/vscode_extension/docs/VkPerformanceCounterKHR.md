# VkPerformanceCounterKHR(3) Manual Page

## Name

VkPerformanceCounterKHR - Structure providing information about a counter



## [](#_c_specification)C Specification

The `VkPerformanceCounterKHR` structure is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef struct VkPerformanceCounterKHR {
    VkStructureType                   sType;
    void*                             pNext;
    VkPerformanceCounterUnitKHR       unit;
    VkPerformanceCounterScopeKHR      scope;
    VkPerformanceCounterStorageKHR    storage;
    uint8_t                           uuid[VK_UUID_SIZE];
} VkPerformanceCounterKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `unit` is a [VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterUnitKHR.html) specifying the unit that the counter data will record.
- `scope` is a [VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterScopeKHR.html) specifying the scope that the counter belongs to.
- `storage` is a [VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterStorageKHR.html) specifying the storage type that the counter’s data uses.
- `uuid` is an array of size `VK_UUID_SIZE`, containing 8-bit values that represent a universally unique identifier for the counter of the physical device.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerformanceCounterKHR-sType-sType)VUID-VkPerformanceCounterKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_KHR`
- [](#VUID-VkPerformanceCounterKHR-pNext-pNext)VUID-VkPerformanceCounterKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterScopeKHR.html), [VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterStorageKHR.html), [VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterUnitKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceCounterKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0