# VkPhysicalDeviceYcbcrDegammaFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceYcbcrDegammaFeaturesQCOM - Structure describing
Y′C<sub>B</sub>C<sub>R</sub> degamma features that can be supported by
an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceYcbcrDegammaFeaturesQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_ycbcr_degamma
typedef struct VkPhysicalDeviceYcbcrDegammaFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           ycbcrDegamma;
} VkPhysicalDeviceYcbcrDegammaFeaturesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-ycbcr-degamma"></span> `ycbcrDegamma` indicates
  whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-ycbcr-degamma"
  target="_blank" rel="noopener">Y′C<sub>B</sub>C<sub>R</sub> degamma</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceYcbcrDegammaFeaturesQCOM` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceYcbcrDegammaFeaturesQCOM` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceYcbcrDegammaFeaturesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceYcbcrDegammaFeaturesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceYcbcrDegammaFeaturesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_DEGAMMA_FEATURES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_ycbcr_degamma](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_ycbcr_degamma.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceYcbcrDegammaFeaturesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
