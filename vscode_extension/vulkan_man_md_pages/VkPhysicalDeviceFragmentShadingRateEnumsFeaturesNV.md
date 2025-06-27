# VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV - Structure
indicating support for fragment shading rate enums



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_fragment_shading_rate_enums
typedef struct VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           fragmentShadingRateEnums;
    VkBool32           supersampleFragmentShadingRates;
    VkBool32           noInvocationFragmentShadingRates;
} VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-fragmentShadingRateEnums"></span>
  `fragmentShadingRateEnums` indicates that the implementation supports
  specifying fragment shading rates using the `VkFragmentShadingRateNV`
  enumerated type.

- <span id="features-supersampleFragmentShadingRates"></span>
  `supersampleFragmentShadingRates` indicates that the implementation
  supports fragment shading rate enum values indicating more than one
  invocation per fragment.

- <span id="features-noInvocationFragmentShadingRates"></span>
  `noInvocationFragmentShadingRates` indicates that the implementation
  supports a fragment shading rate enum value indicating that no
  fragment shaders should be invoked when that shading rate is used.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_fragment_shading_rate_enums](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
