# vkGetPhysicalDeviceFeatures2(3) Manual Page

## Name

vkGetPhysicalDeviceFeatures2 - Reports capabilities of a physical device



## [](#_c_specification)C Specification

To query supported features defined by the core or extensions, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceFeatures2(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceFeatures2*                  pFeatures);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceFeatures2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceFeatures2*                  pFeatures);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the supported features.
- `pFeatures` is a pointer to a [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure in which the physical device features are returned.

## [](#_description)Description

Each structure in `pFeatures` and its `pNext` chain contains members corresponding to fine-grained features. Each structure in `pFeatures` and its `pNext` chain contains [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) members corresponding to fine-grained features. Each such member is returned with a `VK_TRUE` value indicating that feature is supported on this physical device, or a `VK_FALSE` value indicating it is unsupported.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceFeatures2-physicalDevice-parameter)VUID-vkGetPhysicalDeviceFeatures2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceFeatures2-pFeatures-parameter)VUID-vkGetPhysicalDeviceFeatures2-pFeatures-parameter  
  `pFeatures` **must** be a valid pointer to a [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceFeatures2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0