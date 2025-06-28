# VkPhysicalDevicePresentMeteringFeaturesNV(3) Manual Page

## Name

VkPhysicalDevicePresentMeteringFeaturesNV - Structure describing whether the present metering features can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePresentMeteringFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_present_metering
typedef struct VkPhysicalDevicePresentMeteringFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentMetering;
} VkPhysicalDevicePresentMeteringFeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`presentMetering` indicates whether the implementation supports present metering capability.

## [](#_description)Description

If the `VkPhysicalDevicePresentMeteringFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePresentMeteringFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePresentMeteringFeaturesNV-sType-sType)VUID-VkPhysicalDevicePresentMeteringFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_METERING_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_present\_metering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_present_metering.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePresentMeteringFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0