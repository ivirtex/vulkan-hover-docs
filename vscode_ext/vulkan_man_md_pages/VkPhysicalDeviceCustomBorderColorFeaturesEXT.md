# VkPhysicalDeviceCustomBorderColorFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceCustomBorderColorFeaturesEXT - Structure describing
whether custom border colors can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCustomBorderColorFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_custom_border_color
typedef struct VkPhysicalDeviceCustomBorderColorFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           customBorderColors;
    VkBool32           customBorderColorWithoutFormat;
} VkPhysicalDeviceCustomBorderColorFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-customBorderColors"></span> `customBorderColors`
  indicates that the implementation supports providing a `borderColor`
  value with one of the following values at sampler creation time:

  - `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`

  - `VK_BORDER_COLOR_INT_CUSTOM_EXT`

- <span id="features-customBorderColorWithoutFormat"></span>
  `customBorderColorWithoutFormat` indicates that explicit formats are
  not required for custom border colors and the value of the `format`
  member of the
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
  structure **may** be `VK_FORMAT_UNDEFINED`. If this feature bit is not
  set, applications **must** provide the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  the image view(s) being sampled by this sampler in the `format` member
  of the
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceCustomBorderColorFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceCustomBorderColorFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceCustomBorderColorFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceCustomBorderColorFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceCustomBorderColorFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_custom_border_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_custom_border_color.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCustomBorderColorFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
