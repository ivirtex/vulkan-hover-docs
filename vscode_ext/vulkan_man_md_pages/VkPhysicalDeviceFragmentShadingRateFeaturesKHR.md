# VkPhysicalDeviceFragmentShadingRateFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateFeaturesKHR - Structure indicating
support for variable rate fragment shading



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShadingRateFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkPhysicalDeviceFragmentShadingRateFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineFragmentShadingRate;
    VkBool32           primitiveFragmentShadingRate;
    VkBool32           attachmentFragmentShadingRate;
} VkPhysicalDeviceFragmentShadingRateFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-pipelineFragmentShadingRate"></span>
  `pipelineFragmentShadingRate` indicates that the implementation
  supports the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-pipeline"
  target="_blank" rel="noopener">pipeline fragment shading rate</a>.

- <span id="features-primitiveFragmentShadingRate"></span>
  `primitiveFragmentShadingRate` indicates that the implementation
  supports the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
  target="_blank" rel="noopener">primitive fragment shading rate</a>.

- <span id="features-attachmentFragmentShadingRate"></span>
  `attachmentFragmentShadingRate` indicates that the implementation
  supports the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">attachment fragment shading rate</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentShadingRateFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceFragmentShadingRateFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentShadingRateFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShadingRateFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
