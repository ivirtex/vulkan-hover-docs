# VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV - Structure
describing fragment shading rate limits that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV` structure is
defined as:

``` c
// Provided by VK_NV_fragment_shading_rate_enums
typedef struct VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV {
    VkStructureType          sType;
    void*                    pNext;
    VkSampleCountFlagBits    maxFragmentShadingRateInvocationCount;
} VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxFragmentShadingRateInvocationCount"></span>
  `maxFragmentShadingRateInvocationCount` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value indicating
  the maximum number of fragment shader invocations per fragment
  supported in pipeline, primitive, and attachment fragment shading
  rates.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These properties are related to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate"
target="_blank" rel="noopener">fragment shading rates</a>.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_PROPERTIES_NV`

- <a
  href="#VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-maxFragmentShadingRateInvocationCount-parameter"
  id="VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-maxFragmentShadingRateInvocationCount-parameter"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-maxFragmentShadingRateInvocationCount-parameter  
  `maxFragmentShadingRateInvocationCount` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_fragment_shading_rate_enums](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
