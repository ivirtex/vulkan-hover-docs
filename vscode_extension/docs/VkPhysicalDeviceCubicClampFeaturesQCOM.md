# VkPhysicalDeviceCubicClampFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceCubicClampFeaturesQCOM - Structure describing cubic
clamp features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCubicClampFeaturesQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_filter_cubic_clamp
typedef struct VkPhysicalDeviceCubicClampFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           cubicRangeClamp;
} VkPhysicalDeviceCubicClampFeaturesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-filter-cubic-range-clamp"></span> `cubicRangeClamp`
  indicates that the implementation supports cubic filtering in
  combination with a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-range-clamp"
  target="_blank" rel="noopener">texel range clamp</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceCubicClampFeaturesQCOM` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceCubicClampFeaturesQCOM` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceCubicClampFeaturesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceCubicClampFeaturesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceCubicClampFeaturesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUBIC_CLAMP_FEATURES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_filter_cubic_clamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_filter_cubic_clamp.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCubicClampFeaturesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
