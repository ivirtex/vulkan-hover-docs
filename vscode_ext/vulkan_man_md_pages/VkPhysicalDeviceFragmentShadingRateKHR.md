# VkPhysicalDeviceFragmentShadingRateKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateKHR - Structure returning information
about sample count specific additional multisampling capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShadingRateKHR` structure is defined as

``` c
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkPhysicalDeviceFragmentShadingRateKHR {
    VkStructureType       sType;
    void*                 pNext;
    VkSampleCountFlags    sampleCounts;
    VkExtent2D            fragmentSize;
} VkPhysicalDeviceFragmentShadingRateKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sampleCounts` is a bitmask of sample counts for which the shading
  rate described by `fragmentSize` is supported.

- `fragmentSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) describing the width
  and height of a supported shading rate.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceFragmentShadingRateKHR-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShadingRateKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_KHR`

- <a href="#VUID-VkPhysicalDeviceFragmentShadingRateKHR-pNext-pNext"
  id="VUID-VkPhysicalDeviceFragmentShadingRateKHR-pNext-pNext"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceFragmentShadingRatesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFragmentShadingRatesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
