# VkPhysicalDeviceBorderColorSwizzleFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceBorderColorSwizzleFeaturesEXT - Structure describing
whether samplers with custom border colors require the component swizzle
specified in order to have defined behavior



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceBorderColorSwizzleFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_border_color_swizzle
typedef struct VkPhysicalDeviceBorderColorSwizzleFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           borderColorSwizzle;
    VkBool32           borderColorSwizzleFromImage;
} VkPhysicalDeviceBorderColorSwizzleFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-borderColorSwizzle"></span> `borderColorSwizzle`
  indicates that defined values are returned by sampled image operations
  when used with a sampler that uses a
  `VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK`,
  `VK_BORDER_COLOR_INT_OPAQUE_BLACK`,
  `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`, or
  `VK_BORDER_COLOR_INT_CUSTOM_EXT` `borderColor` and an image view that
  uses a non-<a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity component mapping</a>, when
  either `borderColorSwizzleFromImage` is enabled or the
  [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html)
  is specified.

- <span id="features-borderColorSwizzleFromImage"></span>
  `borderColorSwizzleFromImage` indicates that the implementation will
  return the correct border color values from sampled image operations
  under the conditions expressed above, without the application having
  to specify the border color component mapping when creating the
  sampler object. If this feature bit is not set, applications **can**
  chain a
  [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html)
  structure when creating samplers for use with image views that do not
  have an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a> and, when those
  samplers are combined with image views using the same component
  mapping, sampled image operations that use opaque black or custom
  border colors will return the correct border color values.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceBorderColorSwizzleFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceBorderColorSwizzleFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceBorderColorSwizzleFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceBorderColorSwizzleFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceBorderColorSwizzleFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BORDER_COLOR_SWIZZLE_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_border_color_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_border_color_swizzle.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceBorderColorSwizzleFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
