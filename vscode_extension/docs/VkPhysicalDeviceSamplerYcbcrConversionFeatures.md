# VkPhysicalDeviceSamplerYcbcrConversionFeatures(3) Manual Page

## Name

VkPhysicalDeviceSamplerYcbcrConversionFeatures - Structure describing Y′CBCR conversion features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSamplerYcbcrConversionFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceSamplerYcbcrConversionFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           samplerYcbcrConversion;
} VkPhysicalDeviceSamplerYcbcrConversionFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkPhysicalDeviceSamplerYcbcrConversionFeatures VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`samplerYcbcrConversion` specifies whether the implementation supports [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion). If `samplerYcbcrConversion` is `VK_FALSE`, sampler Y′CBCR conversion is not supported, and samplers using sampler Y′CBCR conversion **must** not be used.

If the `VkPhysicalDeviceSamplerYcbcrConversionFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceSamplerYcbcrConversionFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSamplerYcbcrConversionFeatures-sType-sType)VUID-VkPhysicalDeviceSamplerYcbcrConversionFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES`

## [](#_see_also)See Also

[VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSamplerYcbcrConversionFeatures)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0