# vkGetPhysicalDeviceMultisamplePropertiesEXT(3) Manual Page

## Name

vkGetPhysicalDeviceMultisamplePropertiesEXT - Report sample count specific multisampling capabilities of a physical device



## [](#_c_specification)C Specification

To query additional multisampling capabilities which **may** be supported for a specific sample count, beyond the minimum capabilities described for [Limits](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits) above, call:

```c++
// Provided by VK_EXT_sample_locations
void vkGetPhysicalDeviceMultisamplePropertiesEXT(
    VkPhysicalDevice                            physicalDevice,
    VkSampleCountFlagBits                       samples,
    VkMultisamplePropertiesEXT*                 pMultisampleProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the additional multisampling capabilities.
- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the sample count to query capabilities for.
- `pMultisampleProperties` is a pointer to a [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisamplePropertiesEXT.html) structure in which information about additional multisampling capabilities specific to the sample count is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-physicalDevice-parameter)VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-samples-parameter)VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-samples-parameter  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-pMultisampleProperties-parameter)VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-pMultisampleProperties-parameter  
  `pMultisampleProperties` **must** be a valid pointer to a [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisamplePropertiesEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisamplePropertiesEXT.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceMultisamplePropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0