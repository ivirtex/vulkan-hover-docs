# VkPerformanceCounterDescriptionKHR(3) Manual Page

## Name

VkPerformanceCounterDescriptionKHR - Structure providing more detailed
information about a counter



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceCounterDescriptionKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPerformanceCounterDescriptionFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionFlagBitsKHR.html)
  indicating the usage behavior for the counter.

- `name` is an array of size `VK_MAX_DESCRIPTION_SIZE`, containing a
  null-terminated UTF-8 string specifying the name of the counter.

- `category` is an array of size `VK_MAX_DESCRIPTION_SIZE`, containing a
  null-terminated UTF-8 string specifying the category of the counter.

- `description` is an array of size `VK_MAX_DESCRIPTION_SIZE`,
  containing a null-terminated UTF-8 string specifying the description
  of the counter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceCounterDescriptionKHR-sType-sType"
  id="VUID-VkPerformanceCounterDescriptionKHR-sType-sType"></a>
  VUID-VkPerformanceCounterDescriptionKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_DESCRIPTION_KHR`

- <a href="#VUID-VkPerformanceCounterDescriptionKHR-pNext-pNext"
  id="VUID-VkPerformanceCounterDescriptionKHR-pNext-pNext"></a>
  VUID-VkPerformanceCounterDescriptionKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPerformanceCounterDescriptionFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionFlagsKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterDescriptionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
