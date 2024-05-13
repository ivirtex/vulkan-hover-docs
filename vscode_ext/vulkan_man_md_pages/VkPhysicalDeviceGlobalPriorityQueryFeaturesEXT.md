# VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR - Structure describing
whether global priority query can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_global_priority
typedef struct VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           globalPriorityQuery;
} VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_global_priority_query
typedef VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-globalPriorityQuery"></span> `globalPriorityQuery`
  indicates whether the implementation supports the ability to query
  global queue priorities.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
