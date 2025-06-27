# VkPerformanceCounterKHR(3) Manual Page

## Name

VkPerformanceCounterKHR - Structure providing information about a
counter



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceCounterKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `unit` is a
  [VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterUnitKHR.html)
  specifying the unit that the counter data will record.

- `scope` is a
  [VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterScopeKHR.html)
  specifying the scope that the counter belongs to.

- `storage` is a
  [VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterStorageKHR.html)
  specifying the storage type that the counter’s data uses.

- `uuid` is an array of size `VK_UUID_SIZE`, containing 8-bit values
  that represent a universally unique identifier for the counter of the
  physical device.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceCounterKHR-sType-sType"
  id="VUID-VkPerformanceCounterKHR-sType-sType"></a>
  VUID-VkPerformanceCounterKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_KHR`

- <a href="#VUID-VkPerformanceCounterKHR-pNext-pNext"
  id="VUID-VkPerformanceCounterKHR-pNext-pNext"></a>
  VUID-VkPerformanceCounterKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterScopeKHR.html),
[VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterStorageKHR.html),
[VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterUnitKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
